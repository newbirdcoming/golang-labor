package logger

import (
	"os"
	"time"

	config "zap-test/cons"
	"zap-test/setting"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// func InitLogger() (err error) {
// 	// Create a new logger instance
// 	// Logger, err = zap.NewProduction()
// 	Logger, err = zap.NewDevelopment()
// 	if err != nil {
// 		return err
// 	}
// 	// Return the logger instance
// 	return nil
// }

func zapEncoder(config *setting.ZapConfig) zapcore.Encoder {
	// 新建一个配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "Time",
		LevelKey:      "Level",
		NameKey:       "Logger",
		CallerKey:     "Caller",
		MessageKey:    "Message",
		StacktraceKey: "StackTrace",
		LineEnding:    zapcore.DefaultLineEnding,
		FunctionKey:   zapcore.OmitKey,
	}
	// 自定义时间格式
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(config.Prefix + "\t" + t.Format(config.TimeFormat))
	}

	// 日志级别大写
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// 秒级时间间隔
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	// 简短的调用者输出
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 完整的序列化logger名称
	encoderConfig.EncodeName = zapcore.FullNameEncoder
	// 最终的日志编码 json或者console
	switch config.Encode {
	case "json":
		{
			return zapcore.NewJSONEncoder(encoderConfig)
		}
	case "console":
		{
			return zapcore.NewConsoleEncoder(encoderConfig)
		}
	}
	// 默认console
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func zapWriteSyncer(cfg *setting.ZapConfig) zapcore.WriteSyncer {
	syncers := make([]zapcore.WriteSyncer, 0, 2)
	// 如果开启了日志控制台输出，就加入控制台书写器
	if cfg.Writer == config.WriteBoth || cfg.Writer == config.WriteConsole {
		syncers = append(syncers, zapcore.AddSync(os.Stdout))
	}

	// 如果开启了日志文件存储，就根据文件路径切片加入书写器
	if cfg.Writer == config.WriteBoth || cfg.Writer == config.WriteFile {
		// 添加日志输出器
		for _, path := range cfg.LogFile.Output {
			logger := &lumberjack.Logger{
				Filename:   path,                 //文件路径
				MaxSize:    cfg.LogFile.MaxSize,  //分割文件的大小
				MaxBackups: cfg.LogFile.BackUps,  //备份次数
				Compress:   cfg.LogFile.Compress, // 是否压缩
				LocalTime:  true,                 //使用本地时间
			}
			syncers = append(syncers, zapcore.Lock(zapcore.AddSync(logger)))
		}
	}
	return zap.CombineWriteSyncers(syncers...)
}

func zapLevelEnabler(cfg *setting.ZapConfig) zapcore.LevelEnabler {
	switch cfg.Level {
	case config.DebugLevel:
		return zap.DebugLevel
	case config.InfoLevel:
		return zap.InfoLevel
	case config.ErrorLevel:
		return zap.ErrorLevel
	case config.PanicLevel:
		return zap.PanicLevel
	case config.FatalLevel:
		return zap.FatalLevel
	}
	// 默认Debug级别
	return zap.DebugLevel
}

func InitZap(config *setting.ZapConfig) *zap.Logger {
	// 构建编码器
	encoder := zapEncoder(config)
	// 构建日志级别
	levelEnabler := zapLevelEnabler(config)
	// 最后获得Core和Options
	subCore, options := tee(config, encoder, levelEnabler)
	// 创建Logger
	Logger = zap.New(subCore, options...)
	return Logger
}

// 将所有合并
func tee(cfg *setting.ZapConfig, encoder zapcore.Encoder, levelEnabler zapcore.LevelEnabler) (core zapcore.Core, options []zap.Option) {
	sink := zapWriteSyncer(cfg)
	return zapcore.NewCore(encoder, sink, levelEnabler), buildOptions(cfg, levelEnabler)
}

// 构建Option
func buildOptions(cfg *setting.ZapConfig, levelEnabler zapcore.LevelEnabler) (options []zap.Option) {
	if cfg.Caller {
		options = append(options, zap.AddCaller())
	}

	if cfg.StackTrace {
		options = append(options, zap.AddStacktrace(levelEnabler))
	}
	return
}
