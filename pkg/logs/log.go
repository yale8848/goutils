package logs

import (
	"fmt"
	"github.com/yale8848/goutils/pkg/files"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)

	Debugf(format string, a ...interface{})
	Infof(format string, a ...interface{})
	Warnf(format string, a ...interface{})
	Errorf(format string, a ...interface{})

	DebugDepth(msg string, dep int)
	InfoDepth(msg string, dep int)
	WarnDepth(msg string, dep int)
	ErrorDepth(msg string, dep int)
}
type optionFile struct {
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	LocalTime  bool
	Compress   bool
}
type option struct {
	fileConfig *optionFile
}
type OptionFunc func(op *option)

type ZapLogger struct {
	logDebug *zap.Logger
	logInfo  *zap.Logger
	logWarn  *zap.Logger
	logError *zap.Logger
	ops      *option
}

func (z *ZapLogger) Debugf(format string, a ...interface{}) {
	if z.logDebug==nil {
        fmt.Println(fmt.Sprintf(format, a...))
		return
	}
	z.logDebug.WithOptions(zap.AddCallerSkip(1)).Debug(fmt.Sprintf(format, a...))
}

func (z *ZapLogger) Infof(format string, a ...interface{}) {
	if z.logInfo==nil {
		fmt.Println(fmt.Sprintf(format, a...))
		return
	}
	z.logInfo.WithOptions(zap.AddCallerSkip(1)).Info(fmt.Sprintf(format, a...))
}

func (z *ZapLogger) Warnf(format string, a ...interface{}) {
	if z.logWarn==nil {
		fmt.Println(fmt.Sprintf(format, a...))
		return
	}
	z.logWarn.WithOptions(zap.AddCallerSkip(1)).Warn(fmt.Sprintf(format, a...))
}

func (z *ZapLogger) Errorf(format string, a ...interface{}) {
	if z.logError==nil {
		fmt.Println(fmt.Sprintf(format, a...))
		return
	}
	z.logError.WithOptions(zap.AddCallerSkip(1)).Error(fmt.Sprintf(format, a...))
}

func (z *ZapLogger) DebugDepth(msg string, dep int) {
	if z.logDebug==nil {
		fmt.Println(msg)
		return
	}
	z.logDebug.WithOptions(zap.AddCallerSkip(dep)).Debug(msg)
}

func (z *ZapLogger) InfoDepth(msg string, dep int) {
	if z.logInfo==nil {
		fmt.Println(msg)
		return
	}
	z.logInfo.WithOptions(zap.AddCallerSkip(dep)).Info(msg)
}

func (z *ZapLogger) WarnDepth(msg string, dep int) {
	if z.logWarn==nil {
		fmt.Println(msg)
		return
	}
	z.logWarn.WithOptions(zap.AddCallerSkip(dep)).Warn(msg)
}

func (z *ZapLogger) ErrorDepth(msg string, dep int) {
	if z.logError==nil {
		fmt.Println(msg)
		return
	}
	z.logError.WithOptions(zap.AddCallerSkip(dep)).Error(msg)
}

func (z *ZapLogger) Debug(msg string) {
	if z.logDebug==nil {
		fmt.Println(msg)
		return
	}
	z.logDebug.WithOptions(zap.AddCallerSkip(1)).Debug(msg)
}

func (z *ZapLogger) Info(msg string) {
	if z.logInfo==nil {
		fmt.Println(msg)
		return
	}
	z.logInfo.WithOptions(zap.AddCallerSkip(1)).Info(msg)
}

func (z *ZapLogger) Warn(msg string) {
	if z.logWarn==nil {
		fmt.Println(msg)
		return
	}
	z.logWarn.WithOptions(zap.AddCallerSkip(1)).Warn(msg)
}

func (z *ZapLogger) Error(msg string) {
	if z.logError==nil {
		fmt.Println(msg)
		return
	}
	z.logError.WithOptions(zap.AddCallerSkip(1)).Error(msg)
}
func (z *ZapLogger) init() {
	z.logDebug = z.initOne(zapcore.DebugLevel)
	z.logInfo = z.initOne(zapcore.InfoLevel)
	z.logWarn = z.initOne(zapcore.WarnLevel)
	z.logError = z.initOne(zapcore.ErrorLevel)

}

func WithOptionLogFileConfig(fName string, maxSize, maxBackups int) OptionFunc {

	files.MkDirs(fName)

	return func(op *option) {
		op.fileConfig = &optionFile{
			Filename:   fName,
			MaxBackups: maxBackups,
			MaxSize:    maxSize,
		}
	}

}
func (z *ZapLogger) initOne(level zapcore.Level) *zap.Logger {

	var fileWriter io.Writer
	if z.ops.fileConfig != nil && len(z.ops.fileConfig.Filename) > 0 {
		fn := ""
		if level == zap.DebugLevel {
			fn = "debug"
		}
		if level == zap.InfoLevel {
			fn = "info"
		}
		if level == zap.WarnLevel {
			fn = "warn"
		}
		if level == zap.ErrorLevel {
			fn = "err"
		}
		cfn := z.ops.fileConfig.Filename
		nfn := filepath.Join(filepath.Base(cfn), fn+"_"+filepath.Base(cfn))

		fileWriter = &lumberjack.Logger{
			Filename:   nfn,                         // 日志文件路径
			MaxSize:    z.ops.fileConfig.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: z.ops.fileConfig.MaxBackups, // 日志文件最多保存多少个备份
		}
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别,debug可以打印出info,debug,warn；info级别可以打印warn，info；warn只能打印warn
	// debug->info->warn->error
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	var core zapcore.Core
	if fileWriter != nil {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),                                                // 编码器配置
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), // 打印到控制台和文件
			atomicLevel, // 日志级别
		)
	} else {

		fmt.Println("fileWriter == nil")

		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),                   // 编码器配置
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 打印到控制台和文件
			atomicLevel, // 日志级别
		)
	}
	if core==nil {
		fmt.Println("core == nil")
	}

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	//filed := zap.Fields(zap.String("serviceName", "serviceName"))

	return zap.New(core, caller, development)
}
func NewZopLogger(ops ...OptionFunc) *ZapLogger {

	op := &option{}
	for _, v := range ops {
		v(op)
	}
	zp := &ZapLogger{ops: op}
	zp.init()
	return zp
}

//maxSize mini size is 1Mib
func NewLogger(logDir string, maxSize, maxBackups int) *ZapLogger {

	return NewZopLogger(WithOptionLogFileConfig(logDir, maxSize, maxBackups))
}
func NewLoggerInterface(logDir string, maxSize, maxBackups int) Logger {

	return NewZopLogger(WithOptionLogFileConfig(logDir, maxSize, maxBackups))
}