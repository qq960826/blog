---
title: 使用Clion编写汇编语言
url: 378.html
id: 378
comments: false
categories:
  - programing
date: 2018-12-22 20:32:31
tags:
---

0x1 建立clion项目，创建文件为hello.S，内容为

```assembly
output:
.ascii "hello world!"
.global hello,_main
hello:
    movq $1,%rdi
    leaq output(%rip),%rsi
    movq $11,%rdx
    movq $0x2000000,%rax
    orq $4,%rax
    syscall
    retq
_main:
    call hello
    xorq %rax,%rax
    retq
```

0x2修改CMakeLists.txt为

```cmake
cmake_minimum_required(VERSION 3.10)
project(asmtest)
set(CMAKE_CXX_STANDARD 11)
enable_language(C ASM)
add_executable(asmtest hello.S)
```

0x3重新加载CMakeLists，点击运行即可，还可以打断点