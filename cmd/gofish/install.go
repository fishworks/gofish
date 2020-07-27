package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/fishworks/gofish/receipt"

	"github.com/fishworks/gofish"
	"github.com/fishworks/gofish/pkg/home"
	"github.com/fishworks/gofish/pkg/ohai"
	"github.com/spf13/cobra"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

const installDesc = `
Install fish food.
`

func newInstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install <food...>",
		Short: "install fish food",
		Long:  installDesc,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, fishFood := range args {
				relevantFood := search([]string{fishFood})
				switch len(relevantFood) {
				case 0:
					return fmt.Errorf("no fish food with the name '%s' was found", fishFood)
				case 1:
					fishFood = relevantFood[0]
				default:
					var match bool
					// check if we have an exact match
					for _, f := range relevantFood {
						if strings.Compare(f, fishFood) == 0 {
							fishFood = f
							match = true
						}
					}
					if !match {
						return fmt.Errorf("%d fish food with the name '%s' was found: %v", len(relevantFood), fishFood, relevantFood)
					}
				}
				food, err := getFood(fishFood)
				if err != nil {
					return err
				}
				if len(findFoodVersions(fishFood)) > 0 {
					ohai.Ohaif("%s is already installed. Please use `gofish upgrade %s` to upgrade.\n", fishFood, fishFood)
					return nil
				}
				ohai.Ohaif("Installing %s...\n", fishFood)
				start := time.Now()
				if err := food.Install(); err != nil {
					if errors.Is(err, gofish.ErrCouldNotUnlink{}) {
						return fmt.Errorf("%s could not be 'unlinked' try running 'gofish unlink %s': %s", fishFood, fishFood, err.Error())
					} else if errors.Is(err, gofish.ErrCouldNotLink{}) {
						return fmt.Errorf("%s could not be 'linked' try running 'gofish link %s': %s", fishFood, fishFood, err.Error())
					} else {
						return err
					}
				}
				t := time.Now()
				ohai.Successf("%s %s: installed in %s\n", food.Name, food.Version, t.Sub(start).String())
			}
			return nil
		},
	}
	return cmd
}

func getFood(foodName string) (*gofish.Food, error) {
	var (
		name string
		rig  string
	)
	foodInfo := strings.Split(foodName, "/")
	if len(foodInfo) == 1 {
		name = foodInfo[0]
		rig = home.DefaultRig()
	} else {
		name = foodInfo[len(foodInfo)-1]
		rig = path.Dir(foodName)
	}
	if strings.Contains(name, "./\\") {
		return nil, fmt.Errorf("food name '%s' is invalid. Food names cannot include the following characters: './\\'", name)
	}

	// check if there's an install receipt available to check what rig this was installed from
	receiptFile, err := os.Open(filepath.Join(home.Barrel(), name, receipt.ReceiptFilename))
	if err == nil {
		defer receiptFile.Close()
		installReceipt, err := receipt.NewFromReader(receiptFile)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if installReceipt.Rig != "" {
			rig = installReceipt.Rig
		}
	} else if !os.IsNotExist(err) {
		ohai.Warningf("could not read from install receipt: %v", err)
	}

	l := lua.NewState()
	defer l.Close()
	if err := l.DoFile(filepath.Join(home.Rigs(), rig, "Food", fmt.Sprintf("%s.lua", name))); err != nil {
		return nil, err
	}
	var food gofish.Food
	if err := gluamapper.Map(l.GetGlobal(strings.ToLower(reflect.TypeOf(food).Name())).(*lua.LTable), &food); err != nil {
		return nil, err
	}
	food.Rig = rig
	return &food, nil
}
