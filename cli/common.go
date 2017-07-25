package cli

import (
	//"fmt"
	"os"

	"strings"

	counsulApi "github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func errCheck(err error, out string) {

	errlog := logrus.WithFields(logrus.Fields{"err": err})
	if err != nil {
		//		fmt.Println(out)
		errlog.Error(out)
		os.Exit(1)
	}
}

var Config counsulApi.Config
var Client *counsulApi.Client
var Ctlog *counsulApi.Catalog
var Srv counsulApi.AgentService
var log = logrus.New()

func init() {

	var err error

	log.Out = os.Stdout

	Config = newConfig()
	// config := &counsulApi.Config{
	// 	Address: "192.168.20.4:8500",
	// 	Scheme:  "http",
	// }

	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/consul-cli/")

	err = viper.ReadInConfig()
	errCheck(err, "Config file no find,init config failed!")

	Config.Address = viper.GetString("consulapi.address")
	Config.Scheme = viper.GetString("consulapi.scheme")

	//fmt.Println(viper.GetString("consulapi.address"))

	Client, err = counsulApi.NewClient(&Config)
	errCheck(err, "get client err!")

	Ctlog = Client.Catalog()

}

func split(record string) []string {
	if !strings.Contains(record, "#") {
		log.Warningln("flag -r(record) Not setting the port ,the format is \"servicename#ipaddress#port\" ")
		os.Exit(1)

	}
	ipPorts := strings.Split(record, "#")
	if len(ipPorts) < 3 {
		log.Warningln("flag -r(record) Not setting the port ,the format is \"servicename#ipaddress#port\" ")
		os.Exit(1)
	}
	return ipPorts
}

func newConfig() counsulApi.Config {
	return counsulApi.Config{}
}

func newDereg() counsulApi.CatalogDeregistration {
	return counsulApi.CatalogDeregistration{}
}

func newReg() counsulApi.CatalogRegistration {
	return counsulApi.CatalogRegistration{}
}

func newService() counsulApi.AgentService {
	return counsulApi.AgentService{}
}

func newWriteOpt() counsulApi.WriteOptions {
	return counsulApi.WriteOptions{}
}

func newQueryOpt() counsulApi.QueryOptions {
	return counsulApi.QueryOptions{}
}
