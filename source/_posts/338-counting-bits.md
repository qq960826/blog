---
title: 338. Counting Bits
url: 326.html
id: 326
comments: false
categories:
  - OJ
  - LeetCode
date: 2016-05-04 01:16:12
tags:
---

Given a non negative integer number num. For every numbers i in the range 0 ≤ i ≤ num calculate the number of 1's in their binary representation and return them as an array. 

Example: 
For num = 5 you should return \[0,1,1,2,1,2\].
```c++
class Solution {
public:
    vector<int> countBits(int num) {
        int temp,count;
        vector<int>result;
        for(int i=0;i<=num;i++){
            temp=i;
            count=0;
            while(temp){
                count+=temp&1;
                temp>>=1;
            }
            result.push_back(count);
            
            
        }
        return result;
        
    }
};
```