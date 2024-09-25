package xargs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadLine(arguments *[]string) {
	reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, line := range strings.Split(reader, "\n") {
		for _, arg := range strings.Split(line, " ") {
			strings.TrimSpace(arg)
			if len(arg) == 0 {
				continue
			}
			*arguments = append(*arguments, arg)
		}
	}
}
