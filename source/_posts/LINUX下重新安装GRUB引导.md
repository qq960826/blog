---
title: linux下重新安装grub引导
url: 225.html
id: 225
comments: false
categories:
  - Linux
date: 2016-02-01 15:26:20
tags:
---

之前台式机不记得是什么原因，引导在HDD上，而系统安装在SSD上面。然后今天由于把HDD拆下来后发现不能开机发现的。 进入系统后
```bash
grub-install /dev/sda
grub-mkconfig -o /boot/grub/grub.cfg
```
[![屏幕快照 2016-02-01 下午3.25.02](/images/old/2016/02/屏幕快照-2016-02-01-下午3.25.02.png)](/images/old/2016/02/屏幕快照-2016-02-01-下午3.25.02.png) 
关机，拔掉HDD的SATA线，在开机成功进入系统
