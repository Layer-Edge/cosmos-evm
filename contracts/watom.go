package contracts

import (
	_ "embed"

	contractutils "github.com/cosmos/evm/contracts/utils"
	evmtypes "github.com/cosmos/evm/x/vm/types"
)

var (
	// WEDGENSJSON are the compiled bytes of the WEDGENSContract
	//
	//go:embed solidity/WEDGENS.json
	WEDGENSJSON []byte

	// WEDGENSContract is the compiled wedgens contract
	WEDGENSContract evmtypes.CompiledContract
)

func init() {
	var err error
	if WEDGENSContract, err = contractutils.ConvertHardhatBytesToCompiledContract(
		WEDGENSJSON,
	); err != nil {
		panic(err)
	}
}
