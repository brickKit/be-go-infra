# be/go-infra

基础外壳（Go）：一个进程承载编进来的成员。

## 在项目里使用

在装配项目的 `brickkit.yaml` 里加安装源，`brickkit add be/go-infra@<版本>`；被承载的成员也要在项目里（`shell.members` 的精确版本）。部署文件里把成员放在外壳条目下，BrickKit 生成每个成员的配置。

## 文档

| 问题 | 看哪里 |
| --- | --- |
| 它是什么、承载哪些成员、怎样部署 | `BRICKKIT.md` |
| 怎样改它 | `AGENTS.md` |

## 开发

见 `AGENTS.md` 的“构建与测试”。
