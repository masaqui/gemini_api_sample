# gemini_api_sample

## Setup

### ステップ1: APIキーの取得 🔑
まず、Gemini APIを利用するためのAPIキーが必要です。これは、あなたが正当な利用者であることを証明するためのものです。
[Google AI Studio](https://aistudio.google.com/) にアクセスし、Googleアカウントでログインします。
利用規約に同意し、続行します。
画面の左側にあるメニューから 「Get API key」 をクリックします。
「Create API key in new project」 をクリックすると、新しいAPIキーが生成されます。

### ステップ2: .envファイル作成
dotenvファイルをコピって.envファイルを作成してください
GEMINI_API_KEYにAPIキーを設定してください
MODELを指定してください

## 使い方
```
go run gemini_api_sample.go "日本で2番目に高い山はどこですか"

AIへの質問: 「日本で2番目に高い山はどこですか」
AIにメッセージを送信中...

--- AIからの返事 ---
日本の2番目に高い山は北岳（きただけ）です。

--------------------
```
