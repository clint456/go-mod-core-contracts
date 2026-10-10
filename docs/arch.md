# 架构与目录结构说明

本模块为 EdgeX Foundry 的**共享契约模块**（`github.com/edgexfoundry/go-mod-core-contracts/v4`，Go 1.25）。
它为 EdgeX 各微服务之间通过 HTTP / 消息总线传递的数据提供**统一的数据模型、DTO 与客户端**。
模块本身不实现业务逻辑，只定义"契约"：数据结构、校验规则、路由常量、错误类型和调用客户端。

## 依赖方向

```
dtos    ──► common, errors
models  ──► common, errors
clients/interfaces ──► dtos, errors
clients/http       ──► interfaces, dtos, common, errors, clients/http/utils
```

`common` / `errors` 位于最底层，`dtos` / `models` 依赖它们，`clients` 位于最上层并依赖 DTO 与错误类型。

## 顶层目录

```
go-mod-core-contracts/
├── common/     # 全局常量与校验器（路由、字段名、通用工具、validator）
├── models/     # 领域模型（服务内部持久化/业务实体）
├── dtos/       # 数据传输对象（对外 API 的请求/响应结构）
├── clients/    # 各服务的 REST 客户端（接口 + HTTP 实现 + mock）
├── errors/     # EdgeX 统一错误类型（EdgeX error kinds）
└── docs/       # 文档目录
```

## 各目录职责

### `common/`
跨模块共享的基础设施。

- [constants.go](../common/constants.go)：v3 API 的全部路由常量（`ApiBase = "/api/v3"`、`ApiEventRoute`、`ApiReadingRoute` 等）与字段名常量。
- [validator.go](../common/validator.go)：基于 `go-playground/validator` 的校验封装。
- [utils.go](../common/utils.go)：通用工具函数。
- `fake_validator.go`：供测试使用的校验器替身。

### `models/`
面向服务内部使用的领域模型。

- 实体：`node.go`、`event.go`、`reading.go`、`profile.go`、`subscription.go`、`transmission.go` 等。
- 辅助：`consts.go`（枚举/取值常量）、`dbtimestamp.go`（DB 时间戳）、`log_entry.go`（日志条目）。
- 部分模型（如 Profile）同时支持 YAML 编码，并附带业务工具函数（例如 [profile.go](../models/profile.go) 中的 `BuildCoreCommands`）。

### `dtos/`
数据传输对象，定义对外上线协议，按 API 语义分四层：

- 根目录 `dtos/*.go`：资源实体 DTO（`node.go`、`event.go`、`reading.go`、`profile.go`、`notification.go` 等），以及 `metric.go`、`systemevent.go`、`extensions.go` 等特殊结构。
- `dtos/common/`：所有请求/响应共用的基础结构 —— `BaseRequest`、`BaseResponse`、`BaseWithIdResponse`、`Versionable`，以及 `config.go`、`ping.go`、`version.go`、`count.go`、`secret.go`。
- `dtos/requests/`：入参请求体（`node.go`、`event.go`、`subscription.go` 等）。
- `dtos/responses/`：出参响应体（`node.go`、`event.go`、`reading.go`、`corecommand.go`、`uom.go` 等）。

`models` 与 `dtos` 分离的意义在于：前者供服务内部使用，后者定义对外线上协议，二者解耦、可独立演进。

### `clients/`
服务客户端，采用"接口 + 实现 + mock"三段式。

- [interfaces/](../clients/interfaces)：各服务客户端接口（`NodeClient`、`EventClient`、`CommonClient`、`AuthenticationInjector` 等），业务代码应面向接口编程。
- [interfaces/mocks/](../clients/interfaces/mocks)：由 mockery 生成的接口 mock，供测试注入。
- [http/](../clients/http)：接口的 HTTP 实现（`node.go`、`event.go`、`notification.go`、`secret.go`、`auth.go` 等）。
- [http/utils/](../clients/http/utils)：底层 HTTP 请求封装（`GetRequest` / `PostRequest` 等），统一处理鉴权注入与错误包装。
- [logger/](../clients/logger)：日志客户端（基于 `go-kit/log`），含 mock。
- [common.go](../clients/common.go)：`ClientBaseUrlFunc` 等基线 URL 辅助函数。

### `errors/`
- [types.go](../errors/types.go)：统一错误类型 `EdgeX` 与错误种类（`KindNotFound`、`KindContractInvalid`、`KindDuplicateName`、`KindUnauthorized` 等），并携带 HTTP 状态码语义。

## 依赖与工具链

- 关键依赖：`go-playground/validator`（校验）、`go-kit/log`（日志）、`fxamacker/cbor`（CBOR 编码，用于 SystemEvent）、`google/uuid`、`yaml.v3`。
- [Makefile](../Makefile)：
  - `make test`：单测 + 覆盖率 + lint + `go vet` + `gofmt` 检查。
  - `make lint`：运行 golangci-lint。
  - `make unittest`：仅运行单元测试。
