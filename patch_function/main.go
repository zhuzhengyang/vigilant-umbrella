package main

import (
	"fmt"
	"log"
	"plugin"

	"github.com/fsnotify/fsnotify"
)

var handler func() string

func init() {
	handler = func() string {
		return "Hello, init\n"
	}
}

func httpGet() {
	fmt.Printf("handler: %v %s", handler, handler())
}

func patch(path string) {
	p, err := plugin.Open(path)
	if err != nil {
		log.Fatal("error open plugin: ", err)
	}
	s, err := p.Lookup("PluginHandler")
	if err != nil {
		log.Fatal("error lookup PluginHandler: ", err)
	}
	var ok bool
	handler, ok = s.(func() string)
	if !ok {
		log.Fatal("handler load error")
	}
}

func main() {
	httpGet()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	watch()
}
