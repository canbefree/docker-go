# DOCKER-GO

![](https://img.shields.io/static/v1?label=vscode&message=go&color=yellow)

## 版本详情

| 版本号        | 详情         | 备注  |
| ------------- |:-------------:| -----:|
| 1.1.6 | remote development | export GOPROXY=https://goproxy.cn |
| 1.1.7 | remote development | add autoconf automake libtool |
| 1.1.8 | remote development | add autoconf automake libtool  protobuf |

## 阿里云自动构建过程

#### 阿里云添加镜像

- 添加镜像仓库，选择github自动海外构建

#### 添加标签

```bash
    git tag -a release-v1.1.5 -m "Update ...."
    git push origin --tags
```

#### 删除标签

```bash
# 删除本地
git tag -d release-v1.1.5
#删除远程的
git push origin :refs/tags/release-v1.1.5
```

## 镜像地址

- registry.cn-shenzhen.aliyuncs.com/canbefree/docker-go:1.1.5

---

## 问题

- [ ] gopkgs升级 github.com/uudashr/gopkgs/cmd/gopkgs@latest issue