// Copyright Monax Industries Limited
// SPDX-License-Identifier: Apache-2.0

package rpcinfo

import (
	"net"
	"net/http"

	"github.com/klye-dev/hivesmartchain/logging"
	"github.com/klye-dev/hivesmartchain/logging/structure"
	"github.com/klye-dev/hivesmartchain/rpc"
	"github.com/klye-dev/hivesmartchain/rpc/lib/server"
)

func StartServer(service *rpc.Service, pattern string, listener net.Listener, logger *logging.Logger) (*http.Server, error) {
	logger = logger.With(structure.ComponentKey, "RPC_Info")
	routes := GetRoutes(service)
	mux := http.NewServeMux()
	server.RegisterRPCFuncs(mux, routes, logger)
	srv, err := server.StartHTTPServer(listener, mux, logger)
	if err != nil {
		return nil, err
	}
	return srv, nil
}
