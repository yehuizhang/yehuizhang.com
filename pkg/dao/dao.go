package dao

import (
	"github.com/google/wire"
	"yehuizhang.com/go-webapp-gin/pkg/dao/user/account"
	"yehuizhang.com/go-webapp-gin/pkg/dao/user/info"
)

var WireSet = wire.NewSet(
	account.InitUserAccountQuery,
	info.InitUserInfoQuery,
	transactionSet,
)
