package kademlia

//struct for nodelookup sessions
type LookUp struct {
	klist   *ResultList
	sentmap map[string]bool
	rpctype string
	value   []byte
}

//constructor
func NewLookUp(size int, rpctype string, data []byte) *LookUp {
	lookup := &LookUp{}
	klist := NewResultList(size)
	lookup.klist = klist
	sentmap := make(map[string]bool)
	lookup.sentmap = sentmap
	lookup.rpctype = rpctype
	lookup.value = data
	return lookup
}
