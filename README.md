# golang_merge_triplets_to_form_target_triplet

A **triplet** is an array of three integers. You are given a 2D integer array `triplets`, where `triplets[i] = [ai, bi, ci]` describes the `ith` **triplet**. You are also given an integer array `target = [x, y, z]` that describes the **triplet** you want to obtain.

To obtain `target`, you may apply the following operation on `triplets` **any number** of times (possibly **zero**):

- Choose two indices (**0-indexed**) `i` and `j` (`i != j`) and **update** `triplets[j]` to become `[max(ai, aj), max(bi, bj), max(ci, cj)]`.
    - For example, if `triplets[i] = [2, 5, 3]` and `triplets[j] = [1, 7, 5]`, `triplets[j]` will be updated to `[max(2, 1), max(5, 7), max(3, 5)] = [2, 7, 5]`.

Return `true` *if it is possible to obtain the* `target` ***triplet*** `[x, y, z]` *as an **element** of* `triplets`*, or* `false` *otherwise*.

## Examples

**Example 1:**

```
Input: triplets = [[2,5,3],[1,8,4],[1,7,5]], target = [2,7,5]
Output: true
Explanation: Perform the following operations:
- Choose the first and last triplets [[2,5,3],[1,8,4],[1,7,5]]. Update the last triplet to be [max(2,1), max(5,7), max(3,5)] = [2,7,5]. triplets = [[2,5,3],[1,8,4],[2,7,5]]
The target triplet [2,7,5] is now an element of triplets.

```

**Example 2:**

```
Input: triplets = [[3,4,5],[4,5,6]], target = [3,2,5]
Output: false
Explanation: It is impossible to have [3,2,5] as an element because there is no 2 in any of the triplets.

```

**Example 3:**

```
Input: triplets = [[2,5,3],[2,3,4],[1,2,5],[5,2,3]], target = [5,5,5]
Output: true
Explanation:Perform the following operations:
- Choose the first and third triplets [[2,5,3],[2,3,4],[1,2,5],[5,2,3]]. Update the third triplet to be [max(2,1), max(5,2), max(3,5)] = [2,5,5]. triplets = [[2,5,3],[2,3,4],[2,5,5],[5,2,3]].
- Choose the third and fourth triplets [[2,5,3],[2,3,4],[2,5,5],[5,2,3]]. Update the fourth triplet to be [max(2,5), max(5,2), max(5,3)] = [5,5,5]. triplets = [[2,5,3],[2,3,4],[2,5,5],[5,5,5]].
The target triplet [5,5,5] is now an element of triplets.

```

**Constraints:**

- `1 <= triplets.length <= 105`
- `triplets[i].length == target.length == 3`
- `1 <= ai, bi, ci, x, y, z <= 1000`

## 解析

題目給一個 2D 陣列 triplets , 還有一個 陣列 target

其中 每個 triplets[i] 代表一個 3 個元素 的陣列

target 也是一個 3 個元素的陣列

定義 merge(i, j) = [$max(a_i, a_j), max(b_i, b_j), max(c_i, c_j)$], 對所有 triplets[i] = [$a_i, b_i, c_i$], triplets[j] = [$a_j, b_j, c_j$], i ≠ j

對陣列triplets 任意兩個元素可以使用 merge 這個運算使用 0 到 無限整數次

0 次代表不做 merge

要求實作一個演算法來判斷 triplets 是否可以透過 merge 這個運算來推導出 target

首先，可以知道

當對一個 [$a_i, b_i, c_i]$ 與 [$a_j, b_j, c_j$] 做一次 merge 後 假設結果是 [$maxA_{i,j}, maxB_{i,j}, maxB_{i,j}$]

 [$maxA_{i,j}, maxB_{i,j}, maxB_{i,j}$] 不論對 [$a_i, b_i, c_i]$  或 [$a_j, b_j, c_j$]  做 merge 後

都還是  [$maxA_{i,j}, maxB_{i,j}, maxB_{i,j}$]

所以對兩個triplet 可以 merge 的選擇，其實只需要考慮一次或是不做即可

另外透過上面的特性可以發現，當遇到其中一個陣列元素是大於 target 的另一個元素後

不管如何 merge 都無法完成

所以只要出現 triplet 中有元素大於 target 其中一個 元素就不做 merge

已知所以陣列元素都是 大於 1

所以初始化 maxTriplet = [0,0,0]

透過決策樹可以看到如下圖

![](https://i.imgur.com/DWrAVWY.png)

因為從起點開始 每個元素都可以選擇要 merge 或是不 merge

所以透過 DFS 走訪所有可能會是 O($2^n$)

透過圖可以發現有些點的選擇是重複的 

而從 n 個結點出發 有 n 種可能 又有 merge 或不 merge 兩種選擇

所以透過 cache 的方式 會是 2 *n 種可能 ，但需要走 n 次

時間複雜度可以降低到 O(2n*n) = O($n^2$)

空間複雜度是 O(n)

然而  對於 max 的特性

可以發現 假設是 global max 也一定是 local max

所以透過 這個 greedy 特性

每次只要把 把 小於等於 target 的 triplets 做 merge

當有一個 merge 值等於 target 回傳 true 

當最後如果 merge 直還是不相等 則代表無法做到 回傳 false

這樣就可以把時間複雜度將低到 O(n)

且不需要額外的儲存空間

所以空間複雜度是 O(1)

## 程式碼
```go
package sol

func mergeTriplets(triplets [][]int, target []int) bool {
	possibles := make([]bool, 3)
	for _, triplet := range triplets {
		if triplet[0] <= target[0] && triplet[1] <= target[1] && triplet[2] <= target[2] {
			for idx := range possibles {
				if !possibles[idx] && triplet[idx] == target[idx] {
					possibles[idx] = true
				}
			}
			if possibles[0] && possibles[1] && possibles[2] {
				return true
			}
		}
	}
	return false
}
```
## 困難點

1. 需要看出 當 merge 到過大的數值會無法轉換成 target
2. 需要看出 取出 max 這個特性具有 greedy, 也就是 global max 一定是 local max 所找出來的，所以只要在每個 不大於 target 的數值找出 max ， 當這個 max == target 就代表有可能 

## Solve Point

- [x]  建立一個大小為 3 的 boolean 陣列 possible , default 是 false 用來紀錄每個分量找到可能值的狀態
- [x]  逐步檢查  triplets每個 triplets[i] 當  triplets[i] 3個分量都小於等於 target 分量時 做以下檢查
- [x]  當 triplets[i][0] == target[0] && possible[0] == false 時 ， possible[0] = true
- [x]  當 triplets[i][1] == target[1] && possible[1] == false 時 ， possible[1] = true
- [x]  當 triplets[i][2] == target[2] && possible[2] == false 時 ， possible[2] = true
- [x]  當 possible[0] && possible[1] && possible[0] 時, return true
- [x]  全部 loop 完都沒有 return true ， 代表沒有可以 merge 的可能性 return false