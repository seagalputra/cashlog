CREATE TABLE IF NOT EXISTS user_account(
	id BIGSERIAL PRIMARY KEY,
	user_account_id VARCHAR(100),
	first_name VARCHAR(100),
	last_name VARCHAR(100),
	username VARCHAR(100),
	password VARCHAR(255),
	email VARCHAR(255),
	is_disabled BOOLEAN,
	is_verified BOOLEAN
);
