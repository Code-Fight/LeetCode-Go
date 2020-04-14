package s155



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
		//
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
