---
title: 安卓刷recovery的命令
url: 165.html
id: 165
comments: false
categories:
  - 未分类
date: 2015-12-11 13:28:47
tags:
---

查看设备
```bash
adb devices
```

进入bootloader
```bash
adb reboot bootloader
```
解锁
```bash
fastboot oem unlock 2242221191479xxx
```
查看解锁状态
```bash
fastboot oem get-bootinfo
```
刷入recovery
```bash
fastboot flash recovery recovery.img
```
重启
```bash
fastboot reboot
```