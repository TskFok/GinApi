package logger

import (
	"bytes"
	"fmt"
	"github.com/TskFok/GinApi/app/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitLogger() *zap.SugaredLogger {
	path := bytes.NewBufferString(global.LoggerFilePath)
	path.WriteString(gin.Mode())
	path.WriteString(".log")

	logPath := path.String()

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)

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
	sugarLogger := logger.Sugar()

	return sugarLogger
}
