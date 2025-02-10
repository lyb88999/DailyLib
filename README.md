# Go每日一库

本项目用于收集Go语言中有用的第三方库使用示例，每日更新一个实用库。

---

## 📅 2025-01-26 - fsnotify

### 📖 库介绍
`fsnotify` 是一个跨平台的文件系统监控库，可以监控文件/目录的变化（创建、修改、删除、重命名等）。基于操作系统原生API实现，效率较高。

GitHub地址: https://github.com/fsnotify/fsnotify

### 📦 安装
```bash
go get github.com/fsnotify/fsnotify
```

## 📅 2025-01-27 - cron

### 📖 库介绍
`cron` 是一个功能强大的定时任务调度库，支持标准cron表达式（支持秒级精度），提供灵活的定时任务管理功能。

GitHub地址: https://github.com/robfig/cron

### 📦 安装
```bash
go get github.com/robfig/cron/v3
```

## 📅 2025-01-28 - carbon

### 📖 库介绍
`carbon` 是一个强大易用的时间处理库，提供链式调用、时区支持、时间计算等丰富功能，是替代标准库time的绝佳选择。

GitHub地址: https://github.com/dromara/carbon

### 📦 安装
```bash
go get github.com/dromara/carbon/v2
```


## 📅 2025-01-29 - zap

### 📖 库介绍
`zap` 是Uber开源的超高性能日志库，专为需要高性能日志记录的场景设计，提供结构化日志记录和灵活的配置选项。

GitHub地址: https://github.com/uber-go/zap

### 📦 安装
```bash
go get go.uber.org/zap
```


## 📅 2025-01-30～2025-01-31 - go-cmp

### 📖 库介绍
`go-cmp` 是Google提供的深度比较库，支持复杂类型的比较和差异定位，比标准库的`reflect.DeepEqual`更灵活安全。

GitHub地址: https://github.com/google/go-cmp

### 📦 安装
```bash
go get github.com/google/go-cmp/cmp
```


## 📅 2025-02-01 - go-zero/service

### 📖 库介绍
`go-zero/service` 是go-zero框架中的服务管理组件，提供优雅的服务生命周期管理，支持多服务统一启停和信号处理。

GitHub地址: https://github.com/zeromicro/go-zero

### 📦 安装
```bash
go get github.com/zeromicro/go-zero/core/service
```


## 📅 2025-02-02 - go-sqlmock

### 📖 库介绍
`go-sqlmock` 是SQL数据库模拟测试库，可以在不依赖真实数据库的情况下进行数据库操作测试，完美兼容标准库database/sql。

GitHub地址: https://github.com/DATA-DOG/go-sqlmock

### 📦 安装
```bash
go get github.com/DATA-DOG/go-sqlmock
```


## 📅 2025-02-03～2025-02-04 - zlsgo

### 📖 库介绍
`zlsgo` 是一个全栈Go开发框架，包含轻量高效的Web框架(znet)和功能强大的HTTP客户端(zhttp)等，提供开箱即用的开发体验。

GitHub地址: https://github.com/sohaha/zlsgo

### 📦 安装
```bash
go get github.com/sohaha/zlsgo
```


## 📅 2025-02-06 - cast

### 📖 库介绍
`cast` 是强大的类型转换库，提供安全便捷的类型转换方法，支持基本类型、集合类型和时间类型的相互转换，内置错误处理机制。

GitHub地址: https://github.com/spf13/cast

### 📦 安装
```bash
go get github.com/spf13/cast
```


## 📅 2025-02-07 - excelize

### 📖 库介绍
`excelize` 是功能强大的Excel文档操作库，支持读写XLSX/XLSM/XLTM文件、样式设置、公式计算、图表插入等完整Excel功能。

GitHub地址: https://github.com/xuri/excelize

### 📦 安装
```bash
go get github.com/xuri/excelize/v2
```



## 📅 2025-02-08 - sqlx

### 📖 库介绍
`sqlx` 是扩展标准库database/sql的增强版SQL工具包，提供更简洁的数据库操作API，支持结构体映射、批量操作和命名参数等高级特性。

GitHub地址: https://github.com/jmoiron/sqlx

### 📦 安装
```bash
go get github.com/jmoiron/sqlx

# 根据需要安装数据库驱动（以MySQL为例）
go get github.com/go-sql-driver/mysql
```

