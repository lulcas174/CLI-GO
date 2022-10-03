package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI"
	app.Usage = "CLI search IP and server names in internet"

	app.Commands = []cli.Command{
		{
			Name:  "IP",
			Usage: "Search IP addresses",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},
			Action: searchIP,
		},
	}

	return app
}

func searchIP(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
