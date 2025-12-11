package algorithm

import (
	"clo-engine/internal/models"
	"clo-engine/internal/util"
	"sort"
)

func RunFFD3D(container *models.Container, items []*models.ItemInstance) *models.PackingResult {

	var shelves []*models.Shelf
	var placements []models.Placement
	var unpacked []string

	currentZ := 0

	// sort items by volume
	sort.Slice(items, func(i, j int) bool {
		return util.CalculateVolume(items[i]) > util.CalculateVolume(items[j])
	})

	// iterate from largest to smallest
	for _, item := range items {
		placed := false

		orientations := util.GenerateOrientations(item)

		for _, orient := range orientations {

			// check if orientation fits
			if !util.FitsInContainer(orient, container) {
				continue
			}

			for _, shelf := range shelves {
				if util.FitsInShelf(orient, shelf, container) {
					PlaceItemOnShelf(item, orient, shelf)

					placements = append(placements, models.Placement{
						ItemInstanceID: item.InstanceID,
						Orientation:    orient.Name,

						ShelfIndex:  shelf.Index,
						ShelfStartZ: shelf.StartZ,
						ShelfHeight: shelf.Height,

						X: item.X,
						Y: item.Y,
						Z: item.Z,

						Length: orient.Length,
						Width:  orient.Width,
						Height: orient.Height,

						Volume: orient.Length * orient.Width * orient.Height,
					})

					placed = true
					break
				}
			}

			if placed {
				break
			}

			// try placing in other shelf if none worked
			shelfHeight := orient.Height

			// check vertical overflow
			if currentZ+shelfHeight > container.Height {
				unpacked = append(unpacked, item.InstanceID+" (height overflow)")
				placed = true // handled as unpacked
				break
			}

			// make new shelf
			newShelf := NewShelf(len(shelves), currentZ, shelfHeight)
			shelves = append(shelves, newShelf)

			PlaceItemOnShelf(item, orient, newShelf)

			placements = append(placements, models.Placement{
				ItemInstanceID: item.InstanceID,
				Orientation:    orient.Name,

				ShelfIndex:  newShelf.Index,
				ShelfStartZ: newShelf.StartZ,
				ShelfHeight: newShelf.Height,

				X: item.X,
				Y: item.Y,
				Z: item.Z,

				Length: orient.Length,
				Width:  orient.Width,
				Height: orient.Height,

				Volume: orient.Length * orient.Width * orient.Height,
			})

			// update next shelf z lvl
			currentZ += newShelf.Height

			placed = true
			break
		}

		if !placed {
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

	metrics := models.Metrics{
		ContainerVolume:    containerVolume,
		PackedVolume:       packedVolume,
		UtilizationPercent: utilization,
		TotalItems:         len(items),
		ItemsPacked:        len(placements),
		ItemsUnpacked:      len(unpacked),
	}

	// final result
	return &models.PackingResult{
		Status:        "success",
		Metrics:       metrics,
		Placements:    placements,
		UnpackedItems: unpacked,
	}
}
