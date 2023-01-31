package config

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
	"testing"
	"yehuizhang.com/go-webapp-gin/pkg/flags"
)

func Test_InitConfig_OnSuccess(t *testing.T) {
	observedZapCore, observedLogs := observer.New(zap.InfoLevel)
	observedLogger := zap.New(observedZapCore).Sugar()
	config, err := InitConfig(&flags.FlagParser{Env: "local", ConfigName: ".env", ConfigPath: ".", ConfigType: "env"}, observedLogger)
	assert.Equal(t, nil, err)
	assert.Equal(t, "localhost:8080", config.GetString("GIN_PORT"))
	assert.Equal(t, zap.InfoLevel, observedLogs.All()[0].Level)
}

func Test_InitConfig_Invalid_File_Name(t *testing.T) {
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLogger := zap.New(observedZapCore).Sugar()
	_, err := InitConfig(&flags.FlagParser{Env: "", ConfigName: ".invalid", ConfigPath: ".", ConfigType: "env"}, observedLogger)
	assert.ErrorContains(t, err, "\".invalid\" Not Found")
}

func Test_InitConfig_Incorrect(t *testing.T) {
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLogger := zap.New(observedZapCore).Sugar()
	_, err := InitConfig(&flags.FlagParser{Env: "local", ConfigName: ".env", ConfigPath: ".", ConfigType: "json"}, observedLogger)
	assert.ErrorContains(t, err, "error occurred when initializing config.")
}
