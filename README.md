# DOCKER-GO

## 版本详情

| 版本号        | 详情         | 备注  |
| ------------- |:-------------:| -----:|
| 1.0.0     | go |  vscode本地开发版本 |
| 1.0.1 | 支持 buildgo | |
| 1.0.2 | goconvery mysql | |
| 1.0.6 | remote development| go本地开发| 
| 1.0.8 | remote development| new version | 
| 1.0.9 | remote development| go 支持etcd | 

## 阿里云自动构建过程

### 阿里云添加镜像

- 添加镜像仓库，选择github自动海外构建

### 添加标签

```bash
    git tag -a release-v1.0.8 -m "Update ...."
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
