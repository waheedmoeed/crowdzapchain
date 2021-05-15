package rest

import (
	"fmt"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

const (
	restBasicContractId = "basic_contract_id"
	restYieldContractId = "yield_contract_id"
)

// RegisterRoutes registers smartcontracts-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, moduleName string) {
	registerQueryRoutes(cliCtx, r, moduleName)
	registerTxRoutes(cliCtx, r, moduleName)
}

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router, moduleName string) {
	//to broadcast transactions to chain through the REST points.
	r.HandleFunc(fmt.Sprintf("/%s/createBasicContract", moduleName), CreateBasicContractHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/createYieldContract", moduleName), CreateYieldContractHandler(cliCtx)).Methods("POST")
	//TODO
	//Create  POST request for ivestment
	//REST transactions endpoints for this module

}

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, moduleName string) {
	r.HandleFunc(fmt.Sprintf("/%s/get_basic_contract/{%s}", moduleName, restBasicContractId), queryBasicContract(cliCtx, moduleName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/get_yield_contract/{%s}", moduleName, restYieldContractId), queryYieldContract(cliCtx, moduleName)).Methods("GET")
}
