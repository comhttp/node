package main

import (
	"fmt"
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
)

//// GetBitNodes updates the data about all of the coins in the database
func GetBitNodes(j *jdb.JDB) {
	//var b []string
	//bns := make(map[string]*BitNoded)
	//for _, coin := range coins.C {
	//	bn := &BitNoded{}
	//
	//	if utl.FileExists(filepath.FromSlash(cfg.Path + "nodes/" + coin.Slug)) {
	//		b = append(b, coin.Slug)
	//		bitNodes := BitNodes{}
	//		err := j.Read("nodes", coin.Slug, &bitNodes)
	//		utl.ErrorLog(err)
	//
	//		for _, bitnode := range coin.Nodes {
	//			bitnode.getNode(j, bn, coin.Slug)
	//		}
	//		bns[coin.Slug] = bn
	//		j.Write("nodes", coin.Slug+"_"+"bitnodes", bn)

	//data, err := jdb.JDB.ReadAll(filepath.FromSlash(cfg.C.Out + "/nodes/" + coin))
	//utl.ErrorLog(err)
	//nodes := make([][]byte, len(data))
	//for i := range data {
	//	nodes[i] = []byte(data[i])
	//}

	//ns := make(Nodes, len(nodes))
	//
	//for i := range nodes {
	//	if err := json.Unmarshal(nodes[i], &ns[i]); err != nil {
	//		fmt.Println("Error", err)
	//	}
	//}
	//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/info/nodes/"+coin), "nodes", ns)
	//}
	//}

	//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/info"), "bitnoded", b)
	//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/info"), "bitnodestat", bns)
}

func GetNode(j *jdb.JDB, c, ip string) map[string]interface{} {
	node := make(map[string]interface{})
	err := j.Read("nodes", c+"_"+ip, &node)
	utl.ErrorLog(err)
	return node
}

func (b *BitNode) getNode(j *jdb.JDB, bn *BitNoded, c string) {
	b.Jrc = utl.NewClient(cfg.C.RPC.Username, cfg.C.RPC.Password, b.IP, b.Port)
	s := b.GetBitNodeStatus()
	j.Write("nodes", c+"_"+b.IP, s)

	j.Write("info", c+"_mempool", s.GetRawMemPool)
	j.Write("info", c+"_mining", s.GetInfo)
	j.Write("info", c+"_info", s.GetInfo)
	j.Write("info", c+"_network", s.GetNetworkInfo)
	j.Write("info", c+"_peers", s.GetPeerInfo)

	fmt.Println("GetBitNodeStatus: ", c+"_"+b.IP)
	nds := GetNodes(s)
	for _, n := range nds {
		j.Write("nodes", c+"_"+n.IP, n)
		fmt.Println("Node: ", c+"_"+n.IP)

	}

	bn.Coin = c
	bn.BitNodes = append(bn.BitNodes, *s)
	j.Write("nodes", c+"_"+b.IP, s)
	return
}
