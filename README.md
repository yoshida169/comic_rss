# Comic RSS

React + TypeScript + Vite で構築されたコミック RSS リーダーです。

## 必要環境

- Node.js 18 以上
- npm

## セットアップ

```bash
# 依存パッケージのインストール
npm install
```

## 起動方法

### 開発サーバー

```bash
npm run dev
```

起動後、ブラウザで http://localhost:5173 を開いてください。ファイル変更時にホットリロードが有効です。

### 本番ビルド

```bash
npm run build
```

`dist/` ディレクトリに成果物が出力されます。

### ビルド結果のプレビュー

```bash
npm run preview
```

本番ビルドをローカルで確認できます。

## その他のコマンド

```bash
# ESLint による静的解析
npm run lint
```
