# go-mod-core-contracts
[![Build Status](https://jenkins.edgexfoundry.org/view/EdgeX%20Foundry%20Project/job/edgexfoundry/job/go-mod-core-contracts/job/main/badge/icon)](https://jenkins.edgexfoundry.org/view/EdgeX%20Foundry%20Project/job/edgexfoundry/job/go-mod-core-contracts/job/main/) [![Code Coverage](https://codecov.io/gh/edgexfoundry/go-mod-core-contracts/branch/main/graph/badge.svg?token=s4Y4L22Bs0)](https://codecov.io/gh/edgexfoundry/go-mod-core-contracts) [![Go Report Card](https://goreportcard.com/badge/github.com/edgexfoundry/go-mod-core-contracts)](https://goreportcard.com/report/github.com/edgexfoundry/go-mod-core-contracts) [![GitHub Latest Dev Tag)](https://img.shields.io/github/v/tag/edgexfoundry/go-mod-core-contracts?include_prereleases&sort=semver&label=latest-dev)](https://github.com/edgexfoundry/go-mod-core-contracts/tags) ![GitHub Latest Stable Tag)](https://img.shields.io/github/v/tag/edgexfoundry/go-mod-core-contracts?sort=semver&label=latest-stable) [![GitHub License](https://img.shields.io/github/license/edgexfoundry/go-mod-core-contracts)](https://choosealicense.com/licenses/apache-2.0/) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/edgexfoundry/go-mod-core-contracts) [![GitHub Pull Requests](https://img.shields.io/github/issues-pr-raw/edgexfoundry/go-mod-core-contracts)](https://github.com/edgexfoundry/go-mod-core-contracts/pulls) [![GitHub Contributors](https://img.shields.io/github/contributors/edgexfoundry/go-mod-core-contracts)](https://github.com/edgexfoundry/go-mod-core-contracts/contributors) [![GitHub Committers](https://img.shields.io/badge/team-committers-green)](https://github.com/orgs/edgexfoundry/teams/go-mod-core-contracts-committers/members) [![GitHub Commit Activity](https://img.shields.io/github/commit-activity/m/edgexfoundry/go-mod-core-contracts)](https://github.com/edgexfoundry/go-mod-core-contracts/commits)

## 简介

本模块提供 EDGE-HY 各服务之间以 Request/Response 形式传递数据时所使用的**契约模型（contract models）**，
同时包含 [edgex-go](http://192.168.0.89:3000/edgex-go) 仓库中各个服务对应的**服务客户端**。
各类模型与客户端的定义可分别在对应的顶层目录中查看，详见 [docs/arch.md](./docs/arch.md)。


## 目录结构

```
go-mod-core-contracts/
├── common/     # 全局常量（路由、字段名）与校验器、通用工具
├── models/     # 领域模型（服务内部使用）
├── dtos/       # 数据传输对象（对外 API 的请求/响应结构）
│   ├── common/     # 请求/响应共用基类（BaseRequest、BaseResponse 等）
│   ├── requests/   # 入参请求体
│   └── responses/  # 出参响应体
├── clients/    # 各服务客户端（接口 + HTTP 实现 + mock + logger）
├── errors/     # EdgeX 统一错误类型与错误种类
└── docs/       # 文档
```

## 安装

* 确保使用 Go 1.25.0 或更高版本，并已启用 modules（即已初始化 `go.mod` 文件）
* 若代码位于 GOPATH 中，请确保已设置 `GO111MODULE=on`
* 执行 `go get github.com/clint456/go-mod-core-contracts`
    * 该命令会将 go-mod-core-contracts 加入 go.mod 文件，并下载到模块缓存中

## 改动说明

本分支对契约命名做了一次**全量重命名**，将原先以 `node` 为核心的命名统一调整为更中性的 `node` / `profile` / `resource` / `command` 体系。

### 命名映射

| 原命名 | 新命名 | 说明 |
|---|---|---|
| `Node` | `Node` | 设备实体改为节点 |
| `Profile` | `Profile` | 设备配置文件改为配置 |
| `Resource` | `Resource` | 设备资源改为资源 |
| `Command` | `Command` | 设备命令改为命令 |
| `NodeService` | `NodeService` | 设备服务改为节点服务 |
| `DiscoveredNode` | `DiscoveredNode` | 发现的设备改为发现的节点 |
| `NODECONTROL` | `NODECONTROL` | ScheduleAction 类型枚举值 |

### 改动范围

- **Go 标识符**：类型、字段、函数、常量、方法名均按上表替换，且按"最长匹配优先"处理，例如 `ProfileBasicInfo` → `ProfileBasicInfo`、`MultiProfilesResponse` → `MultiProfilesResponse`、`NodeCoreCommand` → `NodeCoreCommand`。
- **文件名**：相关源文件同步重命名，例如 `models/node.go` → `models/node.go`、`dtos/profile.go` → `dtos/profile.go`、`dtos/command.go` → `dtos/command.go`。
- **常量与路由**：`common/constants.go` 中的路由常量（如 `ApiNodeRoute` → `ApiNodeRoute`）与字段名常量同步更新。
- **常量合并**：原先的 `common.Command = "deHTTP 路由：/api/v3/node → /api/v3/node、/api/v3/profile → /api/v3/profile、/api/v3/resource → /api/v3/resource、/api/v3/nodeservice → /api/v3/nodeservice。
消息总线 topic：events/node/# → events/node/#、node/command/request → node/command/request、validate/node → validate/node 等。
系统事件：系统事件类型 "node" → "node"、"profile" → "profile"、"nodeservice" → "nodeservice"。
JSON/YAML 负载字段：nodeName、resource、command、discoveredNode 等字段名均已改变。
viceCommand"` 与既有的 `common.Command = "command"` 重命名后取值相同，已删除冗余定义并将引用统一到 `common.Command`。
- **协议字段**：JSON/YAML tag 与字符串字面量一并调整，例如 `nodeName` → `nodeName`、`resource` → `resource`、`resources` → `resources`、`commands` → `commands`、`command` → `command`、`discoveredNode` → `discoveredNode`。

### 破坏性变更（重要）

本次改动会**破坏与既有 EdgeX 服务之间的线上协议兼容性**，使用方需同步升级：

- **HTTP 路由**：`/api/v3/node` → `/api/v3/node`、`/api/v3/profile` → `/api/v3/profile`、`/api/v3/resource` → `/api/v3/resource`、`/api/v3/nodeservice` → `/api/v3/nodeservice`。
- **消息总线 topic**：`events/node/#` → `events/node/#`、`node/command/request` → `node/command/request`、`validate/node` → `validate/node` 等。
- **系统事件**：系统事件类型 `"node"` → `"node"`、`"profile"` → `"profile"`、`"nodeservice"` → `"nodeservice"`。
- **JSON/YAML 负载字段**：`nodeName`、`resource`、`command`、`discoveredNode` 等字段名均已改变。

### 验证

```bash
gofmt -l .        # 应无输出
go vet ./...      # 通过
go build ./...    # 通过
go test ./...     # 全部通过
```
