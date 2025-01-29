package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {
	// SimpleLogger()
	// ComplexLog()
	// CustomConf()
	// Sample()
}

func SimpleLogger() {
	// 新建一个生产环境下的logger
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			return
		}
	}(logger)
	logger.Info("程序启动了",
		zap.String("version", "1.0"),
		zap.Int("port", 8080),
	)
}

func ComplexLog() {
	logger, _ := zap.NewDevelopment()
	defer func(logger *zap.Logger) {
		if err := logger.Sync(); err != nil {
			return
		}
	}(logger)
	logger.Debug("Debug level log")
	logger.Info("程序启动",
		zap.String("version", "1.0"),
		zap.Int("port", 8080),
	)
	logger.Warn("警告信息",
		zap.String("warning", "about to hit rate limit"))
}

func CustomConf() {
	config := zap.NewDevelopmentConfig()
	config.Encoding = "console"
	config.Level.SetLevel(zap.InfoLevel) // 设置最低日志级别
	logger, _ := config.Build()
	defer func(logger *zap.Logger) {
		if err := logger.Sync(); err != nil {
			return
		}
	}(logger)
	logger.Info("带着配置跑的日志")
}

func Sample() {
	samplingConfig := zap.SamplingConfig{
		Initial:    50, // 初始阶段：每个时间窗口（此处为 1秒）内，前 50 条日志会被完整记录
		Thereafter: 5,  // 当超过 50 条后，每 5 条日志记录 1 条（即采样率为 20%）。
	}
	samplingOption := zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewSamplerWithOptions(core, time.Second, samplingConfig.Initial, samplingConfig.Thereafter)
	})
	logger, _ := zap.NewProduction(samplingOption)
	defer func(logger *zap.Logger) {
		if err := logger.Sync(); err != nil {
			return
		}
	}(logger)
	// 生成大量日志条目以观察采样效果
	for i := 0; i < 1000; i++ {
		logger.Info("程序启动了",
			zap.String("version", "1.0"),
			zap.Int("port", 8080),
			zap.Int("iteration", i),
		)
		time.Sleep(5 * time.Millisecond) // 控制日志生成速度
	}
}
