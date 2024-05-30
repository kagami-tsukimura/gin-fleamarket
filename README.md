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

### `gin-swagger` Install

```bash
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@latest
```
