package config

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
	"testing"
	"yehuizhang.com/go-webapp-gin/pkg/flag_parser"
)

func Test_InitConfig_OnSuccess(t *testing.T) {
	observedZapCore, observedLogs := observer.New(zap.InfoLevel)
	observedLogger := zap.New(observedZapCore).Sugar()
	config, err := InitConfig(&flag_parser.FlagParser{Env: "local"}, observedLogger)
	assert.Equal(t, nil, err)
	assert.Equal(t, "localhost:8080", config.GetString("server.port"))
	assert.Equal(t, zap.InfoLevel, observedLogs.All()[0].Level)
}

func Test_InitConfig_MissingFail(t *testing.T) {
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLogger := zap.New(observedZapCore).Sugar()
	_, err := InitConfig(&flag_parser.FlagParser{Env: "invalid"}, observedLogger)
	assert.Errorf(t, err, "error on parsing configuration file")
}
