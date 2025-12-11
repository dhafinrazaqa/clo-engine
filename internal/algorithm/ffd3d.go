package algorithm

import (
	"clo-engine/internal/models"
	"clo-engine/internal/util"
	"sort"
)

func RunFFD3D(container *models.Container, items []*models.ItemInstance) *models.PackingResult {

	Debug("Starting FFD-3D algorithm witth %d items", len(items))

	var shelves []*models.Shelf
	var placements []models.Placement
	var unpacked []string

	currentZ := 0

	// sort items by volume
	sort.Slice(items, func(i, j int) bool {
		return util.CalculateVolume(items[i]) > util.CalculateVolume(items[j])
	})

	Debug("Items sorted by volume")

	// iterate from largest to smallest
	for _, item := range items {

		Debug("Placing item %s...", item.InstanceID)

		placed := false

		orientations := util.GenerateOrientations(item)

		for _, orient := range orientations {

			Debug("Trying orientation %s (%dx%dx%d)",
				orient.Name, orient.Length, orient.Width, orient.Height)

			// check if orientation fits
			if !util.FitsInContainer(orient, container) {
				Debug("Orientation %s does NOT fit in container", orient.Name)
				continue
			}

			Debug("Orientation %s fits container", orient.Name)

			for _, shelf := range shelves {
				if util.FitsInShelf(orient, shelf, container) {

					Debug("Fits in shelf %d", shelf.Index)

					PlaceItemOnShelf(item, orient, shelf)

					placements = append(placements, makePlacement(item, orient, shelf))

					placed = true
					break
				}
				Debug("Does not fit in shelf %d", shelf.Index)
			}

			if placed {
				break
			}

			// try placing in other shelf if none worked
			shelfHeight := orient.Height

			// check vertical overflow
			if currentZ+shelfHeight > container.Height {
				Debug("Cannot create shelf: height overflow")
				unpacked = append(unpacked, item.InstanceID+" (height overflow)")
				placed = true // handled as unpacked
				break
			}

			Debug("Creating new shelf at Z=%d", currentZ)

			// make new shelf
			newShelf := NewShelf(len(shelves), currentZ, shelfHeight)
			shelves = append(shelves, newShelf)

			PlaceItemOnShelf(item, orient, newShelf)

			placements = append(placements, makePlacement(item, orient, newShelf))

			// update next shelf z lvl
			currentZ += newShelf.Height

			placed = true
			break
		}

		if !placed {
			Debug("Item %s could not be placed", item.InstanceID)
			unpacked = append(unpacked, item.InstanceID+" (no orientation fits)")
		}
	}

	// metrics
	containerVolume := container.Length * container.Width * container.Height

	packedVolume := 0
	for _, p := range placements {
		packedVolume += p.Volume
	}

	utilization := (float64(packedVolume) / float64(containerVolume)) * 100

	Debug("Packing completed. %d items placed, %d unpacked.", len(placements), len(unpacked))

	// final result
	return &models.PackingResult{
		Status: "success",
		Metrics: models.Metrics{
			ContainerVolume:    containerVolume,
			PackedVolume:       packedVolume,
			UtilizationPercent: utilization,
			TotalItems:         len(items),
			ItemsPacked:        len(placements),
			ItemsUnpacked:      len(unpacked),
		},
		Placements:    placements,
		UnpackedItems: unpacked,
	}
}

func makePlacement(item *models.ItemInstance, orient models.Orientation, shelf *models.Shelf) models.Placement {
	return models.Placement{
		ItemInstanceID: item.InstanceID,
		Orientation:    orient.Name,
		ShelfIndex:     shelf.Index,
		ShelfStartZ:    shelf.StartZ,
		ShelfHeight:    shelf.Height,
		X:              item.X,
		Y:              item.Y,
		Z:              item.Z,
		Length:         orient.Length,
		Width:          orient.Width,
		Height:         orient.Height,
		Volume:         orient.Length * orient.Width * orient.Height,
	}
}
