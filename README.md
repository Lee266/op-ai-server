# op-ai-server

## 概要

Golangを使用したAPIサーバーのプロジェクトです。  
[op-ai-serverのリポジトリ](https://github.com/Lee266/op-ai-server)

## 準備

### Auth0の利用

このプロジェクトではAuth0を使っているため、Auth0のアカウントを作成しバックエンド側のApplicationを作成してください。  
[Auth0のサイト](https://auth0.com/)

### .envファイルの修正

.envファイルがない場合は.example.envからコピーしてください。  
.envファイル内の{}の値を自分が設定したい値に変更してください。

## 始めるには

### 初めに

始めるには、次のURLからプロジェクトレポジトリからクローンしプロジェクトレポジトリのREADMEに従ってください.  
[プロジェクトリポジトリ](https://github.com/Lee266/op-ai-monorepo)

上記が成功した場合、<http://localhost:8002>を開いてください。

### git commitについて

Githubでコードを管理するため.githooksにフォルダを移行しました。従ってGit hooksの使用するために下記のコードを実行してください。

```sh
git config --local core.hooksPath .githooks
```

## 注意

Goの環境がローカルに設定されていない場合は、Makefileなどの実行をDocker内で行ってください。

## 使用している主なパッケージ

- go v1.21.3
- gorm.io/gorm v1.25.5
- gorm/driver/postgres v1.5.4
- gin v1.9.1
- auth0 v2.2.0

### その他

### ローカルでテストをするには

dockerの環境内に入り、以下のコマンドを実行してください。

```docker
make test
```
