package fsnotify

import (
    "log"
    "os"
    "os/exec"

    "github.com/fsnotify/fsnotify"
    "github.com/cloudwego/thriftgo/parser"
)

func WatchIDLFiles(idlDir string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
                // if event.Has(fsnotify.Write) {
				// 	log.Printf("File %s modified, triggering code generation...", event.Name)
				// 	// Call the code generation process here
				// 	cmd := exec.Command("go", "generate", "your/package/containing/generated/code")
				// 	cmd.Stdout = os.Stdout
				// 	cmd.Stderr = os.Stderr
				// 	if err := cmd.Run(); err != nil {
				// 		log.Println("Code generation failed:", err)
				// 	}
				// }
                if event.Has(fsnotify.Create) {
                    log.Printf("File %s created, parsing IDL file...", event.Name)
                    _, err := parser.ParseFile(event.Name, []string{""}, true)
    
                    if err != nil {
                        log.Printf("File %s is invalid. Please ensure your file is correct.", event.Name)
                        continue
                    }
                
                    log.Printf("File %s is valid, performing hot reload...", event.Name)
                    // TODO: Update client here
                    cmd := exec.Command("echo", "New IDL update")
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr

                    if err := cmd.Run(); err != nil {
						log.Println("Code generation failed:", err)
					}
                }
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()
    
    // Add idlDir
	err = watcher.Add(idlDir)
	if err != nil {
		log.Fatal(err)
	}
    log.Println("Watching files at ", idlDir)
	<-done
}