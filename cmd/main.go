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
	for i := range input.Items {
		item := &input.Items[i]

		// generate unique instance id
		item.InstanceID = fmt.Sprintf("%s#%d", item.ParentItemID, i+1)
		items = append(items, item)
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
