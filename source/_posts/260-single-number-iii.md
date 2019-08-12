---
title: 260. Single Number III
url: 330.html
id: 330
comments: false
categories:
  - OJ
  - LeetCode
date: 2016-05-04 01:22:24
tags:
---

Given an array of numbers nums, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once.

For example: 

Given nums = \[1, 2, 1, 3, 2, 5\], return \[3, 5\].

```c++
class Solution {
public:
    vector<int> singleNumber(vector<int>& nums) {
        int temp=0;int a=0,b=0;
        for(int i=0;i<nums.size();i++){
            temp^=nums[i];
        }
        int temp1=temp&~(temp-1);
        for(int i=0;i<nums.size();i++){
            if(temp1&nums[i]){
                a^=nums[i];
            }else{
                b^=nums[i];
            }
        }
        return vector<int>({a,b});
    }
};
```