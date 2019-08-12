---
title: oj之坑－高效筛选素数
url: 150.html
id: 150
comments: false
categories:
  - OJ
  - sdust
date: 2015-11-27 09:21:13
tags:
---

Problem H: 筛选素数

Problem H: 筛选素数
---------------

Time Limit: 1 Sec  Memory Limit: 16 MB
Submit: 3157  Solved: 1158

Description
-----------

在数学上，素数的分布没有任何已知规律，因此检测一个数是否素数，只能用比它小的素数来检测整除性质。如果要求出一定范围内的素数表直接检测的代价就太高了。一般采用筛选法的思想：
把从1开始的、某一范围内的正整数从小到大顺序排列，1不是素数，首先把它筛掉。剩下的数中选择最小的数是素数，然后去掉它的倍数。依次类推，直到筛子为空时结束。
如有：
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30
1不是素数，去掉。剩下的数中2最小，是素数，去掉2的倍数，余下的数是：
3 5 7 9 11 13 15 17 19 21 23 25 27 29
剩下的数中3最小，是素数，去掉3的倍数……如此下去直到所有的数都被筛完，求出的素数为：
2 3 5 7 11 13 17 19 23 29
用筛选的办法求出素数表之后，再进行查询。

求区间m～n之间的素数，其中m,n为整数，且0<=m<=n<=500000。

Input
-----

输入多行，到EOF结束。每行为一个测试用例，两个整数m和n，满足0<=m<=n<=500000。总测试用例数不超过1000个。

Output
------

输出为m～n之间的所有素数，每个数一行，两个测试用例之间用一个空行分隔。

如果m～n之间没有素数，则输出一个空行。

Sample Input
------------

1 2
2 3
8 10
1 11
50 100
900 1000
499900 500000

Sample Output
-------------

2

2
3

2
3
5
7
11

53
59
61
67
71
73
79
83
89
97

907
911
919
929
937
941
947
953
967
971
977
983
991
997

499903
499927
499943
499957
499969
499973
499979

HINT
----

Append Code
-----------

\[[Submit](submitpage.php?cid=1971&pid=7&langmask=1022)\]\[[Status](problemstatus.php?id=1179)\]\[[Web Board](bbs.php?pid=1179&cid=1971)\]


由于博主算法太菜，所以刚开始是百度查的相关代码，用到了数论。 

就是所有的倍数最大达到sqrt(range)+1就行了。然后copy了相关代码进行了修改

```c
#include<stdio.h>
int main(void)
{
	unsigned long long a[500000]={0},i,j,count;
    for(i=2;i<708;i++) /*数学上可证明,只要一直判断到根号下100就可找到100以内的全部素数*/
    { if(a[i-1]==1) continue; /*判断该数是否被删除*/
        for(j=i+1;j<=500000;j++) /*如果没被删除，则将他的整数倍都删除*/
            if(j%i==0) a[j-1]=1;
    }
    int min,max;
    unsigned long long location;
    while(scanf("%d %d",&min,&max)!=EOF){
        location=2;
        int counta=0;
        while(location<=max){
            if (location>=min&&a[location-1]==0)
            {
                printf("%lld\n", location);
                counta++;
            }
            location++;
        }
        printf("\n");
        if (counta==0)
        {
            printf("\n");
        }
    }
    return 0;
}
```

然后博主，提交后得到的结果是 
Time Limit Exceed50% 
然后我就静态打表传上去，然后显示是代码太长无法上传。。。 
最后我修改相关计算素数的算法，提高效率传上去，然后就A了 相关代码在下面
```c
#include<stdio.h>
//#include <time.h>
int main(void)
{
 
    unsigned long long a[500000]={0},i,j,count,count1,tmp;
    int flag;
    for(i=2;i<708;i++)
    { if(a[i-1]==1) continue;
        count1=1;
        flag=0;
        while(1){
            if (flag==0) {
                flag=1;
                count1++;
                continue;
            }
            tmp=count1*i;
            
            if (tmp>500000)
            {
                break;
            }
            a[tmp-1]=1;
            count1++;
        }
    }
 
    int min,max;
    unsigned long long location;
    
    while(scanf("%d %d",&min,&max)!=EOF){
        location=2;
        int counta=0;
        while(location<=max){
            if (location>=min&&a[location-1]==0)
            {
                printf("%lld\n", location);
                counta++;
            }
            location++;
            
        }
        printf("\n");
        if (counta==0)
        {
            printf("\n");
        }
        
        
    }
    return 0;
    
    
    
}
```