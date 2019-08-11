---
title: oj之坑－输入数组未清空导致WA
url: 104.html
id: 104
comments: false
categories:
  - OJ
  - sdust
date: 2015-11-20 17:26:58
tags:
---

Problem I: Sequence Problem (III) : Array Practice

Problem I: Sequence Problem (III) : Array Practice
--------------------------------------------------

Time Limit: 1 Sec  Memory Limit: 4 MB  
Submit: 4733  Solved: 1534  
\[[Submit](submitpage.php?cid=1956&pid=8&langmask=1022)\]\[[Status](problemstatus.php?id=1052)\]\[[Web Board](bbs.php?pid=1052&cid=1956)\]

Description
-----------

整数序列是一串按特定顺序排列的整数，整数序列的长度是序列中整数的个数，不可定义长度为负数的整数序列。  
两整数序列A、B的和定义为一个新的整数序列C，序列C的长度是A、B两者中较长的一个，序列C的每个位置上的整数都是A、B对应位置之和。若序列A、B不等长，不妨假设A比B整数多，那么序列C中多出B的那部分整数视作A的对应位置上的整数与0相加。  
你的任务是计算符合某些要求的整数序列的和，这些序列中的整数都是小于1000的非负整数。

Input
-----

输入的第一行为一个整数M(M>0)，后面有M行输入。每行第一个整数为N(N<=1000)，后接一个长度为N的整数序列。

Output
------

对输入的整数序列两两相加：第1行和第2行相加、第2行和第3行相加……按顺序输出结果：每行输出一个整数序列，每两个整数之间用一个空格分隔。若最后序列不足两个，则视作补一个长度为0的整数序列相加。  
值得注意的是一个长度为0的整数序列也应该有输出，即使没有整数输出，也应该占有一行，因为“每行输出一个整数序列”。

Sample Input
------------

3
3 1 2 3
5 10 15 20 30 50
4 100 200 300 400

Sample Output
-------------

11 17 23 30 50
110 215 320 430 50
100 200 300 400

HINT
----

这里最少要用到两个数组来存储整数序列。

Append Code
-----------

\[[Submit](submitpage.php?cid=1956&pid=8&langmask=1022)\]\[[Status](problemstatus.php?id=1052)\]\[[Web Board](bbs.php?pid=1052&cid=1956)\]

以下是错误代码

```c
//
//  oj.c
//  tets
//
//  Created by wzq on 15/11/20.
//  Copyright © 2015年 wzq. All rights reserved.
//

#include "stdio.h"
#include "string.h"
int getmax(int a,int b){
    if (a>b) {
        return a;
    }
    return b;
}

int main(int argc, char const *argv[])
{
    int num1,num2;
    int k;
    int count=0;
    int array [2] [1005];
    memset(array, 0, sizeof(array));
    int flag;
    int flag1;
    scanf("%d",&num1);
    flag1=0;
    for (int i=0; i<num1; i++) {
        if (count==0) {
            scanf("%d",&num2);
            array[0][0]=num2;
            for (int j=1; j<=num2; j++) {
                scanf("%d",&array[0][j]);
            }
            count=1;
        }else{
            scanf("%d",&num2);
            array[1][0]=num2;
            for (int j=1; j<=num2; j++) {
                scanf("%d",&array[1][j]);
            }
            count=0;
        }
        if (flag1!=0) {//第一个不输出
            
            k=getmax(array[0][0],array[1][0]);
            flag=0;
            for (int j=1; j<=k; j++) {
                if (flag==0) {
                    printf("%d",array[0][j]+array[1][j]);
                    flag=1;
                }else{
                    printf(" %d",array[0][j]+array[1][j]);
                }
            }
            
            printf("\n");
        }else{
            flag1=1;
        }
        
    }
    if(count==0){
        memset(array[0], 0, sizeof(array[0]));
        k=array[1][0];

    
    }else{
        memset(array[1], 0, sizeof(array[0]));
        k=array[0][0];
    }
    flag=0;
    for (int j=1; j<=k; j++) {
        if (flag==0) {
            printf("%d",array[0][j]+array[1][j]);
            flag=1;
        }else{
            printf(" %d",array[0][j]+array[1][j]);
        }
    }
    
    printf("\n");
    
    return 0;
    
    
    
}
```



以下是正确代码

```c
//
//  oj.c
//  tets
//
//  Created by wzq on 15/11/20.
//  Copyright © 2015年 wzq. All rights reserved.
//

#include "stdio.h"
#include "string.h"
int getmax(int a,int b){
    if (a>b) {
        return a;
    }
    return b;
}

int main(int argc, char const *argv[])
{
    int num1,num2;
    int k;
    int count=0;
    int array [2] [1005];
    memset(array, 0, sizeof(array));
    int flag;
    int flag1;
    scanf("%d",&num1);
    flag1=0;
    for (int i=0; i<num1; i++) {
        if (count==0) {
            scanf("%d",&num2);
            memset(array[0], 0, sizeof(array[0]));
            array[0][0]=num2;
            
            for (int j=1; j<=num2; j++) {
                scanf("%d",&array[0][j]);
            }
            count=1;
        }else{
            scanf("%d",&num2);
            memset(array[1], 0, sizeof(array[0]));
            array[1][0]=num2;
            for (int j=1; j<=num2; j++) {
                scanf("%d",&array[1][j]);
            }
            count=0;
        }
        if (flag1!=0) {//第一个不输出
            
            k=getmax(array[0][0],array[1][0]);
            flag=0;
            for (int j=1; j<=k; j++) {
                if (flag==0) {
                    printf("%d",array[0][j]+array[1][j]);
                    flag=1;
                }else{
                    printf(" %d",array[0][j]+array[1][j]);
                }
            }
            
            printf("\n");
        }else{
            flag1=1;
        }
        
    }
    
    if(count==0){
        memset(array[0], 0, sizeof(array[0]));
        k=array[1][0];
        
        
    }else{
        memset(array[1], 0, sizeof(array[0]));
        k=array[0][0];
    }
    flag=0;
    for (int j=1; j<=k; j++) {
        if (flag==0) {
            printf("%d",array[0][j]+array[1][j]);
            flag=1;
        }else{
            printf(" %d",array[0][j]+array[1][j]);
        }
    }
     printf("\n");
    
    
   
    
    return 0;
    
    
    
}
```



原因分析:错误是由于当用户输入长度为0的数组时，由于长度为0，循环没有走进去，然后导致原数组没有清空，最后导致错误的数组相加。所以在赋值前需要清空数组