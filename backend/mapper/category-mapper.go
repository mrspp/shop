package mapper

import (
	"shopee-crawler/dto"
	"shopee-crawler/entity"
)

// CategoryDTOToEntity ...
func CategoryDTOToEntity(categoryDTO dto.CategoryDTO) entity.Category {
	return entity.Category{
		ID:   categoryDTO.CategoryID,
		Name: categoryDTO.Name,
	}
}

// CategoryDTOsToEntities ...
func CategoryDTOsToEntities(categoryDTOs []dto.CategoryDTO) []entity.Category {
	categoryEntities := make([]entity.Category, 0)
	for _, c := range categoryDTOs {
		categoryEntities = append(categoryEntities, CategoryDTOToEntity(c))
	}
	return categoryEntities
}
