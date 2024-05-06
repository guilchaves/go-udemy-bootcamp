package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate returns a CLI App ready for execution
func Generate() *cli.App {
	app := cli.NewApp()

	app.Name = "CLI Application"
	app.Usage = "Search IPs and Server names on the web"

    flags := []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			}

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search IPs on web",
			Flags: flags,
            Action: searchIps,
		},
        {
            Name: "servers",
            Usage: "Search Server names on web",
            Flags: flags,
            Action: searchServers,
        },
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

    for _, ip := range ips {
        fmt.Println(ip)
    }
}

func searchServers(c *cli.Context){
    host := c.String("host")

    servers, err := net.LookupNS(host)
    if err != nil {
        log.Fatal(err)
    }

    for _, server := range servers{
        fmt.Println(server.Host)
    }

}
