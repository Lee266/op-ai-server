# op-ai-server

## 概要

Golangを使用したAPIサーバーのプロジェクトです。
[プロジェクトリポジトリ](https://github.com/Lee266/op-ai-server/settings/secrets/actions)

## 始めるには

始めるには、次のURLからプロジェクトをクローンしてください:
[プロジェクトリポジトリ](https://github.com/Lee266/op-ai-monorepo)

### git commitについて

Githubでコードを管理するため.githooksにフォルダを移行しました。従ってGit hooksの使用するために下記のコードを実行してください。

```sh
git config --local core.hooksPath .githooks
```

## 注意

Goの環境がローカルに設定されていない場合は、Makefileなどの実行をDocker内で行ってください。

## 使用している主なパッケージ

- go v1.21.3
