type Stack struct{
	val []Node
}

type Node struct{
	index int
	hei int
}

type StackMethod interface{
	pop()int
	push(e int)
	look()int
}

func (s *Stack) pop() Node{
	if len(s.val)==0{
		return Node{-1,0}
	}
	res:=s.val[len(s.val)-1]
	s.val=s.val[:len(s.val)-1]
	return res

}

func (s *Stack) push(e Node){
	s.val=append(s.val,e)
	return
}

func (s *Stack) look()Node{
	if len(s.val)==0{
		return Node{-1,0}
	}
	return  s.val[len(s.val)-1]
}

func min(a,b int)int{
	if a>b{
		return b
	}else{
		return a
	}
}

func trap(height []int) int {
	if len(height)<=2{
		return 0
	}
	stack:=Stack{val:make([]Node,0)}
	res:=0
	stack.push(Node{0,height[0]})
	for i,v:=range height[1:]{
		//fmt.Println(stack)
		//fmt.Println(res)
		if stack.look().hei>=v{
			stack.push(Node{i+1,v})
		}else{
			this:=stack.look()
			for this.hei<v&&len(stack.val)>0{
				this=stack.pop()
				prev:=stack.look()
				if prev.index!=-1{
					res+=(min(v,prev.hei)-this.hei)*(i-prev.index)
				}
				this=prev
			}
			stack.push(Node{i+1,v})

		}
	}
	return res
}
