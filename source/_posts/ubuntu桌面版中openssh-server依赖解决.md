---
title: ubuntu桌面版中openssh-server依赖解决
url: 297.html
id: 297
comments: false
categories:
  - Linux
date: 2016-04-06 13:08:51
tags:
---

安装openssh-server，然后发现如下报错 
root@wzq-virtual-machine:/etc/ssh# apt-get install openssh-server
Reading package lists… Done
Building dependency tree
Reading state information… Done
Some packages could not be installed. This may mean that you have
requested an impossible situation or if you are using the unstable
distribution that some required packages have not yet been created
or been moved out of Incoming.
The following information may help to resolve the situation:

The following packages have unmet dependencies:
openssh-server : Depends: openssh-client (= 1:6.6p1-2ubuntu2.4)
unity-control-center : Depends: libcheese-gtk23 (>= 3.4.0) but it is not going to be installed
Depends: libcheese7 (>= 3.0.1) but it is not going to be installed
E: Error, pkgProblemResolver::Resolve generated breaks, this may be caused by held packages.

解决方法： 先安装aptitude  
```bash
sudo apt-get install aptitude
```
然后重新安装openssh-client,指定版本1:6.6p1-2ubuntu1
```bash
sudo aptitude install openssh-client=1:6.6p1-2ubuntu1
```
最后安装openssh-server
```bash
sudo apt-get install openssh-server
```
然后
```bash
ssh localhost
```
显示成功