/*******************************************************************************
 * Copyright 2019 Dell Inc.
 * Copyright (C) 2025 IOTech Ltd
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

/*
包 logger 提供一个将日志输出到 STDOUT 的日志客户端。
*/
package logger

// EdgeX Foundry 的 Go 实现的日志客户端

import (
	"fmt"
	stdLog "log"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

// 定义本地日志级别常量，避免外部依赖
const (
	TraceLog = "TRACE"
	DebugLog = "DEBUG"
	InfoLog  = "INFO"
	WarnLog  = "WARN"
	ErrorLog = "ERROR"
)

type edgeXLogger struct {
	logLevel string
	mu       sync.RWMutex // 保护 logLevel
}

// NewClient 创建一个将日志输出到 STDOUT 的 LoggingClient 实例。
// owningServiceName 用于与既有调用方保持签名兼容。
func NewClient(owningServiceName string, logLevel string) LoggingClient {
	upper := strings.ToUpper(logLevel)
	if !isValidLogLevel(upper) {
		upper = InfoLog
	}

	return &edgeXLogger{
		logLevel: upper,
	}
}

// logLevels 返回按从最到最不详细顺序的可能日志级别的数组
func logLevels() []string { // 不带图标,仅用于比较
	return []string{TraceLog, DebugLog, InfoLog, WarnLog, ErrorLog}
}

func isValidLogLevel(l string) bool {
	l = strings.ToUpper(l)
	for _, name := range logLevels() {
		if name == l {
			return true
		}
	}
	return false
}

var logLevelIconMap = map[string]string{
	TraceLog: "🟣",
	DebugLog: "🟦",
	InfoLog:  "🟩",
	WarnLog:  "🟨",
	ErrorLog: "🟥",
}

// level precedence for filtering
var levelOrder = map[string]int{
	TraceLog: 0,
	DebugLog: 1,
	InfoLog:  2,
	WarnLog:  3,
	ErrorLog: 4,
}

func (l *edgeXLogger) currentLevel() string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.logLevel
}

func (l *edgeXLogger) enabled(target string) bool {
	cur := l.currentLevel()
	return levelOrder[target] >= levelOrder[cur]
}

func caller(skip int) string {
	// 跳过若干层调用，获得文件:行号
	if _, file, line, ok := runtime.Caller(skip); ok {
		// 截断文件路径到最后两级
		parts := strings.Split(file, "/")
		if len(parts) > 2 {
			file = strings.Join(parts[len(parts)-2:], "/")
		}
		return fmt.Sprintf("%s:%d", file, line)
	}
	return "?? ?"
}

func (l *edgeXLogger) output(level string, formatted bool, msg string, args ...interface{}) {
	if !isValidLogLevel(level) { // 非法级别直接忽略
		return
	}
	if !l.enabled(level) { // 级别过滤
		return
	}

	// 固定宽度与布局常量
	const (
		levelWidth  = 5                               // TRACE/DEBUG/INFO/WARN/ERROR 最长5
		sourceWidth = 30                              // 可按需要调整，过长截断左侧
		timeLayout  = "2006-01-02 15:04:05.000000000" // 固定长度时间
	)

	icon := logLevelIconMap[level]
	ts := time.Now().Format(timeLayout)
	src := caller(4)
	// 截断 source 只保留末尾
	if len(src) > sourceWidth {
		src = src[len(src)-sourceWidth:]
	}

	renderedMsg := msg
	var extraKVs []string
	if formatted {
		renderedMsg = fmt.Sprintf(msg, args...)
	} else if len(args) > 0 {
		if len(args)%2 == 1 {
			args = append(args, "")
		}
		for i := 0; i < len(args); i += 2 {
			k := fmt.Sprintf("%v", args[i])
			v := fmt.Sprintf("%v", args[i+1])
			if k == "level" || k == "ts" || k == "source" || k == "msg" {
				k = "extra_" + k
			}
			v = strings.ReplaceAll(v, "\"", "'")
			extraKVs = append(extraKVs, fmt.Sprintf("%s=%s", k, v))
		}
	}

	// 构造对齐行：示例  🟩 [INFO ] [ts=2025-10-15 04:29:02.123456789] (source=negotiation/secretkey.go:192   ) msg="..."
	// level 方括号内固定宽度；source 括号内固定宽度左对齐填空格
	levelField := fmt.Sprintf("[%-*s]", levelWidth, level)
	tsField := fmt.Sprintf("[ts=%s]", ts)
	sourceField := fmt.Sprintf("(source=%-*s)", sourceWidth, src)
	// 替换消息中的双引号
	safeMsg := strings.ReplaceAll(renderedMsg, "\"", "'")
	line := fmt.Sprintf("%s %s %s %s msg=\"%s\"", icon, levelField, tsField, sourceField, safeMsg)
	if len(extraKVs) > 0 {
		line = line + " " + strings.Join(extraKVs, " ")
	}
	line += "\n"
	if _, err := fmt.Fprint(os.Stdout, line); err != nil {
		stdLog.Printf("logger write error: %v", err)
	}
}

// 兼容旧接口内部调用
func (lc *edgeXLogger) log(level string, formatted bool, msg string, args ...interface{}) {
	lc.output(level, formatted, msg, args...)
}

func (lc *edgeXLogger) SetLogLevel(logLevel string) error {
	upper := strings.ToUpper(logLevel)
	if !isValidLogLevel(upper) {
		return fmt.Errorf("invalid log level `%s`", logLevel)
	}
	lc.mu.Lock()
	lc.logLevel = upper
	lc.mu.Unlock()
	return nil
}

func (lc *edgeXLogger) LogLevel() string { return lc.currentLevel() }

func (lc *edgeXLogger) Info(msg string, args ...interface{})  { lc.log(InfoLog, false, msg, args...) }
func (lc *edgeXLogger) Trace(msg string, args ...interface{}) { lc.log(TraceLog, false, msg, args...) }
func (lc *edgeXLogger) Debug(msg string, args ...interface{}) { lc.log(DebugLog, false, msg, args...) }
func (lc *edgeXLogger) Warn(msg string, args ...interface{})  { lc.log(WarnLog, false, msg, args...) }
func (lc *edgeXLogger) Error(msg string, args ...interface{}) { lc.log(ErrorLog, false, msg, args...) }

func (lc *edgeXLogger) Infof(msg string, args ...interface{})  { lc.log(InfoLog, true, msg, args...) }
func (lc *edgeXLogger) Tracef(msg string, args ...interface{}) { lc.log(TraceLog, true, msg, args...) }
func (lc *edgeXLogger) Debugf(msg string, args ...interface{}) { lc.log(DebugLog, true, msg, args...) }
func (lc *edgeXLogger) Warnf(msg string, args ...interface{})  { lc.log(WarnLog, true, msg, args...) }
func (lc *edgeXLogger) Errorf(msg string, args ...interface{}) { lc.log(ErrorLog, true, msg, args...) }
