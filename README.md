<!--
 * @Author: your name
 * @Date: 2020-02-27 12:42:45
 * @LastEditTime: 2020-03-03 15:22:04
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \docker-go\README.md
 -->
# DOCKER-GO

![](https://img.shields.io/static/v1?label=vscode&message=go&color=yellow)

## 版本详情

| 版本号        | 详情         | 备注  |
| ------------- |:-------------:| -----:|
| 1.2.5 | remote development | add autoconf automake libtool  protobuf |
| 1.2.6 | graphviz | go mod 依赖

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