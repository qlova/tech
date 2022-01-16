// Package app provides an application runtime and component-system.
package app

import (
	"reflect"

	"qlova.tech/dsl"
	"qlova.tech/gpu"
)

var (
	// desired width and height for the app.
	Width, Height int = 800, 600
)

type shader interface {
	Vertex(dsl.Core)
	Fragment(dsl.Core)
}

type System interface {
	system()
	Update()
}

type Runtime struct {
	name    string
	systems []System
}

func Open(name string, systems ...System) error {
	open(name)

	if err := gpu.Open(); err != nil {
		return err
	}

	for i, system := range systems {
		if loader, ok := system.(Loader); ok {
			if err := loader.Load(); err != nil {
				return err
			}
		}

		if shader, ok := system.(shader); ok {
			program := reflect.ValueOf(shader).Elem().FieldByName("Program")
			if program.IsValid() {

				p, err := gpu.NewProgram(shader.Vertex, shader.Fragment)
				if err != nil {
					return err
				}

				program.Set(reflect.ValueOf(p))
			}
		}

		systems[i] = system
	}

	return launch(systems...)
}

type Loader interface {
	System

	Load() error
}
