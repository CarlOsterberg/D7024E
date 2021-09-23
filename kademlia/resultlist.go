package kademlia

//ResultList keeps a list of KademliaID with a max size
type ResultList struct {
	List []Contact
	Size int
}

func NewResultList(size int) *ResultList {
	resultList := &ResultList{}
	resultList.Size = size
	return resultList
}

//Insert tries to insert an id to the list. If it is better than an existing element it will replace the worst one.
func (resultList *ResultList) Insert(contact Contact, target KademliaID) {
	shouldInsert := false
	var worstIdx int
	var worstDistance *KademliaID
	idDistance := contact.ID.CalcDistance(&target)

	if len(resultList.List) < resultList.Size {
		resultList.List = append(resultList.List, contact)
		return
	}
	for i, v := range resultList.List {
		dist := v.ID.CalcDistance(&target)
		//One of the elements are worse than id
		if dist.Less(idDistance) && !shouldInsert {
			shouldInsert = true
			worstIdx = i
			worstDistance = dist
		} else if !dist.Less(worstDistance) { //If an even worse element is found
			worstIdx = i
			worstDistance = dist
		}
	}

	if shouldInsert {
		resultList.List[worstIdx] = contact
	}
}

//Merge tries to insert all elements in another list into our current list.
func (resultList *ResultList) Merge(other *ResultList, target KademliaID) {
	for _, v := range other.List {
		resultList.Insert(v, target)
	}
}
