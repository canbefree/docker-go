
FROM registry.cn-shenzhen.aliyuncs.com/canbefree/docker-go:1.3.5

RUN git config --global user.name "canbefree" \
    && git config --global user.email "45198757@qq.com" 

# RUN mkdir /root/.ssh
COPY .ssh /root/.ssh

RUN  chmod -R 600 /root/.ssh

RUN ldconfig

