package main

import (
	"context"
	"fmt"
	"os"

	"github.com/nagisa599/RAG-Study/utils"
	"github.com/sashabaranov/go-openai"
)



func main() {
	
	// データベースと仮定する
	document := []string{
		"僕の名前は那須渚です.",
		"私は東京都に住んでいます.",
		"私はプログラミングが好きです.",
	}
	// 質問と仮定する
	question := "あなたの名前は何ですか？"
    

	client := openai.NewClient(os.Getenv("OPENAIAPIKEY"))
	
	documentVector ,err := utils.GetEmbedding(client, document)
	if err != nil {
		fmt.Println("error")
	}
	questionVector, err := utils.GetEmbedding(client, []string{question})
	if err != nil {
		fmt.Println("error")
	}
	if len(questionVector) == 0 || len(documentVector) == 0 {
		fmt.Println("Error: Document vector or question vector is empty.")
		return
	}
	var maxSimilarity float64 = -1
	var mostSimilarDocIndex int = -1

	for i, vec := range documentVector {
		similarity, err := utils.CosSimilarity(vec, questionVector[0])
		if err != nil {
			fmt.Printf("Error calculating similarity between document %d and question: %v\n", i, err)
			continue
		}

		fmt.Printf("Cosine similarity between document %d and the question: %.4f\n", i, similarity)
		
		if similarity > maxSimilarity {
			maxSimilarity = similarity
			mostSimilarDocIndex = i
		}
	}

	if mostSimilarDocIndex != -1 {
		fmt.Printf("Document with highest cosine similarity: %d -> %.4f\n", mostSimilarDocIndex, maxSimilarity)
		fmt.Println("Most similar document:", document[mostSimilarDocIndex])
	} else {
		fmt.Println("No valid similarities found.")
	}
	
	prompt := fmt.Sprintf(`以下の質問に以下の情報をベースにして回答してください。
	[ユーザの情報]
	%s

	[情報]
	%s
	`, question, document[mostSimilarDocIndex])
	// prompt := fmt.Sprintf(`以下の質問に以下の情報をベースにして回答してください。
	// [ユーザの情報]
	// %s

	// [情報]
	// %s
	// `, question, document[0])


	gptChatResponse, err := client.CreateCompletion(context.Background(), openai.CompletionRequest{
		Model:     "gpt-3.5-turbo-instruct", // GPT-3.5-turbo-instructモデルを指定
		Prompt:    prompt,
		MaxTokens: 200, // 応答の最大トークン数
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// レスポンス出力
	fmt.Println("Response:", gptChatResponse.Choices[0].Text)
}