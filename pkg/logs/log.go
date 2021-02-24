package logs

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"goutils/pkg/files"
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

type zapLogger struct {
	logDebug *zap.Logger
	logInfo  *zap.Logger
	logWarn  *zap.Logger
	logError *zap.Logger
	ops      *option
}

func (z *zapLogger) Debugf(format string, a ...interface{}) {
	z.logDebug.WithOptions(zap.AddCallerSkip(1)).Debug(fmt.Sprintf(format, a...))
}

func (z *zapLogger) Infof(format string, a ...interface{}) {
	z.logInfo.WithOptions(zap.AddCallerSkip(1)).Info(fmt.Sprintf(format, a...))
}

func (z *zapLogger) Warnf(format string, a ...interface{}) {
	z.logWarn.WithOptions(zap.AddCallerSkip(1)).Warn(fmt.Sprintf(format, a...))
}

func (z *zapLogger) Errorf(format string, a ...interface{}) {
	z.logError.WithOptions(zap.AddCallerSkip(1)).Error(fmt.Sprintf(format, a...))
}

func (z *zapLogger) DebugDepth(msg string, dep int) {
	z.logDebug.WithOptions(zap.AddCallerSkip(dep)).Debug(msg)
}

func (z *zapLogger) InfoDepth(msg string, dep int) {
	z.logInfo.WithOptions(zap.AddCallerSkip(dep)).Info(msg)
}

func (z *zapLogger) WarnDepth(msg string, dep int) {
	z.logWarn.WithOptions(zap.AddCallerSkip(dep)).Warn(msg)
}

func (z *zapLogger) ErrorDepth(msg string, dep int) {
	z.logError.WithOptions(zap.AddCallerSkip(dep)).Error(msg)
}

func (z *zapLogger) Debug(msg string) {
	z.logDebug.WithOptions(zap.AddCallerSkip(1)).Debug(msg)
}

func (z *zapLogger) Info(msg string) {
	z.logInfo.WithOptions(zap.AddCallerSkip(1)).Info(msg)
}

func (z *zapLogger) Warn(msg string) {
	z.logWarn.WithOptions(zap.AddCallerSkip(1)).Warn(msg)
}

func (z *zapLogger) Error(msg string) {
	z.logError.WithOptions(zap.AddCallerSkip(1)).Error(msg)
}
func (z *zapLogger) init() {
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
func (z *zapLogger) initOne(level zapcore.Level) *zap.Logger {

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

		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),                   // 编码器配置
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 打印到控制台和文件
			atomicLevel, // 日志级别
		)
	}

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	//filed := zap.Fields(zap.String("serviceName", "serviceName"))

	return zap.New(core, caller, development)
}
func NewZopLogger(ops ...OptionFunc) Logger {

	op := &option{}
	for _, v := range ops {
		v(op)
	}
	zp := &zapLogger{ops: op}
	zp.init()
	return zp
}

//maxSize mini size is 1Mib
func NewLogger(logDir string, maxSize, maxBackups int) Logger {

	return NewZopLogger(WithOptionLogFileConfig(logDir, maxSize, maxBackups))
}
