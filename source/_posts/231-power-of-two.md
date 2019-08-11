---
title: 231. Power of Two
url: 320.html
id: 320
comments: false
categories:
  - leetcode
date: 2016-05-04 01:11:01
tags:
---

Given an integer, write a function to determine if it is a power of two.
```c++
class Solution {
public:
    bool isPowerOfTwo(int n) {
        if(n==1)return true;
        if(n&1||n<=0)return false;
        int count=0;
        while(n){
            count+=n&1;
            n>>=1;
        }
        if(count==1)return true;
        return false;
    }
};
```