package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irisnet/testnets/bifrost/phase-2/count/address"
	counttypes "github.com/irisnet/testnets/bifrost/phase-2/count/types"
)

func main() {
	address.ConfigureBech32Prefix()

	validators := getValidators()
	participants := getRawParticipants()

	for _, participant := range participants {
		for _, validator := range validators {
			if strings.Contains(validator.Description.Identity, participant.PGP) || strings.Contains(validator.Description.Moniker, participant.PGP) {
				addr, err := sdk.ValAddressFromBech32(validator.OperatorAddress)
				if err != nil {
					panic(err)
				}
				participant.Addr = sdk.AccAddress(addr).String()

				bz, err := base64.StdEncoding.DecodeString(validator.ConsensusPubkey.Key)
				if err != nil {
					panic(err)
				}

				edPK := ed25519.PubKey{Key: bz}
				participant.ValAddr = strings.ToUpper(hex.EncodeToString(edPK.Address()))
			}
		}
	}

	bz, _ := json.MarshalIndent(participants, "", "    ")
	if err := ioutil.WriteFile("participants.json", bz, 0666); err != nil {
		panic(err)
	}
}

func getValidators() []*counttypes.Validator {
	bytes, err := ioutil.ReadFile("validators.json")
	if err != nil {
		panic(err)
	}

	var validators []*counttypes.Validator
	if err := json.Unmarshal(bytes, &validators); err != nil {
		panic(err)
	}

	return validators
}

func getRawParticipants() []*counttypes.Participant {
	bytes, err := ioutil.ReadFile("participants_verified.json")
	if err != nil {
		panic(err)
	}

	var participants []*counttypes.Participant
	if err := json.Unmarshal(bytes, &participants); err != nil {
		panic(err)
	}

	return participants
}
