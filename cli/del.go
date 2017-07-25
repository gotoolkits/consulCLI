package cli

import (
	"strconv"

	counsulApi "github.com/hashicorp/consul/api"
)

func Del(record *string) {

	dereg := newDereg()
	err := initDereg(&dereg, *record)
	errCheck(err, "Init Catalog Deregistratation struct failed!")

	wo := newWriteOpt()
	_, err = Ctlog.Deregister(&dereg, &wo)
	errCheck(err, "Deregistratation failed!")

	log.Warningln("Delete service is finished!")

}

func initDereg(drg *counsulApi.CatalogDeregistration, r string) error {
	drg.Address = "192.168.20.2"
	drg.CheckID = ""
	drg.Datacenter = "sh1a"
	drg.Node = "MM-SH1A-20-02"
	drg.ServiceID = getSID(r)

	return nil
}

func getSID(record string) string {

	q := newQueryOpt()
	r := split(record)
	port, err := strconv.Atoi(r[2])
	errCheck(err, "Port parse failed!")

	qms, _, err := Ctlog.Service(r[0], "", &q)
	errCheck(err, "get service info err")

	for _, v := range qms {

		if v.ServiceAddress == r[1] && v.ServicePort == port {
			return v.ServiceID
		}
	}

	return ""
}
