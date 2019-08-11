---
title: 201. Bitwise AND of Numbers Range
url: 316.html
id: 316
comments: false
categories:
  - leetcode
date: 2016-05-04 01:06:19
tags:
---

Given a range \[m, n\] where 0 <= m <= n <= 2147483647, return the bitwise AND of all numbers in this range, inclusive.
For example, given the range \[5, 7\], you should return 4.
``` c++
class Solution {
public:
    int rangeBitwiseAnd(int m, int n) {
        if(m==n)return m;
        int temp=1<<31;
        int result=0;
        for(int i=0;i<31;i++){
            if((m&temp)==(n&temp))result|=m&temp;
            temp>>=1;
        }
        return result; 
    }

};
```