
# 编程算法题（python解题）

## Leetcode

### 数组类题目

|相关问题 | 题意 | 解题思路 | 编程注意
|---|---|---|---|
| [283. 移动零](https://leetcode-cn.com/problems/move-zeroes/) |　把数组的0移动到数组末尾，保持原来相对顺序　| 使用双指针，快指针往后搜索非零元素，慢指针站位等待交换 |　交换前可附加双指针是否相等判断
| [75.　颜色分类](https://leetcode-cn.com/problems/sort-colors/) | 数组中有三种颜色，原地对它排序，使相同颜色相邻，并按红、白、蓝顺序，保持原来相对顺序 | 计数排序 遍历两边； 快速排序patition（只有三种元素简化版），三路快速排序（维护三个指针，一头一尾一遍历）　遍历一遍 |　注意计数排序要保持稳定，最后找位置应从后往前
| [88.合并两个有序数组](https://leetcode-cn.com/problems/merge-sorted-array/) | 把有序数组num2（len=N），合并到有序数组num１(len=M+N, 前M为自身元素) |　归并排序merge，双指针，用更少的辅助空间可从后往前搜索（维护三指针）|　对题意因地适宜，从后往前三指针(最优)
| [215. 数组中的第K个最大(最小)元素](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/)  |  在未排序数组中找到第k个最大元素  | 　可构建最大堆，弹出第k个最大（ 时间N*logk ）; 使用快速排序的patition, 第k最大则其主元右边有k-1数比它大（时间平均 N 最坏N**2）|  注意时间复杂度分析

```python
# 计数排序：范围[0～K), A源数组 => B结果数组, 辅助数组C
def counting_sort(A, B, k):
    C = [0] * k
    for i in range(len(A)):
        C[ A[i] ] += 1
    for i in range(1, k):
        C[i] += C[ i-1 ]
    #从后往前，保证排序的稳定性
    for i in range(len(A)-1, -1, -1):
        B[ C[A[i]] ] = A[i]
        C[A[i]] -= 1

# 快速排序中partition
# A[p, r] =>   左边<=x[i+1]<=右边
def partition(A, p, r):
    x = A[r]
    i = p-1
    for j in range(p, r):
        if A[j] <= x:
            A[i+1], A[j] = A[j], A[i+1]
            i += 1
    A[i+1], A[r] = A[r], A[i+1]
    return i+1

# 归并排序merge
# A[p, q]合并A[q+1, r] => A[p, r]
def merge(A, p, q, r):
    n1 = q - p + 1
    n2 = r - (q+1) +1
    L = [0] * (n1 +1) #最后一位放哨兵
    R = [0] * (n2 +1)
    for i in range(n1):   #L = A[p:q+1]
        L[i] = A[p + (i-1)]
    for j in range(n2):   #R = A[q+1:r+1]
        R[j] = A[(p+1) + (j-1)]
    L[n1] = float('inf') #哨兵
    R[n2] = float('inf')
    i = j = 0
    for k in range(p, r+1):
        if L[i] <= R[j]:
            A[k] = L[i]
            i += 1
        else:
            A[k] = R[j]
            j += 1

# 堆排序 heapq
//父节点下标
Parent(i)
    return 向下取整( i/2 )
//左子节点下标
Left(i)
    return 2i
//右子节点下标
Right(i)
    return 2i+1

//维护最大堆
Max-Heapify(A,i)
    l = Left(i)
    r = Right(i)
    if l <= A.heap-size and A[l] > A[i]
        largest = l
    else
        largest = i
    if r <= A.heap-size and A[r] > A[largest]
        largest = r
    if largest != i
        exchange A[i] with A[largest]
        Max-Heapify(A,largest)  //递归调用

//建最大堆
Build-Max-Heap(A)
    for i = 向下取整( A.length/2 ) downto 1  //为什么要从一半开始？保证它有子节点
        Max-Heapify(A,i)


//堆排序主函数
Heapsort(A)
    Build-Max-Heap(A)
    for i = A.length downto 2
        exchange A[1] with A[i]
        A.heap-size = A.heap-size - 1
        Max-Heapify(A,1)
```

### 查找表

- 利用set、map(dict) 底层要求顺序，用二叉搜索树(平衡)，不要求，使用哈希表（散列表）实现

|相关问题 | 题意 | 解题思路 | 编程注意
|--- | --- |--- | --- |
| [350. 两个数组的交集二](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/) | 给定两个数组，计算它们的交集 | 使用字典(映射)记录出现次数, 如果数组有序，可使用归并思想:双指针 |　注意题意要求输出所有（包括重复）
|　[454.四数相加二](https://leetcode-cn.com/problems/4sum-ii/) | 四个整数数组，有多有元组(它们的元素)累加为０ |  字典１记录AB数组累加，字典２记录CD数组的累加，再查找　|　注意结果累加是　字典１的值乘字典２的值
|　[447. 回旋镖的数量](https://leetcode-cn.com/problems/number-of-boomerangs/) | 回旋镖指元组（i, j, k），其中i和j之间距离与i和k之间距离相等，　给多个点，找有多少组　| 字典记录距离，查找字典  |　注意组成元组个数应累加 x*(x-1)　距离相等点数目为ｘ
| [215.数组中的第K个最大(最小)元素](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/) | 找数组第k个最大元素　| 除了构建堆和快速排序partition, 还能用字典存储然后排序 |

### 链表类

-思路：穿针引线，是否需要设虚头指针（头指针会发送改变都要设置），是否要覆盖赋值，是否能双指针（或快慢指针）,断开链表时注意，

|相关问题 | 题意 | 解题思路 | 编程注意
|--- | --- |--- | --- |
| [206.反转链表](https://leetcode-cn.com/problems/reverse-linked-list/submissions/) | 将一个单向链表顺序反转，只遍历一遍　| 链表表头前插入节　| 　可以构造虚头指针
| [92. 反转链表二(注意点比较多)](https://leetcode-cn.com/problems/reverse-linked-list-ii/submissions/)  | 将一个单向链表的某一处(m到n)反转，只遍历一遍　|　复杂的链表某出插入加点：构造虚头指针，要记录已倒叙列表的最后一个元素（即开始反转第一个元素）　| 注意：链表断开与链接，不能造成环死循环
| [83. 删除排序链表中的重复元素](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/) | 给定有序单项链表，把重复元素删除，只遍历一遍　| 判断当前与下一个是否相同，相同则链表删除下一个，不同则移动当前　| 注意：相同时不用移动当前
| [86. 分割链表](https://leetcode-cn.com/problems/partition-list/) | 给单向链表和x值，使得链表前部分小于x,后部分大于等于x, 要保持相对顺序(稳定)，只遍历一遍 |　(类似归并merge)可维护两链表，最后合并：构造虚头节点，维护四个指针，遍历一遍　| 处理节点要断开，把节点的next = None
| [328. 奇偶链表](https://leetcode-cn.com/problems/odd-even-linked-list/)  | 把单向链表，奇数位和偶数位分别排一起，最后相连，只遍历一遍　| (类似归并merge)和上一题一样，维护两个链表(4个指针) | 　注意节点断开，把节点next=None, 循环的节点要后移
|　[2. 两树相加](https://leetcode-cn.com/problems/add-two-numbers/)  | 两个链表相加，以头为低位，满十进一，得出新链表　| 模拟加法，进位处理（有可能需要自己构造出链表） | 永远怀疑有进位，并处理
|　[445. 两数相加二](https://leetcode-cn.com/problems/add-two-numbers-ii/) | 两个链表相加，以尾为低位，满十进一，得出新链表 | 主要解决进位如何加在前驱, （反转链表，或者用栈感觉不符合出题人思路）应该用递归(系统栈)　|　这两个问题可转化成１．两链表同样长度　２．循环或递归向加
|　[203. 移除链表元素](https://leetcode-cn.com/problems/remove-linked-list-elements/)  |　单向链表中删除等于给定值val的所有节点　| 遍历删除，构造虚头指针可减少判断　|　注意循环节点移位
|　[82. 删除排序链表中的重复元素二](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/) | 排序链表中删除所有重复的节点，只保留每出现过重复的节点　| 留一个标记说明当前重复节点也要删除，　构造虚头指针 | 注意判断的条件
|　[21. 合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/) | 合并两个有序链表 | （类似归并merge), 构造虚头指针　|
| [24. 两两交换链表中的节点](https://leetcode-cn.com/problems/swap-nodes-in-pairs/) | 单向链表实际两两交换相邻节点　|　构造虚头指针，注意节点的断开与链接 |
| [25. K 个一组翻转链表](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/) | 单向链表　每k个元素内倒序　|  构造虚头指针，要维护pre插入头指针、尾指针(连后面待处理元素)、当前循环指针　，这更复杂的链表倒序　| 注意点：链表断开与连接
| [147. 对链表进行插入排序](https://leetcode-cn.com/problems/insertion-sort-list/) | 对单向链表进行插入排序　| 又要找前驱，只好递归或者使用栈, 构造虚拟头指针　|
| [148. 排序链表](https://leetcode-cn.com/problems/sort-list/) | 对链表进行排序，要求时间复杂度O(n logn)和常数级空间复杂度　| 归并排序，　其中找中间点方法：普通遍历两边，计数；更好的方法是快慢指针　|　按照数组中的归并排序即可
| [237. 删除链表中的节点](https://leetcode-cn.com/problems/delete-node-in-a-linked-list/) | 给定单向链表中某一节点，要求用O(1)时间删除节点　| 无法找到前驱节点，故假装删除，把值往前赋一下，删除该节点的下一个节点　|
| [19. 删除链表的倒数第N个节点](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/submissions/)  | 单项链表中删除倒数第N个节点　| 使用快慢指针，一个先走N步　| 注意慢指针应该，走到删除节点前驱，为皆
|　[61. 旋转链表](https://leetcode-cn.com/problems/rotate-list/) | 把单向链表一个元素从后拿到第一叫旋转一次　| 快慢指针，　确定旋转的部分，　特例：K>链表长度，还是要遍历一次知道链表长度，然后取余数　| 慢指针，应该是移除节点的前驱，快指针应该是最后一个元素
|　[143. 重排链表](https://leetcode-cn.com/problems/reorder-list/)  | 按照顺序重新编排单向链表（合并链表、奇偶链表的变形）| 找中点，后半部分倒序，再奇偶合并 |

### 栈和队列、优先乘

|相关问题 | 题意 | 解题思路 | 编程注意
|--- | --- |--- | --- |
| [20. 有效的括号](https://leetcode-cn.com/problems/valid-parentheses/) |  括号是否匹配 | 利用栈，遇到右括号时就要出一个左口号 | 注意遍历完，栈减
| [150. 逆波兰表达式求值](https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/)  | 构造逆波兰表达表达式 |利用栈，把计算值重新存在栈中，最后栈元素为结果 | 注意：两个树是否为负数时，如何整除
| [71. 简化路径](https://leetcode-cn.com/problems/simplify-path/)   | 简化文件路径  |利用栈，等于..如有就弹，等于.就忽略，正常路径则加入 |
树遍历，前中后序遍历 144 94 145     |   树的遍历 | 递归很简单，非递归：利用栈，前序则要把右边先压栈； 中序，不停先把根节点压入，在压根节点的左元素，后序就是前序的转换 | 注意，中序的非递归
| [341. 扁平化嵌套列表迭代器](https://leetcode-cn.com/problems/flatten-nested-list-iterator/)  |  列表内嵌列表->一个列表       | 使用队列存储，如果是列表则递归调用存储函数|
| [102. 二叉树的层次遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)  |    利用队列，注意输出结构
| [107. 二叉树的层次遍历二](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/)  |
| [103. 二叉树的锯齿形层次遍历](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/)  |
| [199. 二叉树的右视图](https://leetcode-cn.com/problems/binary-tree-right-side-view/)  |  找每一层的最后一个元素 | 层次遍历
| [279. 完全平方数](https://leetcode-cn.com/problems/perfect-squares/)          |  用最少的完全平方数构造一个函数 |要构建图，层遍历，BFS广度优先搜索; 可背包
| [127. 单词接龙](https://leetcode-cn.com/problems/word-ladder/)  | 每次只能换一个字母，能不能从begin单词到end单词，能求最短转换序列   |构建图，然后层遍历
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

```python

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
| [17. 电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/submissions/) |　给定数字，返回它能代表字母集合(手机) | 这是一个树形结构，用递归 记录下标，一个当前路径，一个结果集合 | 注意 数字范围和空空字符串
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
| [121. 买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/)| [122. 买卖股票的最佳时机II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)|[123. 买卖股票的最佳时机III](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/)|[188. 买卖股票的最佳时机 IV](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)
| [300. 最长上升子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/) | LIS,数组，求上升子序列，不一定连续 |  `f(n) = max(1, f(i)+1) i<n, 且num[i]<num[n]`  | 初始化所有元素长度为1，以自己结尾, 最优是dp 二分查找一个元素位置，若该元素
| [376. 摆动系列](https://leetcode-cn.com/problems/wiggle-subsequence/submissions/) | 数组，求最长摆动系列  | 维护两个dp，一个序列处于上升，一个序列处于下降。 `f(n) = max(1, f(i)+1) i<n` | 注意down和up区别
| LCS 最长公共子序列   |    LCS | LCS(i,j) = 当a[i]=b[j], 1+LCS(i+1,j+1) 当a[i]!=b[j], max(LCS(i,j+1), LCS(i+1,j)) | 明白子串要求连续，而子序列不要求连续
| Dijkstra单源最短路径 | 一点到另外一点最短路径   | shorttestPath(x)= min(shortestPath(a) + w(a->x)) |
| [416. 分割等和子集](https://leetcode-cn.com/problems/partition-equal-subset-sum/submissions/)| 一个数组，平均分两个sum相等的集合 | dp[0..sum//2],下标表示背包容量，值表示背包是否装满. `dp[j] = dp[j](原本背包已满) or dp[j-nums[i]] (看背包变小时，是否满)` |  注意：数组下标表示背包容量，遍历数组，考虑是否放入，背包容量为j dp[j],j>=nums[i]
| [377. 组合总和 IV](https://leetcode-cn.com/problems/combination-sum-iv/) | 给定一个由正整数组成且不存在重复数字的数组，找出和为给定目标正整数的组合的个数。| f(i) = sum(f(i-num[k])), i>=num[k] | 做之前观察例子,第一个nums[k],剩下序列就是f(i-num[k])
| [474. 一和零](https://leetcode-cn.com/problems/ones-and-zeroes/) | 给多个01串，给M个0和N个1，最多能装多少个01串 | `dp[i][j]`看不同大小的背包，限制i为1的个数，j为0的个数；对每一个字符串，对每个可放的进去的背包：`dp[i][j] = max(dp[i][j], 1+dp[i- count1][j-count1]`| 不要把二维数组看成一个大背包，而是把二维数组每个`dp[i][j]`,看成一个背包
| [139. 单词拆分](https://leetcode-cn.com/problems/word-break/)| 对于一字符串，能否拆成一个或多个在字典出现的单词 | dp[i], 表示长度为i的字符串 是否装了单词，`dp[i] = dp[i] or dp[i-len(k)]`, k表示单前单词，能否装进去 | 一位数组，每个值看成一个背包， 对于一个单词，遍历能完好装得下它的背包

## 剑指office

### ==面试题3==：数组中重复的数字

题目一：找出数组中重复的数字：在一个长度为n的数组里所有数字都在0~n-1的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了多少次。请找出数组中任意一个重复的数字。例如,如果输入长度为7的数组[2,3,1,0,2,5,3],那么对应输出的数字2或者3

- 方法一：排序，遍历一遍。 时间复杂度O（nlgn），空间复杂度O（1）
- 方法二：哈希表（字典法），用数组建立一个哈希表，或者直接用py的dict字典，遍历一遍。时间复杂度O（n），空间复杂度O（n）
- 最优方法：巧妙的。值为i的 数字应该在下标i，除非已被占，意思该值已重复，否则交换值i到下标i，看重新得的值是否满足条件(2个)。时间复杂度O(n),空间复杂度O(1)

```python
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

```python
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

> 测试用例

- 长度为n的数组里包含一个或多个重复的数字
- 数组中不包括重复的数字
- 无效输入测试用例（输入空指针；长度为n的数组中包含0~n-1之外的数字）

### ==面试题4==：二维数组中的查找

题目：在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

- 方法一：二分找上界，二分找下界，再二分查找
- 方法二：利用已排序的性质，对比，缩小范围

```python
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

> 测试用例

- 二维数组中包含查找的数字
- 二维数组中没有查找的数字
- 特殊输入测试（数组为空）

### ==面试题5==：替换空格

题目：请实现一个函数，把字符串中每个空格替换成%20，例如，输入"We are happy"，则输出"We%20are%20happy"

- 方法：数组手动遍历，空格换%20。或者直接调用字符串.replace

```python
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

> 测试用例

- 输入的字符串包含空格（空格位置前中后，空格一个或连续多个）
- 输入字符串没有空格
- 特殊输入（空字符串，空格字符串、空格一个或连续多个）

### ==面试题6==：从尾到头打印链表

题目：输入一个链表的头节点，从尾到头反过来打印出每个节点的值。

- 方法一：存储逆置，打印
- 方法二：利用栈或递归，实现

```python
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

> 测试用例

- 功能测试(输入的链表有多个节点；输入的链表只有一个节点)
- 特殊输入测试(输入的链表头节点指针为nullptr)

### ==面试题7==：重构二叉树

题目：输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。例如，输入前序遍历序列{1，2，4，7，3，5，6，8}和中序遍历序列{4，7，2，1，5，3，8，6}，则重建并输出它的头节点

- 方法：分治法，分左子树右子树递归，
  
```python
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

> 测试用例

- 普通二叉树(完全二叉树；不完全二叉树)
- 特殊二叉树(所有节点都没有右子节点的二叉树；所有节点都没有左子节点的二叉树；只有一个节点的二叉树)
- 特殊输入测试(二叉树的根节点指针为nullptr；输入的前序遍历和中序遍历序列不匹配)

### ==面试题8==：二叉树的下一个节点

题目：给定一课二叉树和其中的一个节点，如何找到中序遍历序列的下一个节点？树中节点除了有两个分别指向左、右子节点的指针，还有一个指向父节点的指针

```python
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

### ==面试题9==：用两个栈实现一个队列

题目：用两个栈实现一个队列。队列的声明如下，请实现它的两个函数appendTail和deleteHead，分别完成在队列尾部插入节点和在队列头部删除节点的功能

- 方法：入队：栈一 push ， 出队：栈二 不为空 pop，为空 栈一pop所有元素 至 栈二push
  
```python
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

### ==面试题10==：斐波那契数列

题目一：求斐波那契数列的第n项。写一个函数，输入n，求斐波那契数列的第n项，斐波那契数列定义如下：f(0) = 0 , f(1) = 1 , f(n) = f(n-1) + f(n-2)

```python
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

```python
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

```python
#数学归纳证明f(n)= 2**(k-1)
#即证：当n>=2,2**0+2**1+...2**(k-2)+1 = 2**(k-1)

def jumpFloorII(number:int)->int:
        if number<1:
            return 0
        else:
            return pow(2,number-1)
```

- 用2x1的小矩形去覆盖2xN 的大矩形，总共有多少种方法

```python
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

### ==面试题11==：旋转数组的最小数字

题目：把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组{3，4，5，1，2}为{1，2，3，4，5}的一个旋转，该数组的最小值为1

```python
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

### ==面试题12==：矩阵中的路径

题目：请设计一个函数，用来判断在一个矩阵中是否存在一条包括某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3x4的矩阵中包含一条字符串“bfce”的路径。但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二格子之后，路径不能再次进入这个格子。

a | b | t | g
---|---
c | f | c | s
j | d | e | h

- 回溯法

```python
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

### ==面试题13==：机器人的运动范围

题目：地上有一个m行n列的方格。一个机器人从坐标(0, 0)的格子开始移动， 它每次可以向左、右、上、下移动一格，但不能进入行坐标和列坐标的位数之和大于k的格子。例如，当k=18时，机器人能够进入方格(35, 37)，因为3+5+3+7=18。但它不能进入方格(35,38), 应为3+5+3+8=19.请问该机器人能够到达多少个格子？

- 回溯法

```python
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

### ==面试题14==：剪绳子

题目：给你一根长度为n的绳子，请把绳子剪成m段（m、n都是整数，n>1 并且m>1)，每段绳子的长度记为k[0],k[1]...k[m].请问k[0]xk[1]x...xk[m]可能的最大乘积是多少？例如，当绳子的长度是8时，我们把他剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

- 动态规划：求最优解 、整体问题最优解依赖于各子问题最优解 、子问题有重叠更小子问题 、从上往下分析问题,从下往上求解问题、子问题最优解存储。  
- 贪心算法：当前最优导致局部最优，一般需要用数学方法证明正确性

```python
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

    for i in range(4,length+1):
        mymax = 0
        for j in range(1,(i//2)+1):
            product = products[j]*products[i-j]
            if mymax < product:
                mymax = product
        products[i] = mymax
    return max(products)

#贪心算法 证明n>=5时，2(n-2)>n、3(n-3)>n 且3(n-3)>2(n-2)
def maxProductAfterCutting_solution(length):
    if length<2:
        return 0
    if length==2:
        return 1
    if length==3:
        return 2

    timesOf3 = length//3
    if length - timesOf3*3 == 1:
        timesOf3 -= 1
    timesOf2 = (length-timesOf3*3)//2
    return (3**timesOf3)*(2**timesOf2)
```

### ==面试题15==：二进制中1的个数

题目：请实现一个函数，输入一个整数，输出该数二进制表示中1的个数。例如，把9表示成二进制是1001，有2位是1。因此，如果输入9，则该函数输出2。s

- 二进制位运算

与(&)| 或( \| ) | 异或(^)
---|---|---
都为一，才为一 | 有一，则一 | 不同，则一

- ==左移==运算符m<<n：把m左移n位，最左边n位将被抛弃，同时最右边补n个0 （比如一个规定8位的数字）
- ==右移==运算符m>>n, 要注意若数字原先为正数，则右移后最左边补n个0，若数字原先为负数，则右移后最左边补n个1  （负数，最左符号位为一）

```python
# python 常规解法 利用左移,32位要遍历一遍
# -*- coding:utf-8 -*-
class Solution:
    def NumberOf1(self, n):
        count = 0
        flag =1
        for i in range(0,32):
            if n&flag:
                count+=1
            flag=flag<<1
        return count

#  减一再和之前的数做位与 ，有多少个一循环多少次
# python负数并不是补码，题目要求负数用补码表示， bin(-1 )=>'-0b1'
# 补码： 负数 & 0xffffffff  

# -*- coding:utf-8 -*-
class Solution:
    def NumberOf1(self, n):
        # write code here
        count = 0
        if n < 0: #负数变成补码
            n = n & 0xffffffff
        while n:
            count += 1
            n = (n - 1) & n
        return count
```

### ==面试题16==：数值的整数次方

题目：实现函数double Power(double base, int exponent), 求base的exponent次方。不得使用库函数，同时不需要考虑大数问题

- 解法：
- a^n = a^(n/2) \* a^(n/2)   n为偶数
- a^n = a^(n-1 /2)　\*　a^(n-1 /2)* a     n为奇数

```python
# -*- coding:utf-8 -*-
class Solution:
    def Power(self, base, exponent):
        if base==0 and exponent<0:
            return 0
        absExponent = exponent   # absExponent = abs(exponent) if exponent < 0 else exponent
        if exponent<0:
            absExponent = abs(exponent)
        result = self.PowerWithUnsignedExponentFast(base,absExponent)
        if exponent<0:          # return 1/result if exponent < 0 else result
            result = 1/result
        return result

    #效率不高
    def PowerWithUnsignedExponent(self,base, exponent):
        result = 1
        for i in range(exponent):
            result *= base
        return result
    #高效率递归求解，最佳解法
    def PowerWithUnsignedExponentFast(self, base, exponent):
        if exponent == 0:
            return 1
        if exponent == 1:
            return base
        result = self.PowerWithUnsignedExponentFast(base,exponent>>1)
        #exponent右移一位，即整除以二，python相当于exponent//2
        result *= result        # return result*result*base if exponent & 0x1 == 1 else result * result
        if exponent & 0x1 == 1: #奇数要累乘base
            result*=base
        return result
```

### ==面试题17==：打印从一到最大的n位数

题目：输入数字n，按顺序打印出从1到最大的n位十进制数。比如输入3，则打印出1、2、3一直到最大的3位数999

- 该题目用c/c++做需要考虑数的大小，用数组来存储该数字，模拟数字加法，
- 该问题能转换成数字排列的解法，递归让代码更简洁

```python
#python3 则不需要考虑数字大小，数值是不限制长度的
def Print1ToMax(n:int)->None:
    number = 1
    i = 0
    while i<n:
        number*=10
        i+=1
    for i in range(1,number):
        print(number)

#c/c++
```

### ==面试题18==：删除链表节点

**题目一：**在O(1)时间内删除链表节点。给定单向链表的头指针和一个节点指针，定义一个函数在O(1)时间内删除该节点。

- 常规方法：遍历找出，然后删除
- 时间0(1)方法：删除节点在中间有后续节点，只要要后续节点值复制过来，把后续节点删，就相当于删除了想要删除的节点。另外还有两种情况：1、链表只有一节点，删除头节点（也是尾节点）2、删除节点为最后一节点，则还是需要遍历一遍.  *（这都建立在该删除节点一定在链表中）*

```python
# -*- coding:utf-8 -*-
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None
class Solution:
    def deleteNextNode(self, pHead,pToBeDeleted):
        if pHead == None or pToBeDeleted == None:
            return
        if pToBeDeleted.next!=None:
            pNext = pToBeDeleted.next
            pToBeDeleted.val = pNext.val
            pToBeDeleted.next = pNext.next
        elif pHead == pToBeDeleted:
            pHead = None
        else:
            pNode = pHead
            while pNode.next != pToBeDeleted:
                pNode = pNode.next
            pNode.next = None
```

**题目二：**删除链表中重复的节点。在一个排序的链表中，如何删除重复的节点，意思有重复的都不要

- 方法：遍历一遍，设置flag，设置双指针，该节点不与下一节点值相同时，双指针后移，否则要进去找到一个不同值的节点A，前指针pPreNode的next就要指向节点A，这里注意前指针为空时，应该头指针指向节点A

```python
# -*- coding:utf-8 -*-
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
class Solution:
    def deleteDuplication(self, pHead):
        if pHead == None:
            return
        pPreNode = None
        pNode = pHead
        while pNode!=None:
            pNext = pNode.next
            needDelete = False
            if pNext!=None and pNext.val == pNode.val:
                needDelete = True
            if not needDelete:
                pPreNode = pNode
                pNode = pNode.next
            else:
                value = pNode.val
                pToBeDel = pNode
                while pToBeDel !=None and pToBeDel.val == value:
                    pNext = pToBeDel.next
                    pToBeDel = pNext
                if pPreNode ==None:
                    pHead = pNext
                else:
                    pPreNode.next = pNext
                pNode = pNext
        return pHead
```

### ==面试题19==：正则表达匹配

题目：请实现一个函数用来匹配包含"."和"\*"的正则表达式。模式中的字符"."表示任意一个字符，而"\*"表示踏前面的字符可以出现任意次（含0次）。在本题中，匹配是指字符串中所有字符匹配整个模式。例如，字符串"aaa",与模式"a.a"和"ab\*ac\*a"匹配，但与"aa.a"和"ab*a"均不匹配

- 方法：第一个字符串为 . 则可匹配任意字符，若为ch 则只可以匹配ch
- 第二个字符串为不是‘\*’，前一字符能否匹配，能则都后移，递归匹配，不能反回false  ， 若为‘\*’,匹配一位：字符串后移一位，正则后移两位，或 字符串后移一位，正则不动； 匹配0位：正则后移两位，字符串不动
- 递归思想

``` python
# -*- coding:utf-8 -*-
class Solution:
    # s, pattern都是字符串
    def match(self, s, pattern):
        if s==None or pattern==None:
            return False
        return self.matchCore(s,pattern)
    def matchCore(self,s,pattern):
        if len(s)==0 and len(pattern)==0:
            return True
        if len(s)!=0 and len(pattern)==0:
            return False

        #正则表达式长度大于1，且第二个字符为*
        if len(pattern)>1 and pattern[1] == '*':
            if len(s)!=0 and (pattern[0] == s[0] or pattern[0]=="."):
                return self.matchCore(s[1:],pattern[2:]) or self.matchCore(s[1:],pattern[:]) or self.matchCore(s[:],pattern[2:])
            else:
                return self.matchCore(s[:],pattern[2:])

        #字符串和正则表达式都不为空
        if len(s)!=0 and (s[0] == pattern[0] or pattern[0]=="." ):
            return self.matchCore(s[1:],pattern[1:])
        return False
```

### ==面试题20==：表示数值的字符串

题目：请实现一个函数用来判断字符串是否表示数值(包含整数和小数)。例如，字符串“+100”、“5e2”、“-123”、“3.1416”及“-1E-16”都表示数值，但“12e”、“1a3.14”、“1.2.3”、“+-5”及“12e+5.4”都不是

```python
# -*- coding:utf-8 -*-
class Solution:
    # s字符串
    def isNumeric(self, s):
        if not isinstance(s, str) or len(s) == 0:
            return False
        s = list(s)
        flag_number = self.scanInterger(s)

        if len(s)!=0 and s[0]=='.':
            s.pop(0)
            flag_number = self.scanUnsignedInterger(s) or flag_number
        if len(s)!=0 and (s[0]=='e' or s[0]=='E'):
            s.pop(0)
            flag_number = flag_number and self.scanInterger(s)
        return flag_number and len(s)==0

    def scanInterger(self,s):
        if len(s)==0:
            return False
        if s[0]=='+' or s[0]=='-':
            s.pop(0)
        return self.scanUnsignedInterger(s)
    def scanUnsignedInterger(self,s):
        f_len = len(s)
        while len(s)!=0 and s[0]>='0' and s[0]<='9':
            s.pop(0)
        if len(s)<f_len:
            return True
        return False
```

### ==面试题21==：调整数组顺序使奇数位于偶数前面

题目：输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

- 方法：前后双指针

```python
# -*- coding:utf-8 -*-
class Solution:
    def reOrderArray(self, array):
        return sorted(array,key=lambda c:c & 0x1,reverse=True)

# -*- coding:utf-8 -*-
class Solution:
    def reOrderArray(self, array):
        if len(array)==0:
            return []
        L,index= self.reOrder(array,lambda x:x & 0x1)  #是否为奇数
        #竟然 牛客还要求排序made
        Lleft = L[:index]
        Lright =L[index:]
        return sorted(Lleft)+sorted(Lright)
    def reOrder(self,array,func):
        begin_index = 0
        end_index = len(array)-1
        while begin_index < end_index:
            while begin_index<end_index and func(array[begin_index]):
                begin_index+=1
            while begin_index<end_index and not func(array[end_index]):
                end_index-=1
            if begin_index<end_index:
                array[begin_index],array[end_index] = array[end_index],array[begin_index]
        return array,begin_index
```

### ==面试题22==:链表中倒数第k个结点

题目：输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾结点是倒数第1个节点。例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6.这个链表的倒数第3个节点是值为4的节点。

```python
# -*- coding:utf-8 -*-
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def FindKthToTail(self, head, k):
        if k<=0 or head == None:
            return None
        pAhead = head
        pBchild = None
        #检查链表 是否至少有k个节点
        for i in range(k-1):
            if pAhead.next:
                pAhead = pAhead.next
            else:
                return None
        pBchild = head
        while pAhead.next:
            pAhead = pAhead.next
            pBchild = pBchild.next
        return pBchild
```

### ==面试题23==：链表中环的入口结点

题目：如果一个链表中包含环，如何找出环的入口节点？例如，在如图3.8所示的链表中，环的入口节点是节点3

```python
# -*- coding:utf-8 -*-
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
class Solution:
    def EntryNodeOfLoop(self, pHead):
        meetingNode = self.MeetingNode(pHead)
        if meetingNode == None:
            return None
        nodesInLoop = 1
        pNode1 = meetingNode
        while pNode1.next != meetingNode:
            pNode1 = pNode1.next
            nodesInLoop+=1
        pNode1 = pHead
        for i in range(nodesInLoop):
            pNode1 = pNode1.next
        pNode2 = pHead
        while pNode1 != pNode2:
            pNode1 = pNode1.next
            pNode2 = pNode2.next
        return pNode1
    def MeetingNode(self,pHead):
        if pHead == None:
            return None
        pSlow = pHead
        if pSlow.next == None:
            return None
        pFast = pSlow.next
        while pFast != None and pSlow != None:
            if pFast == pSlow:
                return pFast
            pSlow = pSlow.next
            pFast = pFast.next
            if pFast !=None:
                pFast = pFast.next
        return None
```

### ==面试题24==：反转链表

题目：定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

```python
# -*- coding:utf-8 -*-
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
class Solution:
    # 返回ListNode
    def ReverseList(self, pHead):
        if pHead == None:
            return None
        pnewHead = None
        pNode = pHead
        while pNode:
            pNext = pNode.next
            pNode.next = pnewHead
            pnewHead = pNode
            pNode = pNext
        return pnewHead

#利用python交换特性
class Solution:
    def ReverseList(self, pHead):
        pre = None
        while pHead:
            pHead.next, pHead, pre = pre, pHead.next, pHead
        return pre
```

### ==面试题25==：合并两个排序的链表

题目：输入两个递增排序的链表，合并这两个链表并使新链表中的节点任然是递增排序的。

```python
# -*- coding:utf-8 -*-
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
class Solution:
    # 这用递归
    def Merge(self, pHead1, pHead2):
        if pHead1 == None:
            return pHead2
        elif pHead2 == None:
            return pHead1
        pnewHead = None
        if pHead1.val<pHead2.val:
            pnewHead = pHead1
            pnewHead.next = self.Merge(pHead1.next,pHead2)
        else:
            pnewHead = pHead2
            pnewHead.next = self.Merge(pHead1,pHead2.next)
        return pnewHead


```

### ==面试题26==：树的子结构

题目：输入两棵二叉树A和B，判断B是不是A的子结构

```python
# -*- coding:utf-8 -*-
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    def HasSubtree(self, pRoot1, pRoot2):
        result = False
        if pRoot1 != None and pRoot2 != None:
            if self.Equal(pRoot1.val,pRoot2.val):
                result = self.DoesTree1HaveTree2(pRoot1,pRoot2)
            if not result:
                result = self.HasSubtree(pRoot1.left,pRoot2)
            if not result:
                result = self.HasSubtree(pRoot1.right,pRoot2)
        return result

    def DoesTree1HaveTree2(self,pRoot1,pRoot2):
        if pRoot2 == None:
            return True
        if pRoot1 == None:
            return False
        if not self.Equal(pRoot1.val,pRoot2.val):
            return False
        return self.DoesTree1HaveTree2(pRoot1.left,pRoot2.left) and self.DoesTree1HaveTree2(pRoot1.right,pRoot2.right)

    def  Equal(self,num1,num2):
        if num1-num2>-0.0000001 and num1-num2<0.0000001:
            return True
        else:
            return False
```

### ==面试题27==：二叉树的镜像

题目：请完成一个函数：输入一棵二叉树，该函数输出它的镜像。

```python
# -*- coding:utf-8 -*-
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    # 返回镜像树的根节点
    def Mirror(self, root):
        if root == None:
            return None
        if root.left == None and root.right == None:
            return root
        root.left,root.right = root.right,root.left
        if root.left:
            self.Mirror(root.left)
        if root.right:
            self.Mirror(root.right)
        return root
```

### ==面试题28==：对称的二叉树

题目：请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

```python
# -*- coding:utf-8 -*-
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    def isSymmetrical(self, pRoot):
        return self.isSymmetricalcore(pRoot,pRoot)
    def isSymmetricalcore(self,pRoot1,pRoot2):
        if pRoot1 == None and pRoot2==None:
            return True
        if pRoot1 == None or pRoot2 == None:
            return False
        if pRoot1.val != pRoot2.val:
            return False
        return isSymmetricalcore(pRoot1.left, pRoot2.right) and isSymmetricalcore(pRoot1.right, pRoot2.left)
```

### ==面试题29==：顺时针打印矩阵

题目：输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字：

```python
# -*- coding:utf-8 -*-
class Solution:
    # matrix类型为二维列表，需要返回列表
    def printMatrix(self, matrix):
        if len(matrix)==0:
            return
        start = 0
        row = len(matrix)
        col = len(matrix[0])
        result = []
        while col>start*2 and row>start*2:
            self.printMatrixCircle(matrix,col,row,start,result)
            start+=1
        return result
    def printMatrixCircle(self,matrix,col,row,start,result):
        endX = col - 1 - start
        endY = row - 1 - start
        for i in range(start,endX+1):
            result.append(matrix[start][i])
        if start<endY:
            for i in range(start+1,endY+1):
                result.append(matrix[i][endX])
        if start<endX and start<endY:
            for i in range(endX-1,start-1,-1):
                result.append(matrix[endY][i])
        if start<endX and start<endY-1:
            for i in range(endY-1,start,-1):
                result.append(matrix[i][start])
```

### ==面试题30==：包含min函数的栈

题目：定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的min函数。在该栈中，调用min、push及pop的时间复杂度都是O(1)

- 思路：用主栈加一个min辅助栈

1. push元素入主栈，push元素比辅助栈小则压入，否则辅助栈压入辅助栈顶元素
2. pop两个栈都弹出
3. min返回辅助栈栈顶元素

```python
# -*- coding:utf-8 -*-
class Solution:
    def __init__(self):
        self.stack = []
        self.min_stack = [] #辅助

    def push(self, node):
        self.stack.append(node)
        if self.min_stack == [] or node<=self.min_stack[-1]:
            self.min_stack.append(node)
        else:
            self.min_stack.append(self.min_stack[-1])
    def pop(self):
        self.stack.pop()
        self.min_stack.pop()
    def top(self):
        return self.stack[-1]
    def min(self):
        return self.min_stack[-1]
```

### ==面试题31==：栈的压入、弹出序列

题目：输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个系列是否为该栈的弹出序列。假设压入栈的所有数字均不相等。

- 思路：需要一个栈

1. 依次取出弹栈序列的第一个元素A，如果它在栈中，看弹栈
2. 不栈：依次把压入序列的第一个，且不等于A的压入。循环结束 判断压入栈是否为空，是则返回False，否则删
3. 外层循环结束，则返回True

```python
# -*- coding:utf-8 -*-
class Solution:
    def IsPopOrder(self, pushV, popV):
        stack = [] #栈
        while popV:
            popnumber = popV.pop(0)
            if len(stack)!=0 and stack[-1] == popnumber:
                stack.pop()
            else:
                while pushV and pushV[0]!=popnumber:
                    stack.append(pushV.pop(0))
                if pushV==[]:
                    return False
                else:
                    pushV.pop(0)
        return True
```

### ==面试题32==：从上到下打印二叉树

题目一：不分行从上到下打印二叉树
类似图算法-广度搜索BFS

```python
from collections import deque
# -*- coding:utf-8 -*-
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    # 返回从上到下每个节点值列表，例：[1,2,3]
    def PrintFromTopToBottom(self, root):
        if root == None:
            return []
        d =  deque()
        result = []
        d.appendleft(root)
        while d:
            Node = d.pop()
            result.append(Node.val)
            if Node.left:
                d.appendleft(Node.left)
            if Node.right:
                d.appendleft(Node.right)
        return result
```

题目二：分行从上到下打印二叉树

- 处理一：需要一个变量记录当前行 还没打印的元素数量
- 处理二：BFS 压入队列的元素，增加一个层数标记
  
```python
from collections import deque
#-*- coding:utf-8 -*-
#class TreeNode:
#    def __init__(self, x):
#        self.val = x
#        self.left = None
#        self.right = None
class Solution:
    # 返回从上到下每个节点值列表，
    # 例：[[8],[6,10],[5,7,9,11]]
    def PrintFromTopToBottom(self, root):
        if root == None:
            return []
        d =  deque()
        result = []
        d.appendleft(root)
        nextLevel = 0
        toBePrinted = 1
        hang = []
        while d:
            Node = d.pop()
            toBePrinted -= 1
            hang.append(Node.val)
            if Node.left:
                d.appendleft(Node.left)
                nextLevel+=1
            if Node.right:
                d.appendleft(Node.right)
                nextLevel+=1
            if toBePrinted == 0:
                result.append(hang)
                hang = list()
                toBePrinted = nextLevel
                nextLevel = 0
        return result
```

题目三：之字形打印二叉树

- BFS 压入队列的元素，增加一个层数标记
- 增加一个变量flag=0,1 表示当前行是正向打印，还是逆向打印
  
```python
#-*- coding:utf-8 -*-
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
class Solution:
    # 返回从上到下每个节点值列表，
    # 例：[[8],[6,10],[5,7,9,11]]
    def PrintFromTopToBottom(self, root):
        if root == None:
            return None
        levels = [[],[]]
        current = 0     # 0,1   互逆的艺术
        next = 1
        levels[current].append(root)
        while levels[0] or levels[1]:
            pNode = levels[current].pop()
            print(pNode.val,end=" ")

            if current == 0:
                if pNode.left:
                    levels[next].append(pNode.left)
                if pNode.right:
                    levels[next].append(pNode.right)
            else:
                if pNode.right:
                    levels[next].append(pNode.right)
                if pNode.left:
                    levels[next].append(pNode.left)
            if levels[current]==[]:
                print("\n",end="")
                current = 1- current
                next = 1- next
```

### ==面试题33==：二叉搜索树的后序遍历序列

题目：输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。

- 方法一：直接后序遍历，对比数组
- 方法二：数组前部分（比根小）为左子树的，数组后部分(比根大)为右子树的

```python
# -*- coding:utf-8 -*-
class Solution:
    def VerifySquenceOfBST(self, sequence):
        if sequence==None or len(sequence)==0:
            return False
        root = sequence[-1]

        flag_index = len(sequence)-1
        for i in range(0,len(sequence)-1):
            if sequence[i]>root:
                flag_index = i
                break
        for j in range(flag_index,len(sequence)-1):
            if sequence[j]< root :
                return False
        left = True
        if flag_index>0:
            left = self.VerifySquenceOfBST(sequence[:flag_index])
        right = True
        if flag_index<len(sequence)-1:
            right = self.VerifySquenceOfBST(sequence[flag_index:len(sequence)-1])
        return left and right
```

### ==面试题34==：二叉树中和为某一值得路径

题目：输入一棵二叉树和一个整数，打印出二叉树中节点值得和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的几点形成一条路径。

- 思路：前序遍历，加path存储路径，res存储最终结果
- 判断

```python
# -*- coding:utf-8 -*-
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    # 返回二维列表，内部每个列表表示找到的路径,要求在返回值的list中，数组长度大的数组靠前
    def FindPath(self, root, expectNumber):
        if root == None:
            return []
        result = []
        path = []
        currentSum = 0
        self.FindRightPath(root,expectNumber,path,currentSum,result)
        result.sort(key = lambda x:len(x),reverse = True)
        return result

    def FindRightPath(self,root,expectNumber,path,currentSum,result):
        currentSum += root.val
        path.append(root.val)
        isLeaf = (root.left==None and root.right==None)
        if currentSum == expectNumber and isLeaf:
            result.append(list(path))    #list是可变类型，注意
        if root.left:
            self.FindRightPath(root.left,expectNumber,path,currentSum,result)
        if root.right:
            self.FindRightPath(root.right,expectNumber,path,currentSum,result)
        path.pop()
```

### ==面试题35==：复杂链表的复制

题目：请实现一个函数复制一个复杂链表。在复杂链表中，每个节点除了有一个next指针，还有一个pSibling指向链表中的任意节点或者nullptr。

- 思路：
- 方法一： 优化定位pSibling，遍历复制链表，同时，用hash表存储
  
```python
# -*- coding:utf-8 -*-
# class RandomListNode:
#     def __init__(self, x):
#         self.label = x
#         self.next = None
#         self.random = None
class Solution:
    # 返回 RandomListNode
    def Clone(self, pHead):
        self.CloneNodes(pHead)
        self.ConnectSiblingNodes(pHead)
        return self.ReconnectNodes(pHead)

    def CloneNodes(self, pHead):
        pNode = pHead
        while pNode:
            pCloned = RandomListNode(pNode.label)
            pCloned.next = pNode.next
            pCloned.random = None
            pNode.next = pCloned
            pNode = pCloned.next

    def ConnectSiblingNodes(self,pHead):
        pNode = pHead
        while pNode:
            pCloned = pNode.next
            if pNode.random :
                pCloned.random = pNode.random.next
            pNode = pCloned.next

    def ReconnectNodes(self,pHead):
        pNode = pHead
        pClonedHead = None
        pClonedNode = None
        if pNode:
            pClonedHead = pClonedNode = pNode.next
            pNode.next = pClonedHead.next
            pNode = pNode.next
        while pNode:
            pClonedNode.next = pNode.next
            pClonedNode = pClonedNode.next
            pNode.next = pClonedNode.next
            pNode = pNode.next
        return pClonedHead
```

### ==面试题36==：二叉搜索树与双向链表

```python
# -*- coding:utf-8 -*-
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    def __init__(self):
        self.pLastNodeInList = None
    def Convert(self, pRootOfTree):
        if pRootOfTree==None:
            return
        self.ConvertNode(pRootOfTree)
        pHeadOfList = pRootOfTree
        while pHeadOfList and pHeadOfList.left :
            pHeadOfList = pHeadOfList.left
        return pHeadOfList

    def ConvertNode(self,pNode):
        if pNode == None:
            return
        pCurrent = pNode
        if pCurrent.left:
            self.ConvertNode(pCurrent.left)
        pCurrent.left = self.pLastNodeInList
        if self.pLastNodeInList:
            self.pLastNodeInList.right = pCurrent
        self.pLastNodeInList = pCurrent
        if pCurrent.right:
            self.ConvertNode(pCurrent.right)
```

### ==面试题37==：序列化二叉树

```python
#利用下标
# -*- coding:utf-8 -*-
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    def __init__(self):
        self.flag = -1
    def Serialize(self, root):
        # write code here
        if root == None :
            return '#'
        return str(root.val) + ',' + self.Serialize(root.left) + ',' + self.Serialize(root.right)
    def Deserialize(self, s):
        # write code here
        self.flag += 1

        l = s.split(',')
        if self.flag >= len(s):
            return None

        root = None
        if l[self.flag] != '#':
            root = TreeNode(int(l[self.flag]))
            root.left = self.Deserialize(s)
            root.right = self.Deserialize(s)
        return root

#利用pop()
class Solution:
    def Serialize(self, root):
        # write code here
        if root == None :
            return '#'
        return str(root.val) + ',' + self.Serialize(root.left) + ',' + self.Serialize(root.right)
    def Deserialize(self, s):
        return self.myDeserialize(s.split(','))

    def myDeserialize(self,L):
        if len(L)==0:
            return None
        if L[0]=='#':
            L.pop(0)
            return None
        root = TreeNode(int(L[0]))
        L.pop(0)
        root.left = self.myDeserialize(L)
        root.right = self.myDeserialize(L)
        return root
```

### ==面试题38==：字符串的排序

```python
#利用库函数
# -*- coding:utf-8 -*-
import itertools
class Solution:
    def Permutation(self, ss):
        # write code here
        if not ss:
            return []
        return sorted(list(set(map(''.join, itertools.permutations(ss)))))


# -*- coding:utf-8 -*-
class Solution:
    def __init__(self):
        self.result = set()
    def Permutation(self, ss):
        if not ss:
            return []
        self.myPermutation(list(ss),0)
        return sorted(list(self.result))


    def myPermutation(self,L,pbegin_index):
        if pbegin_index==len(L):
            self.result.add(''.join(L))
        else:
            for pch_index in range(pbegin_index,len(L)):
                L[pch_index],L[pbegin_index] = L[pbegin_index],L[pch_index]
                self.myPermutation(L,pbegin_index+1)
                L[pch_index],L[pbegin_index] = L[pbegin_index],L[pch_index]

```

### ==面试题39==：数组中出现次数超过一半的数字

- 方法一：基于Partition函数，出现次数超过一半的数字一定会位于数组中间位置

```python
# -*- coding:utf-8 -*-
import random
class Solution:
    def MoreThanHalfNum_Solution(self, numbers):
        if type(numbers)!=list or len(numbers) == 0:
            return 0
        middle = len(numbers)>>1
        start = 0
        end = len(numbers)-1
        index = self.Randomize_Partition(numbers,start,end)
        while index != middle:
            if index > middle:
                end = index - 1
                index = self.Randomize_Partition(numbers,start,end)
            else:
                start = index + 1
                index = self.Randomize_Partition(numbers,start,end)
        result = numbers[middle]
        if not self.CheckMoreThanHalf(numbers,result):
            result = 0
        return result

    def CheckMoreThanHalf(self,numbers,result):
        if numbers.count(result)*2<= len(numbers):
            return False
        else:
            return True

    def Partition(self,A,p,r):
        x = A[r]
        i = p-1
        for j in range(p,r):
            if A[j] <= x:
                i+=1
                A[i], A[j] = A[j], A[i]
        A[i+1],A[r] = A[r],A[i+1]
        return i+1
    def Randomize_Partition(self,A,p,r):
        i = random.randint(p,r)
        A[r],A[i] = A[i],A[r]
        return self.Partition(A,p,r)


# -*- coding:utf-8 -*-
class Solution:
    def MoreThanHalfNum_Solution(self, numbers):
        if type(numbers)!=list or len(numbers) == 0:
            return 0
        result = numbers[0]
        times = 0
        for item in numbers:
            if times == 0:
                result = item
                times = 1
            elif item == result:
                times += 1
            else:
                times -= 1
        if not self.CheckMoreThanHalf(numbers,result):
            result = 0
        return result

    def CheckMoreThanHalf(self,numbers,result):
        if numbers.count(result)*2<= len(numbers):
            return False
        else:
            return True
```

### ==面试题40==：最小的k个数

题目：输入n个整数，找出其中最小的k个数。例如，输入的4、5、1、6、2、7、3、8 这8个数字，则最小的4个数字是1、2、3、4.

- 注意要先弄清题目的要求，包裹输入的数据量多大，能否一次性载入内存，是否允许交换输入数据中数字的顺序

```python
#时间复杂度为o(nlgn)
# -*- coding:utf-8 -*-
class Solution:
    def GetLeastNumbers_Solution(self, tinput, k):
        if len(tinput)==0 or len(tinput)<k:
            return []
        return sorted(tinput)[0:k]

#时间复杂度0(n)
import random
# -*- coding:utf-8 -*-
class Solution:
    def GetLeastNumbers_Solution(self, tinput, k):
        if len(tinput)==0 or len(tinput)<k or k<=0:
            return []

        start = 0
        end = len(tinput)-1
        index = self.Randomize_Partition(tinput,start,end)
        while index != k-1:
            if index > k-1:
                end = index - 1
                index = self.Randomize_Partition(tinput,start,end)
            else:
                start = index + 1
                index = self.Randomize_Partition(tinput,start,end)
        return sorted(tinput[:k])
    def Partition(self,A,p,r):
        x = A[r]
        i = p-1
        for j in range(p,r):
            if A[j] <= x:
                i+=1
                A[i], A[j] = A[j], A[i]
        A[i+1],A[r] = A[r],A[i+1]
        return i+1
    def Randomize_Partition(self,A,p,r):
        i = random.randint(p,r)
        A[r],A[i] = A[i],A[r]
        return self.Partition(A,p,r)

# 时间复杂度O(nlogk)  特别适合处理海量数据

#创建一个大小为k的容器来存储最小的k个数字
#输入的n中读取一个数组，如果容器数字个数 < k，直接存放
#                       如果容器数字个数 已有 k个
#当容器满了时，我们要做三件事：一：k个整数中找最大数
                               二：可能在这容器中删除
```

### ==面试题41==：数据流中的中位数

```python
# -*- coding:utf-8 -*-
class Solution:
    def __init__(self):
        self.data=[]
    def Insert(self, num):
        self.data.append(num)
        self.data.sort()

    def GetMedian(self,data):
        length=len(self.data)
        return self.data[length//2] if length & 0x1 else (self.data[length//2]+self.data[length//2-1])/2.0
```

### ==面试题42==：连续子数组的最大和

```python
#数组规律
# -*- coding:utf-8 -*-
class Solution:
    def FindGreatestSumOfSubArray(self, array):
        if type(array) != list or len(array)==0:
            return None

        nCurSum = 0
        nGreatestSum = 0
        maxitem = array[0]
        for item in array:
            if item>maxitem:
                maxitem=item

            if nCurSum<=0:
                nCurSum = item
            else:
                nCurSum += item
            if nCurSum > nGreatestSum:
                nGreatestSum = nCurSum
        return nGreatestSum if nGreatestSum>0 else maxitem
```

### ==面试题43==：1~n整数中1出现的次数

```python
# 时间复杂度 o(nlogn)
# -*- coding:utf-8 -*-
class Solution:
    def NumberOf1Between1AndN_Solution(self, n):
        number = 0
        for i in range(1,n+1):
            number += self.NumberOf1(i)
        return number
    def NumberOf1(self,n):
        number = 0
        while n:
            if n % 10 == 1:
                number+=1
            n = n //10
        return number
        # write code here


# 更好方法 时间复杂度 o(lgn)
# -*- coding:utf-8 -*-
class Solution:
    def NumberOf1Between1AndN_Solution(self, n):
        if n < 0:
            return 0
        return self.NumberOf1(str(n))

    def NumberOf1(self,strN):
        if strN == None or len(strN) == 0:
            return 0

        first = int(strN[0])
        length = len(strN)

        if length == 1 and first ==0:
            return 0
        if length == 1 and first ==1:
            return 1

        #假设strN 是"21345"
        #numFirstDifit 是 数字10000 ~ 19999 的第一位（万位）出现1的数目
        numFirstDigit = 0
        if first>1:
            numFirstDigit = pow(10,length-1)
        elif first == 1:
            numFirstDigit = int(strN[1:])+1

        #numOtherDigits 是 1346 ~ 21345 除了第一位，其他位出现1的数目 （第一位可以分first段，每段中剩下位 选一位为1，其他位可以为0~9中任意一个）
        numOtherDigits = first * (length-1) * pow(10,length-2)
        #numRecursive 是 1~1345 出现 1的数目，递归
        numRecursive = self.NumberOf1(strN[1:])

        return numFirstDigit + numOtherDigits + numRecursive
```

### ==面试题44==：数字序列中某一位的数字

```python
class Solution:
    def digitAtIndex(self, index):
        if index < 0:
            return -1
        digits = 1
        while True:
            numbers = self.countOfIntergers(digits)
            if index < numbers * digits:
                return self.digitAtTheIndex(index,digits)
            index -= (digits * numbers)
            digits+=1
        return -1

    def countOfIntergers(self,digits):
        if digits == 1:
            return 10
        count = pow(10,digits-1)
        return 9 * count

    def digitAtTheIndex(self,digits):
        number = self.beginNumber(digits) + index/digits
        intdexFromRight = digits - index%digits
        for i in range(1,intdexFromRight):
            number /= 10
        return number % 10

    def beginNumber(digits):
        if digits == 1:
            return 0
        return pow(10,digits-1)
```

### ==面试题45==：把数组排成最小的数

```python
# -*- coding:utf-8 -*-
class Solution:
    def PrintMinNumber(self, numbers):
        if type(numbers) != list or len(numbers)==0:
            return ""
        lmb = lambda n, m:int(str(n)+str(m)) - int(str(m)+str(n))
        array = sorted(numbers,cmp=lmb)
        return "".join([str(i) for i in array])

#python3
import functools
class Solution:
    def PrintMinNumber(self, numbers):
        if type(numbers) != list or len(numbers)==0:
            return ""
        lmb = lambda n, m:int(str(n)+str(m)) - int(str(m)+str(n))
        example.sort(key = functools.cmp_to_key(lmb))
        return "".join([str(i) for i in example])
```

### ==面试题46==：把数字翻译成字符串

f(i) = f(i+1)+g(i,i+1)*f(i+2) 注：f(i)表示从第i位数字开始的不同翻译的数目

```python
class Solution:
    def GetTranslationCount(self, number):
        if number<0:
            return 0
        return self.subGetTranslationCount(str(number))

    def subGetTranslationCount(self,str_number):
        length = len(str_number)
        counts = [0]*length
        for i in range(length-1,-1,-1):
            count = 0
            if i<length-1:
                count = counts[i+1]
            else:
                count+=1
            if i<length-1:
                digit1 = int(str_number[i])
                digit2 = int(str_number[i+1])
                converted = digit1 * 10 + digit2
                if converted >=10 and converted <=25:
                    if i<length-2:
                        count += counts[i+2]
                    else:
                        count += 1
            counts[i] = count
        resultcount = counts[0]
```

### ==面试题47==：礼物的最大价值

```python
class Solution:
    def getMaxvalue(self,values):
        rows = len(values)
        cols = len(values[0])
        if not values or rows <= 0 or cols <= 0:
            return 0
        maxValues = [0]*cols

        for i in range(rows):
            for j in range(cols):
                left = 0
                up = 0
                if i>0:
                    up = maxValues[j]
                if j>0:
                    left = maxValues[j-1]
                maxValues[i][j] = max(left,up) + values[i][j] #values[i*cols+j]
        return maxValues[cols-1]
```

### ==面试题48==：最长不含重复字符的子字符串

```python
class Solution:
    def longestSubstrWithoutDuplication(self,SS):
        if not isinstance(SS,str) or len(SS) == 0:
            return 0
        curLength = 0
        maxLength = 0

        position = [-1]*26
        for i in range(len(SS)):
            preIndex = position[ord(SS[i])-ord('a')]
            if preIndex < 0 or i - preIndex > curLength:
                curLength += 1
            else:
                if curLength > maxLength:
                    maxLength = curLength
                curLength = i - preIndex
            position[ord(SS[i])-ord('a')] = i
            if curLength > maxLength:
                maxLength = curLength
        return maxLength

#双指针滑动的思想
class Solution:
    def longest_substring_without_deplication(self, ss):
        if not isinstance(ss, str) or len(ss) == 0:
            return 0

        count = [0] * 26
        left = 0
        right = -1
        res = 0

        while left<len(ss):
            if right+1 <len(ss) and  count[ord(ss[right+1]) - ord('a')] == 0:
                right += 1
                count[ord(ss[right]) - ord('a')] += 1
            else:
                count[ord(ss[left]) - ord('a')] -= 1
                left += 1

            res = max(res, right-left+1)
        return res
```

### ==面试题49==：丑数

```python
# -*- coding:utf-8 -*-
class Solution:
    def GetUglyNumber_Solution(self, index):
        if index <= 0:
            return 0
        UglyNumbers = [1]
        nextUglyIndex = 1

        Multiply2 = 0
        Multiply3 = 0
        Multiply5 = 0

        while nextUglyIndex < index:
            minNumber = min(UglyNumbers[Multiply2]*2,UglyNumbers[Multiply3]*3,UglyNumbers[Multiply5]*5)
            UglyNumbers.append(minNumber)

            while UglyNumbers[Multiply2]*2 <= UglyNumbers[nextUglyIndex]:
                Multiply2 += 1
            while UglyNumbers[Multiply3]*3 <= UglyNumbers[nextUglyIndex]:
                Multiply3 += 1
            while UglyNumbers[Multiply5]*5 <= UglyNumbers[nextUglyIndex]:
                Multiply5 += 1

            nextUglyIndex +=1

        return UglyNumbers[nextUglyIndex-1]


```

### ==面试题50==:第一次只出现一次的字符

题目一：字符串中第一个只出现一次的字符

```python
# -*- coding:utf-8 -*-
class Solution:
    def FirstNotRepeatingChar(self, s):
        if not s or len(s) == 0:
            return -1
        hashTable = {}

        for item in s:
            if item in hashTable.keys():
                hashTable[item] += 1
            else:
                hashTable[item] = 1
        for index in range(len(s)):
            if hashTable[s[index]] == 1:
                return index
        return -1
```

相关题目

- 定义一函数，输入两字符串，从第一个字符串中删除在第二个字符串中出现过的所有字符
- 定义一函数，删除字符串中所有重复出现的字符
- 在英语中，如果两个单词中出现的字母相同，并且每个字母出现的次数也相同，那个这两个单词互为变位词。(字母无位置要求)。如何判断两字符串，是否为变位词

题目二：字符流中第一个值出现一次的字符

```python
# -*- coding:utf-8 -*-
class Solution:
    def __init__(self):
        self.occurrence = {}
        self.index = 0

    def insert(self,char):
        if char in self.occurrence.keys():
            self.occurrence[char] = -1
        else:
            self.occurrence[char] = self.index

        self.index += 1
    def firstAppearing(self):
        result = ''
        resultIndex = self.index
        for k in self.occurrence.keys():
            if self.occurrence[k] != -1 and self.occurrence[k] < resultIndex:
                result = k
                resultIndex = self.occurrence[k]
        return result
```

### ==面试题51==：数组中的逆序对

```python
# -*- coding:utf-8 -*-
import copy
class Solution:
    def InversePairs(self, data):
        if type(data) != list or len(data) == 0:
            return 0
        datacopy = copy.copy(data)
        return self.InversePairsCore(data,datacopy, 0, len(data)-1)

    def InversePairsCore(self, data, datacopy, start, end):
        if start == end:
            datacopy[start] = data[start]
            return 0
        length = (end - start) >> 1

        left = self.InversePairsCore(datacopy, data, start, start + length)
        right = self.InversePairsCore(datacopy, data, start + length + 1, end)

        i = start + length
        j = end
        indexCopy = end
        count = 0

        while i >= start and j >= start + length + 1:
            if data[i] > data[j]:
                datacopy[indexCopy] = data[i]
                count += j - start - length
                indexCopy -= 1
                i -= 1
            else:
                datacopy[indexCopy] = data[j]
                indexCopy -= 1
                j -= 1
        while i>= start:
            datacopy[indexCopy] = data[i]
            indexCopy -= 1
            i -= 1
        while j>= start + length + 1:
            datacopy[indexCopy] = data[j]
            indexCopy -= 1
            j -= 1
        return left + right + count
```

### ==面试题52==：两个链表的第一个公共节点

```python
# -*- coding:utf-8 -*-
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
class Solution:
    def FindFirstCommonNode(self, pHead1, pHead2):
        if not pHead1 or not pHead2:
            return None
        nLength1 = self.GetListLength(pHead1)
        nLength2 = self.GetListLength(pHead2)
        nLengthDif = nLength1 - nLength2
        pListHeadLong = pHead1
        pListHeadShort = pHead2
        if nLength2 > nLength1:
            pListHeadLong = pHead2
            pListHeadShort = pHead1
            nLengthDif = nLength2 - nLength1

        for i in range(nLengthDif):
            pListHeadLong = pListHeadLong.next

        while pListHeadLong and pListHeadShort and pListHeadLong != pListHeadShort:
            pListHeadLong = pListHeadLong.next
            pListHeadShort = pListHeadShort.next
        return pListHeadLong

    def GetListLength(self,pHead):
        nLength = 0
        pNode = pHead
        while pNode:
            nLength += 1
            pNode = pNode.next
        return nLength

```

### ==面试题53==：在排序数组中查找数字

题目一：数字在排序数组中出现的次数

```python
# -*- coding:utf-8 -*-
class Solution:
    def GetNumberOfK(self, data, k):
        result = 0
        if isinstance(data,list) and len(data) > 0:
            first = self.getdown(data, 0, len(data)-1, k, lambda a,b: a>=b)
            last = self.getup(data, 0, len(data)-1, k, lambda a,b:a<=b)
            if first == last:
                if data[first] == k:
                    result = 1
            else:
                result = last - first + 1
        return result

    def getdown(self, arr, left, right, k, cmp):  # 找下界
        while(left < right):
            mid = (left + right) >> 1
            if cmp(arr[mid], k):
                right = mid
            else:
                left = mid + 1
        return left

    def getup(self, arr, left, right, k, cmp):  # 找上界
        while(left < right):
            mid = (left + right + 1) >> 1
            if cmp(arr[mid], k):
                left = mid
            else:
                right = mid - 1
        return left
```

题目二：0~n-1中缺失的数字

```python
# -*- coding:utf-8 -*-
class Solution:
    def GetMissingNumber(self, data):
        if not isinstance(data,list) or len(data) == 0:
            return -1

        left = 0
        right = len(data) - 1
        while left <= right:
            middle = (right + left) >> 1
            if data[middle] != middle:
                if middle == 0 or data[middle-1] == middle - 1:
                    return middle
                right = middle - 1
            else:
                left = middle + 1
        if left == len(data):
            return  left

        return -1

```

题目三：数组中数值和下标相等的元素

```python
# -*- coding:utf-8 -*-
class Solution:
    def GetNumberSameAsIndex(self,data):
        if not isinstance(data,list) or len(data) == 0:
            return -1

        left = 0
        right = len(data)
        while left <= right:
            middle = left + ((right - left)>>1)
            if data[middle] == middle:
                return middle
            elif data[middle] > middle:
                right = middle - 1
            else:
                left = middle + 1
        return -1
```

### ==面试题55==：二叉搜索树的第K大节点

```python
# -*- coding:utf-8 -*-
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    # 返回对应节点TreeNode
    def __init__(self):
        self.k = 0
    def KthNode(self, pRoot, k):
        if pRoot == None or k == 0:
            return None
        self.k = k
        return self.KthNodeCore(pRoot)

    def KthNodeCore(self, pRoot):
        target = None
        if pRoot.left:
            target = self.KthNodeCore(pRoot.left)
        if not target:
            if self.k == 1:
                target = pRoot
            self.k -= 1
        if not target and pRoot.right:
            target = self.KthNodeCore(pRoot.right)
        return target
```

### ==面试题55==：二叉树的深度

题目一：二叉树的深度

```python
# -*- coding:utf-8 -*-
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    def TreeDepth(self, pRoot):
        if pRoot == None:
            return 0
        nLeft = self.TreeDepth(pRoot.left)
        nRight = self.TreeDepth(pRoot.right)

        return nLeft + 1 if nLeft > nRight else nRight + 1
```

题目二：平衡二叉树

```python
# -*- coding:utf-8 -*-
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    def IsBalanced(self, pRoot):
        if pRoot == None:
            pDepth = 0
            return True,pDepth

        flag1, left = self.IsBalanced(pRoot.left)
        flag2, right = self.IsBalanced(pRoot.right)

        if flag1 and flag2:
            diff = left - right
            if diff <= 1 and diff >= -1:
                pDepth = 1 +  (left if left >right else right)
                return True, pDepth

        return False,0


# 牛客网False并不能返回两个值
# -*- coding:utf-8 -*-
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    def IsBalanced_Solution(self, pRoot):
        if pRoot == None:
            pDepth = 0
            return True,pDepth
        flag1 = self.IsBalanced_Solution(pRoot.left)
        flag2 = self.IsBalanced_Solution(pRoot.right)

        if isinstance(flag1,tuple) and isinstance(flag2,tuple):
            if flag1[0] and flag2[0]:
                left = flag1[1]
                right = flag2[1]
                diff = left - right
                if diff <= 1 and diff >= -1:
                    pDepth = 1 +  (left if left >right else right)
                    return True, pDepth
        return False
```

### ==面试题56==：数组中数字出现的次数

题目一：数组中只出现一次的两个数字

```python
# -*- coding:utf-8 -*-
class Solution:
    # 返回[a,b] 其中ab是出现一次的两个数字
    def FindNumsAppearOnce(self, array):
        if not isinstance(array,list) or len(array) < 2:
            return

        resultExclusiveOR = 0
        for item in array:
            resultExclusiveOR ^= item

        indexOf1 = self.FindFirstBitIs1(resultExclusiveOR)

        num1 = num2 = 0
        for item in array:
            if self.IsBit1(item,indexOf1):
                num1 ^= item
            else:
                num2 ^= item

        return [num1,num2]

    def FindFirstBitIs1(self, num):
        indexBit = 0
        while num & 1 == 0:
            num = num >> 1
            indexBit += 1
        return indexBit

    def IsBit1(self, num, indexBit):
        num = num >> indexBit
        return num & 1
```

题目二：数组中唯一只出现一次的数字

```python
# -*- coding:utf-8 -*-
#python2解 int 在32位机器上长度为32，在64位机器上长度为64
#python3解 int 无限制长度，以最大的数字为标准
class Solution:
    def FindNumsAppearOnce(self, array):
        if not  isinstance(array, list) or len(array) < 1:
            return
        length  = len(str(bin(max(array)))) - 2
        bitSum = [0] * length

        for item in array:
            bitMask = 1
            for i in range(length-1,-1,-1):
                bit = item & bitMask
                if bit != 0:
                    bitSum[i] += 1
                bitMask = bitMask << 1

        print(bitSum)
        result = 0
        for item in bitSum:
            result = result << 1
            result += item % 3

        return result

```

### ==面试题57==：和为s的两个数字

```python
# -*- coding:utf-8 -*-
class Solution:
    def FindNumbersWithSum(self, array, tsum):
        if not isinstance(array, list) or len(array) < 2:
            return []
        result = []
        behind = 0
        ahead = len(array) - 1

        while ahead > behind:
            curSum = array[behind] + array[ahead]

            if curSum == tsum:
                result.append( [array[behind], array[ahead]])
                ahead -=1
            elif curSum > tsum :
                ahead -= 1
            else:
                behind += 1
        if len(result) != 0:
            #牛客网要求
            return sorted(result,key = lambda x: x[0] * x[1])[0]
        return result
```

题目二：和为s的连续正数序列

```python
# -*- coding:utf-8 -*-
class Solution:
    def FindContinuousSequence(self, tsum):
        if tsum < 3:
            return []
        result = []
        small = 1
        big = 2
        middle = (1+tsum) >> 1
        curSum = small + big
        while small < middle:
            if curSum == tsum:
                result.append(range(small, big+1))
                big += 1
                curSum += big
            elif curSum > tsum:
                curSum -= small
                small += 1
            else:
                big += 1
                curSum += big
        return result
```

### ==题目58==：翻转字符串

题目一：翻转单词顺序

```python
# -*- coding:utf-8 -*-
class Solution:
    def ReverseSentence(self, s):
        s = s.split(' ')
        return ' '.join(reversed(s))
```

题目二：左旋转字符串

```python
# -*- coding:utf-8 -*-
class Solution:
    def LeftRotateString(self, s, n):
        s = ''.join(reversed(s[:n]))+''.join(reversed(s[n:]))
        return ''.join(reversed(s))

# -*- coding:utf-8 -*-
class Solution:
    def LeftRotateString(self, s, n):
        return s[n:]+s[:n]
```

### ==面试题59==：队列的最大值

题目一：滑动窗口的最大值

```python
# -*- coding:utf-8 -*-
import collections
class Solution:
    def maxInWindows(self, num, size):
        result = []
        if isinstance(num, list) and len(num) >= size and size >= 1:
            d = collections.deque()
            for i in range(size):
                while len(d) != 0 and num[i] >= num[d[0]]:
                    d.popleft()
                d.appendleft(i)
            for i in range(size, len(num)):
                result.append(num[d[-1]])

                while len(d) !=0 and num[i] >= num[d[0]]:
                    d.popleft()
                if len(d) != 0 and  d[-1] <= (i-size):
                    d.pop()
                d.appendleft(i)
            result.append(num[d[-1]])
        return result
```

题目二：队列的最大值

```python
# -*- coding:utf-8 -*-
import collections
class Solution:
    def __init__(self):
        self.data = collections.deque()
        self.maximums = collections.deque()  #滑动窗口看出一个队列
        self.currentIndex = 0

    def push_back(self, number):
        while len(self.maximums) != 0 and number >= self.maximums[0][0]:
            self.maximums.popleft()
        numberdata = (number, self.currentIndex)
        self.data.appendleft(numberdata)
        self.maximums.appendleft(numberdata)
        self.currentIndex += 1

    def pop_front(self, number):
        if len(self.maximums) == 0:
            raise Exception('queue is empty')
        if self.maximums[-1][1] == self.data[-1][1]:
            self.maximums.pop()
        self.data.pop()

    def max(self):
        if len(self.maximums) == 0:
            raise Exception('queue is empty')
        return self.maximums[-1][0]
```

### ==面试题60==：n个骰子的点数

```python
# -*- coding:utf-8 -*-
class Solution:
    def __init__(self, g_maxValue = 6):
        self.g_maxValue = g_maxValue

    def PrintProbability(self, number):
        if number < 1:
            return

        pProbabilities = [[0]*(self.g_maxValue * number + 1), [0]*(self.g_maxValue * number + 1)]
        flag = 0

        for i in range(1, self.g_maxValue+1):
            pProbabilities[flag][i] = 1

        for k in range(2, number+1):
            for i in range(k):
                pProbabilities[1 - flag][i] = 0
            for i in range(k, self.g_maxValue*k + 1):
                pProbabilities[1 - flag][i] = 0
                j = 1
                while j<=i and j<=self.g_maxValue:
                    pProbabilities[1-flag][i] += pProbabilities[flag][i-j]
                    j+=1
            flag = 1 - flag

        total = self.g_maxValue ** number
        for i in range(number, self.g_maxValue+1):
            raito = pProbabilities[flag][i] / total
            print(f'{i} : {raito}')
```

### ==面试题61==：扑克牌中的顺子

```python
# -*- coding:utf-8 -*-
class Solution:
    def IsContinuous(self, numbers):
        if not isinstance(numbers, list) or len(numbers) == 0:
            return False

        numbers = sorted(numbers)

        numberOfZero = numbers.count(0)
        numberOfGap = 0

        small = numberOfZero
        big = small + 1
        while big < len(numbers):
            if numbers[small] == numbers[big]:
                return False
            numberOfGap += (numbers[big] - numbers[small] - 1)
            small = big
            big += 1
        return False if numberOfGap > numberOfZero else True
```

### ==面试题62==：圆圈中最后剩下的数字

```python
#使用数学寻找简单的方式
# -*- coding:utf-8 -*-
class Solution:
    def LastRemaining_Solution(self, n, m):
        if n<1 or m<1:
            return -1
        last = 0
        for i in range(2, n+1):
            last = (last + m) % i
        return last
```

### ==面试题63==：股票的最大利润

```python
# -*- coding:utf-8 -*-
class Solution:
    def MaxDiff(self, numbers):
        if not isinstance(numbers, list) or len(numbers) < 2:
            return 0

        minOfNumbers = numbers[0]
        maxDiff = numbers[1] - minOfNumbers

        for i in range(2, len(numbers)):
            if numbers[i-1] < minOfNumbers:
                minOfNumbers = numbers[i-1]

            currentDiff = numbers[i] - minOfNumbers
            if currentDiff > maxDiff:
                maxDiff = currentDiff

        return maxDiff
```

### ==面试题64==：求 1+2+...+n

```python
# x and y ， 如果x 为False， x and y 返回 False，否则它返回y的计算值
# 利用python的逻辑运算符，and 重复调用函数已达到累加的效果
# -*- coding:utf-8 -*-
class Solution:
    def __init__(self):
        self.sum = 0
    def Sum_Solution(self, n):
        def qiusum(n):
            self.sum += n
            n -= 1
            return n>0 and self.Sum_Solution(n)

        qiusum(n)
        return self.sum
```

### ==面试题65==：不用加减乘除做加法

```python
# -*- coding:utf-8 -*-
class Solution:
    def Add(self, num1, num2):
        if num1 + num2 == 1:
            return 1
        su = num1^num2
        carry = (num1 & num2)<<1
        num1,num2 = su,carry
        while str(num2) != '0' :
            num1,num2 = num1^num2,(num1 & num2)<<1
        return num1
```

### ==面试题68==：树中两个节点的最低公共祖先

| 情况 | 解题思路|
|---|---|
| 树为二叉搜索树，找最低公共祖先 | 当前节点 都大于p,q，则最低公共祖先在它的左子树；当前节点 都小于p,q 则最低公共祖先在它的右子树；当前节点在 p，q之间（包括等于) 则该节点就是最低公共祖先节点
| 树包含父节点 | 这个问题，可以转换为 两个链表的第一个公共节点
| 树是一课普通的树| 从根节点出发，看当前节点其子树是否包含p，q； 若当前节点子树包含p,q；而当前节点的所有子节点对应的子树都不包含p,q
| 树是一课普通的树| 利用辅助空间，把根节点到p,q路径记录下成两个数组，找两个路径数组的最后一个相同节点；路径用递归+回溯
