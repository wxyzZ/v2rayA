package conf

import (
	"sync"
	"time"
)

var (
	Version                  = "debug"
	FoundNew                 = false
	RemoteVersion            = ""
	TickerUpdateGFWList      *time.Ticker
	TickerUpdateSubscription *time.Ticker
	TickerUpdateServer       *time.Ticker
	UpdatingMu               sync.Mutex
	UpdatingMu2              sync.Mutex
)

func IsDebug() bool {
	return Version == "debug"
}
