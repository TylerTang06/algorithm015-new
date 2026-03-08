# 学习笔记

## 位运算

```golang
// XOR
// x ^ 0 = x
// x ^ 1s = ~x // 1s = ~0
// x ^ (~x) = 1s
// x ^ x = 0
// a ^ b = c, a ^ c = b, b ^ c = a // swap
// a ^ b ^ c = a ^ (b ^ c) = (a ^ b) ^ c // associative
func XOR(a, b int) int {
	return a ^ b
}

// Clean n bits to 0 on the right
func ToZeroRight(x, n int) int {
	return x & (^int(0) << n)
}

// get value of the nth bit
func GetBitOfN(x, n int) int {
	return (x >> n) & 1
}

// get power value of the nth bit
func GetPowerOfN(x, n int) int {
	return x & (1 << (n - 1))
}

// update value of the nth bit to 1
func UpdateToOneOfN(x, n int) int {
	return x | (1 << n)
}

// update value of the nth bit to 0
func UpdateToZeroOfN(x, n int) int {
	return x & (^(1 << n))
}

// update value from nth to mth bits to 0,
// mth is the Hightest bit
func UpdateToZeroFromN(x, n int) int {
	return x & ((1 << n) - 1)
}

// upate value from 0th to nth bits
func UpdateToZeroToN(x, n int) int {
	return x & (^((1 << (n + 1)) - 1))
}

// clean the first 1 on the right
func ToZeroFirstOneRight(x int) int {
	return x & (x - 1)
}

// get the first 1 on the right
func GetFirstOneRight(x int) int {
	return x & -x
}

//  Number of 1 Bits
func HammingWeight(num uint32) int {
	times := 0
	for num != 0 {
		num = num & (num - 1)
		times++
	}
	return times
}

// Power of Two
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return (n & (n - 1)) == 0
}

// Given a non negative integer number num.
// For every numbers i in the range 0 ≤ i ≤ num
// calculate the number of 1's in their binary representation
// and return them as an array.
func CountBits(num int) []int {
	counts := make([]int, num+1)
	for i := 1; i <= num; i++ {
		counts[i] += counts[i&(i-1)] + 1
	}
	return counts
}
```

## 排序

