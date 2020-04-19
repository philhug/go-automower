package main

import (
	"flag"
	"log"
	"os"

	"github.com/jamiealquiza/envy"
	client "github.com/philhug/go-automower/pkg/automower"
)

func main() {
	var username = flag.String("username", "", "Username for AMC")
	var password = flag.String("password", "", "Password for AMC")
	envy.Parse("AUTOMOWER") // Expose environment variables.
	flag.Parse()
	if *username == "" || *password == "" {
		flag.Usage()
		os.Exit(1)
	}

	c, err := client.NewClientWithUserAndPassword(*username, *password)
	if err != nil {
		log.Println(err)
	}
	mowers, err := c.Mowers()
	if err != nil {
		log.Println(err)
	}
	log.Println(mowers)
	if len(mowers) == 0 {
		log.Println("no mowers available")
		return
	}
	mower, err := c.Status(&mowers[0])
	if err != nil {
		log.Println(err)
	}
	log.Println(mower)
}
