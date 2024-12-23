# backend

## 目的

- バックエンドのコードをまとめる。


## 構成

### コアドメイン

- 単語帳の復習

### サブドメイン

- 復習タイミングのスケジューリング
  - 復習間隔の算出
  - お知らせ

- オリジナルの単語帳の作成
  - 単語の翻訳

### 外部システム

- DeepL
  - 単語作成時に翻訳
- LINE notify
  - 復習周期により、通知


```
backend
|
├── main.go
|
├── src
|    ├── service
|    ├── repositories
|    └── controllers
|
├── db
|    └── migrations/schema.sql
```

## DB

```mermaid
erDiagram
    User {
        int id PK "AutoIncrement"
        varchar user_name "not Null, max length: 50"
        varchar password "hash"
        timestamp create_at
    }

    Vocabulary_book {
        int id PK "AutoIncrement"
        int user_id FK
        varchar book_name "max lingth: 10"
    }

    Word {
        int id PK "AutoIncrement"
        int vocabulary_book_id FK
        varchar word "max length: 50"
        varchar translated_word "max length: 50"
        int study_count
        timestamp study_at
        timestamp create_at
    }

    User ||--o{ Vocabulary_book : "has"
    Vocabulary_book ||--o{ Word : "contains"
```

### 接続方法

- dockerのビルド: docker compose up
- docker内に入る: docker exec -it db bash
  - mysql内に入る : mysql -u root -p
    - password                     : root
    - DBの確認                      : show databases;
    - my-vocabulary-book へアクセス : use my-vocabulary-book;
      - tableの確認                 : show tables;
  - mysql外に出る : exit

- dockerに入ると同時にmysqlへ : docker exec -it db mysql -uroot -proot
