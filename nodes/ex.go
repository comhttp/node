package nodes

import (
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
)

type JORMnode struct {
	Coin string
	JDB  *jdb.JDB
}

func NewJORMnode(coin string) *JORMnode {
	err := cfg.CFG.Read("conf", "conf", &cfg.C)
	utl.ErrorLog(err)
	j := &JORMnode{
		Coin: coin,
		JDB:  jdb.NewJDB(cfg.C.JDBservers),
	}
	return j
}
