# Rag の実装

## Rag (Retrieval-Augment GEneration == 検索拡張性)とは

### 手順

1. 特定のデータベースに情報を検索
2. その検索結果を元に大規模言語モデル(LLM)にインプット
3. 質問を投げかける。

## 利用場面

- 社内に溜まっているクローズドなデータに対して質問を返してくれる chatbot を作りたい

## メリット

- ファインチューニングとは違い、再学習の手間がない
- model の変化、進化を考えなくていい。(ファインチューニングだと再学習する手間がかかる)

## これから　

手順１でも示したが、検索エンジンのパフォーマンスが重要である。
今回は cos 類似度を使って一番近いものを検索結果としてが、他にも

1. BM25
   BM25 は情報検索で広く使われているアルゴリズムで、文書とクエリ間の関連度を評価します。これは TF-IDF（Term Frequency-Inverse Document Frequency）に基づいたスコアリング方法であり、特に自然言語の文書に適しています。

2. Jaccard 係数
   Jaccard 係数は、セット間の類似度を測定する方法で、共通の要素の割合に基づいています。特にカテゴリデータやバイナリデータに適しており、シンプルで解釈が容易です。

3. ハミング距離
   ハミング距離は、二つの文字列の間の位置ごとの違いの数を計測します。主に固定長のデータに適用され、ビット文字列などで有用です。

4. ユークリッド距離
   ユークリッド距離は、ベクトル空間内の点間の「直線距離」を計測します。連続的な数値データに適しており、多次元データの類似性を計算するのに使われます。

5. ドット積
   ドット積は、ベクトルの類似度を測定するもう一つの方法です。特に、ベクトルが同じ方向を向いているかどうかを評価するのに適しています。これは cos 類似度と関連が深いが、ベクトルの長さも考慮します。

などがある。結果として何がどの場面でどの手段で使うべきかどうかがわからない。
# nislab_chatBot