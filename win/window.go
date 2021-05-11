//Package win provides a simple way to open a gpu-compatible window in Go.
package win

import "errors"

var CurrentDriver Driver

func Open(name string) error {
	if CurrentDriver == nil {
		return errors.New("no driver was imported for win")
	}
	return CurrentDriver.Open(name)
}

func Update() bool {
	if CurrentDriver == nil {
		return false
	}
	return CurrentDriver.Update()
}

func Close() {
	if CurrentDriver == nil {
		return
	}
	CurrentDriver.Close()
}

func Button(id string) bool {
	if CurrentDriver == nil {
		return false
	}
	return CurrentDriver.Button(id)
}
