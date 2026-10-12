# 🚀 iLoveAntigravity

> **专为 Google Antigravity 与自主 AI 编码智能体打造的终极实战级增强工具箱。**  
> *开箱即用、一键跨设备迁移、零上下文稀释、极速原生 Go 编译 MCP 服务、本地向量记忆与确定性工程规范。*

---

### 🌐 多语言文档导航
[ English ](README.md) • [ Tiếng Việt ](README.vi.md) • [ 简体中文 ](README.zh.md) • [ 日本語 ](README.ja.md)

---

## ⚡ 什么是 iLoveAntigravity？

**iLoveAntigravity** 是基于数千小时高强度实战开发经验打造的综合型优化套件，旨在将 [Google Antigravity](https://github.com/google/antigravity) 从常规对话智能体彻底进化为顶级**自主全栈软件工程师**。

核心能力包括：
- **跨机器一键无缝迁移**：更换电脑或重装系统时，仅需运行单行安装命令，60 秒内完整还原整个高效调优开发环境。
- **Hermes 协议（Bounded Memory & Context Compaction）**：严格控制系统提示词大小（<2,500 字符），当上下文窗口达到 50% 时主动紧凑压缩巨量工具输出，结合 SQLite FTS5 实现亚毫秒级（<2ms）本地精准召回，彻底杜绝上下文稀释。
- **原生 Go 语言编译 MCP 核心服务**：淘汰笨重缓慢的 Node.js/Python 进程，采用 Go 编译的纯机器码提供极速 Git、Mem0 桥接、SQLite 和知识图谱 MCP 服务。
- **Mem0 本地向量记忆（零稀释多项目隔离）**：常驻后台运行于 `8765` 端口，采用内置嵌入式 Qdrant 数据库与 768 维向量，直接调用 **Google Antigravity Python SDK (`google.antigravity.Agent`)** 进行轻量推理，**0% 占用显存**，多项目双重隔离机制确保记忆绝不混淆。
- **阿里 Open Code Review (OCR) 委托代码审查体系**：结合阿里 OCR 5 步委托机制过滤无用文件与获取规则集，由智能体执行逐行深度推理，彻底解决行号偏差并免去外部 API 成本。
- **原生生命周期钩子与代码爆炸半径（Blast Radius）防御**：PreToolUse AST 语法分析自动拦截影响范围超过阈值的危险改动；PostToolUse 自动清理语言服务器臃肿内存并校验语法。
- **生产级精英技能与插件库**：内置隐身反指纹浏览器自动化 (`invisible-playwright`)、自主技能合成 (`skill-synthesizer`)、闭环 TDD 测试驱动与 Git Worktree 隔离机制。

---

## 🚀 极速安装与跨设备迁移（1-Click Setup）

### Windows 系统 (PowerShell):
```powershell
# 1. 克隆代码库
git clone https://github.com/mirsevenabcg2-source/iloveantigravity.git
cd iloveantigravity

# 2. 运行一键配置脚本
.\install.ps1
```

### Linux / macOS 系统 (Bash):
```bash
# 1. 克隆代码库
git clone https://github.com/mirsevenabcg2-source/iloveantigravity.git
cd iloveantigravity

# 2. 赋予执行权限并运行
chmod +x install.sh
./install.sh
```

---

## 🏛️ 核心架构要点

1. **Hermes 协议**：精简提示词，50% 上下文自动压缩，解决模型遗忘与幻觉。
2. **Go 语言原生 MCP 服务**：`mcp_git`、`mcp_mem0_bridge`、`mcp_memory`、`mcp_sqlite`，微秒级启动，极低内存消耗。
3. **Mem0 本地记忆与零稀释隔离**：强制 `project="<ProjectName>"` 属性与 `[Project: Category]` 双层标记，永不混淆项目背景。
4. **阿里 OCR 委托代码审查**：确定性工程提取 Diff 与规则，模型精准定界与自动修复。
5. **AST 影响范围保护与生命周期钩子**：拦截高耦合文件修改，监控并裁剪 Language Server 内存泄漏。
6. **Git Worktree 隔离机制**：大型重构与高风险迁移自动在独立工作树分支运行，测试 100% 通过后方可合并。

---

## 🙏 致谢与来源 (Credits & Acknowledgements)

**iLoveAntigravity** 融合了全球开源社区与前沿研究成果的智慧结晶。我们由衷感谢并致敬以下项目与原创团队：
- **阿里巴巴 (Alibaba Group)** ([open-code-review](https://github.com/alibaba/open-code-review))：确定性代码审查委托机制与完备规则库。
- **Herrington Darkholme** ([ast-grep](https://github.com/ast-grep/ast-grep))：基于语法树的结构化 AST 重构与扫描。
- **TheUpsider** ([@theupsider/lsp-mcp](https://github.com/theupsider/lsp-mcp))：Language Server Protocol 与 MCP 协议深度桥接。
- **微软 (Microsoft Corporation)** ([MarkItDown](https://github.com/microsoft/markitdown), [Playwright](https://github.com/microsoft/playwright))：通用文档转 Markdown 与自动化测试套件。
- **Mem0.ai 与 Qdrant 团队** ([mem0](https://github.com/mem0ai/mem0), [qdrant](https://github.com/qdrant/qdrant))：向量记忆范式与嵌入式高性能向量数据库。
- **Google DeepMind**：Antigravity SDK、Gemini 基础大模型及 AlphaFold/AlphaGenome 系统。
- **Paul Gauthier** ([Aider](https://github.com/paul-gauthier/aider))：代码语法树防御与开发工作流灵感。
- **Nous Research**：Hermes 协议受限记忆与自主技能反思哲学。
- **Vercel、Cloudflare、OpenSpec 以及 gstack / Claude Code 社区**：高质量开发技能、设计规范及工程流程。

完整致谢清单、论文引用及上游许可证见 [CREDITS.md](CREDITS.md)。

---

## 📄 开源许可证

本项目基于 [MIT 许可证](LICENSE) 发布。全球范围内免费供个人、商业及科研使用。
