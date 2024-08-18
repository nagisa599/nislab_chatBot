package utils

import (
	"context"
	"fmt"
	"math"

	"github.com/sashabaranov/go-openai"
)

// getEmbedding 関数は指定されたテキストのリストに対して埋め込みを生成し、その埋め込みのリストを返します。
func GetEmbedding(client *openai.Client, inputTexts []string) ([][]float64, error) {
	resp, err := client.CreateEmbeddings(context.Background(), openai.EmbeddingRequest{
		Model:  "text-embedding-ada-002", // モデルを指定
		Input: inputTexts,               // 埋め込みを生成したいテキストのリストを入力
	})

	if err != nil {
		return nil, fmt.Errorf("CreateEmbedding error: %v", err)
	}

	// 生成された埋め込みを抽出して返す
	if len(resp.Data) > 0 {
		embeddings := make([][]float64, len(resp.Data))
		for i, data := range resp.Data {
			embedding := make([]float64, len(data.Embedding))
			for j, val := range data.Embedding {
				embedding[j] = float64(val) // float32からfloat64に変換
			}
			embeddings[i] = embedding
		}
		return embeddings, nil
	}

	return nil, fmt.Errorf("no embedding data returned")
}
// cosSimilarity 関数は二つのベクトル間のコサイン類似度を計算します。
func CosSimilarity(vecA, vecB []float64) (float64, error) {
	if len(vecA) != len(vecB) {
		return 0, fmt.Errorf("vectors must be of the same length")
	}

	var dotProduct float64
	var normA, normB float64

	for i, v := range vecA {
		dotProduct += v * vecB[i]
		normA += v * v
		normB += vecB[i] * vecB[i]
	}

	if normA == 0 || normB == 0 {
		return 0, fmt.Errorf("norm of vector cannot be zero")
	}

	cosine := dotProduct / (math.Sqrt(normA) * math.Sqrt(normB))
	return cosine, nil
}
