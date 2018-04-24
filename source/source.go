package source

import (
	"github.com/unicolony/hayes/source/db"
)

// Init app
func Init() {
	configInit()
	db.Init()
}
