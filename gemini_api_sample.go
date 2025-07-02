package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv" // .env ファイルを読み込む
	"google.golang.org/api/option"
)

/*
 * go run gemini_api_sample.go "あなたの好きな食べ物は何ですか？"
 */
func main() {
	// .env ファイルから環境変数を読み込みます
	err := godotenv.Load()
	if err != nil {
		// .env ファイルがなくてもエラーにせず、次の処理に進みます
		// （CI/CD環境など、直接環境変数が設定される場合を考慮）
		log.Println("ヒント: .envファイルが見つかりません。")
	}

	// コマンドラインからプロンプトを取得します
	// os.Args[0] はプログラム名なので、引数は os.Args[1] から始まります
	if len(os.Args) < 2 {
		fmt.Println("エラー: プロンプトをコマンドライン引数として指定してください。")
		fmt.Println("例: go run gemini_api_sample.go \"日本の首都はどこですか？\"")
		return // 引数がない場合はプログラムを終了
	}
	// 最初の引数をプロンプトとして使用します
	promptText := os.Args[1]

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("エラー: APIキーが設定されていません。.envファイルに GEMINI_API_KEY を設定してください。")
	}
	modelName := os.Getenv("MODEL")
	if modelName == "" {
		log.Fatal("エラー: AIモデルが設定されていません。.envファイルに MODEL を設定してください。")
	}

	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel(modelName)

	// コマンドライン引数から受け取ったテキストをプロンプトとして設定
	prompt := genai.Text(promptText)

	fmt.Printf("AIへの質問: 「%s」\n", promptText)
	fmt.Println("AIにメッセージを送信中...")

	resp, err := model.GenerateContent(ctx, prompt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n--- AIからの返事 ---")
	printResponse(resp)
	fmt.Println("--------------------")
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
}
