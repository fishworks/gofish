package main

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/fishworks/gofish"
)

func TestSearch(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	gofish.HomePath = filepath.Join(cwd, "testdata")

	expectedFoodList := []string{
		"github.com/customorg/fish-food/hugo",
		"hugo",
	}

	foodList := search([]string{})
	if !reflect.DeepEqual(foodList, expectedFoodList) {
		t.Errorf("expected fish food lists to be equal; got '%v', wanted '%v'", foodList, expectedFoodList)
	}
}
