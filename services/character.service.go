package services

import (
	"pokemon/api"
	"pokemon/models"
	"strconv"
	"strings"
)

// SearchCharacters renvoie les personnages de la première page,
// filtrés par l'API sur le nom si une recherche est faite.
func SearchCharacters(name string) ([]models.Character, error) {
	page, err := api.GetCharacters()
	if err != nil {
		return nil, err
	}
	needle := strings.ToLower(strings.TrimSpace(name))
	characters := make([]models.Character, 0, len(page.Results))
	for _, character := range page.Results {
		if character.ID <= 0 {
			continue
		}
		if needle != "" && !strings.Contains(strings.ToLower(character.Name.French), needle) {
			continue
		}
		characters = append(characters, character)
	}
	return characters, nil
}

// GetCharacter vérifie l'identifiant reçu dans l'URL, puis appelle l'API.
func GetCharacter(idParam string) (models.Character, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		return models.Character{}, models.ErrBadRequest
	}
	return api.GetCharacter(id)
}
