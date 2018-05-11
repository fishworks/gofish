package main

import (
	"reflect"
	"testing"

	"github.com/fishworks/gofish"
)

func TestFindFishFood(t *testing.T) {
	gofish.HomePath = "testdata"

	expectedFoodList := []string{
		"hugo",
		"github.com/myorg/fish-food/hugo",
	}

	foodList := findFishFood()
	if !reflect.DeepEqual(findFishFood(), expectedFoodList) {
		t.Errorf("expected fish food lists to be equal; got '%v', wanted '%v'", foodList, expectedFoodList)
	}
}
