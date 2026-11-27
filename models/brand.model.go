package models

type Brand struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	Origin        string `json:"origin"`
	MaterialCount int    `json:"material_count"`
	LogoSlug      string `json:"logo_slug"`
}

type BrandResponse struct {
	Brands []Brand `json:"brands"`
}
