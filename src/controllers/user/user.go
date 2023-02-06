package user

import (
	"github.com/google/wire"
	"yehuizhang.com/go-webapp-gin/pkg/dao/user/account"
	"yehuizhang.com/go-webapp-gin/pkg/dao/user/info"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
)

var WireSet = wire.NewSet(wire.Struct(new(Controller), "*"))

type Controller struct {
	Log          *logger.Logger
	AccountQuery account.IUserAccountQuery
	InfoQuery    info.IUserInfoQuery
}

const (
	UID string = "uid"
)
