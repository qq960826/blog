---
title: os x下终端使用http代理
url: 172.html
id: 172
comments: false
categories:
  - 网络
date: 2015-12-15 22:08:30
tags:
---

进入终端 配置代理
```bash
export http_proxy="http://USER:PWD@proxy.whatever:port"
export https_proxy="http://USER:PWD@proxy.whatever:port"
```
然后使用代理 在要代理的程序前面加:sudo -E 例如
```bash
sudo -E wget https://www.google.com
```
成功如图
[![QQ20151215-0@2x](/images/old/2015/12/QQ20151215-0@2x.png)](/images/old/2015/12/QQ20151215-0@2x.png)
