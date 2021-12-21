package config

import (
	"github.com/salasberryfin/accumulate-sdk-go/api"
	"github.com/salasberryfin/accumulate-sdk-go/db"
)

var (
	Client         = api.NewAPIClient()
	Db             db.DB
	WantJsonOutput = false
)

func ConfigureAPIClient(server string) {
	Client = api.CustomAPIClient(server)
}

func ConfigureDb(database db.DB) {
	Db = database
}
