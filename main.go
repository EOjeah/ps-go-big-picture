package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// f is the file handle and an optional err is the error state
	// if we fail to open this fail, it's not going to through an execption, there are NO exceptions in go
	f, err := os.Open("myapp.log")
	if err != nil {
		log.Fatal(err)
	}
	// after the function that defer is called in (main), it will run the close function at the end.
	//  housekeeping to clean up
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, "ERROR") {
			fmt.Println(s)
		}
	}
}
