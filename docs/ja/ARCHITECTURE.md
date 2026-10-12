# 🏛️ iLoveAntigravity アーキテクチャ解説 (Japanese)

## 1. コア設計思想
従来の AI コーディングアシスタントにおける長時間のセッション劣化や忘却を防ぐため、**iLoveAntigravity** は確定的エンジニアリングによる解決策を提供します：
- **境界付きメモリ（Bounded Memory）**: 初期システムプロンプトをコンパクトに保ち、詳細知識はオンデマンドで検索。
- **超高速ローカル検索**: Go 言語で書かれたネイティブ MCP と組み込み Qdrant によるミリ秒単位の情報召喚。
- **決定論的コードレビュー**: Alibaba Open Code Review の決定論的差分抽出を活用し、行番号の幻覚を完全排除。

---

## 2. 主要コンポーネント一覧

| モジュール | 実装言語 | 実行メカニズム | 役割 |
| :--- | :--- | :--- | :--- |
| `mcp_git` | Go | スタンドアロンバイナリ | Git 状態、差分、ブランチ、履歴をゼロオーバーヘッドで取得 |
| `mcp_mem0_bridge` | Go | SSE クライアント | `127.0.0.1:8765` のローカル Mem0 デーモンと通信（自動起動・復旧機能付き） |
| `mcp_memory` | Go | スレッドセーフな JSON グラフ | プロジェクト内のエンティティと関係性グラフの永続管理 |
| `mcp_sqlite` | Go | 高速 SQLite ドライバ | データベース構造の自動検知、スキーマ検査、安全なクエリ実行 |
| `native_impact_guard` | Python | PreToolUse AST 解析 | ファイル変更前に影響範囲（Blast Radius）を解析し、危険な変更をブロック |
| `aider_guard` | Go | PostToolUse Win32 API | Language Server のメモリ肥大化抑制と構文即時検証 |
| `mem0_service` | Python | FastAPI / Starlette SSE | Google Antigravity SDK と Qdrant を連携したローカル永続記憶 |

---

## 3. プロジェクト混同防止（Anti-Dilution）原則
プロジェクト間での記憶の混濁を根絶するため、Mem0 へのアクセスは以下の規則を遵守します：
1. `project="<ProjectName>"`（全体共通は `project="Global"`）の明示的指定が必須。
2. 保存テキストには必ず `[<ProjectName>: <Category>] 詳細...` 形式のプレフィックスを付与。