```golang
// BubbleSort return sorted array by use Bubble Sort Method
func BubbleSort(sli []int) []int {
	if sli == nil || len(sli) <= 1 {
		return sli
	}

	flag := 1
	for i := 0; i < len(sli) && flag == 1; i++ {
		flag = 0
		for j := 0; j < len(sli)-i-1; j++ {
			if sli[j] > sli[j+1] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
				flag = 1
			}
		}
	}

	return sli
}

// SelectSort return sorted array by use Select Sort Method
func SelectSort(sli []int) []int {
	if sli == nil || len(sli) <= 1 {
		return sli
	}

	for i := 0; i < len(sli)-1; i++ {
		min := i
		for j := i + 1; j < len(sli); j++ {
			if sli[min] > sli[j] {
				min = j
			}
		}
		if min != i {
			sli[i], sli[min] = sli[min], sli[i]
		}
	}

	return sli
}

// QuickSort return sorted array by use Quick Sort Method
func QuickSort(sli []int) []int {
	if sli == nil || len(sli) <= 1 {
		return sli
	}

	left, right, min := 0, len(sli)-1, sli[0]
	for left < right {
		for left < right {
			if min > sli[right] {
				sli[right], sli[left] = sli[left], sli[right]
				break
			}
			right--
		}
		for left < right {
			if min < sli[left] {
				sli[left], sli[right] = sli[right], sli[left]
				break
			}
			left++
		}
	}
	//sli[left] = min && left == right
	QuickSort(sli[:left])
	QuickSort(sli[left+1:])

	return sli
}

// InsertSort return sorted array by use Insert Sort Method
func InsertSort(sli []int) []int {
	if sli == nil || len(sli) <= 1 {
		return sli
	}

	for i := 0; i < len(sli)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if sli[j] < sli[j-1] {
				sli[j], sli[j-1] = sli[j-1], sli[j]
			} else {
				break
			}
		}
	}

	return sli
}

// ShellSort return sorted array by use Shell Sort Method
func ShellSort(sli []int) []int {
	if sli == nil || len(sli) <= 1 {
		return sli
	}

	k := len(sli)
	for k > 0 {
		for i := 0; i < len(sli)-k; i++ {
			for j := i + k; j >= k; j -= k {
				if sli[j] < sli[j-k] {
					sli[j], sli[j-k] = sli[j-k], sli[j]
				} else {
					break
				}
			}
		}
		k /= 2
	}

	return sli
}

// MergeSort return sorted array by use Merge Sort Method
func MergeSort(sli []int) []int {
	if sli == nil || len(sli) <= 1 {
		return sli
	}

	mid := len(sli) / 2
	sliLeft := MergeSort(append([]int{}, sli[:mid]...))
	sliRight := MergeSort(append([]int{}, sli[mid:]...))
	for i := 0; i < len(sli); {
		left, right := 0, 0
		for right < len(sliRight) && left < len(sliLeft) {
			if sliLeft[left] < sliRight[right] {
				sli[i] = sliLeft[left]
				left++
			} else {
				sli[i] = sliRight[right]
				right++
			}
			i++
		}
		for left < len(sliLeft) {
			sli[i] = sliLeft[left]
			left++
			i++
		}
		for right < len(sliRight) {
			sli[i] = sliRight[right]
			right++
			i++
		}
	}

	return sli
}

// HeapSort return sorted array by use Heap Sort Method
func HeapSort(sli []int) []int {
	if sli == nil || len(sli) <= 1 {
		return sli
	}

	for i := (len(sli) - 2) / 2; i >= 0; i-- {
		sli = downAjust(sli, i, len(sli))
	}
	for i := len(sli) - 1; i >= 0; i-- {
		sli[i], sli[0] = sli[0], sli[i]
		sli = downAjust(sli, 0, i)
	}

	return sli
}

func downAjust(sli []int, root, n int) []int {
	if root >= len(sli) {
		return sli
	}

	child := 2*root + 1
	rootValue := sli[root]
	for child < n {
		if child+1 < n && sli[child] < sli[child+1] {
			child++
		}
		if sli[child] < rootValue {
			break
		}
		sli[root] = sli[child]
		root = child
		child = 2*child + 1
	}
	sli[root] = rootValue

	return sli
}

// Bucket struct
type Bucket struct {
	list *list.List
}

// NewBucket return *Bucket
func NewBucket() *Bucket {
	list := list.New()
	return &Bucket{
		list: list,
	}
}

// Size rerurn size of the bucket
func (b *Bucket) Size() int {
	return b.list.Len()
}

// Push the value to back of bucket
func (b *Bucket) Push(value int) {
	b.list.PushBack(value)
}

// Pop return the first value of bucket
func (b *Bucket) Pop() int {
	if b.Size() <= 0 {
		panic("no value to pop")
	}
	value := b.list.Front().Value.(int)
	b.list.Remove(b.list.Front())
	return value
}

// BucketSort ...
func BucketSort(sli []int) []int {
	if sli == nil || len(sli) <= 1 {
		return sli
	}

	var buckets [10]*Bucket
	for i := 0; i < 10; i++ {
		buckets[i] = NewBucket()
	}
	max := sli[0]
	for i := 0; i < len(sli); i++ {
		if max < sli[i] {
			max = sli[i]
		}
	}

	bit := 1
	for max > 0 {
		for i := 0; i < len(sli); i++ {
			buckets[(sli[i]/bit)%10].Push(sli[i])
		}
		count := 0
		for b := 0; b < 10; b++ {
			for s := buckets[b].Size() - 1; s >= 0; s-- {
				sli[count] = buckets[b].Pop()
				count++
			}
		}
		bit *= 10
		max /= 10
	}

	return sli
}
```

## 布隆过滤器

待续...

## LRUCache

```golang
// 双链表+map(hash)实现
// 可以自己实现双链表，节约内存
type LRUCache struct {
	cache *list.List
	val   map[int]*list.Element
	cap   int
}

func Constructor(capacity int) LRUCache {
	l := list.New()
	return LRUCache{cache: l, cap: capacity, val: make(map[int]*list.Element, capacity)}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.val[key]
	if !ok {
		return -1
	}

	this.cache.Remove(v)
	val := v.Value.([2]int)
	this.cache.PushBack(val)
	this.val[key] = this.cache.Back()
	return val[1]
}

func (this *LRUCache) Put(key int, value int) {
	if val := this.Get(key); val == -1 || val != value {
		if val != -1 && val != value {
			this.cache.Remove(this.val[key])
		}
		if this.cache.Len() == this.cap {
			v := this.cache.Front()
			this.cache.Remove(v)
			delete(this.val, v.Value.([2]int)[0])
		}
		this.cache.PushBack([2]int{key, value})
		this.val[key] = this.cache.Back()
	}
}
```

## 总结

- 位运算、排序、布隆过滤器、LRUCache都是很重要的基础，需加强