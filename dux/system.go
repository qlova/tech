//Package dux provides a data-oriented unified component system.
package dux

import (
	"reflect"
)

//System inspired by ECS.
type System interface {
	component()

	//Update the system.
	Update()
}

//Component needs to be embedded inside of a struct to create an automatic component.
//Acts as a tag to indicate that something is a component.
type Component struct{}

func (component *Component) component() {}

//ComponentSystem is a component or a system.
type ComponentSystem interface {
	component()
}

//Set is a set of systems.
type Set struct {
	storage map[reflect.Type]reflect.Value
	updates []func()
}

//New returns a new set of systems.
func New(components ...ComponentSystem) Set {
	var set Set

	set.storage = make(map[reflect.Type]reflect.Value)

	for _, component := range components {
		if system, ok := component.(System); ok {
			set.updates = append(set.updates, system.Update)
		}
		set.storage[reflect.TypeOf(component)] = reflect.ValueOf(component)
	}

	for _, component := range components {
		var T = reflect.TypeOf(component)
		var value = reflect.ValueOf(component)

		if value.Kind() != reflect.Ptr {
			continue
		}

		T = T.Elem()
		value = value.Elem()

		if value.Kind() == reflect.Struct {
			for i := 0; i < value.NumField(); i++ {
				var field = T.Field(i)

				//Automatic system referencing.
				if field.Type.Implements(reflect.TypeOf([0]System{}).Elem()) && field.Type.Kind() == reflect.Ptr && field.PkgPath == "" {
					if system, ok := set.storage[field.Type]; ok {
						value.Field(i).Set(system)
					}
				} else {
					//Automatic component allocation.
					if field.Type.Implements(reflect.TypeOf([0]ComponentSystem{}).Elem()) && field.Type.Kind() == reflect.Ptr && field.PkgPath == "" {
						component, ok := set.storage[field.Type]
						if !ok {
							component = reflect.New(field.Type.Elem())
							set.storage[field.Type] = component
						}
						value.Field(i).Set(component)
					}
				}
			}
		}
	}

	return set
}

//Update runs the update method of each system in the set.
func (set Set) Update() {
	for _, update := range set.updates {
		update()
	}
}
