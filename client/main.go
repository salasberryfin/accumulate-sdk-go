package accsdk

import (
	"log"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"reflect"

	"github.com/AccumulateNetwork/accumulate/client"
	"github.com/AccumulateNetwork/accumulate/cmd/cli/cmd"
	"github.com/AccumulateNetwork/accumulate/cmd/cli/db"

	"github.com/salasberryfin/accumulate-sdk-go/api"
)

var currentUser = func() *user.User {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr
}()

// Options for Session creation
type Options struct {
	ServerURL  string `default:"https://testnet.accumulatenetwork.io/v1"`
	JSONOutput bool   `default:"true"`
}

// Session is the basic unit required for interaction
// with the Accumulate Network
type Session struct {
	API        *api.Client
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

// NewSession creates a new session required for any Sdk interaction
func NewSession(options Options) (session Session, err error) {
	log.Println("Initializing session.")

	typ := reflect.TypeOf(options)
	if options.ServerURL == "" {
		f, _ := typ.FieldByName("ServerURL")
		options.ServerURL = f.Tag.Get("default")
	}

	url, err := url.Parse(options.ServerURL)
	if err != nil {
		return
	}

	cmd.WantJsonOutput = options.JSONOutput
	cmd.Db = initDB(filepath.Join(currentUser.HomeDir, ".accumulate"))
	cmd.Client = &client.APIClient{
		Server: url.String(),
	}

	session = Session{
		API:        api.NewAPIClient(url.String()),
		Db:         cmd.Db,
		JSONOutput: options.JSONOutput,
	}

	return
}
