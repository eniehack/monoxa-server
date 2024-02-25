# monoxa-backend

## install

### required packages

- golang
- sqlc
- dbmate

### step

1. install required packages
2. git clone & cd
3. sqlc generate -f sqlc.yml
4. dbmate up
7. fetch `credential.json` which contains firebase configurations from firebase console
5. cp config.example.toml config.toml
6. vim config.toml
7. go build bin/main.go -o monoxa
8. ./monoxa -f config.toml
