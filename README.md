# DOCKER-GO

## 版本详情

| 版本号        | 详情         | 备注  |
| ------------- |:-------------:| -----:|
| 1.1.0 | remote development | export GOPROXY=https://goproxy.cn |

## 阿里云自动构建过程

### 阿里云添加镜像

- 添加镜像仓库，选择github自动海外构建

### 添加标签

```bash
    git tag -a release-v1.1.0 -m "Update ...."
    git push origin --tags
```

### 删除标签

```bash
# 删除本地
git tag -d release-v1.0.5
#删除远程的
git push origin :refs/tags/release-v1.0.5
```

### 镜像地址

- registry.cn-shenzhen.aliyuncs.com/canbefree/docker-go:1.0.5

### 修改日志

vscode 更新升级 1.1.0 版本