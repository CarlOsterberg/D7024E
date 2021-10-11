package kademlia

import "sort"

//ResultList keeps a list of KademliaID with a max size
type ResultList struct {
	List []Contact
	Size int
}

func NewResultList(size int) *ResultList {
	resultList := &ResultList{}
	reslist := make([]Contact, 0, size)
	resultList.List = reslist
	resultList.Size = size
	return resultList
}

var sortingTarget KademliaID

type byTargetDistance []Contact

func (b byTargetDistance) Len() int{
	return len(b)
}

func (b byTargetDistance) Swap(i, j int){
	b[i], b[j] = b[j], b[i]
}

func (b byTargetDistance) Less(i, j int) bool{
	return b[i].ID.CalcDistance(&sortingTarget).Less(b[j].ID.CalcDistance(&sortingTarget))
}

//Insert tries to insert an id to the list. If it is better than an existing element it will replace the worst one.
func (resultList *ResultList) Insert(contact Contact, target KademliaID) {
	shouldInsert := false
	var worstIdx int
	//var worstDistance *KademliaID
	worstDistance := NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")
	idDistance := contact.ID.CalcDistance(&target)
	worstDistance = NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")
	//	fmt.Println("klist längd: ", len(resultList.List))
	//	fmt.Println(resultList.List)

	for _, v := range resultList.List {
		if v.ID.Equals(contact.ID) {
			return
		}
	}

	if len(resultList.List) < resultList.Size {
		resultList.List = append(resultList.List, contact)
		return
	}

	for i, v := range resultList.List {
		dist := v.ID.CalcDistance(&target)
		//One of the elements are worse than id
		if idDistance.Less(dist) && !shouldInsert {
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

func (resultList *ResultList) Delete(contact Contact){
	idx := -1
	for i, v := range resultList.List{
		if v.ID.Equals(contact.ID) && v.Address == contact.Address{
			idx = i
		}
	}
	if idx != -1{
		resultList.List = append(resultList.List[:idx], resultList.List[idx+1:]...)
	}
}

//Sort sorts the list by distance to the target, only one thread may use this
func (resultList *ResultList) Sort(target KademliaID){
	sortingTarget = target
	sort.Sort(byTargetDistance(resultList.List))
}