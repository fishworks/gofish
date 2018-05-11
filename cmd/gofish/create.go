package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/fishworks/gofish"
	"github.com/spf13/cobra"
)

const createTpl = `local name = "{{ .Name }}"
local version = "0.1.0"

food = {
    name = name,
    description = "enter description here",
    homepage = "https://gofi.sh",
    version = version,
    packages = {
        {
            os = "darwin",
            arch = "amd64",
            url = "https://github.com/example/" .. name .. "/releases/download/v" .. version .. "/" .. name .. "-v" .. version .. "-darwin-amd64.tar.gz",
            -- shasum of the release archive
            sha256 = "",
            resources = {
                {
                    path = name,
                    installpath = "bin/" .. name,
                    executable = true
                }
            }
        },
        {
            os = "linux",
            arch = "amd64",
            url = "https://github.com/example/" .. name .. "/releases/download/v" .. version .. "/" .. name .. "-v" .. version .. "-linux-amd64.tar.gz",
            -- shasum of the release archive
            sha256 = "",
            resources = {
                {
                    path = name,
                    installpath = "bin/" .. name,
                    executable = true
                }
            }
        },
        {
            os = "windows",
            arch = "amd64",
            url = "https://github.com/example/" .. name .. "/releases/download/v" .. version .. "/" .. name .. "-v" .. version .. "windows-amd64.zip",
            -- shasum of the release archive
            sha256 = "",
            resources = {
                {
                    path = name .. ".exe",
                    installpath = "bin\\" .. name .. ".exe"
                }
            }
        }
    }
}
`

func newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <food>",
		Short: "generate fish food and open it in the editor",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			home := gofish.Home(gofish.HomePath)
			destPath := filepath.Join(home.Rigs(), home.DefaultRig(), "Food", fmt.Sprintf("%s.lua", args[0]))
			f, err := os.Create(destPath)
			if err != nil {
				return nil
			}
			defer f.Close()
			t := template.Must(template.New("create").Parse(createTpl))
			if err := t.Execute(f, struct{ Name string }{args[0]}); err != nil {
				return err
			}
			fmt.Println(destPath)
			return nil
		},
	}
	return cmd
}
