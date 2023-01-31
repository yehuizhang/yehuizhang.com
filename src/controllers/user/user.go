package user

import (
	"github.com/google/wire"
	"yehuizhang.com/go-webapp-gin/pkg/database"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
	"yehuizhang.com/go-webapp-gin/src/dao/user/account"
	"yehuizhang.com/go-webapp-gin/src/dao/user/info"
)

var WireSet = wire.NewSet(wire.Struct(new(Controller), "*"))

type Controller struct {
	Log          *logger.Logger
	Db           *database.Database
	AccountQuery account.IUserAccountQuery
	InfoQuery    info.IUserInfoQuery
}

const (
	UID string = "uid"
)
