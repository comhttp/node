package nodes

import (
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
)

type JORMnode struct {
	Coin   string
	JDB    *jdb.JDB
	config cfg.Config
}

func NewJORMnode(coin string) *JORMnode {
	n := new(JORMnode)
	c, _ := cfg.NewCFG(n.config.Path, nil)
	n.config = cfg.Config{}
	err := c.Read("conf", "conf", &n.config)
	utl.ErrorLog(err)
	n.Coin = coin
	// n.JDB=  jdb.NewJDB(cfg.C.JDBservers)

	return n
}
