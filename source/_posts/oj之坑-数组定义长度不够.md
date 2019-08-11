---
title: oj之坑-数组定义长度不够
url: 102.html
id: 102
comments: false
categories:
  - OJ
  - sdust
date: 2015-11-20 12:13:53
tags:
---

Problem H: Sequence Problem (II) : Array Practice

Problem H: Sequence Problem (II) : Array Practice
-------------------------------------------------

Time Limit: 1 Sec  Memory Limit: 4 MB  
Submit: 5170  Solved: 1790  
\[[Submit](submitpage.php?cid=1956&pid=7&langmask=1022)\]\[[Status](problemstatus.php?id=1051)\]\[[Web Board](bbs.php?pid=1051&cid=1956)\]

Description
-----------

整数序列是一串按特定顺序排列的整数，整数序列的长度是序列中整数的个数，不可定义长度为负数的整数序列。

两整数序列A、B的和定义为一个新的整数序列C，序列C的长度是A、B两者中较长的一个，序列C的每个位置上的整数都是A、B对应位置之和。若序列A、B不等长，不妨假设A比B整数多，那么序列C中多出B的那部分整数视作A的对应位置上的整数与0相加。

你的任务是计算符合某些要求的整数序列的和，这些序列中的整数都是小于1000的非负整数。

Input
-----

输入的第一行为一个整数M(M>0)，后面有M行输入。每行输入为不超过1000个整数的整数序列，每个整数序列的输入均以0结束。

Output
------

对输入的整数序列两两相加：第1行和第2行相加、第3行和第4行相加……按顺序输出结果：每行输出一个整数序列，每两个整数之间用一个空格分隔。若序列数目不为偶数，则视作补一个长度为0的整数序列相加。

值得注意的是一个长度为0的整数序列也应该有输出，即使没有整数输出，也应该占有一行，因为“每行输出一个整数序列”。

Sample Input
------------

3
1 2 3 0
10 15 20 30 50 0
100 200 300 400 0

Sample Output
-------------

11 17 23 30 50
100 200 300 400

HINT
----

这里最少要用到一个数组来存数整数序列或整数序列的和。一个省事的做法是把数组定义的稍微大一点，因为有时你的程序可能会边界处理的不是太好。

Append Code
-----------

\[[Submit](submitpage.php?cid=1956&pid=7&langmask=1022)\]\[[Status](problemstatus.php?id=1051)\]\[[Web Board](bbs.php?pid=1051&cid=1956)\]

```c
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
    int num;
    int max;
    int array1 [1000];
    int array2 [1000];
    int array_total [1000];
    int array1_len,array2_len;
    int flag;
    scanf("%d",&num);
    if(num==0){
        return 0;
    }
    for (int j=0; j<num; j+=2) {
        if (num-j>1) {
            flag=0;
            memset(array1, 0, sizeof(array1));
            memset(array2, 0, sizeof(array2));
            memset(array_total, 0, sizeof(array_total));
            array1_len=0;
            while (1) {
                scanf("%d",&array1[array1_len]);
                if(array1[0]==0){
                    //printf("\n");
                    break;
                    
                }else{
                    if (array1[array1_len]!=0) {
                        array1_len++;
                    }else{
                        break;
                    }}
            }
            array2_len=0;
            while (1) {
                scanf("%d",&array2[array2_len]);
                if(array2[0]==0){
                    //printf("\n");
                    break;
                    
                }else{
                    if (array2[array2_len]!=0) {
                        array2_len++;
                    }else{
                        break;
                    }}
            }
            max=getmax(array1_len,array2_len);
            if (max!=0) {
                for (int i=0; i<max; i++) {
                    if (flag==0) {
                        flag=1;
                        printf("%d",array1[i]+array2[i]);
                    }else{
                        printf(" %d",array1[i]+array2[i]);}
                }
            }
            
            printf("\n");
        }else{
            flag=0;
            memset(array1, 0, sizeof(array1));
            memset(array2, 0, sizeof(array2));
            memset(array_total, 0, sizeof(array_total));
            array1_len=0;
            while (1) {
                scanf("%d",&array1[array1_len]);
                if(array1[0]==0){
                    //printf("\n");
                    break;
                    
                }else{
                    if (array1[array1_len]!=0) {
                        array1_len++;
                    }else{
                        break;
                    }
                }
            }
            
            for (int i=0; i<array1_len; i++) {
                if (flag==0) {
                    flag=1;
                    printf("%d",array1[i]);
                }else{
                    printf(" %d",array1[i]);}
            }
            printf("\n");
        }
        
        
    }
    
}
```

刚开始代码这么写，然后怎么改，都是85%wrong answer，过不去。 然后博主一条一条去分析，最后发现数组应该稍微调大点 所以最后的代码是
```c
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
    int num;
    int max;
    int array1 [1002];
    int array2 [1002];
    int array_total [1002];
    int array1_len,array2_len;
    int flag;
    scanf("%d",&num);
    if(num==0){
        return 0;
    }
    for (int j=0; j<num; j+=2) {
        if (num-j>1) {
            flag=0;
            memset(array1, 0, sizeof(array1));
            memset(array2, 0, sizeof(array2));
            memset(array_total, 0, sizeof(array_total));
            array1_len=0;
            while (1) {
                scanf("%d",&array1[array1_len]);
                if(array1[0]==0){
                    //printf("\n");
                    break;
                    
                }else{
                    if (array1[array1_len]!=0) {
                        array1_len++;
                    }else{
                        break;
                    }}
            }
            array2_len=0;
            while (1) {
                scanf("%d",&array2[array2_len]);
                if(array2[0]==0){
                    //printf("\n");
                    break;
                    
                }else{
                    if (array2[array2_len]!=0) {
                        array2_len++;
                    }else{
                        break;
                    }}
            }
            max=getmax(array1_len,array2_len);
            if (max!=0) {
                for (int i=0; i<max; i++) {
                    if (flag==0) {
                        flag=1;
                        printf("%d",array1[i]+array2[i]);
                    }else{
                        printf(" %d",array1[i]+array2[i]);}
                }
            }
            
            printf("\n");
        }else{
            flag=0;
            memset(array1, 0, sizeof(array1));
            memset(array2, 0, sizeof(array2));
            memset(array_total, 0, sizeof(array_total));
            array1_len=0;
            while (1) {
                scanf("%d",&array1[array1_len]);
                if(array1[0]==0){
                    //printf("\n");
                    break;
                    
                }else{
                    if (array1[array1_len]!=0) {
                        array1_len++;
                    }else{
                        break;
                    }
                }
            }
            
            for (int i=0; i<array1_len; i++) {
                if (flag==0) {
                    flag=1;
                    printf("%d",array1[i]);
                }else{
                    printf(" %d",array1[i]);}
            }
            printf("\n");
        }
        
        
    }
    
}
```