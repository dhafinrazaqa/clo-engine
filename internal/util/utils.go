package util

import (
	"clo-engine/internal/models"
	"errors"
)

func CalculateVolume(it *models.ItemInstance) int {
	return it.OriginalLength * it.OriginalWidth * it.OriginalHeight
}

func RotateXYZ(it *models.ItemInstance) models.Orientation {
	return models.Orientation{
		Name:   "XYZ",
		Length: it.OriginalLength,
		Width:  it.OriginalWidth,
		Height: it.OriginalHeight,
	}
}

func RotateYXZ(it *models.ItemInstance) models.Orientation {
	return models.Orientation{
		Name:   "YXZ",
		Length: it.OriginalWidth,
		Width:  it.OriginalLength,
		Height: it.OriginalHeight,
	}
}

func GenerateOrientations(it *models.ItemInstance) []models.Orientation {
	if it.AllowRotation {
		return []models.Orientation{
			RotateXYZ(it),
			RotateYXZ(it),
		}
	}
	return []models.Orientation{
		RotateXYZ(it),
	}
}

func FitsInContainer(orient models.Orientation, c *models.Container) bool {
	return orient.Length <= c.Length &&
		orient.Width <= c.Width &&
		orient.Height <= c.Height
}

func FitsInShelf(orient models.Orientation, shelf *models.Shelf, c *models.Container) bool {
	if shelf.CurrentX+orient.Length > c.Length {
		return false
	}

	if orient.Width > c.Width {
		return false
	}

	return true
}

func CheckBounds(x, y, z int, orient models.Orientation, c *models.Container) error {
	if x+orient.Length > c.Length {
		return errors.New("item exceeds container length")
	}
	if y+orient.Width > c.Width {
		return errors.New("item exceeds container width")
	}
	if z+orient.Height > c.Length {
		return errors.New("item exceeds container height")
	}
	return nil
}
