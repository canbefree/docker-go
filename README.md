# DOCKER-GO

![](https://img.shields.io/static/v1?label=vscode&message=go&color=yellow)

## 版本详情

last_release : v1.3.6
## 阿里云自动构建过程

#### 阿里云添加镜像

- 添加镜像仓库，选择github自动海外构建

#### 添加标签

```bash
    git tag -a release-v1.3.6 -m "Update ...."
    git push origin --tags
```

#### 删除标签

```bash
# 删除本地
git tag -d release-v1.3.6
#删除远程的
git push origin :refs/tags/release-v1.3.6
```

## 镜像地址

- registry.cn-shenzhen.aliyuncs.com/canbefree/docker-go:1.3.6

---

## 问题

- [ ] gopkgs升级 github.com/uudashr/gopkgs/cmd/gopkgs@latest issueA



## 工作区 volumes备份
- 建议使用镜像备份的方式