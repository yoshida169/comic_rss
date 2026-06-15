# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
npm run dev      # 開発サーバー起動 (http://localhost:5173)
npm run build    # プロダクションビルド
npm run lint     # ESLint実行
npm run preview  # ビルド成果物のプレビューサーバー起動
npx tsc --noEmit # 型チェックのみ実行（ビルドなし）
```

## Architecture

**React 19 + TypeScript + Vite** のシングルページアプリケーション。

エントリーポイントは `src/main.tsx` → `src/App.tsx` → `src/components/Dashboard.tsx` の順にレンダリングされる。

- `src/App.tsx` — ルートコンポーネント。現在は `Dashboard` を直接レンダリングするシンプルな構成。
- `src/components/Dashboard.tsx` — メインビュー。今後コミックRSS機能を追加していく画面。
- `src/components/Message.tsx` — メッセージ入力UIコンポーネント（ローカルstateで管理）。
- `src/App.css` / `src/index.css` — スタイル定義。

## TypeScript設定

`tsconfig.json` で `strict: true`、`noUnusedLocals: true`、`noUnusedParameters: true` が有効。未使用の変数・パラメータはコンパイルエラーになる。

ESLintは `.js`/`.jsx` ファイルのみ対象（`.ts`/`.tsx` はTypeScriptコンパイラが型チェックを担う）。
