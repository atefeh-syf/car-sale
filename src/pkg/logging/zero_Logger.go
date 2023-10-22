package logging

import (
	"os"
	"sync"

	"github.com/atefeh-syf/car-sale/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var once sync.Once
var zeroSinLogger *zerolog.Logger 

type zeroLogger struct {
	cfg *config.Config
	Logger *zerolog.Logger 
}

var ZeroLogLevelMap = map[string] zerolog.Level {
	"debug": zerolog.DebugLevel,
	"info": zerolog.InfoLevel,
	"warn": zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
}

func newZeroLogger(cfg *config.Config) * zeroLogger {
	logger :=&zeroLogger{cfg:  cfg}
	logger.Init()
	return logger
}

func (l *zeroLogger) getLogLevel () zerolog.Level {
	level,  exist := ZeroLogLevelMap[ l.cfg.Logger.Level]
	if !exist {
		return  zerolog.DebugLevel
	}
	return level
}

func (l *zeroLogger) Init() {
	once.Do(func() {

		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

		file , err := os.OpenFile(l.cfg.Logger.FilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic("could not open file")		
		}
		var logger = zerolog.New(file).
									With().
									Timestamp().
									Str("AppName", "MyApp").
									Str("LoggerName", "Zerolog").
									Logger()

		zerolog.SetGlobalLevel(l.getLogLevel())
		zeroSinLogger = &logger
	})
	l.Logger = zeroSinLogger
}


func (l *zeroLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.Logger.
		Debug().
		Str("Category" , string(cat)).
		Str("SubCategory" , string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)

}
 
func (l *zeroLogger)  Debugf(template string, args ...interface{}) {
	l.Logger.Debug().Msgf(template, args...)
}

func (l *zeroLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.Logger.
		Info().
		Str("Category" , string(cat)).
		Str("SubCategory" , string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)

}
 
func (l *zeroLogger)  Infof(template string, args ...interface{}) {
	l.Logger.Debug().Msgf(template, args...)
}

func (l *zeroLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.Logger.
		Warn().
		Str("Category" , string(cat)).
		Str("SubCategory" , string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)

}
 
func (l *zeroLogger)  Warnf(template string, args ...interface{}) {
	l.Logger.Warn().Msgf(template, args...)
}


func (l *zeroLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.Logger.
		Error().
		Str("Category" , string(cat)).
		Str("SubCategory" , string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)

}
 
func (l *zeroLogger)  Errorf(template string, args ...interface{}) {
	l.Logger.Error().Msgf(template, args...)
}

func (l *zeroLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.Logger.
		Fatal().
		Str("Category" , string(cat)).
		Str("SubCategory" , string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)

}
 
func (l *zeroLogger)  Fatalf(template string, args ...interface{}) {
	l.Logger.Fatal().Msgf(template, args...)
}



