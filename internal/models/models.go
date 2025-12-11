package models

type Container struct {
	Length int    `json:"length"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Unit   string `json:"unit"`
}

type Orientation struct {
	Name   string `json:"name"` // xyz/yxz
	Length int    `json:"length"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type ItemInstance struct {
	InstanceID        string      `json:"instance_id"`
	ParentItemID      string      `json:"parent_item_id"`
	OriginalLength    int         `json:"original_length"`
	OriginalWidth     int         `json:"original_width"`
	OriginalHeight    int         `json:"original_height"`
	Quantity          int         `json:"quantity"`
	AllowRotation     bool        `json:"allow_rotation"`
	ChosenOrientation Orientation `json:"chosen_orientation"`

	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Shelf struct {
	StartZ   int
	Height   int
	CurrentX int
	Items    []*ItemInstance
}

type Placement struct {
	ItemInstanceID string `json:"item_instance_id"`
	Orientation    string `json:"orientation"`
	X              int    `json:"x"`
	Y              int    `json:"y"`
	Z              int    `json:"z"`
	Length         int    `json:"length"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
}

type PackingResult struct {
	Status        string      `json:"status"`
	Metrics       Metrics     `json:"metrics"`
	Placements    []Placement `json:"placements"`
	UnpackedItems []string    `json:"unpacked_items"`
}

type Metrics struct {
	ContainerVolume    int     `json:"container_volume"`
	PackedVolume       int     `json:"packed_volume"`
	UtilizationPercent float64 `json:"utilization_percent"`
	TotalItems         int     `json:"total_items"`
	ItemsPacked        int     `json:"items_packed"`
	ItemsUnpacked      int     `json:"items_unpacked"`
}
