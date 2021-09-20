package kademlia


//ResultList keeps a lit of KademliaID with a max size
type ResultList struct {
	List    []KademliaID
	Size	int
}

func NewResultList(size int) *ResultList{
	resultList := &ResultList{}
	resultList.Size = size
	return resultList
}


//Insert tries to insert an id to the list. If it is better than an existing element it will replace the worst one.
func (resultList *ResultList) Insert(id KademliaID, target KademliaID){
	shouldInsert := false
	var worstIdx int
	var worstDistance *KademliaID
	idDistance := id.CalcDistance(&target)

	if len(resultList.List) < resultList.Size {
		resultList.List = append(resultList.List, id)
		return
	}
	for i, v := range resultList.List{
		dist := v.CalcDistance(&target)
		//One of the elements are worse than id
		if dist.Less(idDistance) && !shouldInsert{
			shouldInsert = true
			worstIdx = i
			worstDistance = dist
		} else if !dist.Less(worstDistance){ //If an even worse element is found
			worstIdx = i
			worstDistance = dist
		}
	}

	if shouldInsert {
		resultList.List[worstIdx] = id
	}
}

//Merge tries to insert all elements in another list into our current list.
func (resultList *ResultList) Merge(other *ResultList, target KademliaID){
	for _, v:= range other.List{
		resultList.Insert(v, target)
	}
}