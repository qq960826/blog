---
title: 190. Reverse Bits
url: 318.html
id: 318
comments: false
categories:
  - leetcode
date: 2016-05-04 01:08:55
tags:
---

Reverse bits of a given 32 bits unsigned integer. 

For example, given input 43261596 (represented in binary as 00000010100101000001111010011100), return 964176192 (represented in binary as 00111001011110000010100101000000).
```c++
class Solution {
public:
    uint32_t reverseBits(uint32_t n) {
        uint32_t temp=n,result=0;
        for(int i=0;i<32;i++){
            result=(result<<1)+(temp&1);
            temp>>=1;
        }
        return result;
    
    }

};
```