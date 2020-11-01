/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    if root==nil{
        return nil
    }
    queue:=make([]*Node,0)
    queue=append(queue, root)
    result:=make([][]int,0)
    thisLayer:=1
    nextLayer:=0
    thisLine:=make([]int,0)
    for len(queue)!=0{
        if thisLayer>0{
            thisNode:=queue[0]
            queue=queue[1:]
            thisLine=append(thisLine, thisNode.Val)
            thisLayer--
            for _,e:=range thisNode.Children{
                queue=append(queue,e)
                nextLayer++
            }
        }else{
            thisLayer=nextLayer
            nextLayer=0
            result=append(result,thisLine)
            thisLine=make([]int,0)
        }
    }
    result=append(result, thisLine)
    return result
}
