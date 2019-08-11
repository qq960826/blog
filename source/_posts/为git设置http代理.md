---
title: 为git设置http代理
url: 92.html
id: 92
comments: false
categories:
  - 网络
date: 2015-11-17 21:01:51
tags:
---

在国内github上传很蛋疼，经常上传不上去，然而我mac下在系统偏好设置里面设置的http代理无法使用。于是我查阅相关资料，最后找到了解决方案。

 1.清除以前的代理设置(没有的不用清除):

```bash
git config --global --unset https.proxy
git config --global --unset http.proxy
```

2.设置代理:

```bash
git config --global https.proxy https://USER:PWD@proxy.whatever:80
git config --global http.proxy http://USER:PWD@proxy.whatever:80
```

注意：如果你的账号或密码含有@，请自行URL转码将@替换为%40 3.验证设置:

```bash
git config --get https.proxy
git config --get http.proxy
```

