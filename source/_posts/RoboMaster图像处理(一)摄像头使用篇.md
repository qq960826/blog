---
title: 'RoboMaster图像处理(一):摄像头使用篇'
url: 345.html
id: 345
comments: false
categories:
  - 图像识别
date: 2017-06-05 22:01:10
tags:
---

相信大家在使用OpenCV读取摄像头的时候会发现几点问题： 
1.读取图像的速度比较慢 
2.OpenCV的图像读取是为了保证视频流的连续性，而不是实时性，从而会导致识别算法的滞后性 
3.把摄像头的IO操作与算法混合到一起的时候，IO操作加长了整体识别的速度。 
于是在YPW的指引下，采用了共享内存，就是写一个程序用V4L2的库将摄像头的图像读进来，然后将它写入共享内存，然后其他需要摄像头图像的程序直接从共享内存读取图片。 
这样设计有如下的好处: 
1.摄像头操作直接用比较底层的V4L2的库进行读取，IO操作速度相对比较快 
2.miniPC的CPU是四核，开一个进程专门负责摄像头的读取，能提高CPU的利用率 
3.识别算法部分和底层IO操作分离以后，就直接从共享内存中把图片复制过来，速度快，还可以多个程序共用一个摄像头 
4.共享内存中的图片是实时采集的，放弃了视频流的连续性，保证了实时性，从而不会导致识别算法的滞后 
源代码地址在：[https://github.com/qq960826/RoboMaster/tree/master/VideoService](https://github.com/qq960826/RoboMaster/tree/master/VideoService) 
里面分两部分，一个是Server负责摄像头数据的采集，写入共享内存，还有一个是Client负责共享内存的读取然后转换成OpenCV的Mat的形式 
Server的使用方式: 
1.修改Server里面的main.cpp里面的devicefile这个常量，改成自己的摄像头地址，一般都是和我文件里面一样的(/dev/video0) 
2.编译Server:
```bash
cmake .
make
```
3.加入到开机启动项
```bash
nano /etc/rc.local
```
在exit 0前面加入刚才编译出来的二进制文件地址 
然后启动这个服务 
Client的使用方式: 这个的使用例子我写在了Client里面的main.cpp文件，首先先包含Structure.cpp和Structure.h这个两个文件,然后调用 
uint8_t getImageFromMemory(Mat &image) 
读取成功会返回0，并把图像拷入传入的Mat类型的变量中，读取失败则会返回-1。

最后大家如果要把这个移植到自己写的里面，按上述操作执行，然后把读取摄像头的那一部分改成getImageFromMemory这个函数即可