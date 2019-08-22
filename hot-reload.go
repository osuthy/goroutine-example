package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var currentFileTime int64
	channel := make(chan string)


	for {
		file, error := os.Stat("src/main.go")
		
		if error != nil {
			fmt.Println(error)
		}

		modTime := file.ModTime().Unix()

		if (modTime != currentFileTime) {

			// fmt.Println(modTime)
			
			go func(ch chan string) {
				cmd := exec.Command("go", "run", "src/main.go")
				out, _ := cmd.Output()
				ch <- string(out)
			}(channel)
		}

			go func(ch chan string) {
    			fmt.Println(<-ch)
			}(channel)
			
		currentFileTime = modTime
	}

}