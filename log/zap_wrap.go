package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

type Logger struct {
	*zap.Logger
}

func Sync() {
	_ = std.Logger.Sync()
}

var std *Logger

func init() {
	logger, err := stdCfg.Build()
	if err != nil {
		panic(err)
	}
	std = &Logger{Logger: logger}
}

var stdCfg = &zap.Config{
	Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
	Development:       false,
	DisableCaller:     false,
	DisableStacktrace: true,
	Sampling: &zap.SamplingConfig{
		Initial:    100,
		Thereafter: 100,
	},
	Encoding: "json",
	EncoderConfig: zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	},
	OutputPaths:      []string{"stdout"},
	ErrorOutputPaths: []string{"stderr"},
	InitialFields:    map[string]interface{}{"host": getHostname()},
}

func getHostname() string {
	host, _ := os.Hostname()
	return host
}

func (l *Logger) Debug(msg string, fields ...zap.Field)  { l.Logger.Info(msg, fields...) }
func (l *Logger) Info(msg string, fields ...zap.Field)   { l.Logger.Info(msg, fields...) }
func (l *Logger) Warn(msg string, fields ...zap.Field)   { l.Logger.Info(msg, fields...) }
func (l *Logger) Error(msg string, fields ...zap.Field)  { l.Logger.Info(msg, fields...) }
func (l *Logger) DPanic(msg string, fields ...zap.Field) { l.Logger.Info(msg, fields...) }
func (l *Logger) Panic(msg string, fields ...zap.Field)  { l.Logger.Info(msg, fields...) }
func (l *Logger) Fatal(msg string, fields ...zap.Field)  { l.Logger.Info(msg, fields...) }

func (l *Logger) Infof(msg string, args ...interface{})  { l.Logger.Info(fmt.Sprintf(msg, args...)) }
func (l *Logger) Errorf(msg string, args ...interface{}) { l.Logger.Error(fmt.Sprintf(msg, args...)) }

func (l *Logger) Print(v ...interface{})                 { l.Logger.Debug(fmt.Sprint(v...)) }
func (l *Logger) Printf(format string, v ...interface{}) { l.Logger.Debug(fmt.Sprintf(format, v...)) }
func (l *Logger) Println(v ...interface{})               { l.Logger.Debug(fmt.Sprintln(v...)) }

func (l *Logger) Write(p []byte) (n int, err error) { l.Logger.Debug(string(p)); return len(p), nil }

func Writer() io.Writer                      { return std }
func Debug(msg string, fields ...zap.Field)  { std.Logger.Debug(msg, fields...) }
func Info(msg string, fields ...zap.Field)   { std.Logger.Info(msg, fields...) }
func Warn(msg string, fields ...zap.Field)   { std.Logger.Warn(msg, fields...) }
func Error(msg string, fields ...zap.Field)  { std.Logger.Error(msg, fields...) }
func DPanic(msg string, fields ...zap.Field) { std.Logger.DPanic(msg, fields...) }
func Panic(msg string, fields ...zap.Field)  { std.Logger.Panic(msg, fields...) }
func Fatal(msg string, fields ...zap.Field)  { std.Logger.Fatal(msg, fields...) }
