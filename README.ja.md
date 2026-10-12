# 🚀 iLoveAntigravity

> **Google Antigravity & 自律型 AI コーディングエージェントのための実践的スーパーチャージャー。**  
> *ワンクリック導入、PC間即時移行、コンテキスト希釈ゼロ、Go 言語による高速ネイティブ MCP、ローカルベクトル記憶、決定論的エンジニアリング規範。*

---

### 🌐 多言語ドキュメント
[ English ](README.md) • [ Tiếng Việt ](README.vi.md) • [ 简体中文 ](README.zh.md) • [ 日本語 ](README.ja.md)

---

## ⚡ iLoveAntigravity とは？

**iLoveAntigravity** は、[Google Antigravity](https://github.com/google/antigravity) を単なる対話アシスタントから、**最高峰の自律型シニアソフトウェアエンジニア** へと変貌させるための総合最適化スイートです。

実戦開発から導き出された主な機能：
- **環境移行の完全自動化（1-Click Migration）**: 新しいマシンや環境への移行時、単一のコマンドを実行するだけで 60 秒以内に完全な最適化環境を復元。
- **Hermes プロトコル**: システムプロンプトを 2,500 文字未満に抑え、コンテキストウィンドウが 50% に達した段階でツール出力を能動的に要約・圧縮。SQLite FTS5 による 2ms 未満のローカル高速検索でトークン浪費と忘却を防止。
- **Go 言語コンパイル済みネイティブ MCP サーバー**: 重く遅い Node.js/Python サーバーを完全に排除し、Git、Mem0、SQLite、知識グラフを高速な Go ネイティブバイナリで提供。
- **Mem0 ローカル永続記憶（プロジェクト完全分離）**: ポート `8765` 上でバックグラウンド常駐。組み込み Qdrant データベースと **Google Antigravity Python SDK (`google.antigravity.Agent`)** を直接使用し、**VRAM 消費 0%** でプロジェクト間の記憶混同を完全防止。
- **Alibaba Open Code Review (OCR) 委任レビュー**: 確定的なファイル抽出とルール取得を活用し、行番号のズレを解消して安全かつ高精度なコードレビューと自動修正を実施。
- **ライフサイクルフック & 影響範囲（Blast Radius）ガード**: PreToolUse フックによる AST 依存関係解析で影響範囲の広いファイルを自動保護。PostToolUse フックで Language Server のメモリ肥大化を抑制。
- **高度なスキル & プラグイン群**: アンチ検出ステルスブラウザ自動化 (`invisible-playwright`)、自律型スキル自己合成 (`skill-synthesizer`)、TDD 駆動開発および Git Worktree 隔離検証。

---

## 🚀 クイックスタート（簡単セットアップ）

### Windows (PowerShell):
```powershell
# 1. リポジトリをクローン
git clone https://github.com/mirsevenabcg2-source/iloveantigravity.git
cd iloveantigravity

# 2. 自動インストーラーを実行
.\install.ps1
```

### Linux / macOS (Bash):
```bash
# 1. リポジトリをクローン
git clone https://github.com/mirsevenabcg2-source/iloveantigravity.git
cd iloveantigravity

# 2. 実行権限を付与して実行
chmod +x install.sh
./install.sh
```

---

## 🏛️ 主要アーキテクチャ

1. **Hermes プロトコル**: 最小限のプロンプト、50% コンテキスト能動圧縮、超高速ローカル検索。
2. **高速 Go ネイティブ MCP**: `mcp_git`, `mcp_mem0_bridge`, `mcp_memory`, `mcp_sqlite` による超低遅延動作。
3. **Mem0 ローカル記憶 & プロジェクト完全分離**: 厳格な `project="<ProjectName>"` 指定と 2 重プレフィックスによるコンテキスト汚染防止。
4. **Alibaba OCR 委任コードレビュー**: 決定論的差分抽出と Gemini 深層推論による高精度なレビュー・修正。
5. **AST 影響範囲ガード**: 結合度の高いコードへの安全弁と Language Server メモリトリミング。
6. **Git Worktree 委任**: 大規模リファクタリングを独立ワークツリーで安全に検証。

---

## 🙏 謝辞とクレジット (Credits & Acknowledgements)

**iLoveAntigravity** は、世界中の卓越したオープンソースプロジェクト、研究機関、開発者コミュニティの成果に基づき構築されています。以下のオリジナルチームおよびプロジェクトに深く感謝いたします：
- **Alibaba Group** ([open-code-review](https://github.com/alibaba/open-code-review)): 決定論的コードレビュー委任アーキテクチャおよび検査ルールセット。
- **Herrington Darkholme** ([ast-grep](https://github.com/ast-grep/ast-grep)): 構文構造に基づく AST コード検索・リファクタリング。
- **TheUpsider** ([@theupsider/lsp-mcp](https://github.com/theupsider/lsp-mcp)): Language Server Protocol と MCP の統合ブリッジ。
- **Microsoft Corporation** ([MarkItDown](https://github.com/microsoft/markitdown), [Playwright](https://github.com/microsoft/playwright)): 万能ドキュメント Markdown 変換およびブラウザ自動化基盤。
- **Mem0.ai & Qdrant チーム** ([mem0](https://github.com/mem0ai/mem0), [qdrant](https://github.com/qdrant/qdrant)): ベクトル記憶パラダイムおよびローカル組み込みベクトル DB。
- **Google DeepMind**: Antigravity SDK、Gemini 基盤モデル、AlphaFold/AlphaGenome システム。
- **Paul Gauthier** ([Aider](https://github.com/paul-gauthier/aider)): 構文検証ガードおよび開発者ワークフローの着想。
- **Nous Research**: Hermes プロトコルの制約記憶と自律的スキル自己合成哲学。
- **Vercel, Cloudflare, OpenSpec, gstack / Claude Code コミュニティ**: 実践的スキル群、高品位デザインシステムおよび設計仕様標準。

詳細な著作権情報、引用論文およびライセンス一覧は [CREDITS.md](CREDITS.md) をご覧ください。

---

## 📄 ライセンス

本プロジェクトは [MIT ライセンス](LICENSE) の下で公開されています。個人利用、商用利用、研究目的を問わず、世界中で自由にご利用いただけます。
