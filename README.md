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

---

## Step 3

```
rm -rf $GOPATH/pkg/mod/cache
→ cat go.mod
module github.com/im-kulikov/echo-issue-1185

require (
        github.com/im-kulikov/helium v0.5.6
)
→ rm go.sum
→ go mod tidy
go: finding github.com/im-kulikov/helium v0.5.6
go: finding github.com/davecgh/go-spew v1.1.0
go: finding github.com/prometheus/common v0.0.0-20180518154759-7600349dcfe1
go: finding github.com/beorn7/perks v0.0.0-20180321164747-3a771d992973
go: finding github.com/smartystreets/goconvey v0.0.0-20170602164621-9e8dc3f972df
go: finding github.com/go-playground/universal-translator v0.16.0
go: finding github.com/chapsuk/wait v0.0.0-20180530144602-2dc40db0485b
go: finding github.com/nats-io/go-nats v1.5.0
go: finding golang.org/x/sys v0.0.0-20180715085529-ac767d655b30
go: finding go.uber.org/dig v1.3.0
go: finding go.uber.org/multierr v1.1.0
go: finding github.com/urfave/cli v1.20.0
go: finding github.com/spf13/cast v1.2.0
go: finding github.com/gopherjs/gopherjs v0.0.0-20180628210949-0892b62f0d9f
go: finding github.com/magiconair/properties v1.8.0
go: finding github.com/stretchr/testify v1.2.2
go: finding google.golang.org/genproto v0.0.0-20180722052100-02b4e9547331
go: finding github.com/labstack/gommon v0.0.0-20180613044413-d6898124de91
go: finding github.com/pelletier/go-toml v1.2.0
go: finding github.com/prometheus/procfs v0.0.0-20180705121852-ae68e2d4c00f
go: finding github.com/bsm/redis-lock v8.0.0+incompatible
go: finding github.com/fsnotify/fsnotify v1.4.7
go: finding github.com/matttproud/golang_protobuf_extensions v1.0.1
go: finding golang.org/x/net v0.0.0-20180719180050-a680a1efc54d
go: finding github.com/chapsuk/mserv v0.0.0-20180706125941-107033951ceb
go: finding github.com/jtolds/gls v4.2.1+incompatible
go: finding github.com/smartystreets/assertions v0.0.0-20180301161246-7678a5452ebe
go: finding golang.org/x/text v0.3.0
go: finding github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a
go: finding github.com/golang/protobuf v1.1.0
go: finding github.com/chapsuk/worker v0.4.0
go: finding github.com/prometheus/client_golang v0.8.0
go: finding gopkg.in/yaml.v2 v2.2.1
go: finding go.uber.org/zap v1.9.0
go: finding github.com/mattn/go-colorable v0.0.9
go: finding github.com/valyala/bytebufferpool v0.0.0-20160817181652-e746df99fe4a
go: finding github.com/labstack/echo v0.0.0-20180412143600-6d227dfea4d2
go: finding github.com/pmezard/go-difflib v1.0.0
go: finding github.com/prometheus/client_model v0.0.0-20180712105110-5c3871d89910
go: finding github.com/go-pg/pg v6.14.3+incompatible
go: finding github.com/pkg/errors v0.8.0
go: finding github.com/bouk/monkey v0.0.0-20180215074647-5df1f207ff77
go: finding github.com/robfig/cron v0.0.0-20180505203441-b41be1df6967
go: finding github.com/valyala/fasttemplate v0.0.0-20170224212429-dcecefd839c4
go: finding github.com/mitchellh/mapstructure v0.0.0-20180715050151-f15292f7a699
go: finding github.com/spf13/viper v1.0.2
go: finding github.com/go-playground/locales v0.12.1
go: finding github.com/spf13/afero v1.1.1
go: finding github.com/mattn/go-isatty v0.0.3
go: finding github.com/go-redis/redis v6.13.0+incompatible
go: finding github.com/spf13/pflag v1.0.1
go: finding gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
go: finding gopkg.in/go-playground/validator.v9 v9.20.2
go: finding google.golang.org/grpc v1.13.0
go: finding github.com/hashicorp/hcl v0.0.0-20180404174102-ef8a98b0bbce
go: finding github.com/nats-io/nuid v1.0.0
go: finding github.com/spf13/jwalterweatherman v0.0.0-20180109140146-7c0cea34c8ec
go: finding go.uber.org/atomic v1.3.2
go: finding golang.org/x/crypto v0.0.0-20180718160520-a2144134853f
go: finding github.com/smartystreets/gunit latest
go: finding github.com/golang/glog latest
go: finding github.com/BurntSushi/toml v0.3.0
go: finding gopkg.in/go-playground/assert.v1 v1.2.1
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