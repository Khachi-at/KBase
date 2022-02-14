package embed

import (
	"kbase.com/server/config"
	"kbase.com/server/kbaseserver"
)

type KBase struct {
	server *kbaseserver.KBaseServer
}

func StartKBase() (k *KBase, err error) {
	srvConfig := config.ServerConfig{}
	if k.server, err = kbaseserver.NewServer(srvConfig); err != nil {
		return k, err
	}
	k.server.Start()
	return k, nil
}
