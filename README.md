# docker-go
阿里云自动构建 



### Release:

- 1.0.0:初始版本


### 使用示例：
  
```Dockerfile
FROM registry.cn-shenzhen.aliyuncs.com/canbefree/docker-go:1.0.0
```

# tag添加方式

``` 
    git tag -a release-v1.0.0 -m "Update ...."
    git push origin --tags
```

```bash
# 删除本地
git tag -d release-v1.0.0
#删除远程的
git push origin :refs/tags/release-v1.0.0
```