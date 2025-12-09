package algorithm

import "clo-engine/internal/models"

func NewShelf(startZ int, height int) *models.Shelf {
	return &models.Shelf{
		StartZ:   startZ,
		Height:   height,
		CurrentX: 0,
		Items:    []*models.ItemInstance{},
	}
}

func PlaceItemOnShelf(
	item *models.ItemInstance,
	orient models.Orientation,
	shelf *models.Shelf,
) {
	item.ChosenOrientation = orient

	item.X = shelf.CurrentX
	item.Y = 0
	item.Z = shelf.StartZ // shelf floor level

	shelf.Items = append(shelf.Items, item)

	shelf.CurrentX += orient.Length

	if orient.Height > shelf.Height {
		shelf.Height = orient.Height
	}
}
