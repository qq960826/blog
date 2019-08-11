---
title: 102. Binary Tree Level Order Traversal My Submissions Question
url: 232.html
id: 232
comments: false
categories:
  - leetcode
date: 2016-02-22 13:51:56
tags:
---

Given a binary tree, return the level order traversal of its nodes’ values. (ie, from left to right, level by level).

For example:
Given binary tree {3,9,20,#,#,15,7},
```
    3
   / \
  9  20
    /  \
   15   7
```
return its level order traversal as:
```
[
  [3],
  [9,20],
  [15,7]
]
```
confused what "{1,#,2,3}" means? [\> read more on how binary tree is serialized on OJ.](#)

**OJ's Binary Tree Serialization:** 
The serialization of a binary tree follows a level order traversal, where '#' signifies a path terminator where no node exists below. Here's an example:
```
   1
  / \
 2   3
    /
   4
    \
     5
```

The above binary tree is serialized as "{1,2,3,#,#,4,#,#,5}"

```c++
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
vector<vector<int>> result;
    vector<vector<int>> levelOrder(TreeNode* root) {
        find(root);
        return result;
    }
    void find(TreeNode* root){
        if(!root) return;
        queue<TreeNode *> nodeQueue;
        nodeQueue.push(root);
        vector<int> temp;
        temp.push_back(root->val);
        result.push_back(temp);
        TreeNode* node;
        int size;
		while(!nodeQueue.empty()){
            size=nodeQueue.size();
            temp.clear();

            for(int i=0;i<size;i++){
                node = nodeQueue.front();
                nodeQueue.pop();
                if(node->left){
                nodeQueue.push(node->left);  //先将左子树入队
                temp.push_back(node->left->val);
            }
            if(node->right){
                nodeQueue.push(node->right);  //再将右子树入队
                temp.push_back(node->right->val);
            }

        }
        if(!temp.empty())result.push_back(temp);
    }
};
```