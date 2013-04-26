package main

import (
	"fmt"
)

func (client *so9pc) attach(name string, file fid) (FileInfo, error) {
	args := &Nameargs{Name:name, Fid:file}
	var reply Nameresp
	err := client.Client.Call("So9ps.Attach", args, &reply)
	fi := reply.FI
	if debugprint {
		fmt.Printf("clientattach: %v gets %v, %v\n", name, fi, err)
	}
	return fi, err
}

func (client *so9pc) open(name string) (fid, error)  {
	fid := clientfid
	clientfid = clientfid + 1
	args := &Nameargs{Name: name, Fid: fid}
	var reply Nameresp
	err := client.Client.Call("So9ps.Open", args, &reply)
	if debugprint {
		fmt.Printf("clientopen: %v gets %v, %v\n", name, err)
	}
	return fid, err
}

func (client *so9pc) read(file fid, Len int, Off int64) ([]byte, error) {
	args := &Ioargs{Fid: file, Len: Len, Off: Off}
	var reply Ioresp
	err := client.Client.Call("So9ps.Read", args, &reply)
	if debugprint {
		fmt.Printf("clientopen: %v gets %v, %v\n", reply, err)
	}
	return reply.Data, err
}

func (client *so9pc) write(file fid, Data []byte, Off int64) (int, error) {
	args := &Ioargs{Fid: file, Data: Data, Off: Off}
	var reply Ioresp
	err := client.Client.Call("So9ps.Write", args, &reply)
	if debugprint {
		fmt.Printf("clientopen: %v gets %v, %v\n", reply, err)
	}
	return reply.Len, err
}

func (client *so9pc) close(file fid) error {
	args := &Ioargs{Fid: file}
	var reply Ioresp
	err := client.Client.Call("So9ps.Close", args, &reply)
	if debugprint {
		fmt.Printf("clientopen: %v gets %v, %v\n", reply, err)
	}
	return err
}

func (client *so9pc) readdir(file fid) ([]FileInfo, error) {
	args := &Ioargs{Fid: file}
	var reply FIresp
	err := client.Client.Call("So9ps.ReadDir", args, &reply)
	if debugprint {
		fmt.Printf("clientopen: %v gets %v, %v\n", reply, err)
	}
	return reply.FI, err
}
