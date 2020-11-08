package main

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
)

// This file generates mongodb-friendly structs by adding the `bson` tag
// to use it, first generate your project normally with gqlgen using
//
// to gen:
// 		go run github.com/99designs/gqlgen init
//
// to regen:
// 		go run github.com/99designs/gqlgen -v
//
// then run this file.
// more here: https://github.com/99designs/gqlgen/issues/865#issuecomment-573043996

func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		for _, field := range model.Fields {
			name := field.Name
			if name == "id" {
				name = "_id"
			}
			field.Tag += ` bson:"` + name + `"`
		}
	}
	return b
}

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	p := modelgen.Plugin{
		MutateHook: mutateHook,
	}

	err = api.Generate(cfg,
		api.NoPlugins(),
		api.AddPlugin(&p),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
