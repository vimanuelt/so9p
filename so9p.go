package main

import (
       "fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"time"
)

var debugprint = false

func main() {

	if len(os.Args) > 1 && os.Args[1] == "s" {
		servermap = make(map[fid]*sfid, 128)
		S := new(So9ps)
		S.Path = "/"
		rpc.Register(S)
		l, err := net.Listen("tcp", ":1234")
		if err != nil {
			log.Fatal(err)
		}
		rpc.Accept(l)
	} else {
		var client so9pc
		var err error
		rootfid := fid(1)
		client.Client, err = rpc.Dial("tcp", "localhost"+":1234")
		if err != nil {
			log.Fatal("dialing:", err)
		}

		fi, err := client.attach("/", rootfid)
		if err != nil {
			log.Fatal("attach", err)
		}
		if debugprint {
		   fmt.Printf("attach fi %v\n", fi)
		   }
		if len(os.Args) == 1 {
			return
		}
		hostfid, err := client.open(os.Args[1])
		if debugprint {
			fmt.Printf("open %v: %v, %v\n", os.Args[1], hostfid, err)
		}
		for i := 1; i < 1048576; i = i * 2 {
			start := time.Now()
			data, err := client.read(hostfid, i, 0)
			cost := time.Since(start)
			if err != nil {
				log.Fatal("read", err)
			}
			fmt.Printf("%v took %v\n", len(data), cost)

		}
/*
		filelist, err := client.readdir(etcfid)
		if err != nil {
			log.Fatal("readdir", err)
		}
		fmt.Printf("etc has %v\n", filelist)

		hostfid, fi, err := client.walk(etcfid, "hosts")
		if err != nil {
			log.Fatal("walk", err)
		}
		if debugprint {
			fmt.Printf("Walk to hosts: %v, %v, %v\n", hostfid, fi, err)
		}
		err = client.open(hostfid)
		if err != nil {
			log.Fatal("open", err)
		}
		data, err := client.read(hostfid, 1<<20, 0)

		if err != nil {
			log.Fatal("read", err)
		}
		if debugprint {
			fmt.Printf("Read: %v, %v\n", data, err)
		}
		if len(os.Args) < 2 {
			return
		}
		err = client.open(hostfid)
		if err != nil {
			log.Fatal("open", err)
		}
		_, err = client.read(hostfid, 1<<20, 0)
		if err != nil {
			log.Fatal("read", err)
		}
		err = client.close(hostfid)
		if err != nil {
			log.Fatal("close", err)
		}
*/

	}

}
