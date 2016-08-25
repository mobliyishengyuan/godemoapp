package main

import (
	fr "github.com/mobliyishengyuan/goframework"
	_ "github.com/mobliyishengyuan/godemoapp/actions"
	_ "github.com/mobliyishengyuan/godemoapp/plugins"
)

func main() {
	fr.Run()
}