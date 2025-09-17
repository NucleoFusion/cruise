package main

import (
	"fmt"
	"log"
	"os"

	"github.com/NucleoFusion/cruise/cmd"
)

func main() {
	f, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	log.SetOutput(f)

	cmd.Execute()
}
