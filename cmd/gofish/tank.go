package main

import "github.com/fishworks/gofish"

type tank map[string]string

func (t tank) fill() {
	fishHome := gofish.Home(gofish.HomePath)
	userHome := gofish.UserHome(gofish.UserHomePath)

	t["FISH_HOME"] = fishHome.String()
	t["FISH_CACHE"] = userHome.Cache()
	t["FISH_BARREL"] = fishHome.Barrel()
	t["FISH_RIGS"] = fishHome.Rigs()
	t["FISH_DEFAULT_RIG"] = fishHome.DefaultRig()
}
