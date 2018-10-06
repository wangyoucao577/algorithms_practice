# [Golang] 排序
学习各种排序的实验代码，主要参考《算法导论 第3版》，包括:     
- 比较排序算法: 插入排序(Insertion Sort)、归并排序(Merge Sort)、堆排序(Heap Sort)、快速排序(Quick Sort)    
- 非比较排序算法: 计数排序(Counting Sort)、基数排序(Radix Sort)、桶排序(Bucket Sort)     
- 其他算法: 选择第`n-th`元素(Select n-th Element)

## Tips
- 比较排序算法的最坏情况下界为 `O(n*log(n))`.     
    - 即在最坏情况下, 任何比较排序算法都需要做至少 `O(n*log(n))` 次比较.     
- 非比较排序算法则不受此`O(n*log(n))`的限制, 可以在满足一些预设条件的前提下达到线性运行时间.         

## 实验代码
- mysorts    
实现各种排序算法的实验代码的包, 其中比较排序算法主要基于`sort.Interface`来实现. (可通过命令`go test -v -bench=. -benchmem`来进行功能测试和Benchmark)     
    - `insertion_sort.go`: **O(n^2)**     
        - 实现《算法导论 第3版》 ch2.1 介绍的插入排序算法, 算法原理为从第2个元素开始遍历整个数组, 对每个元素, 向前进行遍历比较, `Less()`条件为`false`则`Swap()`, 直到`Less()`条件为`true`是退出. 由于是`in-place`实现, 空间复杂度为 **O(1)**.     
    - `merge_sort.go`: **O(n*log(n))**     
        - 实现《算法导论 第3版》 ch2.3 介绍的归并排序算法. 算法原理为分治法(归并排序其实是二分法), 即将问题分解为2个子问题, 求解每个子问题(对于每个子问题可递归地继续分解, 直至不可分割), 然后递归地向上合并子问题的解.    
            - `aux array based implementation`: 书上介绍的方法, 每次`merge`时需要对每个子问题申请一块辅助的子数组内存来暂存子问题, 空间复杂度为 **O(n)**. 时间复杂度为 **O(n*log(n))**, 缺点为需要申请额外的空间, 以及无法使用典型的`Swap`接口来实现排序过程.    
            - `in-place implementation`: 不需要辅助空间的实现, 借鉴了`insertion_sort`的方法来实现`merge`时的`in-place`. 空间复杂度为 **O(1)**. Benchmark 实测的运行时间比`aux array based implementation`慢很多, 接近`insertion_sort`(比它稍快).    
    - `heap_sort.go`: **O(n*log(n))**     
        - 实现《算法导论 第3版》 ch6.1~6.4 介绍的堆排序算法. 算法原理为借助Heap(一般是用二叉堆) 的性质, 即`root`元素总是最大的(`maxHeap`中; 若是`minHeap`则反之). 首先构建出`maxHeap`, 那么`root`元素一定是数组中的最大值. 将`root`元素与数组最后元素交换, 再针对新的`root`节点维护堆的性质则又可以获得剩余堆(除刚刚的最大元素)中的最大值, 如此循环操作直至遍历完整个数组.     
    - `quick_sort.go`: 平均运行时间接近最好运行时间 **O(n*log(n))**, 最坏情况为 **O(n^2)**.     
        - 实现《算法导论 第3版》ch7 介绍的快速排序算法, 包含固定主元(fixed pivot element)和随机主元(randomized pivot element)两种实现. 算法原理依然基于分治法, 即将问题分解为两个子问题分别求解. 而与归并排序不同的是, 在分解问题时总是选择一个`pivot element`, 将原数组分解为`<= pivot lement`和`>= pivot element`这样两个子问题, 并把`pivot element`放在两个子数组中间, 于是合并操作不需要再做任何事情.     
        - 快速排序的运行时间主要取决于分治(`partition`)的过程, 期望能够将问题分割为两个尽量平衡的子问题(其实只要两个子问题为常数比例即可, 如`9:1`, `99:1`均可达到 `O(n*log(n))` ). 当分割的两个子数组其中一个长度为0时, 即最坏的情况 **O(n^2)**.       
    - `counting_sort.go`: **O(k+n)** 其中`k`为输入数组的最大元素值(限定输入数组元素值范围为`0~k`).     
        - 实现《算法导论 第3版》ch8.2 介绍的计数排序. 其原理为引入一个以输入数组元素值为索引的辅助数组, 对于任意一个输入数组上的元素`x`, 直接在辅助数组上计算它前面有多少元素(也即输入数组中有多少元素小于`x`), 从而`x`就可以直接放在对应的位置上了. 其计算过程不需要任何的元素比较.      
        - 不是比较排序, 所以可以不受比较排序的最坏情况下界 **O(n*log(n))** 的限制.    
        - 需要注意的是`k`的值要可控, 即输入数组中的最小值和最大值间的范围不能太大, 否则空间开销太大, 此方法就不适用了.     
        - 当 k=O(n) 时, 其运行时间可简化为 **O(n)**.    
        - 另外, 当前的实现需要限定输入数组元素值范围为`0~k`, 实际可以比较容易地改造下以支持负数.    
    - `radix_sort.go`: **O(d(k+n))** 其中`d`为输入数组元素的最大位数, `k`为每位上的最大值.     
        - 实现《算法导论 第3版》ch8.3 介绍的基数排序, 其原理为将输入数组看做一个`d`列`n`行的矩阵, 依次地从最低位至最高位对每列进行排序, 最终即可得到排序好的数组. 对于每列的排序需要依赖于稳定的排序算法(典型如`counting sort`就很适用).     
        - 基数排序非常适合用来解决多关键字记录的排序问题, 典型如对于日期的排序(将年、月、日分别看做位).     
        - 基数排序虽然看起来是线性排序, 运行低于快速排序等比较排序方法. 但实际使用时由于快速排序的每次循环更快、`in-place`等因素, 通常快速排序更好.     
    - `bucket_sort.go`: 平均情况下 **O(n)**, 且需假设前提为输入数组服从均匀分布.     
        - 实现《算法导论 第3版》ch8.4 介绍的桶排序, 又一种非比较排序的方法. 其原理为将所有的输入元素均匀的放到`n`个桶中, 然后对每个桶内进行排序, 最后按次序遍历桶放回原数组即可得到排序后的结果. 输入数组服从均匀分布的前提下, 每个桶内的元素个数是接近且均匀的, 于是每个桶内排序(一般采用`insertion sort`, 当然也可以采用`counting sort`等方法)就会很快, 从而整体可以做到线性.      
    - `selection_nth.go`: 期望运行时间 **O(n)**, 最坏情况下 **O(n^2)**.    
        - 实现《算法导论 第3版》ch9.2 介绍的选择第`n-th`大的元素的问题, 更多可参考 [Selection Algorithm - Wikipedia](https://en.wikipedia.org/wiki/Selection_algorithm). 其原理是基于`quick sort`的`randomized partition`, 同样采用分治法递归地将数组分割为两个子数组, 与`quick sort`不同的是, 只需要选取到的`pivot element`是第`n-th`元素即可返回, 且递归时两个子数组中只需要处理一个即可. 当取到了第`n-th`元素时, 即意味着数组中在其前面的都已经小于等于它, 而其后面的都已经大于等于它.     
            - 猜测一下, [std::nth_element](https://en.cppreference.com/w/cpp/algorithm/nth_element) 和 [std::partial_sort](https://zh.cppreference.com/w/cpp/algorithm/partial_sort) 均应是基于此原理实现的.    


## References
- 《算法导论 第3版》    
- [Package sort](https://golang.org/pkg/sort/)
- [Package heap](https://golang.org/pkg/container/heap/)
- [Divide and Conquer Alogrithm](https://en.wikipedia.org/wiki/Divide_and_conquer_algorithm)
- [Selection Algorithm - Wikipedia](https://en.wikipedia.org/wiki/Selection_algorithm)
- [std::nth_element](https://en.cppreference.com/w/cpp/algorithm/nth_element)
- [std::partial_sort](https://zh.cppreference.com/w/cpp/algorithm/partial_sort)

