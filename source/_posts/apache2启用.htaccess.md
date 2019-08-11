---
title: apache2启用.htaccess
url: 334.html
id: 334
comments: false
categories:
  - Linux
date: 2016-05-08 01:20:09
tags:
---

首先执行

```bash
sudo a2enmod rewrite
```

然后

```bash
sudo nano /etc/apache2/apache2.conf
```

找到 

Options FollowSymLinks
AllowOverride None
Require all denied

修改为
Options FollowSymLinks
AllowOverride All
Require all denied
保存
最后

```bash
sudo service apaches restart
```

