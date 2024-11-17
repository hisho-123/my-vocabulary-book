# backend

## 目的

- バックエンドのコードをまとめる。


## 構成

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
        int id PK
        varchar user_name
        timestamp create_at
    }

    Vocabulary_book {
        int id PK
        int user_id FK
        varchar book_name
    }

    Word {
        int id PK
        int vocabulary_book_id FK
        varchar word
        varchar translated_word
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
