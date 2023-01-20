package models

import (
	"github.com/google/wire"
	"yehuizhang.com/go-webapp-gin/src/models/user"
)

var ModelsSet = wire.NewSet(
	user.HandlerSet)
