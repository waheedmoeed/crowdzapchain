package rest

import (
	"fmt"
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers smartcontracts-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, moduleName string) {
	registerQueryRoutes(cliCtx, r, moduleName)
	registerTxRoutes(cliCtx, r, moduleName)
}

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router, moduleName string) {
	//to broadcast transactions to chain through the REST points.
	r.HandleFunc(fmt.Sprintf("/%s/createBasicContract", moduleName), CreateBasicContractHandler(cliCtx)).Methods("POST")
	//REST transactions endpoints for this module

}

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, moduleName string) {
	r.HandleFunc(fmt.Sprintf("/%s/minted_distributed_coins", moduleName), queryBasicContract(cliCtx, moduleName)).Methods("GET")
}
