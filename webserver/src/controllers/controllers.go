package controllers

import (
	"github.com/google/wire"
	"yehuizhang.com/go-webapp-gin/src/controllers/admin"
	"yehuizhang.com/go-webapp-gin/src/controllers/user"
)

var ControllerSet = wire.NewSet(
	admin.WireSet,
	user.WireSet,
)
