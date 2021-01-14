package manager

import (
	"fmt"
	"os"
	"time"
)

//Start is simple loop wich prints uptime in hours
func Start() {
	var (
		uptime int64
	)
	for range time.Tick(time.Hour) {
		uptime++
		fmt.Println("Uptime is ", uptime, " hours")
	}
	return
}

//CriticalHandler is default error handling function wich ruins program
func CriticalHandler(err *error) {
	if *err != nil {
		fmt.Println(*err)
		os.Exit(1)
	}
	return
}

//StandardHandler is default error handling function wich only prints error
func StandardHandler(err *error) {
	if *err != nil {
		fmt.Println(*err)
	}
	return
}
