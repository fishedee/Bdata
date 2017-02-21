package main

import (
	. "github.com/fishedee/web"
)

type MainAoModel struct {
	Model
}

func (this *MainAoModel) Go() {
	this.Log.Debug("Hello World")
}

func main() {
	RunMain((*MainAoModel).Go)
}
