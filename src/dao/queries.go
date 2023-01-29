package dao

import (
	"github.com/google/wire"
	"yehuizhang.com/go-webapp-gin/src/dao/user/account"
	"yehuizhang.com/go-webapp-gin/src/dao/user/info"
)

var QuerySet = wire.NewSet(
	account.WireSet,
	info.WireSet)
