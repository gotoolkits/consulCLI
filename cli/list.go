package cli

import (
	"fmt"
)

func List() {

	q := newQueryOpt()

	m, _, err := Ctlog.Services(&q)
	errCheck(err, "get services err!")

	fmt.Println("------------------------------------------------------------------------")
	fmt.Printf("服务名\t\t 集群IP:Port \t\tNodeName \t\tNodeIP \t\t\n")
	fmt.Println("------------------------------------------------------------------------")

	for srv := range m {

		qms, _, err := Ctlog.Service(srv, "", &q)
		errCheck(err, "get service info err")

		for i := 0; i < len(qms); i++ {

			if i == 0 {
				fmt.Printf("%-15s :%-20d %-20s %-15s\n", srv, qms[i].ServicePort, qms[i].Node, qms[i].Address)
			} else {
				if len(qms[i].ServiceAddress) < 1 {
					fmt.Printf("\t\t:%-20d %-20s %-15s\n", qms[i].ServicePort, qms[i].Node, qms[i].Address)
				} else {
					fmt.Printf("\t\t%-15s:%-5d %-20s %-15s\n", qms[i].ServiceAddress, qms[i].ServicePort, qms[i].Node, qms[i].Address)
				}

			}

		}

	}

}
