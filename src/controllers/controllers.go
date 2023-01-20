package controllers

import "github.com/google/wire"

var ControllerSet = wire.NewSet(
	HealthControllerSet,
	UserAuthControllerSet,
	UserInfoControllerSet,
)
