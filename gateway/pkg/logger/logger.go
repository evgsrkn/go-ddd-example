package logger

import (
	"fmt"
	"os"

	"github.com/evgsrkn/go-ddd-example/gateway/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type fileSyncer struct {
	f  string
	fd *os.File
}

func New(appEnv config.AppEnv) *zap.Logger {
	core := newZapCore(appEnv)
	return zap.New(core)
}

func mustNewFileSyncer(f string) *fileSyncer {
	fs := &fileSyncer{
		f: f,
	}

	fd, err := os.OpenFile(fs.f, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err.Error())
	}

	fs.fd = fd
	return fs
}

func (fs *fileSyncer) Write(p []byte) (int, error) {
	return fs.fd.Write(p)
}

func (fs *fileSyncer) Sync() error {
	return fs.fd.Close()
}

func newZapCore(appEnv config.AppEnv) zapcore.Core {
	if err := os.MkdirAll("./tmp", 0775); err != nil {
		panic(err.Error())
	}

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	errLogPath := fmt.Sprintf("./tmp/%v.log", appEnv.String())
	fsErrors := zapcore.Lock(mustNewFileSyncer(errLogPath))

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	fsEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	return zapcore.NewTee(
		zapcore.NewCore(fsEncoder, fsErrors, zap.ErrorLevel),
		zapcore.NewCore(consoleEncoder, consoleErrors, zap.ErrorLevel),
		zapcore.NewCore(consoleEncoder, consoleDebugging, zap.DebugLevel),
	)
}
