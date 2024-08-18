package utils


func ChunkText(text string, chunkSize, overlap int) []string {
	var chunks []string
	runes := []rune(text) // マルチバイト文字を正しく扱うためにruneスライスに変換
	length := len(runes)

	for i := 0; i < length; i += chunkSize - overlap {
		end := i + chunkSize
		if end > length {
			end = length
		}
		chunks = append(chunks, string(runes[i:end]))
	}

	return chunks
}