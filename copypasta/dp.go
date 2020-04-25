package copypasta

import (
	"sort"
)

/*
参考书籍推荐：
《算法竞赛进阶指南》- 介绍了大量且全面的 DP 内容，是目前市面上讲解 DP 最好的一本书

视频讲解：
https://www.bilibili.com/video/av70148899 DP 入门，01 背包，完全背包，多重背包
https://www.bilibili.com/video/av77393700 LCS LIS
https://www.bilibili.com/video/av83939419 区间 DP
https://www.bilibili.com/video/av93356551 状态压缩 DP
https://www.bilibili.com/video/av98090640 树形 DP
https://www.bilibili.com/video/av85636122 动态规划 · 零 - Introduction
https://www.bilibili.com/video/av86983419 动态规划 · 一 - 序列型
https://www.bilibili.com/video/av89052674 动态规划 · 二 - 坐标、双序列、划分 & 状态压缩

套题/总结：
重要 DP 技巧总结 https://codeforces.ml/blog/entry/47764
状压 DP https://codeforces.ml/blog/entry/45223
CSES DP section editorial https://codeforces.ml/blog/entry/70018
CF 全部 DP 题  https://codeforces.ml/problemset?order=BY_RATING_ASC&tags=dp
力扣上的 DP 问题
    分类汇总 https://zhuanlan.zhihu.com/p/126546914
    https://leetcode.com/discuss/general-discussion/458695/dynamic-programming-patterns
    https://github.com/CyC2018/CS-Notes/blob/master/notes/Leetcode%20%E9%A2%98%E8%A7%A3%20-%20%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92.md
    https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/discuss/108870/Most-consistent-ways-of-dealing-with-the-series-of-stock-problems
    https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/solution/yi-ge-tong-yong-fang-fa-tuan-mie-6-dao-gu-piao-w-5/
    https://leetcode.com/problemset/all/?topicSlugs=dynamic-programming
AT 经典 DP 场 https://atcoder.jp/contests/dp
    题解 https://www.cnblogs.com/shanxieng/p/10232228.html
    题解（日语）https://www.hamayanhamayan.com/entry/2019/01/12/163853
信息学奥赛一本通 第二部分 基础算法 --> 第九章 动态规划 http://ybt.ssoier.cn:8088/index.php

其他资料：
https://github.com/hzwer/shareOI/tree/master/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92
https://oi-wiki.org/dp/
https://cp-algorithms.com/dynamic_programming/divide-and-conquer-dp.html

EXTRA: 如何定义子问题，从而简化思考的复杂性、降低计算量？
https://codeforces.ml/problemset/problem/553/A
https://codeforces.ml/blog/entry/47764 也有提到

NOTE: 若使用滚动数组，复用时可能要初始化
NOTE: 实际情况是使用滚动数组仅降低了内存开销，整体运行效率与不使用滚动数组时无异
NOTE:（区间 DP）正向计算不易时，试着反向计算
TIPS: 如果反复在同一个序列上点权汇合，可以考虑用前缀和或者差分来优化
      - 若转移是若干相邻项之和，可以考虑 f(p) - f(p-1) 的值，用滑动窗口来维护区间和，从而优化转移
      https://leetcode-cn.com/problems/new-21-game/
*/
func dpCollections() {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// EXTRA: 由于数据范围的原因，采用 map 记忆化
	mapDP := func() {
		type pair struct{ x, y int }
		dp := map[pair]int{}
		var f func(int, int) int
		f = func(x, y int) (res int) {
			p := pair{x, y}
			if v, ok := dp[p]; ok {
				return v
			}
			defer func() { dp[p] = res }()

			return
		}
		_ = f
	}

	/* 线性 DP：前缀/后缀之间的转移
	数字三角形 https://www.luogu.com.cn/problem/P1216
	todo 最长公共上升子序列 (LCIS) https://codeforces.ml/problemset/problem/10/D
	todo 两个排列的 LCS https://www.luogu.com.cn/problem/P1439
	*/

	// 最长公共子序列 (LCS)
	// 有向无环图：s1[i] == s2[j] (i-1,j-1) -> (i,j) $ 1
	//            s1[i] != s2[j] (i-1,j) -> (i,j) $ 0
	//                           (i,j-1) -> (i,j) $ 0
	// 例题 https://leetcode-cn.com/problems/longest-common-subsequence/
	// EXTRA: 最短公共超序列 (SCS) https://leetcode-cn.com/problems/shortest-common-supersequence/
	lcs := func(s1, s2 string) int {
		n, m := len(s1), len(s2)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, m+1)
		}
		for i, b1 := range s1 {
			for j, b2 := range s2 {
				if b1 == b2 {
					dp[i+1][j+1] = dp[i][j] + 1
				} else {
					dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
				}
			}
		}
		return dp[n][m]
	}
	lcsPath := func(s1, s2 string) []byte {
		n, m := len(s1), len(s2)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, m+1)
		}
		fa := make([][]int8, n+1)
		for i := range fa {
			fa[i] = make([]int8, m+1)
		}
		for i, b1 := range s1 {
			for j, b2 := range s2 {
				if b1 == b2 {
					dp[i+1][j+1] = dp[i][j] + 1
					fa[i+1][j+1] = 1
				} else {
					if dp[i][j+1] > dp[i+1][j] {
						dp[i+1][j+1] = dp[i][j+1]
						fa[i+1][j+1] = 2
					} else {
						dp[i+1][j+1] = dp[i+1][j]
						fa[i+1][j+1] = 3
					}
				}
			}
		}
		lcs := make([]byte, 0, dp[n][m])
		var makeLCS func(i, j int)
		makeLCS = func(i, j int) {
			if i == 0 || j == 0 {
				return
			}
			if fa[i][j] == 1 {
				makeLCS(i-1, j-1)
				lcs = append(lcs, s1[i-1])
			} else if fa[i][j] == 2 {
				makeLCS(i-1, j)
			} else {
				makeLCS(i, j-1)
			}
		}
		makeLCS(n, m)
		return lcs
	}

	// 最长上升子序列 (LIS)
	// O(n^2) - 定义 dp[i] 为以 a[i] 为末尾的 LIS 的长度
	//          可以把此问题想象成一个「跳跃游戏」，任选一个初始位置向右跳跃，每次只能跳到比当前位置更高的位置，问最多能跳多少次（最后答案加一）
	//          这样能更容易地看出转移的顺序，然后变成一个 DAG 上求最长路的问题
	// O(nlogn) - 定义 dp[i] 为长度为 i+1 的 LIS 末尾元素的最小值
	// https://oi-wiki.org/dp/basic/#_12
	// 例题 https://leetcode-cn.com/problems/longest-increasing-subsequence/
	// 方案数 https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/
	//       https://www.zhihu.com/question/34905638
	lis := func(arr []int) int {
		dp := make([]int, 0, len(arr))
		for _, v := range arr {
			if i := sort.SearchInts(dp, v); i < len(dp) {
				dp[i] = v
			} else {
				dp = append(dp, v)
			}
		}
		return len(dp)
	}

	// 本质不同子序列个数
	// 定义 dp[i][j] 表示前 i 个字符中长度为 j 的本质不同子序列个数
	// 转移 dp[i][j] = dp[i-1][j]（不选第 i 个字符）+ dp[i-1][j-1] - dp[prev[i]-1][j-1]（选第 i 个字符）
	// 其中 prev[i] 为 s[i] 的上一个相同字符位置
	// https://ac.nowcoder.com/discuss/394080 C 题
	// https://codeforces.ml/problemset/problem/1183/H
	distinctSubsequence := func(s string) int64 {
		n := len(s)
		prev := [26]int{}
		dp := make([][]int64, n+1)
		for i := range dp {
			dp[i] = make([]int64, n+1)
		}
		dp[0][0] = 1
		for i := 1; i <= n; i++ {
			c := s[i-1] - 'a'
			dp[i][0] = 1
			for j := 1; j <= i; j++ {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
				if p := prev[c]; p > 0 {
					dp[i][j] -= dp[p-1][j-1]
				}
			}
			prev[c] = i
		}
		sum := int64(0)
		for _, cnt := range dp[n][1:] { // 不计入空字符串
			sum += cnt
		}
		return sum
	}

	/* 背包问题
	https://en.wikipedia.org/wiki/Knapsack_problem
	https://codeforces.ml/blog/entry/59606
	套题 https://www.acwing.com/problem/
	*/

	// 0-1 背包 (n 个物品，背包容量为 maxW)
	// 基本状态：前 i 个物品  i∈[0,n]
	// 附加状态：容量(上限)为 j  j∈[0,maxW]
	// 点权：最大价值
	//     初始值：(0,j)=0  j∈[0,maxW]
	// 有向无环图：不选第 i 个物品，对各个容量 j，连一条横边，即 (i-1,j) -> (i,j) $ 0
	//             选第 i 个物品，对各个容量 j (j≥wi)，连边 (i-1,j-wi) -> (i,j) $ vi
	//     起点：(0,j)  j∈[0,maxW]
	//     终点：(n,maxW)
	// 核心函数：最大价值（最长路），即 max
	// https://oi-wiki.org/dp/knapsack/
	// 模板题 https://atcoder.jp/contests/dp/tasks/dp_d
	// EXTRA: 恰好装满 https://leetcode-cn.com/problems/partition-equal-subset-sum/
	zeroOneKnapsack := func(values, weights []int, maxW int) int {
		n := len(values)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, maxW+1)
		}
		for i, vi := range values {
			wi := weights[i]
			for j, dpij := range dp[i] {
				if j < wi {
					dp[i+1][j] = dpij // 入度为 1，直接转移
				} else {
					dp[i+1][j] = max(dpij, dp[i][j-wi]+vi) // 入度为 2，取核心函数转移
				}
			}
		}
		return dp[n][maxW]
	}

	// 从 a 中选出若干个数，总和为 sum 的方案数
	// 基本状态：前 i 个数  i∈[0,n]
	// 附加状态：和为 j  j∈[0,sum]
	// 点权：方案数
	//     初始值：(0,0)=1
	// 有向无环图：不选第 i 个数，对各个和 j，连一条横边，即 (i-1,j) -> (i,j)
	//             选第 i 个数，对各个和 j (j≥ai)，连边 (i-1,j-ai) -> (i,j)
	//     起点：(0,j)  j∈[0,sum]
	//     终点：(n,sum)
	// 核心函数：方案数（点权汇合），即 +
	// 例题（需要转换）https://leetcode-cn.com/problems/target-sum/
	waysToSum := func(a []int, sum int) int {
		n := len(a)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, sum+1)
		}
		dp[0][0] = 1
		for i, v := range a {
			for j, dpij := range dp[i] {
				if j < v {
					dp[i+1][j] = dpij // 入度为 1，直接转移
				} else {
					dp[i+1][j] = dpij + dp[i][j-v] // 入度为 2，取核心函数转移
				}
			}
		}
		return dp[n][sum]
	}

	// 完全背包
	unboundedKnapsack := func(values, weights []int, maxW int) int {
		n := len(values)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, maxW+1)
		}
		for i, vi := range values {
			wi := weights[i]
			for j, dpij := range dp[i] {
				if j < wi {
					dp[i+1][j] = dpij // 入度为 1，直接转移
				} else {
					dp[i+1][j] = max(dpij, dp[i+1][j-wi]+vi) // 入度为 2，取核心函数转移
				}
			}
		}
		return dp[n][maxW]
	}

	// 恰好装满背包至少需要多少个物品，物品无限。无法装满时返回 -1
	// 基本状态：容量 i  i∈[0,amount]
	// 点权：最少物品数
	//     初始值：0=0
	// 有向无环图：i-wj (wj≤i) -> i $ 1
	//     起点：0
	//     终点：amount
	// 核心函数：最少物品数（最短路），即 min
	// https://leetcode-cn.com/problems/coin-change/
	minCoinChange := func(coins []int, amount int) int {
		const inf int = 1e9
		dp := make([]int, amount+1)
		for i := range dp {
			dp[i] = inf
		}
		dp[0] = 0
		// 按容量遍历以满足拓扑序
		for cur := range dp {
			for _, c := range coins {
				if c <= cur {
					dp[cur] = min(dp[cur], dp[cur-c]+1)
				}
			}
		}
		if dp[amount] < inf {
			return dp[amount]
		}
		return -1
	}

	// EXTRA: 完全背包 - 求方案数
	// https://leetcode-cn.com/problems/coin-change-2/

	// todo 二维费用背包 https://leetcode-cn.com/problems/ones-and-zeroes/

	// 多重背包
	// todo 方法 1：二进制优化

	/* 区间 DP / 环形 DP
	最优三角剖分 https://leetcode-cn.com/problems/minimum-score-triangulation-of-polygon/
	戳气球 https://leetcode-cn.com/problems/burst-balloons/
	打印机 https://leetcode-cn.com/problems/strange-printer/
	todo 石子合并：相邻 k 堆 https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/
	todo 石子合并：环形，相邻 2 堆 https://www.luogu.com.cn/problem/P1880
	todo https://atcoder.jp/contests/abc159/tasks/abc159_f

	博弈类 DP
	    转移：让「自己与对手的分差」最大
	https://leetcode.com/problems/stone-game/ https://nanti.jisuanke.com/t/48
	https://leetcode.com/problems/stone-game-ii/
	https://leetcode.com/problems/stone-game-iii/
	CF tag https://codeforces.ml/problemset?order=BY_RATING_ASC&tags=dp%2Cgames
	*/

	/* 状压 DP
	NOTE: 若问题无法划分成小问题，必须考虑各种可能的情况，则可能是 NP 完全问题
	CF tag https://codeforces.ml/problemset?order=BY_RATING_ASC&tags=dp%2Cbitmasks
	todo 汉密尔顿路径/回路 Hamiltonian path
	*/

	// 旅行商问题 (TSP)
	// https://en.wikipedia.org/wiki/Travelling_salesman_problem
	// 模板题 https://www.luogu.com.cn/problem/P1171 https://www.luogu.com.cn/problem/P1433
	// 建模转换题 https://leetcode-cn.com/problems/find-the-shortest-superstring/
	// EXTRA: 固定起点终点的问题，视问题情况有两种方法：
	//        添加一个节点 https://stackoverflow.com/questions/14527815/how-to-fix-the-start-and-end-points-in-travelling-salesmen-problem
	//        设置距离 https://stackoverflow.com/questions/36086406/traveling-salesman-tsp-with-set-start-and-end-point
	tsp := func(dist [][]int) int {
		n := len(dist)
		const inf int = 1e9
		dp := make([][]int, 1<<n)
		for i := range dp {
			dp[i] = make([]int, n)
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}

		// 记忆化：已经访问的集合 s，当前位置 v
		var f func(s, v int) int
		f = func(s, v int) (res int) {
			dv := &dp[s][v]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
			if s == 1<<n-1 && v == 0 {
				return
			} // 访问了所有节点并回到了 0
			res = inf
			for w := 0; w < n; w++ {
				if s>>w&1 == 0 {
					res = min(res, f(s|1<<w, w)+dist[v][w])
				}
			}
			return
		}
		f(0, 0)

		// DP
		dp = make([][]int, 1<<n)
		for i := range dp {
			dp[i] = make([]int, n)
			for j := range dp[i] {
				dp[i][j] = inf
			}
		}
		dp[1<<n-1][0] = 0
		for s := 1<<n - 2; s >= 0; s-- {
			for v := 0; v < n; v++ {
				for w := 0; w < n; w++ {
					if s>>w&1 == 0 {
						dp[s][v] = min(dp[s][v], dp[s|1<<w][w]+dist[v][w])
					}
				}
			}
		}
		return dp[0][0]
	}

	/* 数位 DP
	入门题 https://atcoder.jp/contests/abc154/tasks/abc154_e
	入门题 https://atcoder.jp/contests/dp/tasks/dp_s
	https://leetcode-cn.com/problems/number-of-digit-one/
	https://leetcode-cn.com/problems/numbers-at-most-n-given-digit-set/
	好题 LC182D https://leetcode-cn.com/problems/find-all-good-strings/
	todo 套题 https://codeforces.ml/blog/entry/53960
	*/
	digitDP := func(lower, upper string) int {
		const mod int = 1e9 + 7

		// <=s 的符合要求的数字/字符串数目
		calc := func(s string) int {
			// 有些题 lowerC 要从 1 开始，而 0 的部分单独计算（由于 0 后面可以填所有数字，这部分可以用 ∑_p>0 f(p, false) 来算）
			const lowerC, upperC byte = '0', '9'
			n := len(s)
			dp := make([][]int, n)
			for i := range dp {
				dp[i] = make([]int, n)
				for j := range dp[i] {
					dp[i][j] = -1
				}
			}
			var f func(p, sum int, isUpper bool) int
			f = func(p, sum int, isUpper bool) (cnt int) {
				//if sum... { return 0 }
				if p >= n {
					return 1 // 0
				}
				dv := &dp[p][sum]
				if !isUpper && *dv >= 0 {
					return *dv
				}
				defer func() {
					if !isUpper {
						*dv = cnt
					}
				}()
				up := upperC
				if isUpper {
					up = s[p]
				}
				for digit := lowerC; digit <= up; digit++ {
					tmp := sum
					// do tmp...
					c := f(p+1, tmp, isUpper && digit == up)
					// do c...
					cnt = (cnt + c) % mod
				}
				return
			}
			return f(0, 0, true)
		}
		ansLower := calc(lower)
		ansUpper := calc(upper)
		ans := ansUpper - ansLower
		// lower 是否算上
		//if lowerIsValid {
		//	ans++
		//}
		ans = (ans%mod + mod) % mod
		return ans
	}

	// 单调队列/单调栈优化
	// https://oi-wiki.org/dp/opt/monotonous-queue-stack/

	// 斜率优化 / 凸包优化 (CHT)
	// https://oi-wiki.org/dp/opt/slope/
	// https://codeforces.ml/blog/entry/63823
	// todo https://blog.csdn.net/weixin_43914593/article/details/105560357
	// todo https://luckyglass.github.io/2019/19Dec21stArt1/
	// 一类单调问题的求解(宋新波) http://www.doc88.com/p-2953873379975.html
	// 题目 https://qiita.com/drken/items/9b311d553aa434bb26e4#%E4%BE%8B%E9%A1%8C-4-4-4k-anonymous-sequence-poj-no3709

	// 四边形不等式优化
	// https://oi-wiki.org/dp/opt/quadrangle/
	// todo https://blog.csdn.net/weixin_43914593/article/details/105150937

	/* 树形 DP
	https://codeforces.ml/blog/entry/20935
	https://codeforces.ml/blog/entry/63257
	CF tag https://codeforces.ml/problemset?order=BY_RATING_ASC&tags=dp%2Ctrees
	https://codeforces.ml/problemset/problem/982/C
	*/

	// todo 公司聚会/树上最大独立集
	// 进阶指南 p.289-290
	// https://stackoverflow.com/questions/13544240/algorithm-to-find-max-independent-set-in-a-tree

	// 树上最大匹配
	// g[v] = ∑{max(f[son],g[son])}
	// f[v] = max{1+g[son]+g[v]−max(f[son],g[son])}
	// https://codeforces.ml/blog/entry/2059
	maxMatchingOnTree := func(n int, g [][]int) int {
		cover, nonCover := make([]int, n), make([]int, n)
		var f func(int, int)
		f = func(v, fa int) {
			for _, w := range g[v] {
				if w != fa {
					f(w, v)
					nonCover[v] += max(cover[w], nonCover[w])
				}
			}
			for _, w := range g[v] {
				cover[v] = max(cover[v], 1+nonCover[w]+nonCover[v]-max(cover[w], nonCover[w]))
			}
		}
		f(0, -1)
		return max(cover[0], nonCover[0])
	}

	// 换根 DP
	// 进阶指南 p.292-295
	// https://codeforces.ml/blog/entry/20935
	// http://poj.org/problem?id=3585
	rerootDP := func(n int) {
		type edge struct{ to, cap int }
		g := make([][]edge, n)
		// read...

		subCap := make([]int, n)
		var f func(v, fa int) int
		f = func(v, fa int) (c int) {
			for _, e := range g[v] {
				if w := e.to; w != fa {
					if len(g[w]) == 1 {
						c += e.cap
					} else {
						c += min(e.cap, f(w, v))
					}
				}
			}
			subCap[v] = c
			return
		}
		f(0, -1)

		ans := make([]int, n)
		var reroot func(v, fa, ansV int)
		reroot = func(v, fa, ansV int) {
			ans[v] = ansV
			for _, e := range g[v] {
				if w, c := e.to, e.cap; w != fa {
					if sc := subCap[w]; len(g[v]) == 1 {
						reroot(w, v, sc+c)
					} else {
						reroot(w, v, sc+min(c, ansV-min(sc, c)))
					}
				}
			}
		}
		reroot(0, -1, subCap[0])
	}

	// 插头 DP / 轮廓线动态规划
	//《训练指南》6.1

	_ = []interface{}{
		mapDP,
		lcs, lcsPath, lis, distinctSubsequence,
		zeroOneKnapsack, waysToSum, unboundedKnapsack, minCoinChange,
		tsp,
		digitDP,
		maxMatchingOnTree, rerootDP,
	}
}
