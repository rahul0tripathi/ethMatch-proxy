package common

import (
	"crypto/ecdsa"
	"fmt"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core"
	"github.com/orgs/ethMatch/p2pmatch/types"
	"go.uber.org/zap"
)

type SignedDataV4 struct {
	Domain struct {
		ChainId uint16 `json:"chainId"`
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"domain"`
	Message     interface{} `json:"message"`
	PrimaryType string      `json:"primaryType"`
	Types       interface{} `json:"types"`
}

func VerifySig(signer ethcommon.Address, sigHex string, input []byte) (isSigner bool) {
	sig, err := hexutil.Decode(sigHex)
	if err != nil {
		return false
	}
	if sig[64] != 27 && sig[64] != 28 {
		return false
	}
	sig[64] -= 27
	var pubKey *ecdsa.PublicKey
	pubKey, err = crypto.SigToPub(input, sig)
	if err != nil {
		return false
	}
	signerAddr := crypto.PubkeyToAddress(*pubKey)
	return signerAddr == signer
}
func NewSignedDataV4(ticketId string, lobby types.Lobby) ([]byte, error) {
	signedData := core.TypedData{
		Domain: core.TypedDataDomain{
			ChainId: math.NewHexOrDecimal256(137),
			Name:    "EthMatch",
			Version: "1",
		},
		PrimaryType: "Proposal",
		Message: map[string]interface{}{
			"ticket_id":           ticketId,
			"id":                  ethcommon.HexToHash(lobby.Id).Hex(),
			"operators_share":     math.NewHexOrDecimal256(1),
			"operators_address":   lobby.OperatorAddress.String(),
			"entry_fee":           math.NewHexOrDecimal256(20),
			"operators_signature": lobby.OperatorSignature.String(),
		},
		Types: core.Types{
			"EIP712Domain": []core.Type{
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "version",
					Type: "string",
				},
				{
					Name: "chainId",
					Type: "uint256",
				},
			},
			"Proposal": []core.Type{
				{
					Name: "id",
					Type: "bytes32",
				},
				{
					Name: "ticket_id",
					Type: "string",
				},
				{
					Name: "entry_fee",
					Type: "uint256",
				},
				{
					Name: "operators_share",
					Type: "uint256",
				},
				{
					Name: "operators_address",
					Type: "address",
				},
				{
					Name: "operators_signature",
					Type: "bytes32",
				},
			},
		},
	}
	domainSeparator, err := signedData.HashStruct("EIP712Domain", signedData.Domain.Map())
	if err != nil {
		Logger.Error("EIP712Domain", zap.Error(err))
		return nil, err
	}
	typedDataHash, err := signedData.HashStruct(signedData.PrimaryType, signedData.Message)
	if err != nil {
		Logger.Error("typedDataHash", zap.Error(err))
		return nil, err
	}
	rawData := crypto.Keccak256([]byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash))))
	return rawData, nil
}
