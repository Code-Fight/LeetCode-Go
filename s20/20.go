package s20


/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

思路：
首先判定边界，有两部分
1.长度为0 返回true。
2.长度为1 返回false

然后准备symbol，是一个括号对称的hashtable，为了快速找到对应的
从第一个元素开始遍历，如果是左括号就加入到待匹配的leftString
碰到右括号就从leftString中取最后一个来匹配，直接循环结束

如果结束之后，有剩余的leftString说明不对称，否则对称

*/


func P20(s string) bool {

	// 先判断两边的边界
	if len(s) ==0{
		return true
	}
	if len(s)==1{
		return false
	}

	// 构造一个map 方便快速找到右括号对应的左括号
	// 方便后面去找对应的左括号
	symbol := map[byte]byte {
		')' : '(',
		']' : '[',
		'}' : '{',
	}

	// 读到所有的左括号
	var leftString []byte

	for  _,value:=range s{
		// 0x7B 0x5B 0x28 是上面那几个左括号
		if byte(value)==0x7B ||byte(value)==0x5B || byte(value)==0x28 {
			leftString = append(leftString,byte(value))
		}else {
			// 找左括号对应的右括号
			if needstr,ok:=symbol[byte(value)];ok {
				temp := leftString

				// 如果 leftString 长度为0
				// 或者 最后一个字符不是我们需要左括号 也就是皮配不上
				// 直接结束 false
				if len(temp)-1 < 0 || (temp[len(temp)-1]) != needstr {
					return false
				}

				// 删掉一个最后加入的
				// 然后继续遍历查看后面的
				leftString = leftString[:len(temp)-1]
			}

		}

	}
	// 遍历s结束后
	// 如果 还有没匹配上的左括号
	// 说明 没有匹配
	if len(leftString)!=0{
		return false
	}
	return true



}