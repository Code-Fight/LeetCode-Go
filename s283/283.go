package s283

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
通过次数136,380提交次数

思路：
最简单暴力法复杂度高O(n^2)
优化使用双指针，慢指针遇到0停下，快指针越过0，但是遇到非0停下，每次判断慢指针是否为0，如果为不是0，就快慢指针同时前进一次，
如果是0，那么快指针移动一次，如果判断快指针不是0，就交换快慢指针
*/


func moveZeroes(nums []int)  {
	slowIndex :=0
	fastIndex :=0

	for  fastIndex<len(nums){


		if nums[slowIndex] !=0{
			slowIndex++
			fastIndex++
		}else {


			if nums[fastIndex] ==0{
				fastIndex ++
			}else {
				nums[fastIndex],nums[slowIndex]=nums[slowIndex],nums[fastIndex]
			}
		}

	}
}


//暴力双层循环
//func moveZeroes(nums []int)  {
//	moveCount := 0
//	for i :=0;i+moveCount<len(nums);{
//		if nums[i]==0{
//
//			for j:=i;j<i+len(nums)-moveCount-i-1;j++{
//				nums[j+1],nums[j]=nums[j],nums[j+1]
//			}
//			moveCount++
//		}else {
//			i++
//		}
//
//	}
//}
