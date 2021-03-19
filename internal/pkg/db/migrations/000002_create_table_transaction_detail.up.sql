CREATE TYPE transaction_status AS ENUM ('INCOME', 'OUTCOME', 'WAITING');
CREATE TABLE IF NOT EXISTS transaction_detail (
	id BIGSERIAL PRIMARY KEY,
	transaction_detail_id VARCHAR(100),
    needs NUMERIC(19,4),
    wants NUMERIC(19,4),
    invest NUMERIC(19,4),
	description TEXT,
	status transaction_status
);
