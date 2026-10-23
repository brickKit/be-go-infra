# be/go-infra

## 组件定位

基础外壳：一个 Go 进程承载编进来的成员，每个成员有自己的端口与运行时（be-sdk-go 的外壳启动器）；除了进程，什么都不合并。这次承载哪些成员由部署决定（BrickKit 的 `BRICKKIT_SERVED_MEMBERS_CONFIG`）；编进来的成员或版本与要求不符时外壳拒绝启动并点名成员。

## 部署前准备

- 外壳登录角色 `shell_go_infra`（`PG_SHELL_USER`、`PG_SHELL_PASSWORD`）：持有每个被承载成员的运行角色（`WITH INHERIT FALSE`），每个事务里切换；承载的成员带数据库时必需。
- 成员的迁移用成员自己的镜像执行，外壳不迁移。
- 成员的其余准备（schema、消息系统、证书……）见各成员的 `BRICKKIT.md`。

## 依赖说明

不声明组件依赖；编进来的成员见“外壳声明”。

## 配置指南

外壳自己的键见 `component.yaml` 的 `configSchema`（都可选）；成员的配置在各成员自己的配置里，由 BrickKit 经 `BRICKKIT_SERVED_MEMBERS_CONFIG` 交给外壳。进程级的键（消息系统、身份与授权地址……）在同一外壳的成员之间必须一致。

## 契约索引

外壳自己的端口只有平台端点：`/healthz`、`/readyz`（全部承载成员就绪才就绪）、`/.well-known/be-component`（带成员）、`/metrics`（全部成员，带 `component` 标签）。成员的接口见各成员。

## 外壳声明

编进来的成员：`test/probe-go@3.2.19`、`test/probe-go2@3.2.19`、`test/probe-go3@3.2.19`、`test/probe-go4@3.2.19`；第 5 步的安全成员 `infra/iam-casdoor@3.0.4`、`infra/authz@3.0.5`、`infra/audit@3.0.3`；第 6 步的主数据 `mdm/org@3.0.4`、`mdm/currency@3.0.5`、`mdm/partner@3.0.5`、`mdm/product@3.0.5`。
