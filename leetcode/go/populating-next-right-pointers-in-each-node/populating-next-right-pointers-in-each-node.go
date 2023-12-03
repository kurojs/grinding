/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Item struct {
    Node *Node
    Level int
}

func connect(root *Node) *Node {
    queue := make([]*Item, 0)
    mNode := make(map[int][]*Node)
    queue = append(queue, &Item{
        Node: root,
        Level: 0,
    })
    
    depth := 0
    
    for len(queue) > 0 {
        item := queue[0]
        queue = queue[1:]
        
        if item.Node == nil {
            continue
        }
        
        if item.Level > depth {
            depth = item.Level
        }
        
        if _, found := mNode[item.Level]; !found {
            mNode[item.Level] = make([]*Node, 0)
        }
                
        mNode[item.Level] = append(mNode[item.Level], item.Node)
        queue = append(queue, &Item{
            Node: item.Node.Left,
            Level: item.Level+1,
        }, &Item{
            Node: item.Node.Right,
            Level: item.Level+1,
        })
    }
    
    
    
    for level := 0; level <= depth; level++ {
        nodes := mNode[level]
        var prev *Node
        
        for _, node := range nodes {
            if prev == nil {
                prev = node
            } else {
                prev.Next = node
                prev = node
            }
        }
    }
   
    return root
}
