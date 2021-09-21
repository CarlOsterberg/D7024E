package kademlia

//struct for nodelookup sessions
type LookUp struct {
	klist   *ResultList
	sentmap map[string]bool
}

//constructor
func NewLookUp(size int) *LookUp {
	lookup := &LookUp{}
	klist := NewResultList(size)
	lookup.klist = klist
	sentmap := make(map[string]bool)
	lookup.sentmap = sentmap
	return lookup
}
