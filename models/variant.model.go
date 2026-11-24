package models

type Filament struct {
	UUID         string `json:"uuid"`
	ID           string `json:"id"`
	Name         string `json:"name"`
	ColorHex     string `json:"color_hex"`
	FilamentID   string `json:"filament_id"`
	Slug         string `json:"slug"`
	Discontinued bool   `json:"discontinued"`
	Sizes        []Size `json:"sizes"`
}

type Size struct {
	UUID              string         `json:"uuid"`
	FilamentWeight    int            `json:"filament_weight"`
	Diameter          float64        `json:"diameter"`
	Refill            bool           `json:"spool_refill"`
	ContainerMaterial string         `json:"container_material"`
	GTIN              string         `json:"gtin"`
	ArticleNumber     string         `json:"article_number"`
	ID                string         `json:"id"`
	VariantID         string         `json:"variant_id"`
	Discontinued      bool           `json:"discontinued"`
	PurchaseLinks     []PurchaseLink `json:"purchase_links"`
}

type PurchaseLink struct {
	StoreID     string `json:"store_id"`
	URL         string `json:"url"`
	ID          string `json:"id"`
	SizeID      string `json:"size_id"`
	SpoolRefill bool   `json:"spool_refill"`
}
