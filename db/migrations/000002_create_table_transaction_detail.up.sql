CREATE TYPE transaction_status AS ENUM ('income', 'outcome', 'waiting');
CREATE TABLE IF NOT EXISTS transaction_detail (
	id BIGSERIAL PRIMARY KEY,
	transaction_detail_id VARCHAR(100),
	amount NUMERIC(19,4),
	description TEXT,
	status transaction_status
);
