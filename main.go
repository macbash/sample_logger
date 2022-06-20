package main

import "log"
import "os"
import "time"
import "strconv"

func main() {
	log.Print("Started logger")
	sec, err := strconv.Atoi(os.Getenv("time_interval"))
	if err != nil {
    log.Fatal(err)
         }
	//time.Duration(sec)*time.Second
	for _ = range time.Tick(time.Second * time.Duration(sec)) {
		log.Print("Logging from: " + os.Getenv("env_name"))
	}
}
