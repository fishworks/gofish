package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/fishworks/gofish/receipt"

	"github.com/fishworks/gofish"
	"github.com/fishworks/gofish/pkg/home"
	"github.com/fishworks/gofish/pkg/ohai"
	"github.com/spf13/cobra"
	ruby "github.com/SeekingMeaning/go-mruby"
)

const installDesc = `
Install fish food.
`

const rubyDsl = `
def food(&block)
  Food.new(&block)
end

class Food
  def initialize(&block)
    instance_eval(&block)
  end

  def name(val = nil); val.nil? ? @name : @name = val; end
  def rig; '@RIG@'; end
  def description(val = nil); val.nil? ? @description : @description = val; end
  def license(val = nil); val.nil? ? @license : @license = val; end
  def homepage(val = nil); val.nil? ? @homepage : @homepage = val; end
  def caveats(val = nil); val.nil? ? @caveats : @caveats = val; end
  def version(val = nil); val.nil? ? @version : @version = val; end
  def packages; @packages ||= []; end
  def package(&block); packages << Package.new(self, &block); end
  def preinstallscript(val = nil); val.nil? ? @preinstallscript : @preinstallscript = val; end
  def postinstallscript(val = nil); val.nil? ? @postinstallscript : @postinstallscript = val; end
end

class Package
  def initialize(food, &block)
    @food = food
    instance_eval(&block)
  end

  def name; @food.name; end
  def description; @food.description; end
  def license; @food.license; end
  def homepage; @food.homepage; end
  def version; @food.version; end

  def os(val = nil); val.nil? ? @os : @os = val; end
  def arch(val = nil); val.nil? ? @arch : @arch = val; end
  def url(val = nil); val.nil? ? @url : @url = val; end
  def sha256(val = nil); val.nil? ? @sha256 : @sha256 = val; end
  def resources; @resources ||= []; end
  def resource(&block); resources << Resource.new(@food, self, &block); end
  def mirrors; @mirrors ||= []; end
  def mirror(val); @mirrors << mirror; end
end

class Resource
  def initialize(food, package, &block)
    @food = food
    @package = package
    instance_eval(&block)
  end


  def name; @food.name; end
  def description; @food.description; end
  def license; @food.license; end
  def homepage; @food.homepage; end
  def version; @food.version; end

  def os; @package.os; end
  def arch; @package.arch; end
  def url; @package.url; end
  def sha256; @package.sha256; end

  def path(val = nil); val.nil? ? @path : @path = val; end
  def installpath(val = nil); val.nil? ? @installpath : @installpath = val; end
  def executable(val = nil); val.nil? ? @executable : @executable = val; end
end
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

	foodBytes, err := ioutil.ReadFile(filepath.Join(home.Rigs(), rig, "Food", fmt.Sprintf("%s.rb", name)))
	if err != nil {
		return nil, err
	}
	mrb := ruby.NewMrb()
	defer mrb.Close()
	if _, err := mrb.LoadString(strings.ReplaceAll(rubyDsl, "@RIG@", rig)); err != nil {
		return nil, err
	}
	value, err := mrb.LoadString(string(foodBytes))
	if err != nil {
		return nil, err
	}
	var food gofish.Food
	if err := ruby.Decode(&food, value); err != nil {
		return nil, err
	}
	return &food, nil
}
