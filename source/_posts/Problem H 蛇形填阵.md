---
title: 'Problem H: 蛇形填阵'
url: 181.html
id: 181
comments: false
categories:
  - sdust
date: 2015-12-18 16:55:41
tags:
---

博主c语言渣了。。。写了1个多小时。。。

Problem H: 蛇形填阵

Problem H: 蛇形填阵
---------------

Time Limit: 1 Sec  Memory Limit: 128 MB  
Submit: 1856  Solved: 833  
\[[Submit](submitpage.php?cid=2021&pid=7&langmask=1022)\]\[[Status](problemstatus.php?id=1184)\]\[[Web Board](bbs.php?pid=1184&cid=2021)\]

Description
-----------

将1～n\*n填入一个n\*n的矩阵中，并要求成为蛇形。蛇形即是从右上角开始向下，向左，向上，向右，循环填入数字。

比如n=5时矩阵为：

13 14 15 16  1
12 23 24 17  2
11 22 25 18  3
10 21 20 19  4
  9  8   7   6   5

Input
-----

﻿输入有多行，每行为一个整数n（1<=n<=50)，每组答案用空行隔开。

Output
------

输出一个n*n的矩阵，n行n列每个数字用一个空格隔开，不能有多余空格。

Sample Input
------------

5

Sample Output
-------------

13 14 15 16 1 12 23 24 17 2 11 22 25 18 3 10 21 20 19 4 9 8 7 6 5

HINT
----

Append Code
-----------
```c
#include <stdio.h>
#include <stdlib.h>
#define len_max 51
void hua(int ** juzheng,int n){
    int ceng=1,b=n;
    int count=1;
    while (b>0) {
        int i=ceng-1;
        while (i<b) {//最右边一列
            juzheng\[i\]\[b-1\]=count++;
            i++;
        }
        i=n-ceng-1;
        while (i>=(ceng-1)) {//最下面一行
            juzheng\[b-1\]\[i\]=count++;
            i--;
        }
        i=0;
        while (i<(b-ceng)) {//最左边一列
            juzheng\[b-i-2\]\[ceng-1\]=count++;
            i++;
        }
        i=ceng;
        while (i<(b-1)) {//最上面一行
            juzheng\[ceng-1\]\[i\]=count++;
            i++;
        }
        b--;
        ceng++;
    }
}
int main(){
    int n;
    int flag;
    int \*\*s=(int \*\*)calloc(len_max, sizeof(int*));
    for (int i=0; i<len_max; i++) {
        *(s+i)=(int *)calloc(len_max, sizeof(int));
    }
    while (scanf("%d",&n)!=EOF) {
        hua(s, n);
        for (int i=0; i<n; i++) {
            flag=0;
            for (int j=0; j<n; j++) {
                if (flag==0) {
                    flag=1;
                    printf("%d",s\[i\]\[j\]);
                    continue;
                }
                printf("% d",s\[i\]\[j\]);
            }
            printf("\\n");
        }
        printf("\\n");
    }
}
```