package accsdk

import (
	"log"
	"net/url"
	"os"
	"os/user"
	"path/filepath"

	"github.com/AccumulateNetwork/accumulate/client"
	"github.com/AccumulateNetwork/accumulate/cmd/cli/cmd"
	"github.com/AccumulateNetwork/accumulate/cmd/cli/db"
)

const (
	defaultTestNet = "https://testnet.accumulatenetwork.io/v1"
)

var currentUser = func() *user.User {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr
}()

// Session is the basic configuration required for interaction
// with the Accumulate Network
type Session struct {
	API *client.APIClient
	Db  db.DB
}

func initDB(defaultWorkDir string) db.DB {

	err := os.MkdirAll(defaultWorkDir, 0600)
	if err != nil {
		log.Fatal(err)
	}

	db := new(db.BoltDB)
	err = db.InitDB(filepath.Join(defaultWorkDir, "wallet.db"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// MakeSession creates a new session required for any Sdk interaction
func MakeSession(address string) (session Session, err error) {
	url, err := url.Parse(address)
	if err != nil {
		return
	}

    cmd.WantJsonOutput = true
	cmd.Db = initDB(filepath.Join(currentUser.HomeDir, ".accumulate"))
    cmd.Client = &client.APIClient{
        Server: url.String(),
    }

	session = Session{
		API: cmd.Client,
		Db: cmd.Db,
	}

	return
}
