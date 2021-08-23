package cordyn

import (
	"github.com/google/wire"
)

// WireSet ...
var WireSet = wire.NewSet(NewDB)
