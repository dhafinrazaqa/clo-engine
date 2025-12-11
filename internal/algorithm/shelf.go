package algorithm

import "clo-engine/internal/models"

func NewShelf(index int, startZ int, height int) *models.Shelf {

	Debug("Creating shelf %d at Z=%d (initial height=%d)", index, startZ, height)

	return &models.Shelf{
		Index:    index,
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

	Debug("Placing %s on shelf %d at X=%d, Z=%d (dims=%dx%dx%d)",
		item.InstanceID, shelf.Index, shelf.CurrentX, shelf.StartZ,
		orient.Length, orient.Width, orient.Height)

	item.ChosenOrientation = orient

	item.X = shelf.CurrentX
	item.Y = 0
	item.Z = shelf.StartZ // shelf floor level

	shelf.CurrentX += orient.Length

	if orient.Height > shelf.Height {
		Debug("Shelf %d height updated from %d -> %d", shelf.Index, shelf.Height, orient.Height)
		shelf.Height = orient.Height
	}

	shelf.Items = append(shelf.Items, item)
}
