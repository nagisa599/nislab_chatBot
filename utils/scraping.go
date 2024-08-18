package utils

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// 指定されたURLからHTMLを取得し、divタグ内のテキストを整形して返す関数
func FetchAndProcessURL(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", fmt.Errorf("error parsing HTML: %w", err)
	}

	var divTexts []string
	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		text := strings.Join(strings.Fields(s.Text()), " ")
		if text != "" {
			divTexts = append(divTexts, text)
		}
	})

	consolidatedText := strings.Join(divTexts, " ")
	return consolidatedText, nil
}

// 複数のURLからテキストをフェッチして連結する関数
func FetchAndProcessMultipleURLs(urls []string) (string, error) {
	var wg sync.WaitGroup
	results := make([]string, len(urls))
	errors := make([]error, len(urls))

	for i, url := range urls {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			result, err := FetchAndProcessURL(url)
			if err != nil {
				errors[i] = err
				return
			}
			results[i] = result
		}(i, url)
	}

	wg.Wait()

	// エラーをチェックし、最初のエラーを報告
	for _, err := range errors {
		if err != nil {
			return "", err
		}
	}

	// 全ての結果を連結
	finalText := strings.Join(results, " ")
	return finalText, nil
}
