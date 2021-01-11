package manager

import (
	"fmt"
	"os"
	"time"
)

func Start() {
	uptime := 0
	for range time.Tick(time.Hour) {
		uptime++
		fmt.Println("Uptime is ", uptime, " hours")
	}
	return
}

func CriticalHandler(err *error) {
	if *err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}

func StandardHandler(err *error) {
	if *err != nil {
		fmt.Println(err)
	}
	return
}
