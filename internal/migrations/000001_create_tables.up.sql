CREATE TABLE app_user (
	id bigserial PRIMARY KEY,
	username varchar NOT NULL UNIQUE,
	email varchar NOT NULL UNIQUE,
	password_hash TEXT NOT NULL,
	created_at timestamp DEFAULT NOW()
);

CREATE TABLE cuisine ( 
  id bigserial PRIMARY KEY,
  user_id bigint NOT NULL REFERENCES app_user(id) ON DELETE CASCADE,
  name varchar NOT NULL,
  description text,
  created_at timestamp DEFAULT NOW()
);

CREATE TABLE category ( 
  id bigserial PRIMARY KEY,
  name varchar UNIQUE NOT NULL
);

CREATE TABLE ingredient ( 
	id bigserial PRIMARY KEY,
	name varchar UNIQUE NOT NULL
);

CREATE TABLE recipe ( 
	id bigserial PRIMARY KEY,
	name varchar NOT NULL,
	user_id bigint NOT NULL REFERENCES app_user(id) ON DELETE CASCADE,
	description text,
	cooking_time int,
	price decimal(10, 2),
	expires_after int,
	store_in_freezer bool DEFAULT FALSE,
	favorite bool DEFAULT FALSE,
	fridgeless_store DEFAULT FALSE,
	is_public bool DEFAULT FALSE,	
	category_id bigint REFERENCES category(id) ON DELETE SET NULL,
	cuisine_id bigint REFERENCES cuisine(id) ON DELETE SET NULL,
	created_at timestamp DEFAULT NOW()
);

CREATE TABLE recipe_ingredient (
	recipe_id bigint NOT NULL REFERENCES recipe(id) ON DELETE CASCADE,
	ingredient_id bigint NOT NULL REFERENCES ingredient(id) ON DELETE CASCADE,
	quantity varchar(50) NOT NULL ,
	PRIMARY KEY (recipe_id, ingredient_id)
);

CREATE TABLE shopping_list ( 
    id bigserial PRIMARY KEY,
    ration_id bigint NOT NULL REFERENCES ration(id) ON DELETE CASCADE,
    created_at timestamp DEFAULT NOW()
);

CREATE TABLE shopping_list_recipe (
	shopping_list_id bigint NOT NULL REFERENCES shopping_list(id) ON DELETE CASCADE, 
	recipe_id bigint NOT NULL REFERENCES recipe(id) ON DELETE CASCADE,
	PRIMARY KEY (shopping_list_id, recipe_id)
);

CREATE TABLE ration (
	id bigserial PRIMARY KEY,
	user_id bigint NOT NULL REFERENCES app_user(id) ON DELETE CASCADE,
	duration int DEFAULT 7,
	created_at timestamp DEFAULT NOW()
);

CREATE TABLE ration_recipe (
	ration_id bigint NOT NULL REFERENCES ration(id) ON DELETE CASCADE,
	recipe_id bigint NOT NULL REFERENCES recipe(id) ON DELETE CASCADE,
	PRIMARY KEY (ration_id, recipe_id)
);

