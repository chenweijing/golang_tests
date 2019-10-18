
## 在协同开发中使用  go module 版本管理


### [基本教程](https://juejin.im/post/5d0b865c6fb9a07f050a6f45)
### gomod 版本管理
1. tag 版本: google.golang.org/grpc v1.12.2 (x.y.z   ，分别为主要版本号，次要版本号，修订版本号
2. pseudo-version: go module 使用伪版本号(pseudo-version) 来标记，伪版号的格式为 [之前最近一次的版本号]-[提交时间`yyyymmddhhmmss`]-[提交哈希值]，比如 `v0.0.0-20180611022520-ca850f594eaa`，这样伪版本号也可以用于比较，伪版本号不需要手动输入，会由 go module 在冻结版本时自动生成。



###  命令解析
1. `go mod tidy` 增加缺失的依赖(module)，丢掉没用的依赖(module)，go.mod不存在时，会新建所有引用，指向 master 分支的lastest commit 
2. `go get -u` 升级到最新的提交，只会到次要版本或者修订版本，不会到主要版本号
3. `go get -u=patch` 升级到最新的修订版本
4. `go get package@version` 升级到指定的版本号
5. `go get ./...` 查找所有依赖


### 引用某个 commit 版本 
1. 引用新版本， 先确认需要使用的文件的commit hash版本，比如要引用 github.com/crispto/utils的4f65ba7cb108e37741c6a5f1ea4204ebeab3f2d1，然后运行 
`go get github.com/crispto/utils@4f65ba7` (7位以上即可)
该命令会自动更新 go.mod 文件 ，然后更新自己的代码，使之匹配新的接口，版本采用 pseudo-version （伪版本）的形式记录

2. 引用旧版本， 如果该版本已经在本地存在直接更改 go.mod，
比如原来的go.mod为
`require github.com/crispto/utils v0.0.0-20191010032225-5158376ee5d8
`
在项目后面的 `v0.0.0-20191010032225-5158376ee5d8` 替换为需要的 commit hash， 运行 `go mod tidy`会直接更新  go.mod
如替换为 `5158376ee5d84570a0d4221b3bd45e0f25d843a3` 后，gomod 变为
`require github.com/crispto/utils v0.0.0-20191010033150-4f65ba7cb108`
 
### 引用某个分支
`go get github.com/crispto/utils@dev`
会将dev的 latest commit 作为当前引用的版本，并且更新 go.mod 文件（仍然使用伪版本的形式）

### 开发项目引用了不同版本

只需要在不同项目根目录下（go.mod所在的目录）执行以上命令即可, 亲测可行。

### 参考资料

[stackoverflow](https://stackoverflow.com/questions/53682247/how-to-point-go-module-dependency-in-go-mod-to-a-latest-commit-in-a-repo/53682399)