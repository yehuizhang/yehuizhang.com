package admin

import "github.com/google/wire"

var WireSet = wire.NewSet(wire.Struct(new(Controller), "*"))

type Controller struct {
}
