# 自身の所属する研究室の chatbot を作成

## 利用技術

| 技術                 | 利用しているライブラリやツール等               |
| -------------------- | ---------------------------------------------- |
| 言語                 | go1.22.45                                      |
| ローカル環境構築     | Docker                                         |
| スクレイピング       | goquery                                        |
| Open AI 外部 package | sashabaranov                                   |
| Open AI モデル       | text-embedding-ada-002, gpt-3.5-turbo-instruct |

## 利用 open AI モデル

### gpt-3.5-turbo-instruct

- 特にインストラクションに基づいて最適化された応答を生成するように設計されている。このモデルは、特定の指示に基づいた応答生成に特化しており、ユーザーからの明確な要求に対して効率的かつ正確に答えることができる

### text-embedding-ada-002

- テキストデータを固定長のベクトル表現に変換するために設計されている。このエンベディングモデルは、さまざまな自然言語処理タスクでの使用を目的としており、テキストの意味的な内容を数値的な形式で捉えることができる。

### 検索エンジン

研究室の web をスクレイピングして、文字列を羅列して 400 文字ごとに区切る(チャンクを作成)
質問とそれぞれのチャンクに対して、cosin 類似度を計算。cosin 類似度が高い 2 つのチャンクを取り出す

## index

- [環境構築](./docs/up-local.md)
- [実結果]()
