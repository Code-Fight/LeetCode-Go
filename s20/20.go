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

	if len(s) ==0{
		return true
	}
	if len(s)==1{
		return false
	}

	symbol := map[byte]byte {
		')' : '(',
		']' : '[',
		'}' : '{',
	}

	var leftString []byte

	for  _,value:=range s{

		if byte(value)==0x7B ||byte(value)==0x5B || byte(value)==0x28 {
			leftString = append(leftString,byte(value))
		}else {

			if needstr,ok:=symbol[byte(value)];ok{
				temp:=leftString
				if len(temp)-1<0||(temp[len(temp)-1])!=needstr{
					return false
				}
				leftString=leftString[:len(temp)-1]
			}

		}
	}
	if len(leftString)!=0{
		return false
	}
	return true



}