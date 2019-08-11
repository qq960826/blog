---
title: 使用C语言写网页
url: 338.html
id: 338
comments: false
categories:
  - Linux
date: 2016-07-22 22:02:57
tags:
---

今天刚好做完Linux的socket通信的实验，然后突发奇想如何用C语言写个网页，然后在网上查询相关资料后得知需要让apache开启CGI的支持就可以。 方法如下
```bash
sudo nano /etc/apache2/httpd.conf
```
找到#LoadModule cgi\_module libexec/apache2/mod\_cgi.so 把前面的注释给去除 然后找到 在options后面加上ExecCGI 找到AddHandler，加入cgi 然后用C语言写个hello world测试
```c
#include <stdio.h>
int main(){
        printf("Content-Type: text/html\n\n");
        printf("hello world");
        return 0;
}
```
然后编译
```bash
gcc hello.c -o hello.cgi -std=c99
```
然后进网页http://127.0.0.1/hello.cgi
[![QQ20160722-0@2x](/images/old/2016/07/QQ20160722-0@2x-1024x681.png)](/images/old/2016/07/QQ20160722-0@2x.png)
