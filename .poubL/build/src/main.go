package main

import (
	"golang.org/x/exp/mmap"
	"libvirt.org/go/libvirt"
)

func main() {
	// get a connection from the hypervisor store. Should be called first
	conn, err := libvirt.NewConnect("qemu:///session?socket=/Users/jemercie/.cache/libvirt/libvirt-sock")
	if err != nil {
		println(err.Error())
		return
	}
	defer conn.Close()
	// check if connection is alive
	_, err = conn.IsAlive()
	if err != nil {
		println(err.Error())
		return
	}
	// get all the domains
	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		println(err.Error())
		return
	}
	// print the domain names
	for _, dom := range doms {
		name, err := dom.GetName()
		if err == nil {
			println(name)
		}
		dom.Free()
	}

	xmlData, err := read_file("./debianfromgo.xml")
	if err != nil {
		println(err.Error())
		return
	}

	domain, err := conn.DomainDefineXML(xmlData)
	if err != nil {
		println("aaa", err.Error())
		return
	}
	err = domain.Create()
	if err != nil {
		println(err.Error())
		return
	}
	println("vm launched")

	ok, err := domain.IsActive()
	if err != nil {
		println(err.Error())
		return
	}
	println("domain status: ", ok)
	name, err := domain.GetName()
	if err != nil {
		println(err.Error())
		return
	}
	println("name:", name)
	info, err := domain.GetInfo()
	if err != nil {
		println(err.Error())
		return
	}
	println("state:", info.State, "maxmem:", info.MaxMem, "mem:", info.Memory, "cputime:", info.CpuTime)

	// stream, err := conn.NewStream(0)
	// if err != nil {
	// 	println(err.Error())
	// 	return
	// }

	// domain.OpenConsole("", stream, libvirt.DOMAIN_CONSOLE_SAFE)

	// n, err := stream.Send([]byte("ls -la \n"))
	// if err != nil {
	// 	println(err.Error())
	// 	return
	// }
	// println("aaa", n)

	// testQMP(domain)
	// time.Sleep(15 * time.Second)

	// cmd := `{"execute":"guest-ping"}`
	// result, err := domain.QemuAgentCommand(cmd, libvirt.DOMAIN_QEMU_AGENT_COMMAND_DEFAULT, uint32(10))
	// if err != nil {
	// 	println(err.Error())
	// 	return
	// }
	// fmt.Println("Result:", result)

	// time.Sleep(2 * time.Second)

	// println("shutting down everything")
	// err = domain.Shutdown()
	//
	//	if err != nil {
	//		println(err.Error())
	//		return
	//	}
	//
	// err = domain.Destroy()
	//
	//	if err != nil {
	//		println(err.Error())
	//		return
	//	}
	//
	// println("everything is shut down okkk")
}

// func testQMP(domain *libvirt.Domain) error {
// 	cmd := `{"execute":"query-status"}`
// 	result, err := domain.QemuMonitorCommand(cmd, libvirt.DOMAIN_QEMU_MONITOR_COMMAND_DEFAULT)
// 	if err != nil {
// 		return fmt.Errorf("QMP failed: %v", err)
// 	}
// 	fmt.Println("QMP Result:", result)
// 	return nil
// }

// func executeCommand(domain *libvirt.Domain) error {
// 	// Utiliser QMP pour exécuter via guest agent

// 	cmd := `{"execute":"guest-ping"}`
// 	result, err := domain.QemuAgentCommand(cmd, libvirt.DOMAIN_QEMU_AGENT_COMMAND_DEFAULT, uint32(0))
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("Result:", result)
// 	return nil
// }

func read_file(filename string) (string, error) {

	mappedFile, err := mmap.Open(filename)
	if err != nil {
		println(err.Error())
		return "", err
	}
	buf := make([]byte, mappedFile.Len())
	_, err = mappedFile.ReadAt(buf, 0)
	if err != nil {
		println(err.Error())
		return "", err
	}
	return string(buf), nil
}
