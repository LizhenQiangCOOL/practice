# [283. 移动零](https://leetcode-cn.com/problems/move-zeroes/)

## 1. 题目描述
给定一个数组 nums ，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

## 2. 示例
```
输入： nums = [0, 1, 0, 3, 12]
输出： [1, 3, 12, 0, 0]
```
```
输入： nums = [0]
输出： [0]
```

## 3. 思考
1. 题目类型为 数组、快慢指针
2. 非零元素是要保持顺序的，零元素是不需要保持顺序