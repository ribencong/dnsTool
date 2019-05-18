package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"net"
	"os"
	"time"
)

var rootCmd = &cobra.Command{
	Use: "DNS",

	Short: "DNS [name|ip] args",

	Long: `""`,

	//Run: mainRun,

	//Args:  cobra.MinimumNArgs(2),
}

var nameCmd = &cobra.Command{
	Use: "name",

	Short: "DNS name args",

	Long: `"DNS name args"`,

	Run: nameRun,

	Args: cobra.MinimumNArgs(1),
}

var ipCmd = &cobra.Command{
	Use: "ip",

	Short: "DNS ip args",

	Long: `"DNS ip args"`,

	Run: ipRun,

	Args: cobra.MinimumNArgs(1),
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var IPType = 0
var DNSAddr = "8.8.8.8：53"

func init() {
	rootCmd.AddCommand(nameCmd)
	rootCmd.AddCommand(ipCmd)

	nameCmd.Flags().StringVarP(&DNSAddr, "dns-address", "d", "8.8.8.8：53",
		"default dns address")

	ipCmd.Flags().IntVarP(&IPType, "duration", "t", 0,
		"IP type, 0:tcp; 1:udp")
}

func mainRun(_ *cobra.Command, _ []string) {

}

func nameRun(_ *cobra.Command, addr []string) {

	a := addr[0]

	fmt.Print("LookupAddr:")
	fmt.Println(net.LookupAddr(a))

	fmt.Print("LookupCNAME:")
	fmt.Println(net.LookupCNAME(a))

	fmt.Print("LookupHost:")
	fmt.Println(net.LookupHost(a))

	fmt.Print("LookupIP:")
	fmt.Println(net.LookupIP(a))

	fmt.Print("LookupMX:")
	fmt.Println(net.LookupMX(a))

	fmt.Print("LookupNS:")
	fmt.Println(net.LookupNS(a))

	fmt.Print("LookupSRV:")
	fmt.Println(net.LookupSRV(addr[0], addr[1], addr[2]))

	fmt.Print("LookupTXT:")
	fmt.Println(net.LookupTXT(a))
}
func ipRun(_ *cobra.Command, addr []string) {

	var network = "tcp"
	if IPType == 1 {
		network = "udp"
	}
	conn, err := net.Dial(network, addr[0])
	if err != nil {
		panic(err)
	}

	if err := conn.SetDeadline(time.Now().Add(time.Second * 4)); err != nil {
		panic(err)
	}
	n, e := conn.Write([]byte("hello!"))
	if e != nil {
		panic(e)
	}
	fmt.Println("write data：", n)

	buff := make([]byte, 1024)
	n, e = conn.Read(buff)
	if e != nil {
		panic(e)
	}
	fmt.Printf("Read (%d) data:%2x", n, buff[:n])
}
