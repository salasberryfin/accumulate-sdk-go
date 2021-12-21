package accsdk

import (
	"log"
	"net/url"
	"os"
	"os/user"
	"path/filepath"

	"github.com/salasberryfin/accumulate-sdk-go/api"
	"github.com/salasberryfin/accumulate-sdk-go/config"
	"github.com/salasberryfin/accumulate-sdk-go/db"
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
	Client     *api.APIClient
	Db         db.DB
	JSONOutput bool
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

	sessionJSON := true
	sessionDb := initDB(filepath.Join(currentUser.HomeDir, ".accumulate"))

	// configure parameters for Sdk
	config.ConfigureAPIClient(address)
	config.ConfigureDb(sessionDb)
	config.WantJsonOutput = sessionJSON

	session = Session{
		Client: &api.APIClient{
			Server: url.String(),
		},
		Db:         sessionDb,
		JSONOutput: sessionJSON,
	}

	return
}
