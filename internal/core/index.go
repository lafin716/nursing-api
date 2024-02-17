package core

import (
	"github.com/google/wire"
	"nursing_api/pkg/mono"
)

var Set = wire.NewSet(
	mono.NewMono,
)
