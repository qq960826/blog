---
title: PYTHON3下读写文件编码问题
url: 220.html
id: 220
comments: false
categories:
  - python
date: 2016-01-18 16:26:30
tags:
---

以前写爬虫啥的，直接用的
```python
file=open("data1","r+")
a=file.read()
```
如果有中文的话会报以下错误

> Traceback (most recent call last): File "/Users/wzq/Desktop/tice/generatemysql.py", line 13, in <module> a=file.read() File "/Library/Frameworks/Python.framework/Versions/3.5/lib/python3.5/encodings/ascii.py", line 26, in decode return codecs.ascii_decode(input, self.errors)\[0\] UnicodeDecodeError: 'ascii' codec can't decode byte 0xe7 in position 13: ordinal not in range(128)

然后改编吗啥的没有用，依然报错，最后谷歌查到 文件应该这样打开
```python
file=open("data1","r+",encoding="utf-8")
```
然后问题就解决了