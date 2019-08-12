---
title: 342. Power of Four
url: 324.html
id: 324
comments: false
categories:
  - OJ
  - LeetCode
date: 2016-05-04 01:13:48
tags:
---

Given an integer (signed 32 bits), write a function to check whether it is a power of 4. 
Example: Given num = 16, return true. Given num = 5, return false.
```c++
class Solution {
public:
    bool isPowerOfFour(int num) {
        if(num==1)return true;
        if(num<0)return false;
        if(num<=2)return false;
        int count=0;int location=0;
        for(int i=0;i<32;i++){
            count+=(num>>i)&1;
            location=location==0?(((num>>i)&1)==1?i:0):location;
        }
        if(count!=1)return false;
        if(location%2)return false;
        return true;
    }
};
```