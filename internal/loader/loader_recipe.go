package loader

import (
	"fmt"
	"github.com/gofrs/uuid/v5"
	"github.com/hashicorp/go-memdb"
	"github.com/nitwhiz/svapi/internal/data"
	"github.com/nitwhiz/svapi/internal/storage"
	"github.com/nitwhiz/svapi/pkg/model"
	"strings"
)

func newRecipeIngredientGroup(id string) *model.RecipeIngredientGroup {
	return &model.RecipeIngredientGroup{
		ID:          id,
		Items:       []*model.Item{},
		Ingredients: []*model.RecipeIngredient{},
	}
}

func loadRecipes(txn *memdb.Txn) error {
	recipes, err := data.Parse[Recipes]("recipes.json")

	if err != nil {
		return err
	}

	categoryIdToGroup := map[string]*model.RecipeIngredientGroup{}

	for _, recipe := range recipes.Recipes {
		recipeModel := &model.Recipe{
			ID:          uuid.NewV5(namespaceId, recipe.Name).String(),
			Ingredients: []*model.RecipeIngredient{},
			Name:        recipe.Name,
			IsCooking:   recipe.IsCooking,
			OutputItems: []*model.Item{},
			OutputYield: recipe.Output.Amount,
		}

		for idx, ingredient := range recipe.Ingredients {
			if ingredient.ItemID == "" {
				continue
			}

			var ingredientGroupModel *model.RecipeIngredientGroup

			if strings.HasPrefix(ingredient.ItemID, "-") {
				// is a category group

				cat, err := first[model.Category](txn, "internalId", ingredient.ItemID)

				if err != nil {
					return err
				}

				if cat == nil {
					continue
				}

				var ok bool

				ingredientGroupModel, ok = categoryIdToGroup[cat.ID]

				if !ok {
					ingredientGroupModel = newRecipeIngredientGroup(uuid.NewV5(namespaceId, cat.ID).String())

					items, err := storage.SearchAll(txn, model.Item{Category: &model.Category{ID: cat.ID}})

					if err != nil {
						return err
					}

					for _, i := range items {
						item := i.(*model.Item)

						ingredientGroupModel.Items = append(ingredientGroupModel.Items, item)
						item.IngredientGroups = append(item.IngredientGroups, ingredientGroupModel)
					}

					if err := txn.Insert(ingredientGroupModel.TableName(), ingredientGroupModel); err != nil {
						return err
					}

					categoryIdToGroup[cat.ID] = ingredientGroupModel
				}
			} else {
				recipeIngredientGroupId := uuid.NewV5(
					namespaceId,
					fmt.Sprintf(
						"%d_%s",
						idx,
						recipe.Name,
					),
				).String()

				ingredientGroupModel = newRecipeIngredientGroup(recipeIngredientGroupId)

				if err := storage.Insert(txn, ingredientGroupModel); err != nil {
					return err
				}

				ingredientItem, err := first[model.Item](txn, "internalId", ingredient.ItemID)

				if err != nil {
					return err
				}

				if ingredientItem == nil {
					continue
				}

				ingredientGroupModel.Items = append(ingredientGroupModel.Items, ingredientItem)
				ingredientItem.IngredientGroups = append(ingredientItem.IngredientGroups, ingredientGroupModel)
			}

			recipeIngredientModel := &model.RecipeIngredient{
				ID:              uuid.NewV5(namespaceId, fmt.Sprintf("%s_%s", recipeModel.ID, ingredient.ItemID)).String(),
				Recipe:          recipeModel,
				IngredientGroup: ingredientGroupModel,
				Quantity:        ingredient.Quantity,
			}

			ingredientGroupModel.Ingredients = append(ingredientGroupModel.Ingredients, recipeIngredientModel)

			if err := storage.Insert(txn, recipeIngredientModel); err != nil {
				return err
			}

			recipeModel.Ingredients = append(recipeModel.Ingredients, recipeIngredientModel)
		}

		for _, outputItemId := range recipe.Output.ItemIDs {
			if outputItemId == "" {
				continue
			}

			outputItem, err := first[model.Item](txn, "internalId", outputItemId)

			if err != nil {
				return err
			}

			if outputItem == nil {
				continue
			}

			recipeModel.OutputItems = append(recipeModel.OutputItems, outputItem)
		}

		if err := storage.Insert(txn, recipeModel); err != nil {
			return err
		}
	}

	return nil
}
