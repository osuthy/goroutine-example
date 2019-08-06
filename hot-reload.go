package main

import (
	"fmt"
	"os"
	"os/exec"
	// "sync"
	// "io/ioutil"
	// "reflect"
)

func main() {
	var currentFileTime int64

	for {
		file, error := os.Stat("src/main.go")
		
		if error != nil {
			fmt.Println(error)
		}

		modTime := file.ModTime().Unix()
		if (modTime != currentFileTime) {

			fmt.Println(modTime)
			
			go func() {

				cmd := exec.Command("go", "run", "src/main.go")
				out, _ := cmd.Output()
				fmt.Println(string(out))
				
				}()
			}
			
		currentFileTime = modTime
	}

}