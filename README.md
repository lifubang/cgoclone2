## What did you see happen?
```
root@acmcoder:/home/acmcoder/cgo# go version
go version go1.22.0 linux/amd64
root@acmcoder:/home/acmcoder/cgo# go run main.go
STAGE_PARENT
STAGE_CHILD
STAGE_INIT
This from nsexec
```

## What did you expect to see?
```
root@acmcoder:/home/acmcoder/cgo# go version
go version go1.21.1 linux/amd64
root@acmcoder:/home/acmcoder/cgo# go run main.go
STAGE_PARENT
STAGE_CHILD
STAGE_INIT
This from nsexec
From main!
```

## Difference
The last line `From main!` is gone in `go1.22.0`.