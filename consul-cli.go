package main

import (
	"fmt"
	"os"

	//	counsulApi "github.com/hashicorp/consul/api"
	"github.com/gotoolkits/consulCLI/cli"
	"github.com/spf13/cobra"
)

func errCheck(err error, out string) {
	if err != nil {
		fmt.Println(out)
		os.Exit(1)
	}
}

var RootCmd = &cobra.Command{
	Use:   "consul-cli",
	Short: "Consul console tool",
	Long: `Consul-cli for console command tool, Add/List/Del service
in Consul Servers Cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		//RootCmd.Help()
		//cli.Usage(cmd)
		cmd.Usage()
	},
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List the dns record",
	Long: `Consul-cli for console command tool, List the services
in Consul Servers Cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		cli.List()
	},
}

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a dns record",
	Long: `Consul-cli for console command tool, Add register a service
in Consul Servers Cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(dnsRecord) < 1 {
			cmd.Usage()
		}

		cli.Add(&dnsRecord)
	},
}

var DelCmd = &cobra.Command{
	Use:   "del",
	Short: "Del a dns record",
	Long: `Consul-cli for console command tool, Del a service
in Consul Servers Cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff
		if len(dnsRecord) < 1 {
			cmd.Usage()
		}

		cli.Del(&dnsRecord)
	},
}

var outType, dnsRecord string

func init() {
	RootCmd.PersistentFlags().StringVarP(&outType, "Out Type", "t", "shell", "Out type Json/shell")
	RootCmd.PersistentFlags().StringVarP(&dnsRecord, "DNS record", "r", "", "ServiceName#ipaddress#port")
	RootCmd.AddCommand(ListCmd)
	RootCmd.AddCommand(AddCmd)
	RootCmd.AddCommand(DelCmd)
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
