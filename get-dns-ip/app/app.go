package app

import (
	"github.com/urfave/cli"
	"net"
	"fmt"
)

func Start() *cli.App {
	c := cli.NewApp()
	c.Name = "whois"
	c.Usage = "Procurar ips/dns"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
		},
	}
	
	c.Commands = []cli.Command{
		{
			Name: "dns",
			Usage: "pegar dns",
			Flags: flags,
			Action: findDns, 
		},
		{
			Name: "ip",
			Usage: "pegar ip",
			Flags: flags,
			Action: findIp, 
		},
	}

	return c
}

func findIp(c *cli.Context){
	host := c.String("host")
	ips, err := net.LookupIP(host) 
	if err != nil {
		fmt.Println("findIp error", err)
		return
	}

	for _, value := range ips{
		fmt.Println(value)
	}
}

func findDns(c *cli.Context){
	host := c.String("host")
	ns, err := net.LookupAddr(host) 
	if err != nil {
		fmt.Println("findDns error", err)
		return
	}

	for _, value := range ns {
		fmt.Println(value)
	}
}