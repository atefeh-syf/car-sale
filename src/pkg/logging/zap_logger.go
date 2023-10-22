package logging

import (
	"github.com/atefeh-syf/car-sale/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var LogLevelMap = map[string] zapcore.Level {
	"debug": zapcore.DebugLevel,
	"info": zapcore.InfoLevel,
	"warn": zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}
type ZapLogger struct {
	cfg *config.Config
	Logger *zap.SugaredLogger
}

func newZapLogger(cfg *config.Config) *ZapLogger {
	logger := &ZapLogger{cfg:  cfg}
	logger.Init()
	return logger
}

func (l *ZapLogger) getLogLevel() zapcore.Level {
	level, exist := LogLevelMap[l.cfg.Logger.Level]
	if !exist {
		return zapcore.DebugLevel
	}
	return level
}

func (l *ZapLogger) Init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:  l.cfg.Logger.FilePath,
		MaxSize: 1,
		MaxAge: 5,
		LocalTime: true,
		MaxBackups: 10,
		Compress: true,
	})
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		l.getLogLevel(),
	)

	Logger := zap.New(core,  zap.AddCaller() , zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
	Logger = Logger.With("AppName", "MyApp", "LoggerName", "zap")
	l.Logger  = Logger
}

func (l *ZapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, sub)
	l.Logger.Debugw(msg, params...)
}
 
func (l *ZapLogger)  Debugf(template string, args ...interface{}) {
	l.Logger.Debugf(template, args)
}

func (l *ZapLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, sub)
	l.Logger.Infow(msg, params...)
}
 
func (l *ZapLogger)  Infof(template string, args ...interface{}) {
	l.Logger.Infof(template, args)
}

func (l *ZapLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, sub)
	l.Logger.Warnw(msg, params...)
}
 
func (l *ZapLogger)  Warnf(template string, args ...interface{}) {
	l.Logger.Warnf(template, args)
}

func (l *ZapLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, sub)
	l.Logger.Errorw(msg, params...)
}
 
func (l *ZapLogger)  Errorf(template string, args ...interface{}) {
	l.Logger.Errorf(template, args)
}

func (l *ZapLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, sub)
	l.Logger.Fatalw(msg, params...)
}
 
func (l *ZapLogger)  Fatalf(template string, args ...interface{}) {
	l.Logger.Fatalf(template, args)
}

func prepareLogKeys(extra map[ExtraKey]interface{}, cat Category, sub SubCategory)  []interface{}{
	if extra == nil {
		extra = make(map[ExtraKey]interface{}, 0)
	}
	extra["Category"] = cat
	extra["SubCategory"] = sub
	params := mapToZapParams(extra)
	return params
}