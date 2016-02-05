# Echo Loggers
[![GoDoc](https://godoc.org/github.com/o1egl/elogs?status.svg)](https://godoc.org/github.com/o1egl/elogs) [![Build Status](http://img.shields.io/travis/o1egl/elogs.svg?style=flat-square)](https://travis-ci.org/o1egl/elogs) [![Build status](https://ci.appveyor.com/api/projects/status/em3qo2mils0kj0mv/branch/master?svg=true)](https://ci.appveyor.com/project/o1egl/elogs/branch/master) [![Coverage Status](http://img.shields.io/coveralls/o1egl/elogs.svg?style=flat-square)](https://coveralls.io/r/o1egl/elogs) [![License](http://img.shields.io/:license-mit-blue.svg)](LICENSE)

## Overview

Echo Loggers is a bundle of third party loggers implementation for [Echo](https://github.com/labstack/echo) web framework

## Supported loggers

- [Logrus](https://github.com/Sirupsen/logrus)
- [glog](https://github.com/golang/glog)
- [go-logger](github.com/apsdehal/go-logger)

## Installation

```
$ go get -u github.com/o1egl/echo-loggers
```

## Usage
#### Logrus
```go
    import (
     "github.com/labstack/echo"
     "github.com/o1egl/elogs/logrus"
      log "github.com/Sirupsen/logrus"
    )
    // From default logger
    l := logrus.New()
    // From custom logger
    ls := log.New()
    l := logrus.FromLogger(ls)
    e := echo.New()
    e.SetLogger(l)
```

#### glog
```go
    import (
     "github.com/labstack/echo"
     "github.com/o1egl/elogs/glog"
    )
    l := glog.New()
    e := echo.New()
    e.SetLogger(l)
```

#### go-logger
```go
    import (
     "github.com/labstack/echo"
     "github.com/o1egl/elogs/gologger"
     log "github.com/apsdehal/go-logger"
     "os"
    )
    // From default logger
    l := gologger.New()
    // From custom logger
    lg, _ := log.New("test", 1, os.Stdout)
    l := gologger.FromLogger(lg)
    e := echo.New()
    e.SetLogger(l)
```

## Copyright, License & Contributors

### Submitting a Pull Request

1. Fork it.
2. Create a branch (`git checkout -b my_branch`)
3. Commit your changes (`git commit -am "Added new awesome logger"`)
4. Push to the branch (`git push origin my_branch`)
5. Open a [Pull Request](https://github.com/o1egl/echo-loggers/pulls)
6. Enjoy a refreshing Diet Coke and wait

echo-loggers is released under the MIT license. See [LICENSE](LICENSE)
