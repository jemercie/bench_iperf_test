package main

import (
	"golang.org/x/exp/mmap"
	"libvirt.org/go/libvirt"
)

func main() {
	println("bonjour je vais essayer de creer une vm depuis du code avec libvirt.org/go/libvirt j'ai aucune idee de comment on fait salut")
	// get a connection from the hypervisor store. Should be called first
	conn, err := libvirt.NewConnect("qemu:///session")
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
	if err != nil{
		println(err.Error())
		return
	}

	domain, err := conn.DomainDefineXML(xmlData)
	if err != nil{
		println("aaa", err.Error())
		return
	}
	err = domain.Create()
	if err != nil{
		println(err.Error())
		return
	}
	println("vm launched")
	info, err := domain.GetInfo()
	if err != nil{
		println(err.Error())
		return
	}

	println("state:", info.State)
	println("maxmem:", info.MaxMem)
	println("mem:", info.Memory)
	println("cputime:", info.CpuTime)

	err = domain.Destroy()
	if err != nil{
		println(err.Error())
		return
	}
}

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

