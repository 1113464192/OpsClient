//go:build linux
// +build linux

package main

import (
	"ops_client/configs"
	"ops_client/internal/controller"
)

func main() {
	configs.Init()
	r := controller.NewRoute()
	r.Run(":9080")
}
