// Package main provides ...
package main

/*
* 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
* 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
*
* 准备一个空栈存储柱子
* 1. 查找栈底和栈顶
* 2. 如果栈底小于栈顶，sum = bottom*len-中间柱子高度
* 3. 栈底大于栈顶 sum=top*len-中间柱子高度
*/

type Stack struct {
}

type Pillar struct {
	height
	index
}

func main() {
	stack := Stack{}
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	sum := 0

	for i, h := range height {
		if i == 0 && h == 0 {
			height = height[1:]
			continue
		}
		stack.Push(Pillar{height: h, index: index})
		if stack.Top() > stack.Bottom() && (stack.Bottom().index-stack.Top().index) > 1 {
			sum += stack.Bottom()*(stack.Bottom().index-stack.Top().index-1) - sum(h[stack.Bottom().index+1:stack.Top().index])

		}
	}
}

func push(h) int {

}
