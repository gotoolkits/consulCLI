package cli

import (
	"strconv"
	//counsulApi "github.com/hashicorp/consul/api"
)

func Add(record *string) {

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
	ctReg.Datacenter = "sh1a"
	ctReg.Node = "MM-SH1A-20-02"
	ctReg.Address = "192.168.20.2"
	ctReg.Service = &Srv

	wm, err := Ctlog.Register(&ctReg, &wo)

	errCheck(err, "Add service failed,Please check defined for the Catalogs !")

	//errCheck(err, "Service Register err!")

	log.Infoln("Service Register at: %v ", wm.RequestTime)

}
