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

func main() {
	// Create new watcher. 
    tempIdlPath := "./examples/hello/idl_temp"
    watcher := newWatcher(tempIdlPath)
    defer watcher.Close()

    idlPath := "./examples/hello/http-server/idl"
    w := newWatcher(idlPath)
    defer w.Close()

    // Copy and edit IDL file at /temp_idl/
    // rename IDL file to [idl_name]_temp.thrift
    // Validate with thriftgo
    // Copy over new IDL file to /http-server/idl/
    // Use new IDL file -> if error, use old IDL file, log error
    // -> if success, rename new IDL file, and replace old IDL file

    // Start listening for events.
    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
				// log.Println("Event:", event)
				// TODO
				switch {
					case event.Has(fsnotify.Write):
						log.Println("Modified file:" + event.Name)
					case event.Has(fsnotify.Create):
						log.Println("Created new file:" + event.Name)
                        // TODO: Use the parser to validate
                        _, err := parser.ParseFile(event.Name, []string{""}, true)

                        if err != nil {
                            log.Println("Handle Error here")
                        }
                    
                        log.Println("Valid IDL")

					case event.Has(fsnotify.Remove):
						log.Println("Removed file: " + event.Name)
					case event.Has(fsnotify.Rename):
						log.Println("Renamed file: " + event.Name)
					case event.Has(fsnotify.Chmod):
						log.Printf("Chmod: " + event.Name)
				}
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Println("error:", err)
            }
        }
    }()

    go func() {
        for {
            select {
            case event, ok := <-w.Events:
                if !ok {
                    return
                }
				// log.Println("Event:", event)
				// TODO
				switch {
					case event.Has(fsnotify.Write):
						log.Println("Modified file:" + event.Name)
					case event.Has(fsnotify.Create):
						log.Println("Created new file: " + event.Name)
                        log.Println("CHECK IF IDL WORKS FINE HERE")
                        // TODO: Check if new IDL file works fine
                        // _, err := check()

                        // if err != nil {
                        //     return 
                        // }
					case event.Has(fsnotify.Remove):
						log.Println("Removed file: " + event.Name)
					case event.Has(fsnotify.Rename):
						log.Println("Renamed file: " + event.Name)
					case event.Has(fsnotify.Chmod):
						log.Printf("Chmod: " + event.Name)
				}
            case err, ok := <-w.Errors:
                if !ok {
                    return
                }
                log.Println("error:", err)
            }
        }
    }()

    // Block main goroutine forever.
    <-make(chan struct{})
}