package cli

import (
	"strconv"
)

func Add(record *string, dnsNode *string) {

	Srv := newService()
	r := split(*record)
	port, err := strconv.Atoi(r[2])
	errCheck(err, "Port parse failed!")

	var tags = make([]string, 20)
	tags = append(tags, r[0])
	tags = append(tags, "Custome Tags")

	wo := newWriteOpt()

	Srv.Service = r[0]
	Srv.Address = r[1]
	Srv.Port = port
	//Srv.ModifyIndex = 0
	//Srv.CreateIndex = 1
	Srv.EnableTagOverride = true
	Srv.Tags = tags

	ctReg := newReg()
	ctReg.Datacenter = getDefaultDC()

	ctReg.Address = getDefaultAddr()
	ctReg.Service = &Srv

	if *dnsNode != "" {
		ctReg.Node = *dnsNode
	} else {
		ctReg.Node = getDefaultNode()
	}

	wm, err := Ctlog.Register(&ctReg, &wo)

	errCheck(err, "Add service failed,Please check defined for the Catalogs !")

	log.Infoln("(", Srv.Service, Srv.Address, Srv.Port, ctReg.Node, ctReg.Address, ctReg.Datacenter, ")", "Service Register spent time:", wm.RequestTime)

}
