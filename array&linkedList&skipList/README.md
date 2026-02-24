# 学习笔记

## 数组、链表、跳表

### 数组

- 内存中一块连续的地址
- 同一种数据类型的固定长度的序列，golang中数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变
- 访问数组时间复杂度O(1)，插入删除操作时间复杂度为O(n)
- golang中slice底层是基于数组实现，len(),cap()时间复杂度均为O(1)
- 切片扩容规则：如果原Slice容量小于1024，则新Slice容量将扩大为原来的2倍；如果原Slice容量大于等于1024，则新Slice容量将扩大为原来的1.25倍

### 链表

- 链表的元素中包含指向临近元素的指针
- 查找操作时间复杂度O(n)
- 插入删除操作时间复杂度O(1)
- 链表的元素在内存中不是连续的，没有容量的概念
- golang标准库实现了container/list双向链表

### 跳表

- 跳表是链表的升维结构，在原始的链表上新增K个层级的索引，类似二分法
- 跳表存储数据必须有序
- 维护成本较高，每次修改元素都要修改每个层级的索引
- 插入、删除、搜索操作时间复杂度O(log n)
- 空间复杂度为O(n)，典型的空间换时间的方法
- 原理简单、容易实现、方便扩展、效率更高，Redis、LevelDB中均有使用
- 维护成本较高，每次修改元素都要修改每个层级的索引

## 栈、队列、优先队列、双端队列

### 栈、队列

- 栈先进后出LIFO(Last In First Out)
- 队列先进先出FIFO(First In Fist Out)
- 栈、队列查询复杂度均为O(n)

### 优先队列

- 插入操作时间复杂度O(1)
- 取值操作时间复杂度O(log n)
- goalng可以使用container/heap接口方便实现优先队列

### 双端队列

- 两端都可以提取、插入操作的队列
- golang可以使用container/list模拟栈、对列、双端队列

### 解题技巧

- Array
    - 双指针法-操作快慢指针，比如移动0到数组最后
    - 双指针法-左右收敛，比如有序数组两数之和，数组三数之和(需先排序)，装水最多的水桶，数组翻转
    - 循环数组，比如LRU缓存设计，数组队列，双端队列设计

    ```golang
    // Adds an item at the front of Deque
    this.front = (this.front - 1 + len(this.q)) % len(this.q)
    
    // Adds an item at the rear of Deque
    this.rear = (this.rear + 1) % len(this.q)
    
    // Deletes an item from the front of Deque
    this.front = (this.front + 1) % len(this.q)
    
    // Deletes an item from the rear of Deque
    this.rear = (this.rear - 1 + len(this.q)) % len(this.q)
    ```

- Linked
    - 递归，链表的结构满足递归性质
    - 快慢指针，比如判断循环链表，输出倒数第k个结点
    - 交换链表nodes的时候，注意前置结点信息的保存
    - 合并两链表code技巧

    ```golang
    for l2 != nil {
        if l1 == nil { // 1.当l1== nil,直接把l2赋给l1
			l1, l2 = l2, nil
			continue
		}
		if l1.Val < l2.Val {
			l.Next, l1 = l1, l1.Next
		} else {
			l.Next, l2 = l2, l2.Next
		}
		l = l.Next
	}
	if l1 != nil { // 2.基于1，最后永远是l1不为nil
		l.Next = l1
	}
    ```

- Stack/Queue/Deque
    - 使用Stack典型的题目包括：括号匹配、接雨水、柱状内最大矩形等
    - 使用Queue典型场景BFS场景
    - 使用Deque最典型的题目：滑动窗口

    ```golang
    // myQue := list.New()
    // 技巧：如果双端队列尾的元素小于新到来的元素，则从队尾移除元素，这样保证了对列元素的有序性
    for myQue.Len() > 0 && nums[myQue.Back().Value.(int)] <= val {
		myQue.Remove(myQue.Back())
	}
    ```

## 总结

本周队常用数据结构进行了复习，对跳表、优先队列等还需结合具体使用场景加强理解；本周仓促中能够完成基本任务，还需继续多投入时间。
学习中比较有疑问的地方，这些数据结构如何保证并发安全？