---
title: 268. Missing Number
url: 328.html
id: 328
comments: false
categories:
  - OJ
  - LeetCode
date: 2016-05-04 01:21:06
tags:
---

Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array. 

For example, 
Given nums = \[0, 1, 3\] return 2.
```c++

class Solution {
public:
    int missingNumber(vector<int>& nums) {
        int sum1=0,sum2=0;
        int result=0;

        for(int i=0;i<nums.size();i++){
            result^=i;
            result^=nums[i];
        }
        result^=nums.size();
        return result;
    }
};
```