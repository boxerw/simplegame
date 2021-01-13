package core

import "github.com/boxerw/simplegame/core/foundation"

type Component = foundation.Component

type ComponentBundle = foundation.ComponentBundle

func NewComponentBundle(name string, component Component) ComponentBundle {
	return ComponentBundle{
		Name:      name,
		Component: component,
	}
}
