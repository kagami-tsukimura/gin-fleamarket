# Gin Fleamarket

[![Go](https://pkg.go.dev/badge/github.com/gin-gonic/gin.svg)](https://pkg.go.dev/github.com/gin-gonic/gin)

## Getting Started

Go, Gin 環境の構築。  
`asdf`を使用。  
`Ubuntu22.04`を使用。

### `asdf` Install

```bash
apt install curl git
git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.13.1
asdf update
. "$HOME/.asdf/asdf.sh"
. "$HOME/.asdf/completions/asdf.bash"
asdf -v
```

### `Go` Install

```bash
asdf plugin add golang https://github.com/asdf-community/asdf-golang.git
asdf install golang latest
asdf global golang latest
asdf current golang
go version
asdf plugin list all | grep -e golang
```

### `Gin` Install

```bash
go get -u github.com/gin-gonic/gin
asdf plugin list all | grep -e gin
```

### `air` Install

```bash
go install github.com/cosmtrek/air@latest
```

```bash
go env GOPATH
```

```bash
vim ~/.bashrc
```

```bash: .bashrc
export PATH="<go env GOPATHで確認したパス>/bin:$PATH"
```

```bash
source ~/.bashrc
```

```bash
air init
```

- `air init` で `.air.toml`が作成されることを確認。

### `air` によるホットリロードで実行

```bash
air
```

- `air` でサーバが実行することを確認。
- パスを変更して、ホットリロードの反映を確認。

### `godotenv`による環境変数の読み込み

```bash
go get github.com/joho/godotenv
```

### `GORM`による ORM の使用

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u gorm.io/driver/postgres
```

### `jwt`トークン生成

```bash
go get -u github.com/golang-jwt/jwt/v5
```

### `cors`設定

```bash
go get github.com/gin-contrib/cors
```

### `pgadmin`接続

```bash
docker-compose up -d
```

- `pgadmin` にログイン。

  - email: gin@example.com
  - password: ginpassword

- `pgadmin` で DB にアクセス。
  - General タブ
    - 名前: fleamarket
  - 接続タブ
    - ホスト名/アドレス: postgres
    - 管理用データベース: fleamarket
    - ユーザー名: ginuser
    - パスワード: ginpassword

### DB Migration

```bash
go run migratinos/migration.go
```
