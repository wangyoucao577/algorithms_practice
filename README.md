# algorithms_practice
Learn and practice algorithms and data structures. 来源包括但不限于《算法导论》, [hankerrank](https://www.hackerrank.com), [leecode](https://leetcode.com/), [北京大学 PKU JudgeOnline](http://poj.org/), etc.     


## 实验环境
编译与运行主要在`Linux`上进行, 但代码都应属于平台无关的, 理论上不限于操作系统/编译器等.     
- `Ubuntu 18.04 LTS`: `WSL(Windows Subsystem for Linux) on Win10 1803-17134.228`
- `cmake version 3.10.2`
- `gcc (Ubuntu 7.3.0-16ubuntu3) 7.3.0`
- `go version go1.10.1 linux/amd64`

## 实验代码
虽说都是实验代码, 但大部分提供了封装好的`golang pkg`, 可直接导入作为基础库使用. 内部进行实验时, 也尽可能地拆分成`pkg`以解耦.    

### [Golang] 排序
学习各种排序的实验代码，主要参考《算法导论 第3版》，
- 包括:     
    - 比较排序算法: 插入排序(Insertion Sort)、归并排序(Merge Sort)、堆排序(Heap Sort)、快速排序(Quick Sort)    
    - 非比较排序算法: 计数排序(Counting Sort)、基数排序(Radix Sort)、桶排序(Bucket Sort)     
    - 其他算法: 选择第`n-th`元素(Select n-th Element)、基于Binary Search Tree的排序(Tree Sort)      
- 详见:    
[pkg - mysorts](./mysorts/)    


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

### [Golang] 二叉搜索树
- binarysearchtree    
实现《算法导论 第3版》ch12 介绍的二叉搜索树的实验代码.     
    - 二叉搜索树性质: 设`x`为二叉搜索树中的一个节点, 若`y`是`x`的左子树中的一个节点, 则`y.key <= x.key`; 若`y`是`x`的右子树中的一个节点, 则`y.key >= x.key`.    
    - 二叉搜索树实现时, 一般每个`node`中都会记录`parent/leftChild/rightChild`三个指针以维护树的结构, 同时`node`中需记录`key`以维护二叉搜索树的性质. 可选的`node`中可能会记录额外的`payload`.    
    - 二叉搜索树一般至少会提供接口: (实现在`binary_search_tree.go`中)    
        - `Minimum()/Maximum()`: `O(h)` 返回最小/最大`key`的节点(同`MinHeap/MaxHeap`中的`Minimum()/Maximum()`)    
            - 故[Binary Search Tree](https://en.wikipedia.org/wiki/Binary_search_tree)也可以作为一个可以同时方便地查询`Max/Min`的[Priority Queue](https://en.wikipedia.org/wiki/Priority_queue)来使用.     
        - `Successor()/Predecessor()`: `O(h)` 返回当前`node`的`Succesor/Predecessor`节点    
            - `Successor`节点: `>= node.key` 的最小`key`节点    
            - `Predecessor`节点: `<= node.key` 的最大`key`节点    
        - `Insert()`: `O(h)` 插入一个新的节点(总是会插入为`tree`的`leaf`)    
        - `Delete()`: `O(h)` 删除一个新的节点. 实现上最复杂的一个接口, 主要是要删除的节点同时存在`leftChild`和`rightChild`时的情况比较复杂.    
        - `Search()`: `O(h)` 搜索一个指定`key`的`node`.    
        - `InorderTreeWalk()`: `O(n)` 中序遍历, 即总是按照 `node.leftChild => node => node.rightChild` 的顺序递归遍历.     
            - 注: 由于二叉搜索树的性质, 中序遍历的结果总是按照`key`升序排序的. 也即可以通过二叉搜索树来实现排序. 我的实验代码见 [mysorts - tree sort](./mysorts/tree_sort.go).    
        - `PreorderTreeWalk()`: `O(n)` 先序遍历, 即总是按照 `node => node.leftChild => node.rightChild` 的顺序递归遍历.     
        - `postorderTreeWalk()`: `O(n)` 后续遍历, 即总是按照 `node.leftChild => node.rightChild => node` 的顺序递归遍历.    
    - 注: 以上分析的`O(n)`中的`n`为树中的总的节点数, `O(h)`中的`h`为树的高度, 最坏情况下`h == n`, 但平均情况接近最好情况即`lg(n)`. 也即平均情况下[Binary Search Tree](https://en.wikipedia.org/wiki/Binary_search_tree)的接口的运行时间为`O(lg(n))`.       

### [Golang] 红黑树
- rbtree    
实现《算法导论 第3版》ch13 介绍的红黑树的实验代码.     
    - 红黑树的性质(满足红黑性质的二叉搜索树即为红黑树):    
        - 每个`node`都有一个`color`的属性, 要么是红色, 要么是黑色;    
        - `root node`总是黑色的;    
        - 每个`leat node`总是黑色的(为了实现的简便, 一般总是用一个`sentinal nil node`作为实现时的`nil`, 也即每个`leaf node`都是这个`sentinal nil node`, 只要它置为黑色即可满足此性质);    
        - 如果一个`node`是红色的, 那么它的两个子`node`总是黑色的;    
        - 对于每个`node`, 从该`node`到其所有后代`leaf node`的简单路径上, 均包含相同数目的黑色`node`.    
    - 黑高(Black-Height)的定义: 从某个`node`出发(不包含该`node`)到达一个`leaf node`的任意一条简单路径上的黑色`node`数目, 即为黑高(Black-Height), 一般记作`bh(node)`.    
        - 定义`root`节点的黑高为红黑树的黑高.    
    - 红黑树的`Insert/Delete`除了类似二叉搜索树的操作外, 都需要去维护红黑性质; 而维护了红黑性质后, 红黑树非常接近平衡树, 从而可以带来较好的平均性能.    
    - 红黑树提供的接口与二叉搜索树完全一致, 复杂度分析也同样一致. 只是由于红黑树几乎为平衡二叉树, 树的高度总是为`log(n)`的, 所以接口的运行时间总是可以做到`O(log(n))`.    
    - 非常好的学习资料(红黑树的`Insert`和`Delete`过程比较复杂, 以下资料个人认为讲的远好于书上. 而可能本质上就是需要枚举出所有的可能情况, 相对来讲书上枚举的不够详细. 跟着教程自己推导一遍, 收货颇丰):     
        - [Youtube 红黑树的插入](https://www.youtube.com/watch?v=axa2g5oOzCE)    
        - [红黑树的删除](https://segmentfault.com/a/1190000012115424)    


### [Golang] 图算法
![golang_pkg_import_graph](golang_pkg_import_graph.mmd.png)

- graph     
通常使用符号 **G(V, E)** 来表示一张图, 其中 **V** 为点数, **E** 为边数. 此`pkg`定义了一堆表示`graph`的类型与通用接口, 支持包括邻接链表(Adjacency List)和邻接矩阵(Adjacency Matrix)两种图的表示方法. 其中点通过从`0`开始的`uint`来表示, 所以邻接链表和邻接矩阵都基于了基础的`slice`来实现. 基础概念可参考《算法导论 第3版》 ch22.1 图的表示.     
注: 图论中的各种涉及到路径的算法, 通常都基于`point-to-point`来讨论, 而不是`edge-to-edge`, 从`graph`的表现形式就可以反映出这一点.     

- graphsamples    
构建好的`graph`的例子, 方便进行各种测试. 详见 [graphsamples](./graphsamples/).     


- bfs    
    - **O(V+E)**    
    - 实现《算法导论 第3版》ch22.2 广度优先搜索 中的算法描述, [Breadth First Search](https://en.wikipedia.org/wiki/Breadth-first_search). 基本思路为搜索过程中从`queue`(借助其先入先出的特性)头上取下一次迭代的初始节点, 并将迭代到的节点存储到`queue`尾, 从而实现**广度优先**. 搜索过程中的`tree`的信息及`depth`等通过节点属性的形式保存在一个节点数组中.    
    - 提供了基于`bfs`的生成`level graph`的实现, 以供`dinic`算法使用.    

- cmd/test_bfs   
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

- cmd/test_maxflow   
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

- shortestpaths    
《算法导论 第3版》ch24 中所讨论的 [Single-Source Shorest Paths](https://en.wikipedia.org/wiki/Shortest_path_problem#Single-source_shortest_paths) 问题, 即在一张每条`edge`都有其`weight`的连通图上(一般是基于`directed graph`, 但其实只要`weight`非负, 也完全可以在`undirected graph`上讨论), 找到从某个`source-node`开始到其他任意`node`的最短`weight`路径.    
    - `bellman-ford.go`: **O(VE)** 迭代`V-1`次, 每次所有的`edge`以更新每个`node`上的`cost`(从`source-node`到当前`node`). 最后再迭代一次所有的`edge`以判断是非存在`negative-cycle`. `bellman-ford`算法实现比较简单, 但迭代次数较多, 主要优势是可以支持`negative weight`(不能有`negative-cycle`).    
    - `dijkstra.go`: **O((V+E)log(V))** 基于`PriorityQueue`每次总是以最小`weight`的`node`开始`relax`, 从而可以大大减少迭代的次数. 此处的复杂度是基于二叉堆实现的`PriorityQueue`来分析的. `Dijkstra`算法必须要求`weight`为非负的.     
        - 注意:     
            - 参照书上伪代码实现的此`Dijkstra`, 流程是首先将所有的`node-priority pair`加入到`PriroityQueue`中, 过程中有`priority`发生变化时再`DecreaseKey`, 所以会有非常大量的`DecreaseKey`动作.     
            - 但实际工程上实现时多半都是先只将`source-node`加入到`PriorityQueue`中, 过程中再将新`relax`到的`node`加入`PriorityQueue`, 所以会有大量的`Push`动作, `DecreaseKey`则大大减少, 同时`PriorityQueue`的规模也是动态增加的而不是一开始加满.     
            - 这也是为什么学术上经常会讨论如何优化`DecreaseKey`效率从而提升`Dijkstra`的性能, 但却并不适用于工程实现.    
    - `directed_acyclic_graph_shortest_paths.go`: **O(V+E)** 先将所有的`node`使用`dfs`拓扑排序, 再按照拓扑排序的顺序从`source-node`开始遍历一次`edge`即可. 适用于有向无环图的特殊情况, 复杂度主要是由于拓扑排序.    
    
## References
- 《算法导论 第3版》    
- [Package sort](https://golang.org/pkg/sort/)
- [Package heap](https://golang.org/pkg/container/heap/)
- [Maximum Subarray Problem](https://en.wikipedia.org/wiki/Maximum_subarray_problem)
- [Divide and Conquer Alogrithm](https://en.wikipedia.org/wiki/Divide_and_conquer_algorithm)
- [Maximum Subarray Problem - Kadane's algorithm](https://en.wikipedia.org/wiki/Maximum_subarray_problem)
- [Priority Queue](https://en.wikipedia.org/wiki/Priority_queue)
- [std::priority_queue](http://www.cplusplus.com/reference/queue/priority_queue/)
- [Fibonacci heap](https://en.wikipedia.org/wiki/Fibonacci_heap)
- [Binary Search Tree](https://en.wikipedia.org/wiki/Binary_search_tree)
- [Tree Traversal](https://en.wikipedia.org/wiki/Tree_traversal)
- [Breadth First Search](https://en.wikipedia.org/wiki/Breadth-first_search)
- [Depth-first search - Wikipedia](https://en.wikipedia.org/wiki/Depth-first_search)
- [Topological Sorting](https://en.wikipedia.org/wiki/Topological_sorting)
- [Directed Acyclic Graph](https://en.wikipedia.org/wiki/Directed_acyclic_graph)
- [Continuous Delivery](https://en.wikipedia.org/wiki/Continuous_delivery)
- [Strongly Connected Component](https://en.wikipedia.org/wiki/Strongly_connected_component)
- [maximum flow problem](https://en.wikipedia.org/wiki/Maximum_flow_problem)
- [Drainage Ditches](http://poj.org/problem?id=1273)
- [Youtube Minimum Spanning Tree](https://www.youtube.com/watch?v=5INWifzqStU)
- [Youtube Kruskal Algorithm](https://www.youtube.com/watch?v=5xosHRdxqHA)
- [Youtube Prim Algorithm](https://www.youtube.com/watch?v=z1L3rMzG1_A)
- [Red Black Tree](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree)
- [Youtube 红黑树的插入](https://www.youtube.com/watch?v=axa2g5oOzCE)    
- [红黑树的删除](https://segmentfault.com/a/1190000012115424)    
- [Single-Source Shorest Paths](https://en.wikipedia.org/wiki/Shortest_path_problem#Single-source_shortest_paths)   
- [Dijkstra Algorithm](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm)



