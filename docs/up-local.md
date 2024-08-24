# 環境構築

> [!WARNING]
> docker が PC に入ってること

## OpenAI apiKey を取得

[以下の記事を参考](https://nicecamera.kidsplates.jp/help/6648/)

- build フォルダーに.env.example を copy して.env という名前に変更する。
- OPEN-AI-API-KEY のパラメータに自身の apiKey を入れる

## 立ち上げ

#### image を作成

```bash
make build-local
```

#### コンテナを作成

```bash
make up-local
```

### スクレインプイングする URL を記述

[url.go を編集](../constants/url.go)

### 質問を記述

[main.go を編入](../cmd/main.go)
