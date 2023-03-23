package main

import (
	"EsoteraFinal/app"
	"os"
	"os/exec"
)

func main() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	app.Run()

}
