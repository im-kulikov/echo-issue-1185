# Echo issue 1185
https://github.com/labstack/echo/issues/1185

```
→ go version
go version go1.11 darwin/amd64
→ export GO111MODULE=on
→ go mod init github.com/im-kulikov/echo-issue-1185
go: creating new go.mod: module github.com/im-kulikov/echo-issue-1185
→ go mod tidy
go: finding github.com/im-kulikov/helium v0.5.6
go: finding github.com/chapsuk/mserv v0.0.0-20180706125941-107033951ceb
go: downloading github.com/im-kulikov/helium v0.5.6
go: finding github.com/smartystreets/gunit latest
go: finding github.com/golang/glog latest
go: finding golang.org/x/sync/errgroup latest
go: finding golang.org/x/sync latest
→ cat go.mod
module github.com/im-kulikov/echo-issue-1185

require (
        github.com/BurntSushi/toml v0.3.0 // indirect
        github.com/chapsuk/mserv v0.3.2
        github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
        github.com/im-kulikov/helium v0.5.6
        github.com/labstack/echo v0.0.0-20180412143600-6d227dfea4d2
        github.com/smartystreets/gunit v0.0.0-20180314194857-6f0d6275bdcd // indirect
        go.uber.org/dig v1.4.0
        go.uber.org/zap v1.9.1
        golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f // indirect
        gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)
```

As you can see
```
github.com/labstack/echo v0.0.0-20180412143600-6d227dfea4d2
```

---

## Step 2:

```
→ rm -rf $GOPATH/pkg/mod/cache/download/github.com/labstack/echo
→ cat go.mod
module github.com/im-kulikov/echo-issue-1185

require (
        github.com/im-kulikov/helium v0.5.6
)
→ go mod tidy
go: finding github.com/labstack/echo v0.0.0-20180412143600-6d227dfea4d2
go: downloading github.com/chapsuk/mserv v0.0.0-20180706125941-107033951ceb
go: finding github.com/smartystreets/gunit latest
go: finding github.com/golang/glog latest
go: finding golang.org/x/sync/errgroup latest
go: finding golang.org/x/sync latest
→ cat go.mod
module github.com/im-kulikov/echo-issue-1185

require (
        github.com/BurntSushi/toml v0.3.0 // indirect
        github.com/chapsuk/mserv v0.0.0-20180706125941-107033951ceb
        github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
        github.com/im-kulikov/helium v0.5.6
        github.com/labstack/echo v0.0.0-20180412143600-6d227dfea4d2
        github.com/smartystreets/gunit v0.0.0-20180314194857-6f0d6275bdcd // indirect
        go.uber.org/dig v1.3.0
        go.uber.org/zap v1.9.0
        golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f // indirect
        gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)
```
As you can see
```
github.com/labstack/echo v0.0.0-20180412143600-6d227dfea4d2
```