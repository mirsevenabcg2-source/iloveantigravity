# 🏛️ iLoveAntigravity 架构全景文档 (Simplified Chinese)

## 1. 核心设计哲学
传统 AI 编码助手在持续多轮会话中经常出现“记忆衰退”与“幻觉偏差”。**iLoveAntigravity** 通过工程化手段彻底解决该痛点：
- **约束内存（Bounded Memory）**：严格控制初始化 Prompt 体积，长尾事实动态外置。
- **本地极速召回**：基于 Go 编写的高性能 MCP 服务与嵌入式 Qdrant 向量库，提供毫秒级事实召回。
- **确定性工程 Review**：结合阿里 Open Code Review 的确定性过滤能力，消除大模型行号幻觉。

---

## 2. 核心组件总览

| 模块 | 实现语言 | 运行机制 | 核心职责 |
| :--- | :--- | :--- | :--- |
| `mcp_git` | Go | 纯二进制命令行适配器 | 极速获取 Git 状态、暂存/未暂存 Diff、分支与提交历史 |
| `mcp_mem0_bridge` | Go | SSE 长连接客户端 | 连接 `127.0.0.1:8765` 本地 Mem0 服务，具备服务守护与自愈能力 |
| `mcp_memory` | Go | 内存 JSON 线程安全读写 | 项目内实体与关系图谱持久化管理 |
| `mcp_sqlite` | Go | Modern Go 纯纯 SQL 驱动 | 自动扫描项目 SQLite 数据库，执行只读检查与架构自省 |
| `native_impact_guard` | Python | PreToolUse AST 语法分析 | 在 Agent 修改核心代码前，分析依赖爆炸半径并执行硬拦截 |
| `aider_guard` | Go | PostToolUse Win32 API 裁剪 | 自动清理 Language Server 泄漏内存并拦截非标 JSON/Python 代码 |
| `mem0_service` | Python | FastAPI / Starlette SSE 后台进程 | 集成 Google Antigravity SDK 与 Qdrant，提供永久本地项目记忆 |

---

## 3. 多项目防稀释（Anti-Dilution）规则
为防止不同仓库间的知识产生混淆，所有对 Mem0 的调用必须严格遵守双层隔离：
1. 强制显式传入 `project="<ProjectName>"` 参数（全局偏好传 `project="Global"`）。
2. 记忆文本统一前缀格式：`[<ProjectName>: <Category>] 详细内容...`。
