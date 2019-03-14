package main

import (
	"log"

	"github.com/nolka/goincrementor/types"
)

func main() {

	//Instantiating IncrementorManager for demonstrating interface implementation
	mgr := types.NewManager(types.NewIncrementor())
	log.Println(mgr.GetNumber())
	mgr.IncrementNumber()
	log.Println(mgr.GetNumber())
	mgr.IncrementNumber()
	log.Println(mgr.GetNumber())
	mgr.SetMaxValue(2)
	log.Println(mgr.GetNumber())
	mgr.IncrementNumber()
	log.Println(mgr.GetNumber())

	mgr.IncrementNumber()
	mgr.SetMaxValue(0)
	log.Println(mgr.GetNumber())
}
