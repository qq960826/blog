---
title: 使用finalspeed加速访问国外VPS速度
url: 307.html
id: 307
comments: false
categories:
  - 网络
date: 2016-04-09 15:36:44
tags:
---

FinalSpeed是高速双边加速软件,可加速所有基于tcp协议的网络服务,在高丢包和高延迟环境下,仍可达到90%的物理带宽利用率,即使高峰时段也能轻松跑满带宽.然后我使用后，基本上快把带宽跑满了。。。
首先下载安装脚本
```bash
wget http://fs.d1sm.net/finalspeed/install_fs.sh
```
然后运行脚本
```bash
chmod +x install_fs.sh
./install_fs.sh 2>&1 | tee install.log
```
然后等待安装
安装成功 [![QQ20160409-0@2x](/images/old/2016/04/QQ20160409-0@2x-1024x900.png)](/images/old/2016/04/QQ20160409-0@2x.png) 
卸载
```bash
sh /fs/stop.sh
rm -rf /fs
```
启动
```bash
sh /fs/start.sh
tail -f /fs/server.log
```
停止
```bash
sh /fs/stop.sh
```
重新启动
```bash
sh /fs/restart.sh
tail -f /fs/server.log
```
查看日志
```bash
tail -f /fs/server.log
```
默认udp 150和tcp 150 ,由于finalspeed的工作原理,请不要在本机防火墙开放finalspeed所使用的tcp端口.
os x客户端下载
链接：[http://pan.baidu.com/s/1hrIKITu](http://pan.baidu.com/s/1hrIKITu) 密码：qoji

打开客户端后输入ip，添加要加速的端口就行了 [![QQ20160409-1@2x](/images/old/2016/04/QQ20160409-1@2x-1024x627.png)](/images/old/2016/04/QQ20160409-1@2x.png)
