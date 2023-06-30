package dto

import (
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-recipe/internal/entity"
)

const maxMacronutrients = 500

func newMacronutrients(req *api.Macronutrients) *entity.Macronutrients {
	var macronutrientsPtr *entity.Macronutrients
	if req != nil {
		macronutrients := entity.Macronutrients{}
		if req.Protein != nil && *req.Protein > 0 {
			macronutrients.Protein = req.Protein
			if *macronutrients.Protein > maxMacronutrients {
				*macronutrients.Protein = maxMacronutrients
			}
		}
		if req.Fats != nil && *req.Fats > 0 {
			macronutrients.Fats = req.Fats
			if *macronutrients.Fats > maxMacronutrients {
				*macronutrients.Fats = maxMacronutrients
			}
		}
		if req.Carbohydrates != nil && *req.Carbohydrates > 0 {
			macronutrients.Carbohydrates = req.Carbohydrates
			if *macronutrients.Carbohydrates > maxMacronutrients {
				*macronutrients.Carbohydrates = maxMacronutrients
			}
		}
		if macronutrients.Protein != nil || macronutrients.Fats != nil || macronutrients.Carbohydrates != nil {
			macronutrientsPtr = &macronutrients
		}
	}
	return macronutrientsPtr
}
