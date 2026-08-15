CREATE INDEX idx_recipe_user_id ON recipe(user_id);
CREATE INDEX idx_recipe_category_id ON recipe(category_id);
CREATE INDEX idx_recipe_cuisine_id ON recipe(cuisine_id);
CREATE INDEX idx_ration_user_id ON ration(user_id);
CREATE INDEX idx_shopping_list_ration_id ON shopping_list(ration_id);
CREATE INDEX idx_recipe_ingredient_recipe_id ON recipe_ingredient(recipe_id);
CREATE INDEX idx_recipe_ingredient_ingredient_id ON recipe_ingredient(ingredient_id);

CREATE INDEX idx_ingredient_name ON ingredient(name);
CREATE INDEX idx_category_name ON category(name);
CREATE INDEX idx_recipe_name ON recipe(name);