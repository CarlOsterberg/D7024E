package kademlia

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"program/udp"
	"sync"
)

func Run(state Kademlia, cliCh chan string) {
	/*Docker should always create the node 10.20.0.2
	use that as bootstrapping node to start the network*/
	var stateMutex = &sync.Mutex{}
	if udp.GetOutboundIP().String() != "10.20.0.2" {
		ip := "10.20.0.2:1234"
		stateMutex.Lock()
		public := NewContact(NewSha1KademliaID([]byte(ip)), ip)
		//Add public node to routing table
		state.routingTable.AddContact(public)
		//The perform the join rpc
		joinlookup := NewLookUp(k, "JOIN", []byte(""))
		myid := NewSha1KademliaID([]byte(state.network.Self))
		convID, _ := uuid.NewV4()
		state.convIDMap[*convID] = *joinlookup
		state.network.SendFindContactMessage(&public, *convID, *myid)
		stateMutex.Unlock()
	}
	debug := false
	for {
		Server_Channel(state, stateMutex, debug)
		Cli_Channel(state, stateMutex, cliCh, debug)
	}
}

func CheckMsgChainDone(state Kademlia, lookup LookUp, ConvID uuid.UUID) {
	count := 0
	for _, responded := range lookup.sentmap {
		if !responded {
			count++
		}
	}

	if count == 0 {
		//All contacts have responded, we are done
		if lookup.rpctype == "STORE" {
			//Instruct the nodes to store
			for _, v := range lookup.klist.List {
				state.network.SendStoreMessage(lookup.value, v.Address)
			}
		}
		if lookup.rpctype == "JOIN" {
			for _, v := range lookup.klist.List {
				joinid := NewSha1KademliaID([]byte(v.Address))
				state.routingTable.AddContact(NewContact(joinid, v.Address))
			}
		}
		if lookup.rpctype == "GET" {
			if !lookup.foundValue {
				fmt.Println("The value for the given key was not found.")
			} else {
				fmt.Println("The value for the given key is: ", string(lookup.value))
			}
		}
		//Delete the lookup when we are done with the conversation
		delete(state.convIDMap, ConvID)
	}

}
