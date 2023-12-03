type LRUCache struct {
    head, tail *Node
    capacity int
    dict map[int]*Node
}


func Constructor(capacity int) LRUCache {
    head := &Node{}
    tail := &Node{}
    head.next = tail
    tail.prev = head
    return LRUCache{
        head: head,
        tail: tail,
        capacity: capacity,
        dict: make(map[int]*Node),
    }
}


func (this *LRUCache) Get(key int) int {
    node, found := this.dict[key]
    if !found {
        return -1
    }
    
    this.remove(node)
    this.insert(&Node{
        val: node.val,
        key: node.key,
    })
    
    return node.val
}


func (this *LRUCache) Put(key int, value int)  {
    if node, found := this.dict[key]; found {
        this.remove(node)
    }
    
    if len(this.dict) == this.capacity {
        this.remove(this.tail.prev)
    }
    
    this.insert(&Node{
        val: value,
        key: key,
    })
}


func (this *LRUCache) insert(node *Node) {
    this.dict[node.key] = node
    
    next := this.head.next
    this.head.next = node
    next.prev = node
    node.prev = this.head
    node.next = next
}

func (this *LRUCache) remove(node *Node) {
    delete(this.dict, node.key)
    node.prev.next = node.next
    node.next.prev = node.prev
}

type Node struct {
    prev, next *Node
    val, key int
}
