package main

import (
	"github.com/fishworks/fish"
)

type tank map[string]string

func (t tank) fill() {
	fishHome := fish.Home(fish.HomePath)
	userHome := fish.UserHome(fish.UserHomePath)

	t["FISH_HOME"] = fishHome.String()
	t["FISH_CACHE"] = userHome.Cache()
	t["FISH_BARREL"] = fishHome.Barrel()
	t["FISH_RIGS"] = fishHome.Rigs()
	t["FISH_DEFAULT_RIG"] = fishHome.DefaultRig()
}
