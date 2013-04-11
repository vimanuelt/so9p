package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func (server *So9ps) Attach(Args *Nameargs, Resp *Nameresp) (err error) {
	fmt.Printf("attach args %v resp %v\n", Args, Resp)
	_, err = os.Stat(Args.Name)
	if err != nil {
		log.Print("Attach", err)
		return err
	}

	n, err := server.Fs.Root()
	Resp.FI, err = n.FI()
	Resp.Fid = Args.Fid
	servermap[Args.Fid] = &sfid{n}
	return err
}

func (server *So9ps) Walk(Args *Nameargs, Resp *Nameresp) (err error) {
	var serverfid *sfid
	var ok bool
	fmt.Printf("Walk args %v resp %v\n", Args, Resp)
	/* ofid valid? */
	ofid := Args.Fid
	if serverfid, ok = servermap[ofid]; !ok {
		return err
	}

	fmt.Printf("ofid %v\n", serverfid)
	nfid := Args.NFid

	/* shortcut: new name is 0 length */
	if len(Args.Name) == 0 {
		servermap[nfid] = servermap[ofid]
		return nil
	}
	n := serverfid.Node
	dirfi, err := n.FI()
	if err != nil {
		return err
	}

	walkTo := path.Join(dirfi.SName, Args.Name)
	/* walk to whatever the new path is -- may be same as old */
	if fs, ok := n.(interface {
		Walk(string) (Node, error)
	}); ok {
		newNode, err := fs.Walk(walkTo)
		fmt.Printf("fs.Walk returns (%v, %v)\n", newNode, err)
		if err != nil {
			log.Print("walk", err)
			return nil
		}
		if stat, ok := newNode.(interface {
			FI() (FileInfo, error)
		}); ok {
			fmt.Printf("stat seems to exist, ...\n");
			Resp.FI, err = stat.FI()
			if err != nil {
				log.Print("walk", err)
				return nil
			}
		} else {
			return nil
		}
		Resp.Fid = Args.Fid
		servermap[Args.Fid] = &sfid{newNode}
	}

	return err
}
