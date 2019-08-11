---
title: 北京联通光猫G-140W-UD超级管理员破解及改桥接
url: 360.html
id: 360
comments: false
categories:
  - 网络
date: 2018-12-11 10:42:21
tags:
---

本文未经许可禁止转载 

**0x1 查看超级管理员的账号密码** 
此步骤如果已经得到超级管理员的账号密码可以省，在网上看北京联通的默认账户名与密码均为CUAdmin。如果密码不对可以再来尝试此步 

1.使用chrome浏览器进入http://192.168.1.1/
[![](https://wzq.io/wp-content/uploads/2018/12/gm1-1024x745.png)](https://wzq.io/wp-content/uploads/2018/12/gm1.png)
2.输入光猫背部的账号与密码登录 
[![](https://wzq.io/wp-content/uploads/2018/12/gm2-1024x745.png)](https://wzq.io/wp-content/uploads/2018/12/gm2.png) 
3.进入管理-用户管理 
[![](https://wzq.io/wp-content/uploads/2018/12/gm3-1024x745.png)](https://wzq.io/wp-content/uploads/2018/12/gm3.png) 
4.在浏览器页面中点击右键，然后点击审查元素 
[![](https://wzq.io/wp-content/uploads/2018/12/gm4-1024x745.png)](https://wzq.io/wp-content/uploads/2018/12/gm4.png) 
5.点击下面的任意元素按ctrl+f(mac下位command+F)搜索login_supcfg
[![](https://wzq.io/wp-content/uploads/2018/12/gm5-1024x745.png)](https://wzq.io/wp-content/uploads/2018/12/gm5.png) 
6.搜索后你会得到var login_supcfg = { Password:'CUAdmin', UserName:'CUAdmin' };这样一个文本，即账号为CUAdmin，密码为CUAdmin 
7.注销登录
**0x2登录超级用户** 
1.使用chrome浏览器进入http://192.168.1.1/ 
[![](https://wzq.io/wp-content/uploads/2018/12/gm1-1024x745.png)](https://wzq.io/wp-content/uploads/2018/12/gm1.png) 
2.在登录按钮上(>>符合)右键点击审查元素，然后找到
```html
<form id="loginForm" method="post" action="login.cgi">
```
[![](https://wzq.io/wp-content/uploads/2018/12/gm7-1024x745.png)](https://wzq.io/wp-content/uploads/2018/12/gm7.png) 
3.双击action="login.cgi"，将action="login.cgi"改为action="cu.html"
[![](https://wzq.io/wp-content/uploads/2018/12/gm8-1024x745.png)](https://wzq.io/wp-content/uploads/2018/12/gm8.png)
4.关闭审查元素的小窗口，输入第一步得到的账号密码进行登录,然后成功进入后台 [![](https://wzq.io/wp-content/uploads/2018/12/gm11-1024x745.png)](https://wzq.io/wp-content/uploads/2018/12/gm11.png)
**0x3修改至桥接模式**
1.进入基本配置-上行线路配置，WAN连接列表选中2\_INTERNET\_B\_VID\_3961
[![](https://wzq.io/wp-content/uploads/2018/12/gm10-1024x745.png)](https://wzq.io/wp-content/uploads/2018/12/gm10.png) 
2.将封装类型改为PPPoE，连接模式由路由模式改为桥接模式即可，最后点击保存/启用 

最后这个技术我是在淘宝花100元买的，然后把人家的技术悟出来改的简易版，希望大家能打一个赏 
![](https://wzq.io/wp-content/uploads/2018/12/IMG_8594-1024x1024.jpg)
[![](https://wzq.io/wp-content/uploads/2018/12/IMG_8592.jpg)](https://wzq.io/wp-content/uploads/2018/12/IMG_8592.jpg)