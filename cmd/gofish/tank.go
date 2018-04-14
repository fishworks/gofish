package main

import "github.com/fishworks/gofish"

type tank map[string]string

func (t tank) fill() {
	fishHome := gofish.Home(gofish.HomePath)
	userHome := gofish.UserHome(gofish.UserHomePath)

	t["GOFISH_HOME"] = fishHome.String()
	t["GOFISH_CACHE"] = userHome.Cache()
	t["GOFISH_BARREL"] = fishHome.Barrel()
	t["GOFISH_RIGS"] = fishHome.Rigs()
	t["GOFISH_DEFAULT_RIG"] = fishHome.DefaultRig()
}
