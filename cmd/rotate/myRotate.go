package main

import (
	"flag"
	"fmt"
	"github.com/fanfaronDo/tools/pkg/rotate"
	"sync"
)

func main() {

	flagA := flag.String("a", "", "Move to next directory")
	flag.Parse()
	files := flag.Args()

	if len(files) == 0 {
		fmt.Printf("Usage: ./myRotate /path/to/logs/some_application.log\n")
		return
	}
	wg := new(sync.WaitGroup)
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			archiveName := rotate.ArchiveNameBuilder(file)
			rotate.CreateArchive(archiveName)
			if *flagA != "" {
				newLoc := *flagA + "/" + rotate.GetFileName(archiveName)
				rotate.MoveTo(archiveName, newLoc)
			}
		}(file)
	}
	wg.Wait()
}
