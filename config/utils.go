package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/AccumulateNetwork/accumulate/protocol"
	"github.com/AccumulateNetwork/accumulate/types"
	acmeapi "github.com/AccumulateNetwork/accumulate/types/api"
	"github.com/AccumulateNetwork/accumulate/types/api/response"
	"github.com/AccumulateNetwork/accumulate/types/synthetic"
	"github.com/AccumulateNetwork/jsonrpc2/v15"
)

type ActionResponse struct {
	Txid      types.Bytes32 `json:"txid"`
	Hash      types.Bytes32 `json:"hash"`
	Log       types.String  `json:"log"`
	Code      types.String  `json:"code"`
	Codespace types.String  `json:"codespace"`
	Error     types.String  `json:"error"`
	Mempool   types.String  `json:"mempool"`
}

func (a *ActionResponse) Print() (string, error) {
	ok := a.Code == "0" || a.Code == ""

	var out string
	if WantJsonOutput {
		if ok {
			a.Code = "ok"
		}
		b, err := json.Marshal(a)
		if err != nil {
			return "", err
		}
		out = string(b)
	} else {
		out += fmt.Sprintf("\n\tTransaction Identifier\t:\t%x\n", a.Txid)
		out += fmt.Sprintf("\tTendermint Reference\t:\t%x\n", a.Hash)
		if !ok {
			out += fmt.Sprintf("\tError code\t\t:\t%s\n", a.Code)
		} else {
			out += fmt.Sprintf("\tError code\t\t:\tok\n")
		}
		if a.Error != "" {
			out += fmt.Sprintf("\tError\t\t:\t%s\n", a.Error)
		}
		if a.Log != "" {
			out += fmt.Sprintf("\tLog\t\t\t:\t%s\n", a.Log)
		}
		if a.Codespace != "" {
			out += fmt.Sprintf("\tCodespace\t\t:\t%s\n", a.Codespace)
		}
	}

	if ok {
		return out, nil
	}
	return "", errors.New(out)
}

func PrintJsonRpcError(err error) (string, error) {
	var e jsonrpc2.Error
	switch err := err.(type) {
	case jsonrpc2.Error:
		e = err
	default:
		return "", fmt.Errorf("error with request, %v", err)
	}

	if WantJsonOutput {
		out, err := json.Marshal(e)
		if err != nil {
			return "", err
		}
		return "", errors.New(string(out))
	} else {
		var out string
		out += fmt.Sprintf("\n\tMessage\t\t:\t%v\n", e.Message)
		out += fmt.Sprintf("\tError Code\t:\t%v\n", e.Code)
		out += fmt.Sprintf("\tDetail\t\t:\t%s\n", e.Data)
		return "", errors.New(out)
	}
}

