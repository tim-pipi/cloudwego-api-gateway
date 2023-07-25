package main

import fsnotify "github.com/tim-pipi/cloudwego-api-gateway/pkg/fsnotify"

func main() {
	idlDir := "./internal/idltest"
	fsnotify.WatchIDLFiles(idlDir)
}