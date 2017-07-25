package cli

import (
	"strconv"
)

func Del(record *string, d *string) {

	drg := newDereg()

	drg.Address = getDefaultAddr()
	drg.CheckID = ""
	drg.Datacenter = getDefaultDC()

	if *d != "" {
		drg.Node = *d
	} else {
		drg.Node = getDefaultNode()
	}

	if getSID(*record) == "" {

		log.Warningln("Get Service ID is Null, Please Check the record !")

	}
	drg.ServiceID = getSID(*record)

	wo := newWriteOpt()

	_, err := Ctlog.Deregister(&drg, &wo)
	errCheck(err, "Deregistratation failed!")

	log.Infoln("(", drg.ServiceID, drg.Node, drg.Address, drg.Datacenter, ")", "Delete service is finished!")

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
