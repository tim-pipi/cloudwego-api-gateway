package fsnotify

import (
    "log"

    "github.com/fsnotify/fsnotify"
    "github.com/cloudwego/thriftgo/parser"
)

func newWatcher(dir string) *fsnotify.Watcher {
    watcher, err := fsnotify.NewWatcher()
    log.Println("Running Watcher for " + dir)
    
    if err != nil {
        log.Fatal(err)
    }

    err = watcher.Add(dir)
    log.Println("Watching files at " + dir)

    if err != nil {
        log.Fatal(err)
    }


    return watcher
}

func runWatcher(watcher *fsnotify.Watcher) {
    for {
        select {
        case event, ok := <-watcher.Events:
            if !ok {
                return
            }

            if event.Has(fsnotify.Create) {
                log.Println("Created new file:" + event.Name)
                _, err := parser.ParseFile(event.Name, []string{""}, true)

                if err != nil {
                    log.Println("IDL file is invalid!")
                    continue
                }
            
                log.Println("IDL file is valid!")
            }

        case err, ok := <-watcher.Errors:
            if !ok {
                return
            }
            log.Println("error:", err)
        }
    }
}

// func main() {
// 	// Create new watcher.
//     idlPath := "./pkg/fsnotify/idltest"
//     watcher := newWatcher(idlPath)
//     defer watcher.Close()

//     // Copy and edit IDL file at /temp_idl/
//     // rename IDL file to [idl_name]_temp.thrift
//     // Validate with thriftgo
//     // Copy over new IDL file to /http-server/idl/
//     // Use new IDL file -> if error, use old IDL file, log error
//     // -> if success, rename new IDL file, and replace old IDL file

//     // Start listening for events.
//     go runWatcher(watcher)
//     // Block main goroutine forever.
//     <-make(chan struct{})
// }