package main

import (
	"encoding/json"
	"go.uber.org/zap"
	"testing"
)

func TestZabConfig(t *testing.T) {
	rawJSON := []byte(`{
    "level":"debug",
    "encoding":"json",
    "outputPaths": ["stdout", "test.log"],
    "errorOutputPaths": ["stderr"],
    "initialFields":{"name":"dj"},
    "encoderConfig": {
      "messageKey": "message",
      "levelKey": "level",
      "levelEncoder": "lowercase"
    }
  }`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("server start work successfully!")
}

func TestZabOption(t *testing.T) {
	// AddStacktrace(lvl zapcore.LevelEnabler)：用来在指定级别及以上级别输出调用堆栈。
	// zap.WithCaller(enabled bool)：指定是否在日志输出内容中增加文件名和行号。
	// zap.AddCaller()：与zap.WithCaller(true)等价，指定在日志输出内容中增加行号和文件名。
	// zap. AddCallerSkip(skip int)：指定在调用栈中跳过的调用深度，否则通过调用栈获得的行号可能总是日志组件中的行号。
	// zap. IncreaseLevel(lvl zapcore.LevelEnabler)：提高日志级别，如果传入的lvl比当前logger的级别低，则不会改变日志级别。
	// ErrorOutput(w zapcore.WriteSyncer)：指定日志组件中出现异常时的输出位置。
	// Fields(fs ...Field)：添加公共字段。
	// Hooks(hooks ...func(zapcore.Entry) error)：注册钩子函数，用来在日志打印时同时调用hook方法。
	// WrapCore(f func(zapcore.Core) zapcore.Core)：替换Logger的zapcore.Core。 - Development()：将Logger修改为Development模式。
	logger, _ := zap.NewProduction(zap.AddCaller())
	logger.Info("hello world")
}

func TestZabWithField(t *testing.T) {
	logger := zap.NewExample(zap.Fields(
		zap.Int("userId", 10),
		zap.String("requestId", "fbf54504"),
	))
	logger.Debug("This is a debug message")
	logger.Info("This is a info message")
}

func TestGlobalLogger(t *testing.T) {
	zap.L().Info("default global logger")
	zap.S().Info("default global sugared logger")

	example := zap.NewExample(zap.AddCaller())
	defer example.Sync()

	zap.ReplaceGlobals(example)
	zap.L().Info("replace global logger")
	zap.S().Info("replace global sugared logger")
}
