---
title: Palindrome Linked List My Submissions Question
url: 170.html
id: 170
comments: false
categories:
  - leetcode
date: 2015-12-13 15:12:42
tags:
---

Given a singly linked list, determine if it is a palindrome.    
```c
bool isPalindrome(struct ListNode* head) {
    int a [1000000];
    //ListNode *p=head;
    int count=0;
    while(head){
        a[count]=head->val;
        
        head=head->next;
        count++;
    }
    int flag=0;
    int cishu=count/2;
    for(int i=0;i<cishu;i++){
        if(a[i]!=a[count-i-1]){
            flag=1;
            break;
        }
    }
    if(flag==1){
        return false;
    }
    return true;
}
```