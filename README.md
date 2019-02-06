# このリポジトリについて
Go言語+CleanArchitecture+TDDでBlogを作成することを目的とした勉強用リポジトリです。

# 開発用のdatabaseについて
docker-composeを立ち上げることで、開発環境用のdatabaseを立ち上げれます。

# migrationについて
https://github.com/pressly/goose  
を使用。

    [ファイルの作成]
    goose create ファイル名 sql
    [up処理]
    goose postgres "$DATABASE" up 
    [down処理]
    goose postgres "$DATABASE" down