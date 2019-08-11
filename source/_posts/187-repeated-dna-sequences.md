---
title: 187. Repeated DNA Sequences
url: 332.html
id: 332
comments: false
categories:
  - leetcode
date: 2016-05-04 01:24:18
tags:
---

All DNA is composed of a series of nucleotides abbreviated as A, C, G, and T, for example: “ACGAATTCCG”. When studying DNA, it is sometimes useful to identify repeated sequences within the DNA.

Write a function to find all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule.

For example,

Given s = “AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT”,

Return:
[“AAAAACCCCC”, “CCCCCAAAAA”].

```c++
class Solution {
public:
vector<string> findRepeatedDnaSequences(string s) {
    string temp;
    set<string> set_all,set_result;
    vector<string> a;
    if(s==""||s.length()<10)return a;
    for(int i=0;i<s.length()-9;i++){
        
        temp=s.substr(i,10);
        if(set_all.find(temp)!=set_all.end()){//exist
            if(set_result.find(temp)==set_result.end()){
                set_result.insert(temp);
                a.push_back(temp);
                
            }
        }
        set_all.insert(temp);
    }
    return a;
}

};
```





