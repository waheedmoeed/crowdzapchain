package rest

import (
	"fmt"
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

const (
	restRelContractorAddress = "contractor"
	restPollId               = "pollId"
)

// RegisterRoutes registers relcontractors-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, moduleName string) {
	registerQueryRoutes(cliCtx, r, moduleName)
	registerTxRoutes(cliCtx, r, moduleName)
}

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router, moduleName string) {
	r.HandleFunc(fmt.Sprintf("/%s/updaterelcontractor", moduleName), updateRelContractorAddressHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/createVotingPoll", moduleName), createVotingPollHandler(cliCtx)).Methods("POST")
}

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, moduleName string) {
	r.HandleFunc(fmt.Sprintf("/%s/minted_distributed_coins", moduleName), queryDistributeMintedCoins(cliCtx, moduleName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/contractors", moduleName), queryRelContractors(cliCtx, moduleName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/contractor_by_address/{%s}", moduleName, restRelContractorAddress), queryRelContractorByAddress(cliCtx, moduleName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/distributed_records", moduleName), queryDistributedRecords(cliCtx, moduleName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/minted_records", moduleName), queryMintedRecords(cliCtx, moduleName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/polls", moduleName), queryPolls(cliCtx, moduleName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/poll_by_id/{%s}", moduleName, restPollId), queryPollById(cliCtx, moduleName)).Methods("GET")
}
