# [105. 从前序与中序遍历序列构造二叉树](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

## 1. 题目描述
给定两个整数数组 preorder 和 inorder, 其中 preorder 是二叉树的先序遍历，inorder 是同一颗树的中序遍历，请构造二叉树并返回其根节点。


## 2. 示例
```
    3
   / \
  9   20
     /  \
    15   7
    
输入： preorder = [3, 9, 20, 15, 7], inorder = [9, 3, 15, 20, 7]
输出： [3, 9, 20, null, null, 15, 7]
```
```
输入： preorder = [-1], inorder = [-1]
输出： [-1]
```


## 3. 思考
1. 该题目主要考二叉树，框架：二叉树的遍历-前序遍历