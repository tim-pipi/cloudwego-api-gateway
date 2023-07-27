package fsnotify

import (
	"log"
	"os"
	"time"

	"github.com/tim-pipi/cloudwego-api-gateway/pkg/clientpool"

	"github.com/cloudwego/thriftgo/parser"
	"github.com/fsnotify/fsnotify"
)

// Watches changes to IDL files in a specified directory
func WatchIDLFiles(idlDir string, cp *clientpool.ClientPool) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		// Using timer to prevent fsnotify from firing twice upon write
		var (
	        timer     *time.Timer
	        lastEvent fsnotify.Event
        )
        timer = time.NewTimer(time.Millisecond)
        <-timer.C // timer should be expired at first
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				lastEvent = event
				timer.Reset(time.Millisecond * 100)

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)

			case <- timer.C:
				// Checking for new IDL files
                if lastEvent.Has(fsnotify.Create) {
                    log.Printf("File %s created, parsing IDL file...", lastEvent.Name)
                    _, err := parser.ParseFile(lastEvent.Name, []string{""}, true)
    
                    if err != nil {
                        log.Printf("File %s is invalid. Please ensure your file is correct.", lastEvent.Name)
                        continue
                    }
                
                    log.Printf("File %s is valid, updating client...", lastEvent.Name)
                    // Update client when IDL is valid
					etcdAddr := os.Getenv("ETCD_ADDR")
					err = cp.UpdateClient(idlDir, etcdAddr)
					
                    if err != nil {
						log.Println("Client update failed: ", err)
					}
                }
				// Upating existing IDL files
				if lastEvent.Has(fsnotify.Write) {
					log.Printf("File %s modified, parsing IDL file...", lastEvent.Name)
                    _, err := parser.ParseFile(lastEvent.Name, []string{""}, true)
    
                    if err != nil {
                        log.Printf("File %s is invalid. Please ensure your file is correct.", lastEvent.Name)
                        continue
                    }
                
                    log.Printf("File %s is valid, performing hot reload...", lastEvent.Name)
                    // Update client when IDL is valid
					etcdAddr := os.Getenv("ETCD_ADDR")
					err = cp.UpdateClient(idlDir, etcdAddr)
					
                    if err != nil {
						log.Println("Client update failed: ", err)
					}
				}
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