package main

import (
	"clo-engine/internal/algorithm"
	"clo-engine/internal/models"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	// ./clo-engine input.json output.json
	if len(os.Args) < 3 {
		fmt.Println("Usage: clo-engine <input.json> <output.jaon>")
		return
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	// read input.json
	inputData, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// decode input
	var input struct {
		Container models.Container      `json:"container"`
		Items     []models.ItemInstance `json:"items"`
	}

	err = json.Unmarshal(inputData, &input)
	if err != nil {
		fmt.Println("Invalid JSON:", err)
		return
	}

	// convert item list to pointers
	items := []*models.ItemInstance{}
	// map to track counts per parent, so numbering is sequential
	counts := make(map[string]int)

	for i := range input.Items {
		src := input.Items[i]

		qty := src.Quantity
		if qty <= 0 {
			qty = 1
		}

		// ensure ParentItemID exists
		parentID := src.ParentItemID
		if parentID == "" {
			parentID = fmt.Sprintf("ITEM-%id", i+1)
		}

		for k := 0; k < qty; k++ {
			counts[parentID]++
			instanceNum := counts[parentID]

			// create new instance copy
			inst := &models.ItemInstance{
				ParentItemID:   parentID,
				OriginalLength: src.OriginalLength,
				OriginalWidth:  src.OriginalWidth,
				OriginalHeight: src.OriginalHeight,
				Quantity:       1, // after expansion each instance is quantity 1
				AllowRotation:  src.AllowRotation,
			}

			inst.InstanceID = fmt.Sprintf("%s#%d", parentID, instanceNum)
			items = append(items, inst)
		}
	}

	// run packing algo
	result := algorithm.RunFFD3D(&input.Container, items)

	// encode result as json
	outputJSON, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		fmt.Println("Error generating output JSON:", err)
		return
	}

	// save output file
	err = os.WriteFile(outputPath, outputJSON, 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		return
	}

	fmt.Println("Packing complete. Output saved to:", outputPath)
}
