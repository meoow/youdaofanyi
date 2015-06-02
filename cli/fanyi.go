package main

import (
	"fmt"
	fy "github.com/meoow/youdaofanyi"
	"log"
	"os"
	"strings"
)

func init() {
	log.SetOutput(os.Stderr)
}

func main() {

	if len(os.Args) < 2 {
		return
	}

	out, err := fy.Fanyi(strings.Join(os.Args[1:], " "), fy.Plain)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(out)
}
