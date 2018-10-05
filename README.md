# algorithms_practice
Learn and practice algorithms and data structures. 来源包括但不限于《算法导论》, [hankerrank](https://www.hackerrank.com), [leecode](https://leetcode.com/), [北京大学 PKU JudgeOnline](http://poj.org/), etc.


## 实验环境
编译与运行主要在`Linux`上进行, 但代码都应属于平台无关的, 理论上不限于操作系统/编译器等.     
- `Ubuntu 18.04 LTS`: `WSL(Windows Subsystem for Linux) on Win10 1803-17134.228`
- `cmake version 3.10.2`
- `gcc (Ubuntu 7.3.0-16ubuntu3) 7.3.0`
- `go version go1.10.1 linux/amd64`

## 实验代码

### [Golang] 排序
- mysorts    
实现各种排序算法的实验代码, 基于`sort.Interface`来实现. (可通过命令`go test -v -bench=. -benchmem`来进行功能测试和Benchmark)     
    - `insertion_sort.go`: **O(n^2)** 实现《算法导论 第3版》 ch2.1 介绍的插入排序算法, 算法原理为从第2个元素开始遍历整个数组, 对每个元素, 向前进行遍历比较, `Less()`条件为`false`则`Swap()`, 直到`Less()`条件为`true`是退出. 由于是`in-place`实现, 空间复杂度为 **O(1)**.     
    - `merge_sort.go`: **O(n*log(n))** 实现《算法导论 第3版》 ch2.3 介绍的归并排序算法. 算法原理为分治法(归并排序其实是二分法), 即将问题分解为2个子问题, 求解每个子问题(对于每个子问题可递归地继续分解, 直至不可分割), 然后递归地向上合并子问题的解.    
        - `aux array based implementation`: 书上介绍的方法, 每次`merge`时需要对每个子问题申请一块辅助的子数组内存来暂存子问题, 空间复杂度为 **O(n)**. 时间复杂度为 **O(n*log(n))**, 缺点为需要申请额外的空间, 以及无法使用典型的`Swap`接口来实现排序过程.    
        - `in-place implementation`: 不需要辅助空间的实现, 借鉴了`insertion_sort`的方法来实现`merge`时的`in-place`. 空间复杂度为 **O(1)**. Benchmark 实测的运行时间比`aux array based implementation`慢很多, 接近`insertion_sort`(比它稍快).    
    - `heap_sort.go`: **O(n*log(n))** 实现《算法导论 第3版》 ch6.1~6.4 介绍的堆排序算法. 算法原理为借助Heap(一般是用二叉堆) 的性质, 即`root`元素总是最大的(`maxHeap`中; 若是`minHeap`则反之). 首先构建出`maxHeap`, 那么`root`元素一定是数组中的最大值. 将`root`元素与数组最后元素交换, 再针对新的`root`节点维护堆的性质则又可以获得剩余堆(除刚刚的最大元素)中的最大值, 如此循环操作直至遍历完整个数组.     
    - `quick_sort.go`: 平均运行时间接近最好运行时间 **O(n*log(n))**, 最坏情况为 **O(n^2)**. 实现《算法导论 第3版》ch7 介绍的快速排序算法, 包含固定主元(fixed pivot element)和随机主元(randomized pivot element)两种实现. 算法原理依然基于分治法, 即将问题分解为两个子问题分别求解. 而与归并排序不同的是, 在分解问题时总是选择一个`pivot element`, 将原数组分解为`<= pivot lement`和`>= pivot element`这样两个子问题, 并把`pivot element`放在两个子数组中间, 于是合并操作不需要再做任何事情. 快速排序的运行时间主要取决于分治(`partition`)的过程, 期望能够将问题分割为两个尽量平衡的子问题(其实只要两个子问题为常数比例即可, 如`9:1`, `99:1`均可达到 `O(n*log(n))` ). 当分割的两个子数组其中一个长度为0时, 即最坏的情况 **O(n^2)**.       
    - `counting_sort.go`: **O(k+n)** 其中`k`为输入数组的最大元素值(限定输入数组元素值范围为`0~k`). 实现《算法导论 第3版》ch8.2 介绍的计数排序. 其原理为引入一个以输入数组元素值为索引的辅助数组, 对于任意一个输入数组上的元素`x`, 直接在辅助数组上计算它前面有多少元素(也即输入数组中有多少元素小于`x`), 从而`x`就可以直接放在对应的位置上了. 其计算过程不需要任何的元素比较.      
        - 不是比较排序, 所以可以不受比较排序的最坏情况下界 **O(n*log(n))** 的限制.    
        - 需要注意的是`k`的值要可控, 即输入数组中的最小值和最大值间的范围不能太大, 否则空间开销太大, 此方法就不适用了.     
        - 当 k=O(n) 时, 其运行时间可简化为 **O(n)**.    
        - 另外, 当前的实现需要限定输入数组元素值范围为`0~k`, 实际可以比较容易地改造下以支持负数.    
    - `radix_sort.go`: **O(d*(k+n))** 其中`d`为输入数组元素的最大位数, `k`为每位上的最大值. 实现《算法导论 第3版》ch8.3 介绍的基数排序, 其原理为将输入数组看做一个`d`列`n`行的矩阵, 依次地从最低位至最高位对每列进行排序, 最终即可得到排序好的数组. 对于每列的排序需要依赖于稳定的排序算法(典型如`counting sort`就很适用).     
        - 基数排序非常适合用来解决多关键字记录的排序问题, 典型如对于日期的排序(将年、月、日分别看做位).     
        - 基数排序虽然看起来是线性排序, 运行低于快速排序等比较排序方法. 但实际使用时由于快速排序的每次循环更快、`in-place`等因素, 通常快速排序更好.     


### [Golang] 最大子数组问题
- maxsubarray    
解决最大子数组问题([Maximum Subarray Problem](https://en.wikipedia.org/wiki/Maximum_subarray_problem))的实验代码.    
    - `divide_and_conquer_algorithm.go`: **O(n*log(n))** 实现《算法导论 第3版》 ch4.1 最大子数组问题 中介绍的基于分治法([Divide and Conquer Alogrithm](https://en.wikipedia.org/wiki/Divide_and_conquer_algorithm))的实现. 其基本原理为将原数组分解为两个子数组, 分别求解两个子数组的最大子数组及同时跨两个子数组的最大子数组, 取最大值为最大子数组. 递归地执行这个过程直至求到最终解.      
    - `violent_solution.go`: **O(n^2)** 实现暴力方法求解, 即遍历所有子数组的可能, 从而取出最大子数组.     
    - `kadane_algorithm.go`: **O(n)** 实现《算法导论 第3版》 习题4.1-5 中所描述的线性算法. 基本原理为同时记录到目前为止的`max subarray`以及以当前为止为`ending`的`max subarray`, 遍历下一个元素时, 要得到最大子数组, 要么为`max subarray so far`, 要么为 max(`max subarray ending here` + `new element`, `new element`). 伪代码可参考[Maximum Subarray Problem - Kadane's algorithm](https://en.wikipedia.org/wiki/Maximum_subarray_problem).    

### [C++] 堆与优先队列
- cc_heaps    
《算法导论 第3版》 ch6.5 介绍的基于最大堆/最小堆的优先队列的设计.     
    - 最大优先队列[Max Priority Queue](https://en.wikipedia.org/wiki/Priority_queue)一般支持`Insert/Maximum/ExtractMax/IncreaseKey/Delete`等操作(最小优先队列对应的为`Insert/Minimum/ExtractMin/DecreaseKey/Delete`). 
    - 其基本原理仍然是基于最大堆/最小堆(与堆排序的原理非常类似), 即`Insert`时将新元素放在数组末尾然后进行`up`操作, 而`ExtractMax/Delete`时将对应元素与末尾元素交换后, 对当前元素进行`down`(i.e. `MaxHeapity`)操作, 从而维护堆的性质. 特殊处在于`IncreaseKey/DecreaseKey`操作, 需要记录针对每一个`HeapNode`的`handle`才能发起此操作. 实验中的`handle`直接用的是`HeapNode`的指针来表示.([std::priority_queue](http://www.cplusplus.com/reference/queue/priority_queue/)干脆就没有暴露这个接口, 所以看起来接口简单很多.)      
    - NOTE:     
        - 此处的实现代码, 基于`BinaryHeap`进行实现的, 且仅实现了`MaxPriorityQueue`. 而实际上同样可以基于`d-ary heap`来实现, 区别只是`2-dry`或`d-ary`.     
        - `Golang`中目前不支持泛型, 也没有`class/object`等概念, 不太好实现一个封装起来的`PriorityQueue`. `Golang`源码库中的实现是给出了一些`heap`的操作函数`Fix/Init/Push/Pop`等, 而把内部的数据结构实现交给了使用者. 实现和使用思路都与一般的面向对象思路不同, 个人感觉是不太直观的, 所以这里用`c++`来实现了此处的实验代码, 看起来或许更直观些.     
            - 另外, 下面的包`minspanningtree`中实现`Prim`算法时用到了`MinHeap`, 除借助`golang`的`"container/heap"`中提供的方法外, 也自己实现了一份在`my_min_binary_heap.go`中, 可供参考.    
        - 当前代码中暂未考虑效率问题(如`object`的多处拷贝等), 实际使用的话应考虑优化, 如多采用`std::move`等. 另可参考`boost::heap::d_ary_heap`, [Fibonacci heap](https://en.wikipedia.org/wiki/Fibonacci_heap)等.     
    - `g++ -std=c++11 main.cc && ./a.out`    


### [Golang] 图算法
![golang_pkg_import_graph](golang_pkg_import_graph.mmd.png)

- graph     
通常使用符号 **G(V, E)** 来表示一张图, 其中 **V** 为点数, **E** 为边数. 此`pkg`定义了一堆表示`graph`的类型与通用接口, 支持包括邻接链表(Adjacency List)和邻接矩阵(Adjacency Matrix)两种图的表示方法. 其中点通过从`0`开始的`uint`来表示, 所以邻接链表和邻接矩阵都基于了基础的`slice`来实现. 基础概念可参考《算法导论 第3版》 ch22.1 图的表示.     
注: 图论中的各种涉及到路径的算法, 通常都基于`point-to-point`来讨论, 而不是`edge-to-edge`, 从`graph`的表现形式就可以反映出这一点.     

- graphsample1     
来自《算法导论 第3版》ch22.2 广度优先搜索 中的示例`Undirected Graph`, 基于上面的`package graph`的定义的实现, 从而方便后续的实验.    

- graphsample2    
来自《算法导论 第3版》ch22.1 中的示例`Directed Graph`, 基于`package graph`的定义的实现, 从而方便后面的实验.    

- graphsample3    
来自《算法导论 第3版》ch22.3 中的示例`Directed Acyclic Graph`, 基于`package graph`的定义的实现, 从而方便后面的实验.    

- graphsample4    
来自《算法导论 第3版》ch22.5 Strongly Connected Component 中的示例`Directed Graph`, 基于`package graph`的定义的实现, 从而方便后面的实验.    

- graphsample5    
来自《算法导论 第3版》ch23.2 Kruskal和Prim算法 中的示例`Undirected Graph`, 基于`package graph`和`package weightedgraph`的定义的实现, 从而方便后面的实验.    


- bfs    
    - **O(V+E)**    
    - 实现《算法导论 第3版》ch22.2 广度优先搜索 中的算法描述, [Breadth First Search](https://en.wikipedia.org/wiki/Breadth-first_search). 基本思路为搜索过程中从`queue`(借助其先入先出的特性)头上取下一次迭代的初始节点, 并将迭代到的节点存储到`queue`尾, 从而实现**广度优先**. 搜索过程中的`tree`的信息及`depth`等通过节点属性的形式保存在一个节点数组中.    
    - 提供了基于`bfs`的生成`level graph`的实现, 以供`dinic`算法使用.    

- bfs_main    
执行`package bfs`代码的`main`.     

- dfs    
    - **O(V+E)**    
    - 实现《算法导论 第3版》ch22.3 深度优先搜索 中的算法描述. 本书章节中的伪码是基于递归的描述, 比较清晰易懂也容易实现. 实际`coding`时同时参考[Depth-first search - Wikipedia](https://en.wikipedia.org/wiki/Depth-first_search)实现了基于栈的实现. 注: 基于栈的实现一般来讲结果会和递归的实现不太一样, 主要是遍历的次序关系.    
    - 《算法导论 第3版》ch22.4 拓扑排序([Topological Sorting](https://en.wikipedia.org/wiki/Topological_sorting)) 也是依赖 [Depth-first search - Wikipedia](https://en.wikipedia.org/wiki/Depth-first_search) 的方法来实现的, 故其实现也放在了`dfs`包中. 由于`graph`中若存在环是无法拓扑排序的, 所以拓扑排序只能基于[Directed Acyclic Graph](https://en.wikipedia.org/wiki/Directed_acyclic_graph), 依据此也可以通过`dfs`判断一个`graph`是否为[Directed Acyclic Graph](https://en.wikipedia.org/wiki/Directed_acyclic_graph).    
        - 拓扑排序的一个典型应用: 比如要做一个[Continuous Delivery](https://en.wikipedia.org/wiki/Continuous_delivery)比较复杂的Pipeline(Pipeline上许多步骤间有依赖关系), 那么便可以先把步骤的依赖关系图先画出来, 算出拓扑排序, 再按照拓扑排序的顺序来梳理Pipeline的流程.    
    - 《算法导论 第3版》ch22.5 强连通分量([Strongly Connected Component](https://en.wikipedia.org/wiki/Strongly_connected_component)) 同样依赖2次 [Depth-first search - Wikipedia](https://en.wikipedia.org/wiki/Depth-first_search) 实现, 于是实验代码也放在`dfs`包中. 第一次基于默认`graph`进行`DFS`, 然后以`timestampF`倒序作为第二次遍历的`node`顺序; 第二次基于`graph`的转置(反转所有`edge`)来进行`dfs`, 然后以`timestampF`正序来进行遍历并输出[Strongly Connected Component](https://en.wikipedia.org/wiki/Strongly_connected_component), 以每个`dfs tree`的`root`作为切分.    
        - 注: 第一次`dfs`后的`timestampF`倒序, 看起来为上述拓扑排序的结果的倒序即可. 但实际上主要区别在于拓扑排序的`graph`必须无环, 而切分强连通分量几乎必然有环(否则就只好每个`node`为一个强连通分量). 所以并不能直接调用拓扑排序的函数实现.    
  
- levelgraph    
`Dinic`算法所要用到的分层图, 即以[Depth-first search - Wikipedia](https://en.wikipedia.org/wiki/Depth-first_search)在`graph`上进行搜索, 以每个`node`的`depth`作为`level`.    

- flownetwork    
描述[maximum flow problem](https://en.wikipedia.org/wiki/Maximum_flow_problem)的流网络, 主要包含基于`directed graph`的图以及图上每两个`node`间(i.e. `edge`)的容量. 需要注意的是两个`node`间只能有单向的`edge`, 不能有反向. 同时为了描述[maximum flow problem](https://en.wikipedia.org/wiki/Maximum_flow_problem)问题, 也记录了入点和出点.    
NOTE: `flownetwork` 与 `weightedgraph` 非常相似, 所以其实完全可以基于`weightedgraph`来实现. 此处仅仅是因为`flownetwork`有点类似于描述[maximum flow problem](https://en.wikipedia.org/wiki/Maximum_flow_problem)的一个专用数据结构, 所以减少了依赖. 有必要的时候可以进行 refactor 以去除重复代码.    

- maxflow    
[maximum flow problem](https://en.wikipedia.org/wiki/Maximum_flow_problem)的算法实现, 包括`FordFulkerson`, `EmondsKarp`, `Dinic`, etc. 在其内部的`_test.go`中以[Drainage Ditches](http://poj.org/problem?id=1273)问题作为典型的测试用例.    
    - `FordFulkerson`: **O(E|f|)**  基础的最大流问题解决方法. 定义了`flow`, `residual network`, `augmenting path`等重要的基础概念, 以及解决问题的一般思路.    
    - `EmondKarp`: **O(V(E^2))** 基于`FordFulkerson`, 在如何寻找`augmenting path`的方法上进行了扩展优化, 即以[Breadth First Search](https://en.wikipedia.org/wiki/Breadth-first_search)来寻找点到点的最短路径, 效率更高.    
    - `Dinic`: **O((V^2)E)** 依然是基于`FordFulkerson`的方法, 最主要的区别在于在生成`residual network`后, 先采用[Breadth First Search](https://en.wikipedia.org/wiki/Breadth-first_search)来生成分层图(`level graph`, 以每个`node`的`depth`作为其层次), 再在`level graph`上寻找`blocking flow`(即直到不能再找到新的`flow`), 以此`blocking flow`作为`residual network`上的`augmenting flow`.     

- maxflow_main    
调用`maxflow`以解决[maximum flow problem](https://en.wikipedia.org/wiki/Maximum_flow_problem), 支持从`stdin`来构造`flownetwork`, 以更容易测试新的问题.     

- weightedgraph    
即在`package graph`定义的`directed/undirected graph`的基础操作上, 为每个`edge`增加一个`weight`值.     

- minspanningtree    
《算法导论 第3版》ch23 中所描述的最小生成树问题, 即在一张每条`edge`都有其`weight`的连通图上(一般是基于`undirected graph`讨论), 找到经过的所有`node`的`sum(weight)`最小的生成树. 
    - 此处实验中实现了书上提到的两种算法: 
        - Kruskal: **O(E*log(V))** 按`weight`排序所有的`edge`, 从最小`weight`的`edge`开始一条一条取出来生成树. 每条取出的`edge`的两个`node`不属于同一个集合, 才是有效的`edge`.      
        - Prim: **O(E*log(V))** 借助了最小堆的概念, 从任意`node`开始构造最小堆, 每次从`ExtractMin`的`node`遍历邻接的`node`, 并将经过的`edge`的`weight`作为遍历到的`node`的`key`以更新维护最小堆的性质, 于是每次`ExtractMin`的`node`与它的`parent`间的`edge`一定为最小生成树的一部分.     
            - 注: 据说借助[Fibonacci heap](https://en.wikipedia.org/wiki/Fibonacci_heap)来实现可以进一步优化运行效率.    
    - 注: 关于最小生成树的问题, 个人认为几个视频讲的比书上要好, FYI:    
        - [Youtube Minimum Spanning Tree](https://www.youtube.com/watch?v=5INWifzqStU)
        - [Youtube Kruskal Algorithm](https://www.youtube.com/watch?v=5xosHRdxqHA)
        - [Youtube Prim Algorithm](https://www.youtube.com/watch?v=z1L3rMzG1_A)


    
## References
- 《算法导论 第3版》    
