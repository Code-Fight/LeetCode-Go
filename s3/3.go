package s3

import "math"

/*
给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。

示例1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
    请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。

思路：
参考：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/hua-jie-suan-fa-3-wu-zhong-fu-zi-fu-de-zui-chang-z/
标签：滑动窗口
暴力解法时间复杂度较高，会达到 O(n^2)，故而采取滑动窗口的方法降低时间复杂度


定义一个 map 数据结构存储 (k, v)，其中 key 值为字符，value 值为字符位置 +1，加 1 表示从字符位置后一个才开始不重复
我们定义不重复子串的开始位置为 start，结束位置为 end
随着 end 不断遍历向后，会遇到与 [start, end] 区间内字符相同的情况，此时将字符作为 key 值，获取其 value 值，并更新 start，此时 [start, end] 区间内不存在重复字符
无论是否更新 start，都会更新其 map 数据结构和结果 。
时间复杂度：O(n)
*/
func lengthOfLongestSubstring(s string) int {
	if len(s)==0{
		return 0
	}
	lastMax :=map[uint8]int{}
	// 用来记录最大的长度
	lastMaxVal :=0
	// 不重复开始的地方
	startIndex :=0


	for endIndex:=0 ;endIndex<len(s);endIndex++{

		// 从map中查找该字符是否存在
		// 如果存在，就判断key对应的val和当前startIndex哪个大，保留大的
		// 因为大的那个肯定是出现在后面的，后面的代表重复字符的最后一个
		if v,ok:=lastMax[s[endIndex]];ok{
			startIndex = int(math.Max(float64(startIndex),float64(v)))
		}


		// 对比最新的不重复串长度 和 之前的最大长度 哪个大 保留 长的
		lastMaxVal = int(math.Max(float64(lastMaxVal),float64(endIndex- startIndex +1)))

		// 每次前进 都会将此map中相应位置的保存下来
		// 更新当前的最大不重复串的其实位置  就是当前s的endIndex，并把这个作为key 保存到map中，方便后面再次对比
		lastMax[s[endIndex]]=endIndex+1
	}
	return lastMaxVal
}


//
//func lengthOfLongestSubstring(s string) int {
//	if len(s)==0{
//		return 0
//	}
//	lastMax :=map[uint8]int{}
//	lastMaxVal :=0
//	curMaxVal :=0
//
//	for i:=0 ;i<len(s);i++{
//		if v,ok:=lastMax[s[i]];ok {
//			if curMaxVal >lastMaxVal{
//				lastMaxVal = curMaxVal
//			}
//			//i = v
//			curMaxVal =0
//			lastMax =make(map[uint8]int)
//		}else {
//			curMaxVal ++
//			lastMax[s[i]] = i
//		}
//	}
//	if curMaxVal >lastMaxVal{
//		lastMaxVal = curMaxVal
//
//	}
//
//	return lastMaxVal
//}
