package plugins

import (
	"github.com/mobliyishengyuan/goframework/library"
	"fmt"
)

func init() {
	library.RegisterPlugin(library.PluginPreOptParse, preOptParse)
	library.RegisterPlugin(library.PluginPreOptParse, postOptParse)
	library.RegisterPlugin(library.PluginPreOptParse, preServerStart)
}

func preOptParse() {
	fmt.Println("pre opt parse")
}

func postOptParse() {
	fmt.Println("post opt parse")
}

func preServerStart() {
	fmt.Println("pre server start")
}