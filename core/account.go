package core

//"bytes"
//"context"
//"crypto/ed25519"
//"encoding/json"

//"io/ioutil"
//"log"
//"time"
//url2 "github.com/AccumulateNetwork/accumulate/internal/url"
//"github.com/AccumulateNetwork/accumulate/protocol"
//"github.com/AccumulateNetwork/accumulate/types"
//acmeapi "github.com/AccumulateNetwork/accumulate/types/api"
//"github.com/mdp/qrterminal"

// GenerateAccount creates a new Accumulate account and returns its attributes
//func GenerateAccount() (string, error) {
//	fmt.Println("Generating new account")
//	return GenerateKey("")
//}

//func GetAccount(url string) (string, error) {
//	var res acmeapi.APIDataResponse
//
//	params := acmeapi.APIRequestURL{}
//	params.URL = types.String(url)
//
//	if err := Client.Request(context.Background(), "token-account", params, &res); err != nil {
//		return PrintJsonRpcError(err)
//	}
//
//	return PrintQueryResponse(&res)
//}
//
//func QrAccount(s string) (string, error) {
//	u, err := url2.Parse(s)
//	if err != nil {
//		return "", fmt.Errorf("%q is not a valid Accumulate URL: %v\n", s, err)
//	}
//
//	b := bytes.NewBufferString("")
//	qrterminal.GenerateWithConfig(u.String(), qrterminal.Config{
//		Level:          qrterminal.M,
//		Writer:         b,
//		HalfBlocks:     true,
//		BlackChar:      qrterminal.BLACK_BLACK,
//		BlackWhiteChar: qrterminal.BLACK_WHITE,
//		WhiteChar:      qrterminal.WHITE_WHITE,
//		WhiteBlackChar: qrterminal.WHITE_BLACK,
//		QuietZone:      2,
//	})
//
//	r, err := ioutil.ReadAll(b)
//	return string(r), err
//}
//
////account create adiActor labelOrPubKeyHex height index tokenUrl keyBookUrl
//func CreateAccount(url string, args []string) (string, error) {
//	actor, err := url2.Parse(url)
//	if err != nil {
//		return "", err
//	}
//
//	args, si, privKey, err := prepareSigner(actor, args)
//	if len(args) < 3 {
//		return "", fmt.Errorf("insufficient number of command line arguments")
//	}
//
//	accountUrl, err := url2.Parse(args[0])
//	if err != nil {
//		return "", fmt.Errorf("invalid account url %s", args[0])
//	}
//	if actor.Authority != accountUrl.Authority {
//		return "", fmt.Errorf("account url to create (%s) doesn't match the authority adi (%s)", accountUrl.Authority, actor.Authority)
//	}
//	tok, err := url2.Parse(args[1])
//	if err != nil {
//		return "", fmt.Errorf("invalid token url")
//	}
//
//	var keybook string
//	if len(args) > 2 {
//		kbu, err := url2.Parse(args[2])
//		if err != nil {
//			return "", fmt.Errorf("invalid key book url")
//		}
//		keybook = kbu.String()
//	}
//
//	//make sure this is a valid token account
//	tokenJson, err := Get(tok.String())
//	if err != nil {
//		return "", err
//	}
//	token := protocol.TokenIssuer{}
//	err = json.Unmarshal([]byte(tokenJson), &token)
//	if err != nil {
//		return "", fmt.Errorf("invalid token type %v", err)
//	}
//
//	tac := &protocol.TokenAccountCreate{}
//	tac.Url = accountUrl.String()
//	tac.TokenUrl = tok.String()
//	tac.KeyBookUrl = keybook
//
//	binaryData, err := tac.MarshalBinary()
//	if err != nil {
//		return "", err
//	}
//
//	jsonData, err := json.Marshal(&tac)
//	if err != nil {
//		return "", err
//	}
//
//	nonce := uint64(time.Now().Unix())
//
//	params, err := prepareGenTx(jsonData, binaryData, actor, si, privKey, nonce)
//	if err != nil {
//		return "", err
//	}
//
//	var res acmeapi.APIDataResponse
//	if err := Client.Request(context.Background(), "token-account-create", params, &res); err != nil {
//		//todo: if we fail, then we need to remove the adi from storage or keep it and try again later...
//		return "", err
//	}
//
//	ar := ActionResponse{}
//	err = json.Unmarshal(*res.Data, &ar)
//	if err != nil {
//		return "", fmt.Errorf("error unmarshalling account create result")
//	}
//
//	return ar.Print()
//}

//func ListAccounts() (string, error) {
//
//	b, err := Db.GetBucket(BucketLabel)
//	if err != nil {
//		//no accounts so nothing to do...
//		return "", err
//	}
//	var out string
//	for _, v := range b.KeyValueList {
//		lt, err := protocol.AnonymousAddress(v.Value, protocol.AcmeUrl().String())
//		if err != nil {
//			continue
//		}
//		if lt.String() == string(v.Key) {
//			out += fmt.Sprintf("%s\n", v.Key)
//		}
//	}
//	//TODO: this probably should also list out adi accounts as well
//	return out, nil
//}
//
//func RestoreAccounts() (out string, err error) {
//	anon, err := Db.GetBucket(BucketAnon)
//	if err != nil {
//		//no anon accounts so nothing to do...
//		return
//	}
//	for _, v := range anon.KeyValueList {
//		u, err := url2.Parse(string(v.Key))
//		if err != nil {
//			out += fmt.Sprintf("%q is not a valid URL\n", v.Key)
//		}
//		key, _, err := protocol.ParseAnonymousAddress(u)
//		if err != nil {
//			out += fmt.Sprintf("%q is not a valid lite account: %v\n", v.Key, err)
//		} else if key == nil {
//			out += fmt.Sprintf("%q is not a lite account\n", v.Key)
//		}
//
//		privKey := ed25519.PrivateKey(v.Value)
//		pubKey := privKey.Public().(ed25519.PublicKey)
//		out += fmt.Sprintf("Converting %s : %x\n", v.Key, pubKey)
//
//		err = Db.Put(BucketLabel, v.Key, pubKey)
//		if err != nil {
//			log.Fatal(err)
//		}
//		err = Db.Put(BucketKeys, pubKey, privKey)
//		if err != nil {
//			return "", err
//		}
//		err = Db.DeleteBucket(BucketAnon)
//		if err != nil {
//			return "", err
//		}
//	}
//	return out, nil
//}