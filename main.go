package main

import (
	"encoding/json"
	"fmt"
	"github.com/doctornick42/gosli-example/types"
)

func getASlice() []types.A {
	return []types.A{
		types.A{FieldInt: 1},
		types.A{FieldInt: 2},
		types.A{FieldInt: 3},
	}
}

func main() {
	aSlice := getASlice()
	
	aFiltered := types.ASlice(aSlice).
		Where(func(a types.A) bool {
			return a.FieldInt%2 != 0
		})

	filteredAJSON, _ := json.Marshal(aFiltered)
	fmt.Println(string(filteredAJSON))
}
