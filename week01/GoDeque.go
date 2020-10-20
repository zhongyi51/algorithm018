type MyCircularDeque struct {
	queue []int
	maxLength int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{make([]int,0),k}
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if len(this.queue)==this.maxLength{
		return false
	}
	this.queue=append([]int{value},this.queue...)
	return true

}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if len(this.queue)==this.maxLength{
		return false
	}
	this.queue=append(this.queue,value)
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if len(this.queue)==0{
		return false
	}
	this.queue=this.queue[1:]
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if len(this.queue)==0{
		return false
	}
	this.queue=this.queue[:len(this.queue)-1]
	return true

}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if len(this.queue)==0{
		return -1
	}
	return this.queue[0]

}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if len(this.queue)==0{
		return -1
	}
	return this.queue[len(this.queue)-1]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if len(this.queue)==0{
		return true
	}
	return false
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if len(this.queue)==this.maxLength{
		return true
	}
	return false
}
