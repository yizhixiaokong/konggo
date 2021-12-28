package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Handle    *zap.Logger
	Encrypted bool // 是否加密
)

type KV map[string]interface{}

func Info(msg ...string) {
	if Encrypted {
		enMsg, _ := EnPwdCode([]byte(ConcatWithPart(", ", msg...)))
		Handle.Info(" " + enMsg)
	} else {
		Handle.Info(ConcatWithPart(", ", msg...))
	}
}

func InfoKV(msg string, param KV) {
	if Encrypted {
		mjson, _ := json.Marshal(param)
		enMsg, _ := EnPwdCode([]byte(msg + " " + string(mjson)))
		Handle.Info(" " + enMsg)
	} else {
		Handle.Info(msg, kvToField(param)...)
	}
}

func Infof(format string, a ...interface{}) {
	if Encrypted {
		enMsg, _ := EnPwdCode([]byte(fmt.Sprintf(format, a...)))
		Handle.Info(" " + enMsg)
	} else {
		Handle.Info(fmt.Sprintf(format, a...))
	}
}

func Error(msg ...string) {
	if Encrypted {
		enMsg, _ := EnPwdCode([]byte(ConcatWithPart(", ", msg...)))
		Handle.Error(" " + enMsg)
	} else {
		Handle.Error(ConcatWithPart(", ", msg...))
	}
}

func ErrorKV(msg string, param KV) {
	if Encrypted {
		mjson, _ := json.Marshal(param)
		enMsg, _ := EnPwdCode([]byte(msg + " " + string(mjson)))
		Handle.Error(" " + enMsg)
	} else {
		Handle.Error(msg, kvToField(param)...)
	}
}

func Errorf(format string, a ...interface{}) {
	if Encrypted {
		enMsg, _ := EnPwdCode([]byte(fmt.Sprintf(format, a...)))
		Handle.Error(" " + enMsg)
	} else {
		Handle.Error(fmt.Sprintf(format, a...))
	}
}
func Errorp(param interface{}, err interface{}) {
	Handle.Error(fmt.Sprintf("param: %+v, err: %v", param, err))
}

func Warn(msg ...string) {
	if Encrypted {
		enMsg, _ := EnPwdCode([]byte(ConcatWithPart(", ", msg...)))
		Handle.Warn(" " + enMsg)
	} else {
		Handle.Warn(ConcatWithPart(", ", msg...))
	}
}

func Warnf(format string, a ...interface{}) {
	if Encrypted {
		enMsg, _ := EnPwdCode([]byte(fmt.Sprintf(format, a...)))
		Handle.Warn(" " + enMsg)
	} else {
		Handle.Warn(fmt.Sprintf(format, a...))
	}
}

func WarnKV(msg string, param KV) {
	if Encrypted {
		mjson, _ := json.Marshal(param)
		enMsg, _ := EnPwdCode([]byte(msg + " " + string(mjson)))
		Handle.Warn(" " + enMsg)
	} else {
		Handle.Warn(msg, kvToField(param)...)
	}
}

func Debug(msg ...string) {
	if Encrypted {
		enMsg, _ := EnPwdCode([]byte(ConcatWithPart(" ", msg...)))
		Handle.Debug(" " + enMsg)
	} else {
		Handle.Debug(ConcatWithPart(" ", msg...))
	}
}

func DebugKV(msg string, param KV) {
	if Encrypted {
		mjson, _ := json.Marshal(param)
		enMsg, _ := EnPwdCode([]byte(msg + " " + string(mjson)))
		Handle.Debug(" " + enMsg)
	} else {
		Handle.Debug(msg, kvToField(param)...)
	}
}

func Debugf(format string, a ...interface{}) {
	if Encrypted {
		enMsg, _ := EnPwdCode([]byte(fmt.Sprintf(format, a...)))
		Handle.Debug(" " + enMsg)
	} else {
		Handle.Debug(fmt.Sprintf(format, a...))
	}
}

// 非加密日志
// func Info(msg ...string) {
//     Handle.Info(ConcatWithPart(", ", msg...))
// }

// func InfoKV(msg string, param KV) {
//    Handle.Info(msg, kvToField(param)...)
// }

// func Infof(format string, a ...interface{}) {
//    Handle.Info(fmt.Sprintf(format, a...))
// }

// func Error(msg ...string) {
//    Handle.Error(ConcatWithPart(", ", msg...))
// }

// func ErrorKV(msg string, param KV) {
//    Handle.Error(msg, kvToField(param)...)
// }

// func Errorf(format string, a ...interface{}) {
//    Handle.Error(fmt.Sprintf(format, a...))
// }

// func Warn(msg ...string) {
//    Handle.Warn(ConcatWithPart(", ", msg...))
// }

// func WarnKV(msg string, param KV) {
//    Handle.Warn(msg, kvToField(param)...)
// }

// func Debug(msg ...string) {
//    Handle.Debug(ConcatWithPart(", ", msg...))
// }

// func DebugKV(msg string, param KV) {
//    Handle.Debug(msg, kvToField(param)...)
// }

// func Debugf(format string, a ...interface{}) {
//    Handle.Debug(fmt.Sprintf(format, a...))
// }

func InitLog(applicationPath string, encrypted bool) {
	Encrypted = encrypted
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:    "time",
		LevelKey:   "level",
		NameKey:    "logger",
		CallerKey:  "caller",
		MessageKey: "msg",
		//StacktraceKey:  "stacktrace",
		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseColorLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		//EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
		EncodeCaller: CustomCallerEncoder,
	}

	// 生成日志目录
	logPath := applicationPath + "/log/"
	if err := os.MkdirAll(logPath, os.ModePerm); err != nil {
		panic("log dir init failed")
	}

	//自定义日志级别：自定义Info级别
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	// 获取io.Writer的实现
	infoWriter := getWriter(logPath + "info.log")
	errorWriter := getWriter(logPath + "error.log")

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(errorWriter), warnLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), infoLevel), //同时将日志输出到控制台
	)

	Handle = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.InfoLevel))

	Info("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	Info("日志初始化成功")
	Info("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}

func kvToField(param KV) (list []zap.Field) {
	for key, value := range param {
		list = append(list, Any(key, value))
	}
	return list
}

func Any(key string, value interface{}) zap.Field {
	switch val := value.(type) {
	case error:
		return zap.Any(key, value.(error).Error())
	default:
		return zap.Any(key, val)
	}
}

func ConcatWithPart(part string, s ...string) (result string) {
	for i, v := range s {
		if i != 0 {
			result += part
		}
		result += v
	}
	return result
}

func getWriter(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    1, // megabytes
		MaxBackups: 30,
		MaxAge:     30,   // days
		LocalTime:  true, // 本地时间
	}
}

func CustomCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	pc, file, line, ok := runtime.Caller(6)
	caller = zapcore.EntryCaller{
		Defined: ok,
		PC:      pc,
		File:    file,
		Line:    line,
	}
	enc.AppendString(caller.TrimmedPath()) // 短路径
	//enc.AppendString(caller.String())      // 长路径
}
