package logger

import (
	"bytes"
	"fmt"
	"github.com/TskFok/GinApi/app/utils/conf"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var SugarLogger *zap.SugaredLogger

func init() {
	path := bytes.NewBufferString(conf.LoggerFilePath)
	path.WriteString(gin.Mode())
	path.WriteString(".log")

	logPath := path.String()

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil {
		fmt.Println(err)
	}

	if nil != err {
		fmt.Println(err.Error())
	}
	sync := zapcore.AddSync(file)

	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	core := zapcore.NewCore(encoder, sync, zap.DebugLevel)

	logger := zap.New(core)

	SugarLogger = logger.Sugar()
}

func Debug(debug interface{}) {
	defer SugarLogger.Sync()
	SugarLogger.Debug(debug)
}

func Error(error interface{}) {
	defer SugarLogger.Sync()
	SugarLogger.Error(error)
}

func Info(info interface{}) {
	defer SugarLogger.Sync()
	SugarLogger.Info(info)
}
