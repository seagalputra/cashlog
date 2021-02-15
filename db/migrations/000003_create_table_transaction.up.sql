CREATE TABLE IF NOT EXISTS transaction (
	id BIGSERIAL PRIMARY KEY,
	transaction_id VARCHAR(100),
	title VARCHAR(255),
	amount NUMERIC(19,4),
	transaction_date DATE,
	transaction_detail_id BIGINT,
	user_account_id BIGINT,
	UNIQUE(transaction_detail_id)
);

ALTER TABLE "transaction" ADD CONSTRAINT fk_transaction_detail FOREIGN KEY (transaction_detail_id) REFERENCES transaction_detail(id) ON DELETE CASCADE;
ALTER TABLE "transaction" ADD CONSTRAINT fk_user_account FOREIGN KEY (user_account_id) REFERENCES user_account(id) ON DELETE CASCADE;
