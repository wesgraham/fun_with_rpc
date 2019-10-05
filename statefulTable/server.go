package main

import (
	"fmt"
	"sync"
)

type Potato struct {
	Key string `json:"key"`
	Type string `json:"type"`
	Colour string `json:"colour"`
	Age int `json:"age"`
}

type PotatoTable struct {
	Table map[string]Potato
	Lock sync.Mutex
}


var PTable = PotatoTable{Table: make(map[string]Potato, 10)}

type API int

// CRUD Methods on Table
func (api *API) GetPotato(PotatoKey string, reply *Potato) error {
		for k := range PTable.Table {
			if PotatoKey == k {
				*reply = PTable.Table[k]
				return nil
			}
		}
		return fmt.Errorf("unable to locate potato with name: %s", PotatoKey)
}

func (api *API) AddPotato(Potato Potato, reply *int) error {
	for k := range PTable.Table {
		if Potato.Key == k {
			*reply = 1
			return fmt.Errorf("key Collision: %s", Potato.Key)
		}
	}
	PTable.Table[Potato.Key] = Potato
	*reply = 1
	return nil
}

func (api *API) EditPotato(Potato Potato, reply *int) error {
	for k := range PTable.Table {
		if Potato.Key == k {
			PTable.Table[Potato.Key] = Potato
			*reply = 0
			return nil
		}
	}
	*reply = 1
	return fmt.Errorf("key not found: %s", Potato.Key)
}

func (api *API) DeletePotato(PotatoKey string, reply *int) error {
	for k := range PTable.Table {
		if PotatoKey == k {
			delete(PTable.Table, k)
			*reply = 0
			return nil
		}
	}
	*reply = 1
	return fmt.Errorf("key Not Found: %s", PotatoKey)
}

func (api *API) DisplayPotatoTable(TableName string, reply *map[string]Potato) error {
	*reply = PTable.Table
	return nil
}