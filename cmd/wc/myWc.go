package main

import (
	"flag"
	"fmt"
	"github.com/fanfaronDo/tools/pkg/wc"
	"os"
	"sync"
)

func main() {
	flagW := flag.Bool("w", false, "Counter word")
	flagL := flag.Bool("l", false, "Counter line")
	flagM := flag.Bool("m", false, "Counter symbols")
	flag.Parse()

	indexFirstFile := 2

	if !(*flagM || *flagW || *flagL) {
		indexFirstFile -= 1
	}

	argc := len(os.Args)

	// Определяем есть ли аргументы в командной строке
	if (argc + 1) <= 2 {
		wc.Reader(os.Stdin, "", *flagW, *flagL, *flagM, false)

	} else {
		files := flag.Args()
		wg := new(sync.WaitGroup)
		for _, file := range files {
			fPtr, err := os.Open(file)
			if err != nil {
				fmt.Println("File openning error")
				return
			}
			wg.Add(1)
			go func(file string) {
				defer wg.Done()
				wc.Reader(fPtr, file, *flagW, *flagL, *flagM, true)
			}(file)
			wg.Wait()
		}
	}

}
