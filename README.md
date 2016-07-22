# restdemo
基于go语言echo框架写的restdemo,mysql、redispool

## go包管理glide
```
$ glide create                            # Start a new workspaces
$ open glide.yaml                         # and edit away!
$ glide get github.com/Masterminds/cookoo # Get a package and add to glide.yaml
$ glide install                           # Install packages and dependencies
# work, work, work
$ go build                                # Go tools work normally
$ glide up                                # Update to newest versions of the package
```

注：这里每次通过glide get ...后只会在vendor路径下生成第三方包，需要手动将该包同步到$GOPATH,然后运行go build
