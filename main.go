package main

import "git-knowledge/app"

var APP *app.BootStrap

func main() {
	APP = app.NewBootstrap()
	APP.Start()
}
