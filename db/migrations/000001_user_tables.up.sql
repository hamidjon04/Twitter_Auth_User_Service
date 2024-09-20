CREATE TYPE role AS ENUM('admin', 'user');

CREATE TABLE IF NOT EXISTS users(
	id uuid PRIMARY KEY,  
	username varchar UNIQUE NOT NULL,
	email varchar UNIQUE NOT NULL,
	password varchar NOT NULL,
	name varchar DEFAULT 'name',
	lastname varchar DEFAULT 'lastname',
	birth_day varchar DEFAULT '2000-01-01',
	image varchar DEFAULT 'avatar',
    role role DEFAULT 'user',
	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamp DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS user_followers(
	id uuid PRIMARY KEY,
	user_id uuid NOT NULL REFERENCES users(id),
	follower_id uuid NOT NULL REFERENCES users(id),
	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamp DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS user_following(
	id uuid PRIMARY KEY,
	user_id uuid NOT NULL REFERENCES users(id),
	following_id uuid NOT NULL REFERENCES users(id),
	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamp DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS refresh_tokens(
	user_id uuid PRIMARY KEY REFERENCES users(id),
	token varchar UNIQUE NOT NULL,
	expires_at int NOT NULL
);