func PrintQueryResponse(res *acmeapi.APIDataResponse) (string, error) {
	if WantJsonOutput {
		data, err := json.Marshal(res)
		if err != nil {
			return "", err
		}
		return string(data), nil
	} else {
		switch res.Type {
		case "anonTokenAccount":
			ata := response.AnonTokenAccount{}
			err := json.Unmarshal(*res.Data, &ata)
			if err != nil {
				return "", err
			}

			amt, err := formatAmount(ata.TokenUrl, &ata.Balance.Int)
			if err != nil {
				amt = "unknown"
			}

			var out string
			out += fmt.Sprintf("\n\tAccount Url\t:\t%v\n", ata.Url)
			out += fmt.Sprintf("\tToken Url\t:\t%v\n", ata.TokenUrl)
			out += fmt.Sprintf("\tBalance\t\t:\t%s\n", amt)
			out += fmt.Sprintf("\tCredits\t\t:\t%s\n", ata.CreditBalance.String())
			out += fmt.Sprintf("\tNonce\t\t:\t%d\n", ata.Nonce)

			return out, nil
		case "tokenAccount":
			ata := response.TokenAccount{}
			err := json.Unmarshal(*res.Data, &ata)
			if err != nil {
				return "", err
			}

			amt, err := formatAmount(ata.TokenUrl, &ata.Balance.Int)
			if err != nil {
				amt = "unknown"
			}
			var out string
			out += fmt.Sprintf("\n\tAccount Url\t:\t%v\n", ata.Url)
			out += fmt.Sprintf("\tToken Url\t:\t%v\n", ata.TokenUrl)
			out += fmt.Sprintf("\tBalance\t\t:\t%s\n", amt)
			out += fmt.Sprintf("\tKey Book Url\t:\t%s\n", ata.KeyBookUrl)

			return out, nil
		case "adi":
			adi := response.ADI{}
			err := json.Unmarshal(*res.Data, &adi)
			if err != nil {
				return "", err
			}

			var out string
			out += fmt.Sprintf("\n\tADI Url\t\t:\t%v\n", adi.Url)
			out += fmt.Sprintf("\tKey Book Url\t:\t%s\n", adi.KeyBookName)

			return out, nil
		case "directory":
			dqr := protocol.DirectoryQueryResult{}
			err := json.Unmarshal(*res.Data, &dqr)
			if err != nil {
				return "", err
			}
			var out string
			out += fmt.Sprintf("\n\tADI Entries\n")
			for _, s := range dqr.Entries {
				data, err := Get(s)
				if err != nil {
					return "", err
				}
				r := acmeapi.APIDataResponse{}
				err = json.Unmarshal([]byte(data), &r)

				chainType := "unknown"
				if err == nil {
					if v, ok := ApiToString[*r.Type.AsString()]; ok {
						chainType = v
					}
				}
				out += fmt.Sprintf("\t%v (%s)\n", s, chainType)
			}
			return out, nil
		case "sigSpecGroup":
			//workaround for protocol unmarshaling bug
			var ssg struct {
				Type      types.ChainType `json:"type" form:"type" query:"type" validate:"required"`
				ChainUrl  types.String    `json:"url" form:"url" query:"url" validate:"required,alphanum"`
				SigSpecId []byte          `json:"sigSpecId"` //this is the chain id for the sig spec for the chain
				SigSpecs  []types.Bytes32 `json:"sigSpecs"`
			}

			err := json.Unmarshal(*res.Data, &ssg)
			if err != nil {
				return "", err
			}

			u, err := url2.Parse(*ssg.ChainUrl.AsString())
			if err != nil {
				return "", err
			}
			var out string
			out += fmt.Sprintf("\n\tHeight\t\tKey Page Url\n")
			for i, v := range ssg.SigSpecs {
				//enable this code when testnet updated to a version > 0.2.1.
				//data, err := GetByChainId(v[:])
				//keypage := "unknown"
				//
				//if err == nil {
				//	r := acmeapi.APIDataResponse{}
				//	err = json.Unmarshal(*data.Data, &r)
				//	if err == nil {
				//		ss := protocol.SigSpec{}
				//		err = json.Unmarshal(*r.Data, &ss)
				//		keypage = *ss.ChainUrl.AsString()
				//	}
				//}
				//out += fmt.Sprintf("\t%d\t\t:\t%s\n", i, keypage)
				//hack to resolve the keypage url given the chainid
				s, err := resolveKeyPageUrl(u.Authority, v[:])
				if err != nil {
					return "", err
				}
				out += fmt.Sprintf("\t%d\t\t:\t%s\n", i+1, s)
			}
			return out, nil
		case "sigSpec":
			ss := protocol.SigSpec{}
			err := json.Unmarshal(*res.Data, &ss)
			if err != nil {
				return "", err
			}

			out := fmt.Sprintf("\n\tIndex\tNonce\tPublic Key\t\t\t\t\t\t\t\tKey Name\n")
			for i, k := range ss.Keys {
				keyName := ""
				name, err := FindLabelFromPubKey(k.PublicKey)
				if err == nil {
					keyName = name
				}
				out += fmt.Sprintf("\t%d\t%d\t%x\t%s", i, k.Nonce, k.PublicKey, keyName)
			}
			return out, nil
		case "tokenTx":
			tx := response.TokenTx{}
			err := json.Unmarshal(*res.Data, &tx)
			if err != nil {
				return "", fmt.Errorf("cannot extract token transaction data from request")
			}

			var out string
			for i := range tx.ToAccount {
				bi := big.Int{}
				bi.SetInt64(int64(tx.ToAccount[i].Amount))
				amt, err := formatAmount("acc://ACME", &bi)
				if err != nil {
					amt = "unknown"
				}
				out += fmt.Sprintf("Send %s from %s to %s\n", amt, *tx.From.AsString(), tx.ToAccount[i].URL.String)
				out += fmt.Sprintf("  - Synthetic Transaction : %x\n", tx.ToAccount[i].SyntheticTxId)
			}

			out += printGeneralTransactionParameters(res)
			return out, nil
		case "syntheticTokenDeposit":
			deposit := synthetic.TokenTransactionDeposit{}
			err := json.Unmarshal(*res.Data, &deposit)

			if err != nil {
				return "", err
			}

			out := "\n"
			amt, err := formatAmount(*deposit.TokenUrl.AsString(), &deposit.DepositAmount.Int)
			if err != nil {
				amt = "unknown"
			}
			out += fmt.Sprintf("Receive %s from %s to %s\n", amt, *deposit.FromUrl.AsString(),
				*deposit.ToUrl.AsString())

			out += printGeneralTransactionParameters(res)
			return out, nil

		default:
		}
	}
	return "", nil
}
