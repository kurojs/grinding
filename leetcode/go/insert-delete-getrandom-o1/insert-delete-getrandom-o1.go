type RandomizedSet struct {
    dict map[int]int
    rand []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
    return RandomizedSet{
        dict: make(map[int]int),
        rand: make([]int, 0),
    }
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
    // defer func() {
    //     fmt.Println("-insert-")
    //     fmt.Println(val)
    //     fmt.Println(this.dict)
    //     fmt.Println(this.rand)
    // }()
    if _, found := this.dict[val]; found {
        return false
    }
    
    this.dict[val] = len(this.rand)
    this.rand = append(this.rand, val)
    

    return true
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
//     defer func() {
//         fmt.Println("-remove-")
//         fmt.Println(val)
//         fmt.Println(this.dict)
//         fmt.Println(this.rand)
//     }()
    
    idx, found := this.dict[val]
    if !found {
        return false
    }
    
    // check if not the last element, swap then remove for best performance
    if idx != len(this.rand) - 1 {
        lastIdxVal := this.rand[len(this.rand) - 1]
        this.rand[idx] = lastIdxVal
        this.dict[lastIdxVal] = idx
        
    }
  
    this.rand = this.rand[:len(this.rand)-1]
    delete(this.dict, val)
    return true
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
    rand.Seed(time.Now().UnixNano())
    if len(this.rand) == 0 {
        return -1
    }
    
    return this.rand[rand.Intn(len(this.rand))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */