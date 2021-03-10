package s155
/*
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素


思路：
栈可以直接存一个最小值来解决getMin的问题，
每次push和pop数据必须先判断最小值是否产生了变化
然后，使用一个切片来保存这个栈即可
*/


type MinStack struct {
	min int
	arr []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{0,make([]int,0)}
}


func (this *MinStack) Push(x int)  {
	if len(this.arr)==0{
		this.min =x
	}

	if x < this.min{
		this.min =x
	}
	this.arr = append(this.arr, x)
}


func (this *MinStack) Pop()  {

	if(len(this.arr)==0){
		return
	}
	//if(len(this.arr)==1){
	//	this.arr =make([]int,10)
	//	return
	//}

	cur := this.arr[len(this.arr)-1]
	this.arr = this.arr[0:len(this.arr)-1]
	if cur == this.min{
		//如果pop后 把最小值pop了，那么需要更换新的最小值
		if len(this.arr)>0{
			this.min = this.arr[0]
			for i:=1 ;i<len(this.arr);i++{
				if this.arr[i] < this.min{
					this.min =this.arr[i]
				}
			}
		}
		
	}
	
	
}


func (this *MinStack) Top() int {
	if len(this.arr)==0{
		return 0
	}
	return this.arr[len(this.arr)-1]
}


func (this *MinStack) GetMin() int {
	if len(this.arr)==0{
		return 0
	}
	return this.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
