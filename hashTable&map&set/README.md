# 学习笔记

## 哈希表、映射、集合

### Hash Table

- 哈希表(Hash Table)/散列表
- 哈希函数(Hash Function)
- 哈希碰撞常用拉链法解决
- 缓存(LRU Cache)、健值对存储(Redis)等
- Average: Search O(1), Insertion O(1), Delete O(1)
- Worst: Search O(n), Insertion O(n), Delete O(1)

### Map: key-value, key不重合

- new HashMap()/new TreeMap()
- map.set(key, value)
- map.get(key)
- map.has(key)
- map.size()
- map.clear()

### Set: 不重合元素的集合

- new HashSet()/new TreeSet()
- set.add(value)
- set.delete(value)
- set.has(value)

## 树、二叉树、二叉搜索树

### 二叉搜索树

- 左子树上所有结点的值均小于它的根结点的值
- 右子树上所有结点的值均大于它的根结点的值
- 左右子树分别为二叉搜索树
- 中序遍历: 升序遍历
- 查询、插入、删除平均时间复杂度 O(logn)

## 堆、二叉堆、图

### 堆 Heap

- 可以迅速找到一堆数中的最大或者最小值的数据结构
- 大顶堆 与 小顶堆
- 大顶堆: find-max O(1); delete-max O(logn); insert O(logn) or O(1)

### 二叉堆

- 通过完全二叉树实现
- 是一棵完全树
- 树中任意结点的值总是>=其子结点的值
- 数组实现二叉树, 假设第一个结点在索引0位置:
    - 索引为i的左子树索引为 2*i+1
    - 索引为i的右子树索引为 2*i+2
    - 所以为i的结点的父结点索引为 floor((i-1)/2)

### 图

- 临接矩阵
- 临接表

### 解题技巧

- Hashset
    - 使用Hashset往往是以空间换时间，golang中使用map实现set，比如异位词、两数之和
- Tree
    - Tree满足递归性质的结构体，比如最近公共祖先、树的遍历、树的最大深度等
    - 树的前中后序遍历递归实现很容易，非递归遍历需借助其他结构体(中序遍历较复杂)

    ```golang
    // 非递归中序遍历二叉树，中序遍历的顺序为：左中右，借助堆栈时，将结点压入的顺序为：右中左
    // myStack := list.New()
    // 树的每个节点需要放入Stack两次
    for myStack.Len() > 0 {
		node := myStack.Back().Value.(*TreeNode)
        myStack.Remove(myStack.Back())
        // 当第一次取到node时
		if 1 == myMap[node] {
			if node.Right != nil {
                // 先将其右结点压入堆栈
				myStack.PushBack(node.Right)
				myMap[node.Right] = 1
            }
            // 将node第二次压入堆栈
			myStack.PushBack(node)
			myMap[node] = 2
			if node.Left != nil {
                // 再将其左子树压入堆栈
				myStack.PushBack(node.Left)
				myMap[node.Left] = 1
			}
		} else {
            // 当第二次从堆栈中取到node时，直接记录node
			res = append(res, node.Val)
		}
	}
    ```
- Heap
    - 最小堆/最大堆，比如前k个高频词、前k个最小(大)数
    ```golang
    // 借助sort接口实现最小堆
    type MinIntHeap []int

    func (a MinIntHeap) Len() int {
        return len(a)
    }

    func (a MinIntHeap) Swap(i, j int) {
        a[i], a[j] = a[j], a[i]
    }

    func (a MinIntHeap) Less(i, j int) bool {
        return a[i] < a[j]
    }

    func (a *MinIntHeap) Push(value interface{}) {
        *a = append(*a, value.(int))
    }

    func (a *MinIntHeap) Pop() interface{} {
        old := *a
        value := old[len(old)-1]
        *a = old[:len(old)-1]

        return value
    }

    // 使用实例，初始化
    myMinHeap := &MinIntHeap{}
    heap.Init(myMinHeap)
    // 将弹出myMinHeap最大的数
    heap.Pop(myMinHeap)
    // 将val压入最小堆
    heap.Push(myMinHeap, val)
    ```

## HashMap

- Golang中可基于Map实现
- 参考链接: https://github.com/emirpasic/gods

### Golang Map实现原理

- 哈希表作为底层数据结构
- 哈希表结点: bucket, 每个bucket含有一对或多对key-value
- 每个bucket可以存储8个key-value
- 哈希冲突: 链地址法
- 负载因子: 负载因子 = 键数量/bucket数量
- 扩容条件：
    - 负载因子 > 6.5时，也即平均每个bucket存储的键值对达到6.5个
    - overflow数量 > 2^15时，也即overflow数量超过32768时
- 扩容方式: 增量扩容、等量扩容

```golang
// map 数据结构
type hmap struct {
    count int // 当前保存的元素个数
    ...
    B unit
    ...
    buckets unsafe.Pointer // bucket数组指针, 数组的大小为2^B
}

// bucket 数据结构
type bucket struct {
    tophash [8]unit // 存储哈希值的高8位
    data byte[1] // key-value数据: key/key/key/.../value/value/value/...;如此存放是为了节省字节对齐带来的空间浪费
    overflow *bmap // 溢出bucket的位置, 指针指向的是下一个bucket, 据此将所有冲突的键连接起来
}
```

# 总结

二叉树是最适合使用递归的数据结构体之一；相对递归，使用循环的方式实现树的遍历，需要借队列、堆、map等结构体。

golang标准库中没有set这样的结构体，但是map可以到达相同的效果；同时，list是双端链表可以实现队列、堆栈、双端队列等结构体；golang中heap不能开箱即用，它提供了相应的接口供开发者根据自己场景实现heap。

