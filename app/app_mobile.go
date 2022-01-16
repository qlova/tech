//go:build android

package app

import (
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/gl"
	"qlova.tech/gpu"
)

var GL gl.Context

var appName string
var appSystems []System

func open(name string, systems ...System) error {
	app.Main(func(a app.App) {
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					GL, _ = e.DrawContext.(gl.Context)
					load(systems)
					a.Send(paint.Event{})
				case lifecycle.CrossOff:
					GL = nil
				}
			case size.Event:
				bounds := e.Bounds()
				Width, Height = bounds.Dx(), bounds.Dy()
			case paint.Event:

				for _, system := range systems {
					system.Update()
				}

				gpu.Sync()
				a.Publish()
				a.Send(paint.Event{}) // keep animating
			}
		}
	})
	return nil
}
