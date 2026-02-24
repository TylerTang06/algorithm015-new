# 学习笔记

## 递归

- 写出正确的递推公式
    - 不要试图用人脑去分解递归的每个步骤
- 警惕堆栈溢出
    - 如果递归求解的数据规模很大，调用层次很深，一直压入栈就会有堆栈溢出的风险
    - 在代码中限制最大递归深度超过深度就抛错
- 警惕重复计算
    - 通过比如散列表保存已求解的结果，空间复杂度从 O(1) 变成 O(n)
- 警惕数据造成死循环
    - 比如A的推荐人是B，B的推荐人是C，C的推荐人是A
    - 解决方式是判断环
- 复杂度分析，若每次分区操作，都是正好分成左右两个大小接近的分区，则时间复杂度为O(nlogn)，比如快排、归并
- 优缺点
    - 逻辑清晰，代码简洁
    - 容易堆栈溢出，重复计算

## 递归模版

### golang

```golang
func recursion(level, n int, params ...interface{}) {
    // recursion terminator
    if level > n {
        // process result
        processResult()

        return 
    }

    // process current logic
    process(level, params)

    // drill down
    recursion(level+1, n, params)

    // reverse current status if needed

    return
}
```

### 思维要点

- 不要人肉进行递归
- 找到最近最简单方法，将其拆分成可重复解决的问题
- 数学归纳法思想

## 分治、回溯

### 分治

- problem->subproblems
    - Divide & Conquer + Optional substructure分治 + 最优子结构
    - 状态的存储
    - 在每一步中淘汰次优解，最保留这一步的最优状态
- split->merge
    - 得到最优解

```golang
func divideConquer(problem interface{}, params ...interface{}){
    // recursion terminator
    if problem == nil {
        processResult()

        return 
    }

    // prepare data
    data := prepareData(problem)
    subproblems := splitProblem(problem, data)

    // conquer subproblems
    subproblem1 := divideConquer(subproblems[0], params)
    subproblem2 := divideConquer(subproblems[1], params)
    ...

    // process and generate final reslut
    result := processResult(subproblem1, subproblem2, ...)

    return
}
```

### 回溯

- 采用试错的思想(尝试分步去解决一个问题)
- 分步解决问题过程中
    - 找到一个可能的答案
    - 找不到答案，取消上一步或几步，通过其他可能分步解答尝试找到答案

## 总结

递归是比分治、回溯更广泛的概念，甚至包括动态规划。回溯一般是要找到所有可能的答案。需注意递归、分治与回溯模版的细节差异之处。