---
title: ubuntu15下启动直接进入命令行模式
url: 228.html
id: 228
comments: false
categories:
  - Linux
date: 2016-02-02 18:05:31
tags:
---

以前网上修改grub的方法已经失效 执行以下命令开机直接进入命令行
```bash
sudo systemctl enable multi-user.target --force
sudo systemctl set-default multi-user.target
```
如果要进入图形界面
```bash
service lightdm start
```