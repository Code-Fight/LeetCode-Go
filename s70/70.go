package s70

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

思路：
很简单的题目
最简单的解法是直接递归 f(n)=f(n-1)+f(n-2)

但是递归太多会造成栈溢出，那么解决的办法就是不用递归
下面是用for循环的方式来实现，就是一次计算每一层的，知道n层，然后返回

或者可以考虑动态规划的思路。
*/


func climbStairs(n int) int {

	if n==0{
		return 0
	}
	if n ==1{
		return 1
	}
	if n ==2 {
		return 2
	}
	last_1:=2
	last_2:=1
	count:=last_1+last_2
	for n>3 {
		temp := last_1
		last_1= count
		count+=temp
		last_2=temp
		n--
	}
	return count
}
