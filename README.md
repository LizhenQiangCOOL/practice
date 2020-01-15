
# 编程算法题（python解题）

# Letcode

### 数组类题目

|相关问题 | 题意 | 解题思路 | 编程注意
|--- | --- |--- | --- |
|283. 移动零                     |               |    方法：使用双指针   |   
|75. 颜色分类                    |               |    计数排序 遍历两边，  三路快速排序 遍历一遍, 使用两个辅助指针                         |
|88.合并两个有序数组             |               |                               |
|215.数组中的第K个最大(最小)元素 |               |                               |


### 查找表      
- 利用set、map(dict) 底层要求顺序，用二叉搜索树(平衡)，不要求，使用哈希表（散列表）实现

|相关问题 | 题意 | 解题思路 | 编程注意
|--- | --- |--- | --- |
|350. 两个数组的交集            |               |              |   
|454.四数相加二                 |               |              |
|447. 回旋镖的数量               |               |                               |
|215.数组中的第K个最大(最小)元素 |               |                               |




### 链表类
-思路：穿针引线，是否需要设虚头指针，是否要覆盖赋值，是否能双指针（或快慢指针）,断开链表时注意，

|相关问题 | 注意点 
|--- | ---
|206.反转链表                        |
|92. 反转链表二(注意点比较多)        |
|83. 删除排序链表中的重复元素        |
|86. 分割链表                        |
|328. 奇偶链表（该处理的都要处理）   |
|2. 两树相加                         |
|445. 两数相加二                     |
|203. 移除链表元素                   |
|82. 删除排序链表中的重复元素二      |
|21. 合并两个有序链表                |
|24. 两两交换链表中的节点            |
|25. K 个一组翻转链表                |
|147. 对链表进行插入排序             |
|148. 排序链表                       | 归并排序
|237. 删除链表中的节点               |
|19. 删除链表的倒数第N个节点         |
|61. 旋转链表                        |
|143. 重排链表                       |



### 栈和队列、优先队列(堆)

