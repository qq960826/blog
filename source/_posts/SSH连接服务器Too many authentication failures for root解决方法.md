---
title: SSH连接服务器Too many authentication failures for root解决方法
url: 222.html
id: 222
comments: false
categories:
  - 网络
date: 2016-01-18 18:56:23
tags:
---

ssh连接我的服务器提示这个 

```bash
sudo ssh wzq.hk -i qingyun
disconnect from 121.201.24.112: 2: Too many authentication failures for root Disconnected from 121.201.24.112 
```
然后我查找资料，要把连接后面加入一个-o IdentitiesOnly=yes的参数即可
```bash
sudo ssh wzq.hk -i qingyun -o IdentitiesOnly=yes
```