package cli

import (
	"bufio"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/waheedmoeed/relchain/x/relcontractors/types"
)

//GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	chainserviceTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	chainserviceTxCmd.AddCommand(flags.PostCommands(
	 //TODO: Add tx based commands
	 	GetCmdUpdateRelContractor(cdc),
	)...)

	return chainserviceTxCmd
}

 func GetCmdUpdateRelContractor(cdc *codec.Codec) *cobra.Command {
 	return &cobra.Command{
 		Use:   "update-rel-contractor [newAddress]",
 		Short: "add new rel contractor in contract",
 		Args:  cobra.ExactArgs(1),
 		RunE: func(cmd *cobra.Command, args []string) error {
 			cliCtx := context.NewCLIContext().WithCodec(cdc)
 			inBuf := bufio.NewReader(cmd.InOrStdin())
 			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			newAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil{
				return err
			}
 			msg := types.NewMsgUpdateRelContractorAddress(cliCtx.GetFromAddress(), newAddr)
 			err = msg.ValidateBasic()
 			if err != nil {
 				return err
 			}

 			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
 		},
 	}
 }
