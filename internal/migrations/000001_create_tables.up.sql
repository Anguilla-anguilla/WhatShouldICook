CREATE TABLE cuisine ( 
  id bigserial PRIMARY KEY,
  name varchar NOT NULL,
  description text
);

CREATE TABLE category ( 
  id bigserial PRIMARY KEY,
  name varchar NOT NULL
);

CREATE TABLE ingredient ( 
	id bigserial PRIMARY KEY,
	name varchar NOT NULL
);

CREATE TABLE recipe ( 
	id bigserial PRIMARY KEY,
	name varchar NOT NULL,
	description text NOT NULL,
	cooking_time int,
	price decimal(10, 2),
	store_in_freezer bool DEFAULT FALSE,
	expires_after int,
	category_id bigint REFERENCES category(id) ON DELETE SET NULL,
	cuisine_id bigint REFERENCES cuisine(id) ON DELETE SET NULL
);

CREATE TABLE recipe_ingredient (
	recipe_id bigint NOT NULL REFERENCES recipe(id) ON DELETE CASCADE,
	ingredient_id bigint NOT NULL REFERENCES ingredient(id) ON DELETE CASCADE,
	quantity varchar(50) NOT NULL ,
	PRIMARY KEY (recipe_id, ingredient_id)
);

CREATE TABLE shopping_list ( 
	id bigserial PRIMARY KEY,
	created_at timestamp DEFAULT NOW()
);

CREATE TABLE shopping_list_recipe (
	shopping_list_id bigint NOT NULL REFERENCES shopping_list(id) ON DELETE CASCADE, 
	recipe_id bigint NOT NULL REFERENCES recipe(id) ON DELETE CASCADE,
	PRIMARY KEY (shopping_list_id, recipe_id)
);

CREATE TABLE ration (
	id bigserial PRIMARY KEY,
	user_id bigint NOT NULL,
	created_at timestamp DEFAULT NOW()
);

CREATE TABLE ration_recipe (
	ration_id bigint NOT NULL REFERENCES ON ration(id) ON DELETE CASCADE,
	recipe_id bigint NOT NULL REFERENCES ON recipe(id) ON DELETE CASCADE,
	duration int,
	PRIMARY KEY (ration_id, recipe_id)
);

CREATE TABLE app_user (
	id bigserial PRIMARY KEY,
	username varchar NOT NULL UNIQUE,
	email varchar NOT NULL UNIQUE
);