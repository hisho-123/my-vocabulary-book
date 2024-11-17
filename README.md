# my-vocabulary-book

## 目的
- 業務上使用している技術について理解を深め、業務効率を上げる
- そもそもの開発業務（プログラミングや git操作）に慣れる

## 概要

### アプリ名
my vocabulary book (暫定)

### 説明
- 自分だけの単語(熟語)帳を作成し、復習できるアプリ
- よくある忘却曲線だか何だかを計算して記憶定着の良い復習頻度で復習させる

### 画面遷移図

![](./画面遷移図.drawio)

## 技術
|           | 使用技術                 |
| ---       | ---                     |
| frontend  | Vite + Vue + TypeScript |
| backend   | go + mySQL              |

## 構造

- フロントエンドはデータの受け渡し・表示のみ
- データのバリデーション・復習周期の計算など、

```
├── README.md   # このREADMEファイル
├── frontend    # フロントエンドのソースコード
├── backend     # バックエンドのソースコード
```

## api

- line api
  - [プッシュメッセージ](https://developers.line.biz/ja/reference/messaging-api/#send-push-message)
- deepl api free
  - [free api](https://support.deepl.com/hc/ja/articles/360021200939-DeepL-API-Free)

## 開発ルール
### 命名規則
- 日本語で単語・熟語を考えて翻訳する
- 誰(何)のためのものかを名前の先頭に持ってくる

|                          | 記法             | 例        |
| ---                      | ---              | ---       |
| ファイル名                | パスカルケース    | UserPage  |
| URL・HTMLのクラス         | ケバブケース      | user-page |
| branch・関数・変数・定数   | キャメルケース    | userId    |
| 不変な定数                | コンスタントケース | GLOBAL_ID |
| db                        | スネークケース    | user_id   |

### git関連
- ブランチ名は「issue/<チケット番号>」
- コミットメッセージのタイトル?(1行目)に、ブランチ名を記入する