|相关问题 | 题意 | 解题思路 | 编程注意
|--- | --- |--- | --- |
| [20. 有效的括号](https://leetcode-cn.com/problems/valid-parentheses/) |  括号是否匹配 | 利用栈，遇到右括号时就要出一个左口号 | 注意遍历完，栈还不为空
| [150. 逆波兰表达式求值 ](https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/)  | 构造逆波兰表达表达式 |利用栈，把计算值重新存在栈中，最后栈元素为结果 | 注意：两个树是否为负数时，如何整除
| [71. 简化路径  ](https://leetcode-cn.com/problems/simplify-path/)   | 简化文件路径  |利用栈，等于..如有就弹，等于.就忽略，正常路径则加入 |
树遍历，前中后序遍历 144 94 145     |   树的遍历 | 递归很简单，非递归：利用栈，前序则要把右边先压栈； 中序，不停先把根节点压入，在压根节点的左元素，后序就是前序的转换 | 注意，中序的非递归
| [341. 扁平化嵌套列表迭代器 ](https://leetcode-cn.com/problems/flatten-nested-list-iterator/)  |  列表内嵌列表->一个列表       | 使用队列存储，如果是列表则递归调用存储函数|
| [102. 二叉树的层次遍历 ](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)  |    利用队列，注意输出结构
| [107. 二叉树的层次遍历二](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/)  |
| [103. 二叉树的锯齿形层次遍历](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/)  |
| [199. 二叉树的右视图](https://leetcode-cn.com/problems/binary-tree-right-side-view/)  |  找每一层的最后一个元素 | 层次遍历
| [279. 完全平方数  ](https://leetcode-cn.com/problems/perfect-squares/)          |  用最少的完全平方数构造一个函数 |要构建图，层遍历，BFS广度优先搜索; 可背包
| [127. 单词接龙 ](https://leetcode-cn.com/problems/word-ladder/)  | 每次只能换一个字母，能不能从begin单词到end单词，能求最短转换序列   |构建图，然后层遍历
| [347. 前 K 个高频元素](https://leetcode-cn.com/problems/top-k-frequent-elements/)|给一个数组，有数字，根据评率返回前k个 | 利用优先队列(堆)来实现 | 模块heapq 而heapq.heapify无key参数，heapq.nlargest有key参数
| 23. 合并k个有序序列 | 将k个有序数组，归并成一个有序数组 | 时间复杂度最小，归并时应使用k大小的优先队列(堆)，每次拿出最小的只要log(k)时间复杂度 |  可以使用heapq.merge(*iterables(sort), key, reverse=False)




### 二叉树和递归
- 是递归(中序遍历)解决，还是 BFS 利用队列 或 DFS 利用栈
- 如果题目要求是有路径，应考虑递归+回溯，用path记录数组

|相关问题 | 题意 | 解题思路 | 编程注意
|--- | --- |--- | --- |
| [111. 二叉树的最小深度](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/) |二叉树的最小深度，深度指根到叶子节点 |递归：用前序遍历，到叶子节点就把深度记录；层序遍历：遍历每一层，遇到None，返回该层 | 注意递归终止条件是叶子节点
| [226. 翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/) | 翻转一棵二叉树。即输入一棵树，输出这个树的镜像 |递归：使用前序遍历进行左右子树交换 | 
| [100. 相同的树](https://leetcode-cn.com/problems/same-tree/submissions/) | 递归：前序遍历判断是否相同 |注意终止条件，都为空返回True
| [101. 对称二叉树](https://leetcode-cn.com/problems/symmetric-tree/submissions/)| 判断一棵树是否对称，即是否镜像 | 递归：前序遍历，左子树与右子树判断  | 注意终止条件
| [222. 完全二叉树的节点个数](https://leetcode-cn.com/problems/count-complete-tree-nodes/submissions/) | 求一个完全二叉树的节点个数 | 递归：前序遍历，累计个数|
| [110. 平衡二叉树](https://leetcode-cn.com/problems/balanced-binary-tree/submissions/) | 判断一棵树是否为平衡二叉树 | 要一个辅助函数，判断高度。主函数：判断左右子树高度差，判断左子树是否平衡，判断右子树是否平衡 | 这里不用前序遍历，只要一个辅助函数求高度
| [112. 路径总和](https://leetcode-cn.com/problems/path-sum/submissions/) | 二叉树，判断是否存在节点和 为 sum的路径 | 递归：前序遍历，sum减去当前的值； 如果要打印路径 用栈来做 | 注意 判断该节点是否为叶子点
| [404. 左叶子之和](https://leetcode-cn.com/problems/sum-of-left-leaves/submissions/)| 二叉树，左叶子的和         | 递归：前序遍历，是左叶子累加，给出标记说明这个节点是上一节点的左子树还是右子树 |  注意：根节点不是左叶子
| [257. 二叉树的所有路径](https://leetcode-cn.com/problems/binary-tree-paths/submissions/)| 给定一个二叉树，返回所有从根节点到叶子节点的路径。| 方法一：要求返回路径，使用递归（前序遍历）+回溯，用一个栈保存节点，某次递归函数结束都要回溯(删除节点)（可转化为循环，stack要保存两项一个是节点，一个该节点对应的路径）；方法二：完全递归，每一递归函数返回它的所有路径，是空节点，直接返回当前只有一个元素的列表，不是空节点，遍历左子树，遍历右子树，把结果都整合 |  注意：回溯要用栈保存，遍历树还是用递归； 方法二完全依赖递归
| [113. 路径总和 II](https://leetcode-cn.com/problems/path-sum-ii/submissions/)| 二叉树，返回所有节点和 为 sum的路径 |方法一：完全递归 ； 方法二：递归（前序遍历）+回溯法 |注意：两种方法的区别
| [129. 求根到叶子节点数字之和](https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/) | 把路径数字拼接，返回其和 | 变成求出所有路径；方法一：完全递归，方法二：递归+回溯 |
| [437. 路径总和 III](https://leetcode-cn.com/problems/path-sum-iii/)| 路径和只要在树中从上到下即可  | 方法一：各个节点都当成根节点，问题转化；方法二 回溯存储，反向依次 看和是否相等，相等结果+1 |
| [235. 二叉搜索树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)| 二分搜索树中，找两个节点的最近公共祖先节点 | 利用二分搜索树的性质，当前节点大于（小于）这两个节点，则公共祖先节点在左边（右边）；当前节点在中间，则就是公共祖先节点| 注意在中间可以是等于
| [98. 验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree/) |验证一个二叉树是否为二分搜索树 | 中序遍历，递归 或者 非递归的中序遍历| 注意：如何写非递归的中序遍历，有点类似深度遍历，但不一样 循环内嵌了循环
| [94. 二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)| 二叉树的中序遍历 | 写非递归写法，栈，一直压当前元素的左边元素 |
| [450. 删除二叉搜索树中的节点](https://leetcode-cn.com/problems/delete-node-in-a-bst/)| 二叉搜索树删除指定节点 | 这如果要求时间复杂O(h)就应该 使用递归，
| [108. 将有序数组转换为二叉搜索树](https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/submissions/)| 有序数组转换成一颗高度平衡的二叉搜索树 | 递归构建树 |
| [230. 二叉搜索树中第K小的元素](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/)|给定二叉搜索树，找第k个小的元素 | 中序遍历，递归；可以记录一个前一个值； 也可以把这个值当函数返回，要参考剑指offic  | 注意：k为全局的

- 删除二叉搜索树中节点，并且无父亲节点
```

class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    #deleteNode(root, key)，返回删除指定节点后的树的根节点。
    #在递归过程中，是一边删除一边构造二叉树
    def deleteNode(self, root, key):
        if root is None:
            return None

        if key < root.val:
            root.left = self.deleteNode(root.left, key)
            return root

        if key > root.val:
            root.right = self.deleteNode(root.right, key)
            return root
        
        # 下面就是 該root要刪除
        if root.left is None:
            new_root = root.right
            root.right = None
            return new_root

        if root.right is None:
            new_root = root.left
            root.left = None
            return new_root

        # 找到右子树中最小的结点，复制它的值
        successor = self.__minimum(root.right)
        successor_copy = TreeNode(successor.val)
        successor_copy.left = root.left
        successor_copy.right = self.__remove_min(root.right)
        root.left = None
        root.right = None
        return successor_copy

    def __remove_min(self, node):
        if node.left is None:
            new_root = node.right
            node.right = None
            return new_root
        node.left = self.__remove_min(node.left)
        return node

    def __minimum(self, node):
        while node.left:
            node = node.left
        return node
```

### 递归与回溯

|相关问题 | 题意 | 解题思路 | 编程注意
|--- | --- |--- | --- |
| [17. 电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/submissions/) | 给定数字，返回它能代表字母集合(手机) | 这是一个树形结构，用递归 记录下标，一个当前路径，一个结果集合 | 注意 数字范围和空空字符串
| [93. 复原IP地址](https://leetcode-cn.com/problems/restore-ip-addresses/submissions/) | 给定一个包含数字的字符串，复原它并返回所有可能的IP地址格式 | 使用递归，一个记录当前，一个结果集合 | 注意 保证判断字符串否为空，‘0’情况，局部字符串大于'255'
| [131. 分割回文串](https://leetcode-cn.com/problems/palindrome-partitioning/submissions/) | 给定一个字符串，分割成子串，使每个子串都是回文串 | 递归，回溯，用index记录下标，pre记录前面
| [46. 全排列](https://leetcode-cn.com/problems/permutations/)  | 给没有重复数字的序列，返其所有可能的全排列 | 交换回溯，当前，后面数字 ；递归，回溯 选当前，记录之前 | 注意：交换回溯要 恢复原来状态


### 动态规划类
- 先思考递归，求递归推导式，优化成动态规划从低往上。主要考虑 第i项怎么依赖0到i-1项
- 状态机：如何维护这个状态多维数组,如dp[i][k][j]
- 0-1背包：F(n, C)表示n个物品放进容量为C的背包，使得价值最大，`F(i, C) = max( F(i-1, C), v(i)+F(i-1, c-w(i) )` 放与不放
- 完全背包：每个物品可以无限使用，
- 多重背包问题： 每个物品不止1个，有num(i)个
- 多维费用背包，背包物品约束

|相关问题 | 题意 | 解题思路 | 编程注意
|--- | --- |--- | --- |
| [70. 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)     | n 阶，每次可以爬1或2台阶，多少种方法          | 推导式：f(n) = f(n-1)+f(n-2)                      | n小于0时特殊处理
| [120. 三角形最小路径和](https://leetcode-cn.com/problems/triangle/) | 给定一个三角形，找出自顶向下的最小路径和      | 非第一和最后时，f(i, j)依赖f(i-1, j-1)和f(i-1, j) | 为第一和最后一个时 f(i,j)依赖项
| [64. 最小路径和](https://leetcode-cn.com/problems/minimum-path-sum/)| MxN网格，从左上角到右下角，最小路径和         | f(i, j) 依赖于f(i, j-1)和f(i-1, j)                | 注意 i,j特殊情况，为0时
| [343. 整数拆分](https://leetcode-cn.com/problems/integer-break/)    |整数N，拆分多个整数，这些整数的最大乘积        | f(N) 依赖于 max(f(i), i) 和 max(f(j), j), i+j=N   | 注意  内循环j(优化)的范围为[1~i>>1
| [279. 完全平方数](https://leetcode-cn.com/problems/perfect-squares/)|正整数 n，找最少个完全平方数使得它们的和等于 n | `dp[i] = MIN(dp[i - j * j] + 1)`，i表示当前数字，j*j表示平方数 dp[0]=0 | 注意的dp[0]
| [322. 零钱兑换](https://leetcode-cn.com/problems/coin-change/)      |一个总金额 amount 换 无限不同面额的硬币 coins，个数最少，无方案返回-1 | `dp[n] = min(dp[n - c1], ... dp[n - cn]) + 1`, dp[0]=0| 注意 最小值初始`float('inf')`
| [91. 解码方法](https://leetcode-cn.com/problems/decode-ways/)       |给一字符串数字 解码成字母（A-Z),有多少种       | `dp[i] = dp[i-1](能单独成为字母) + dp[i-2](能和前一个组成字母)` | 注意：给的字符串以0开头是没结果的
| [198. 打家劫舍](https://leetcode-cn.com/problems/house-robber/submissions/)|一数组，选不相邻元素，最大和            | `dp[i] = max(dp[i-2]+nums[i], dp[i-1])`，是否偷第i个，dp[n]为结果|  
| [213. 打家劫舍II](https://leetcode-cn.com/problems/house-robber-ii/submissions/)|一组数成环，选不相邻元素，最大和    | 和打家劫舍一样                                    | 拿第一个不能拿最后一个，拿最后一个不能拿第一个，把它们分别去掉变成打家劫舍  
| [337 打家劫舍三](https://leetcode-cn.com/problems/house-robber-iii/) | 树形结构，拿父节点，不能拿其直属的孩子节点    | 树形递归解决dp[0,1],dp[0]表示不选父，dp[1]表示选父|可递归动态规划，也可BFS广度搜索，变成数组回归 大家劫舍   
| [309. 最佳买卖股票时机含冷冻期](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)| 卖出股票就可以买，可买多次，最大收益| 状态机：dp[i][k][0]第i天，总买卖k次不持有股票 dp[i][k][1]第i天，总买买k次持有股票 |利用这个状态机可以解决一些列股票买卖问题
| [121. 买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/)| [122. 买卖股票的最佳时机II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)|[123. 买卖股票的最佳时机III ](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/)|[188. 买卖股票的最佳时机 IV](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)
| [300. 最长上升子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/) | LIS,数组，求上升子序列，不一定连续 |  `f(n) = max(1, f(i)+1) i<n, 且num[i]<num[n]`  | 初始化所有元素长度为1，以自己结尾, 最优是dp 二分查找一个元素位置，若该元素
| [376. 摆动系列](https://leetcode-cn.com/problems/wiggle-subsequence/submissions/) | 数组，求最长摆动系列  | 维护两个dp，一个序列处于上升，一个序列处于下降。 `f(n) = max(1, f(i)+1) i<n` | 注意down和up区别
| LCS 最长公共子序列   |    LCS | LCS(i,j) = 当a[i]=b[j], 1+LCS(i+1,j+1) 当a[i]!=b[j], max(LCS(i,j+1), LCS(i+1,j)) | 明白子串要求连续，而子序列不要求连续
| Dijkstra单源最短路径 | 一点到另外一点最短路径   | shorttestPath(x)= min(shortestPath(a) + w(a->x)) |
| [416. 分割等和子集](https://leetcode-cn.com/problems/partition-equal-subset-sum/submissions/)| 一个数组，平均分两个sum相等的集合 | dp[0..sum//2],下标表示背包容量，值表示背包是否装满. `dp[j] = dp[j](原本背包已满) or dp[j-nums[i]] (看背包变小时，是否满)` |  注意：数组下标表示背包容量，遍历数组，考虑是否放入，背包容量为j dp[j],j>=nums[i]
| [377. 组合总和 IV](https://leetcode-cn.com/problems/combination-sum-iv/) | 给定一个由正整数组成且不存在重复数字的数组，找出和为给定目标正整数的组合的个数。| f(i) = sum(f(i-num[k])), i>=num[k] | 做之前观察例子,第一个nums[k],剩下序列就是f(i-num[k])
| [474. 一和零](https://leetcode-cn.com/problems/ones-and-zeroes/) | 给多个01串，给M个0和N个1，最多能装多少个01串 | `dp[i][j]`看不同大小的背包，限制i为1的个数，j为0的个数；对每一个字符串，对每个可放的进去的背包：`dp[i][j] = max(dp[i][j], 1+dp[i- count1][j-count1]`| 不要把二维数组看成一个大背包，而是把二维数组每个`dp[i][j]`,看成一个背包
| [139. 单词拆分](https://leetcode-cn.com/problems/word-break/)| 对于一字符串，能否拆成一个或多个在字典出现的单词 | dp[i], 表示长度为i的字符串 是否装了单词，`dp[i] = dp[i] or dp[i-len(k)]`, k表示单前单词，能否装进去 | 一位数组，每个值看成一个背包， 对于一个单词，遍历能完好装得下它的背包


## 剑指office 

#### ==面试题3==：数组中重复的数字
题目一：找出数组中重复的数字：在一个长度为n的数组里所有数字都在0~n-1的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了多少次。请找出数组中任意一个重复的数字。例如,如果输入长度为7的数组[2,3,1,0,2,5,3],那么对应输出的数字2或者3
- 方法一：排序，遍历一遍。 时间复杂度O（nlgn），空间复杂度O（1）
- 方法二：哈希表（字典法），用数组建立一个哈希表，或者直接用py的dict字典，遍历一遍。时间复杂度O（n），空间复杂度O（n）
> - 最优方法：巧妙的。值为i的 数字应该在下标i，除非已被占，意思该值已重复，否则交换值i到下标i，看重新得的值是否满足条件(2个)。时间复杂度O(n),空间复杂度O(1)
```
#这里输出所有重复的数字，题目中数字范围0~n-1十分关键，这样值和下标能对应

def duplicate(arr):
    if not isinstance(arr, list) or len(arr)==0:
        return False, 0
    for item in arr :
        if item<0 or item>len(arr)-1:
            return False,0
    for i in range(len(arr)):
        while arr[i] != i :
            if arr[i] == arr[ arr[i] ]:
                return True,arr[i]
            print(arr)
            x = arr[i]
            arr[i],arr[x] = arr[x],arr[i]
    return False,0

```
题目二：不修改数组找出重复的数字：在一个长度为n+1的数组里的所有数字都在1~n的范围内，所以数组中至少有一个数字是重复的。请找出数组中任意一个重复的数字，但不能修改输入的数组。例如，如果输入长度为8的数组{2, 3, 5, 4, 3, 2, 6, 7}, 那么对应的输出是重复的数字2或者3
```
#这里重复数字的 次数可能多次， 分组看arr范围中个数 与 范围差的比较
def getDuplication(arr):
    if not isinstance(arr, list) or len(arr)==0:
        return -1
    
    start = 1
    end =len(arr) + 1
    while end>=start:
        middle = ((end-start)>>1) + start
        count = countRange(arr, start, middle)
        if end == start:
            if count>1:
                return start
            else:
                break
        if count > (middle - start + 1):
            end = middle
        else:
            start = middle + 1
    return -1

def countRange(arr, start, end):
    return len([x for x in arr if x>=start and x<=end])
```

###### 测试用例
- 长度为n的数组里包含一个或多个重复的数字
- 数组中不包括重复的数字
- 无效输入测试用例（输入空指针；长度为n的数组中包含0~n-1之外的数字）

#### ==面试题4==：二维数组中的查找
题目：在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

- 方法一：二分找上界，二分找下界，再二分查找
> - 方法二：利用已排序的性质，对比，缩小范围
```
#方法一
class Solution:
    # array 二维列表
    def Find(self, target, array):
        def getup(arr, left, right, k, cmp):  # 找上界
            while(left < right):
                mid = (left + right + 1) // 2
                if cmp(arr[mid][0], k):
                    left = mid
                else:
                    right = mid - 1
            return left
        def getdown(arr, left, right, k, cmp):  # 找下界
            while(left < right):
                mid = (left + right) // 2
                if cmp(arr[mid][len(arr[mid])-1], k):
                    right = mid
                else:
                    left = mid + 1
            return left
        
        def  erfen(arr,left,right,k):
            while left < right:
                mid = (left + right + 1) // 2
                if arr[mid] == k:
                    return k
                elif arr[mid]< k:
                    left = mid+1
                else:
                    right = mid -1
            return -1
        left = getdown(array,0,len(array)-1,target,lambda a,b:a>=b)
        right = getup(array,0,len(array)-1,target,lambda a,b:a<=b)
        for i in range(left,right+1):
            if erfen(array[i],0,len(array[i])-1,target)!=-1:
                return True
        return False
# 方法二
class Solution:
    def Find(self, target, array):
        if not isinstance(target, int) or not isinstance(array, list) or len(array) == 0:
            return False
        col = 0
        row = len(array)-1
        while col<len(array[0]) and row >=0:
            if array[row][col] == target:
                return True
            elif array[row][col] >target:
                row -= 1
            else:
                col += 1
        return False
```
###### 测试用例
- 二维数组中包含查找的数字
- 二维数组中没有查找的数字
- 特殊输入测试（数组为空）

#### ==面试题5==：替换空格
题目：请实现一个函数，把字符串中每个空格替换成%20，例如，输入"We are happy"，则输出"We%20are%20happy"
- 方法：数组手动遍历，空格换%20。或者直接调用字符串.replace
```
def replaceSpace(s):
        return s.replace(" ", "%20")
#手动走
def replaceSpace(self, s):
    L = list()
    for item in s:
        if item ==" ":
            L.append("%20")
        else:
            L.append(item)
    return "".join(L)
```
###### 测试用例
- 输入的字符串包含空格（空格位置前中后，空格一个或连续多个）
- 输入字符串没有空格
- 特殊输入（空字符串，空格字符串、空格一个或连续多个）


#### ==面试题6==：从尾到头打印链表
题目：输入一个链表的头节点，从尾到头反过来打印出每个节点的值。
- 方法一：存储逆置，打印
- 方法二：利用栈或递归，实现
```
# -*- coding:utf-8 -*-
# class ListNode:
#     def __init__(self, x):
#         self.val = x  
#         self.next = None
        # self.pre = None

class Solution:
    # 返回从尾部到头部的列表值序列，例如[1,2,3]
    #方法一存储逆置
    def printListFromTailToHead(self, listNode):
        l = []
        while listNode:
            l.append(listNode.val)
            listNode = listNode.next
        l.reverse()
        return l
    
    #方法二利用栈或递归，实现
    def printListFromTailToHead(self, listNode):
        def ListReverse(listNode,L):
            if listNode:
                if listNode.next:
                    ListReverse(listNode.next,L)
                L.append(listNode.val)
        L = list()
        ListReverse(listNode,L)
        return L
```
###### 测试用例
- 功能测试(输入的链表有多个节点；输入的链表只有一个节点)
- 特殊输入测试(输入的链表头节点指针为nullptr)

####  ==面试题7==：重构二叉树
题目：输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。例如，输入前序遍历序列{1，2，4，7，3，5，6，8}和中序遍历序列{4，7，2，1，5，3，8，6}，则重建并输出它的头节点
- 方法：分治法，分左子树右子树递归，
```
# -*- coding:utf-8 -*-
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    # 返回构造的TreeNode根节点
    def reConstructBinaryTree(self, pre, tin):

        def rcBinaryTree(preL,tinL):
            preL_len = len(preL)
            if preL_len == 0:
                return None
            elif preL_len == 1:
                return TreeNode(preL[0])
            else:
                T = TreeNode(preL[0])
                root_index = tinL.index(preL[0])
                T.left = rcBinaryTree(preL[1:1+(root_index)],tinL[:root_index])
                T.right = rcBinaryTree(preL[1+(root_index):],tinL[root_index+1:])
                return T
        return rcBinaryTree(pre,tin)
```
###### 测试用例
- 普通二叉树(完全二叉树；不完全二叉树)
- 特殊二叉树(所有节点都没有右子节点的二叉树；所有节点都没有左子节点的二叉树；只有一个节点的二叉树)
- 特殊输入测试(二叉树的根节点指针为nullptr；输入的前序遍历和中序遍历序列不匹配)


#### ==面试题8==：二叉树的下一个节点
题目：给定一课二叉树和其中的一个节点，如何找到中序遍历序列的下一个节点？树中节点除了有两个分别指向左、右子节点的指针，还有一个指向父节点的指针
```
class TreeNode:
    def __init__(self,x):
        self.val = x
        self.p = None
        self.left = None
        self.right= None

class Solution:
    # 返回pNode中序遍历中下一个节点
    def NextTreeNode(self,TNode):
        def final_left(TNode):
            while TNode.left:
                TNode = TNode.left
            return TNode
        
        if TNode.right:
            return final_left(TNode.right)
        else:
            while TNode.p and TNode.p.right == TNode:  #找到节点位于left
                TNode = TNode.p
            return TNode.p
```


#### ==面试题9==：用两个栈实现一个队列
题目：用两个栈实现一个队列。队列的声明如下，请实现它的两个函数appendTail和deleteHead，分别完成在队列尾部插入节点和在队列头部删除节点的功能
- 方法：入队：栈一 push ， 出队：栈二 不为空 pop，为空 栈一pop所有元素 至 栈二push
```
# -*- coding:utf-8 -*-
class Solution:
    def __init__(self):
        self.stack1 = []
        self.stack2 = []
       
    def push(self, node):
        self.stack1.append(node)
        
    def pop(self):
        if self.stack2 == []:
            while self.stack1:
                self.stack2.append(self.stack1.pop())
        return self.stack2.pop()
```
**类似题目：**
- 两个队实现一个栈

方法：入栈：不为空的队push ， 出栈：不为空的队逐一出队到位空的队，只剩下一个pop

#### ==面试题10==：斐波那契数列
题目一：求斐波那契数列的第n项。写一个函数，输入n，求斐波那契数列的第n项，斐波那契数列定义如下：f(0) = 0 , f(1) = 1 , f(n) = f(n-1) + f(n-2)
```
def Fibonacci(n:int)->[int,None]:
    if n<0:
        return None
    elif n<2:
        return [0,1][n]
    fibNMinusOne = 1
    fibNMinusTwo = 0
    fibN = 0
    for i in range(2,n+1):
        fibN = fibNMinusOne+fibNMinusTwo
        fibNMinusTwo,fibNMinusOne = fibNMinusOne,fibN
    return fibN
```
题目二：青蛙跳台阶问题：一只青蛙一次可以跳1级台阶，也可以跳上2级台阶。求该青蛙跳上一个n级的台阶总共有多少种跳法
```
def jumpFloor(number:int)->int:
    if number<1:
        return 0
    elif number<3:
        return [1,2][number-1]
    fibOne = 2
    fibTwo = 1
    fibN = 0
    for i in range(3,number+1):
        fibN = fibOne+fibTwo
        fibTwo,fibOne = fibOne,fibN
    return fibN
```
题目三：青蛙变态跳台阶问题：青蛙一次可以跳1或2或3...n级台阶
```
#数学归纳证明f(n)= 2**(k-1)
#即证：当n>=2,2**0+2**1+...2**(k-2)+1 = 2**(k-1)

def jumpFloorII(number:int)->int:
        if number<1:
            return 0 
        else:
            return pow(2,number-1)
```
- 用2x1的小矩形去覆盖2xN 的大矩形，总共有多少种方法
```
def rectCover(number:int)->int:
        if number<1:
            return 0
        elif number<3:
            return [1,2][number-1]
        fibOne = 2
        fibTwo = 1
        fibN = 0
        for i in range(3,number+1):
            fibN = fibOne+fibTwo
            fibTwo,fibOne = fibOne,fibN
        return fibN
```

#### ==面试题11==：旋转数组的最小数字
题目：把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组{3，4，5，1，2}为{1，2，3，4，5}的一个旋转，该数组的最小值为1
```
# -*- coding:utf-8 -*-
class Solution:
    def minNumberInRotateArray(self, rotateArray):
        if len(rotateArray)==0:
            return 0
        left = 0
        right = len(rotateArray)-1
        middle= left
        while rotateArray[left]>= rotateArray[right]:
            if right-left == 1:
                middle = right
                break
            middle = (left+right)//2
            if rotateArray[middle] > rotateArray[left]:
                left = middle
            elif rotateArray[middle] < rotateArray[right]:
                right = middle
            else:#三者相等
                left+=1
        return rotateArray[middle]
```

#### ==面试12==：矩阵中的路径
题目：请设计一个函数，用来判断在一个矩阵中是否存在一条包括某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3x4的矩阵中包含一条字符串“bfce”的路径。但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二格子之后，路径不能再次进入这个格子。

a | b | t | g 
---|---
c | f | c | s
j | d | e | h

- 回溯法
```
# -*- coding:utf-8 -*-
class Solution:
    def hasPath(self, matrix, rows, cols, path):

        visited = [False]*(rows*cols)
        pathLength = 0
        for i in range(rows):
            for j in range(cols):
                if self.hasPathCore(matrix,rows,cols, i,j ,path,pathLength,visited):
                    return True
        return False

    def hasPathCore(self, matrix,rows,cols, i,j ,path,pathLength,visited):
        if pathLength == len(path):
            return True
        
        hasPathFlag = False
        if i>=0 and i<rows and j>=0 and j<cols and matrix[i*cols+j] == path[pathLength] and not visited[i*cols+j]:
            pathLength+=1
            visited[i*cols+j] = True
            hasPathFlag = self.hasPathCore(matrix,rows,cols,i,j-1,path,pathLength,visited) or self.hasPathCore(matrix,rows,cols,i-1,j,path,pathLength,visited) or self.hasPathCore(matrix,rows,cols,i,j+1,path,pathLength,visited) or self.hasPathCore(matrix,rows,cols,i+1,j,path,pathLength,visited)
            if not hasPathFlag:
                pathLength-=1
                visited[i*cols+j] = False
        return hasPathFlag
```
#### ==面试题13==：机器人的运动范围
题目：地上有一个m行n列的方格。一个机器人从坐标(0, 0)的格子开始移动， 它每次可以向左、右、上、下移动一格，但不能进入行坐标和列坐标的位数之和大于k的格子。例如，当k=18时，机器人能够进入方格(35, 37)，因为3+5+3+7=18。但它不能进入方格(35,38), 应为3+5+3+8=19.请问该机器人能够到达多少个格子？
- 回溯法
```
# -*- coding:utf-8 -*-
class Solution:
    def movingCount(self, threshold, rows, cols):
        if threshold<0 or rows<=0 or cols<=0:
            return 0
        visited = [False]*(rows*cols)
        count = self.movingCountCore(threshold,rows,cols,0,0,visited)
        return count
    def movingCountCore(self,threshold,rows,cols,i,j,visited):
        count = 0
        if self.check(threshold,rows,cols,i,j,visited):
            visited[i*cols+j] = True
            count = 1+self.movingCountCore(threshold,rows,cols,i-1,j,visited)+self.movingCountCore(threshold,rows,cols,i,j-1,visited)+self.movingCountCore(threshold,rows,cols,i+1,j,visited)+self.movingCountCore(threshold,rows,cols,i,j+1,visited)
        return count
    def check(self,threshold,rows,cols,i,j,visited):
        if i>=0 and i<rows and j>=0 and j<cols and self.getDigitSum(i)+self.getDigitSum(j)<=threshold and not visited[i*cols+j]:
            return True
        return False
    def getDigitSum(self,number):
        sum = 0
        while number>0:
            sum += number%10
            number= number//10
        return sum
        # write code here
```

#### ==面试题14==：剪绳子
题目：给你一根长度为n的绳子，请把绳子剪成m段（m、n都是整数，n>1 并且m>1)，每段绳子的长度记为k[0],k[1]...k[m].请问k[0]xk[1]x...xk[m]可能的最大乘积是多少？例如，当绳子的长度是8时，我们把他剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
- 动态规划：求最优解 、整体问题最优解依赖于各子问题最优解 、子问题有重叠更小子问题 、从上往下分析问题,从下往上求解问题、子问题最优解存储。  
- 贪心算法：当前最优导致局部最优，一般需要用数学方法证明正确性
```
#动态规划 时间复杂度O(n**2) 空间复杂度O(n)
def maxProductAfterCutting_solution(length):
    if length<2:
        return 0
    if length==2:
        return 1
    if length==3:
        return 2
    
    products = [0]*(length+1)
    products[1]=1
    products[2]=2
    products[3]=3
