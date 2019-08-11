---
title: 191. Number of 1 Bits
url: 322.html
id: 322
comments: false
categories:
  - leetcode
date: 2016-05-04 01:12:29
tags:
---

Write a function that takes an unsigned integer and returns the number of ’1' bits it has (also known as the Hamming weight). 

For example, the 32-bit integer ’11' has binary representation 00000000000000000000000000001011, so the function should return 3.
```c++
class Solution {
public:
    int hammingWeight(uint32_t n) {
        int count=0;
            while(n){
                count+=n&1;
                n>>=1;
            }
        return count;
    }
};
```