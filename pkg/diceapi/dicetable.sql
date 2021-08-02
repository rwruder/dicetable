CREATE TABLE IF NOT EXISTS tables (
	id SERIAL NOT NULL PRIMARY KEY,
	table_name VARCHAR(255) UNIQUE NOT NULL,
	owner int NOT NULL references users(id),
	description VARCHAR(255),
	date_created timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS pools (
	id SERIAL NOT NULL PRIMARY KEY,
	table_id int NOT NULL references tables(id),
	size_of_dice int NOT NULL,
	dice int[] NOT NULL,
	description VARCHAR(255),
	date_created timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
	id SERIAL NOT NULL PRIMARY KEY,
	username VARCHAR(255) UNIQUE NOT NULL,
	password VARCHAR(255),
	date_created timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS characters (
	id SERIAL NOT NULL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	user_id int NOT NULL REFERENCES users(id),
	table_id int NOT NULL REFERENCES tables(id),
	date_created timestamp NOT NULL
);

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'change') THEN
        CREATE TYPE change AS ENUM ('rolled dice', 'set dice', 'added dice', 'removed dice');
    END IF;
END
$$;


CREATE TABLE IF NOT EXISTS rolls(
	id SERIAL NOT NULL PRIMARY KEY,
	table_id int NOT NULL references tables(id),
	user_id int NOT NULL REFERENCES users(id),
	character_id int NOT NULL REFERENCES characters(id),
	pool_id int NOT NULL references pools(id),
	type_of_change change NOT NULL,
	dice_before int[] NOT NULL,
	dice_after int[] NOT NULL,
	date timestamp NOT NULL
);


