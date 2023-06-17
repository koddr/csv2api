//go:build wireinject

package main

import "github.com/google/wire"

// initialize ...
func initialize() (*App, error) {
	wire.Build(newConfig, newInputs, newFiltered, newOutputs, newApp)
	return &App{}, nil
}
