CREATE TABLE transaction_outputs (
  id SERIAL PRIMARY KEY,
  transaction_id VARCHAR(100) NOT NULL,
  address VARCHAR(100) NOT NULL,
  value BIGINT NOT NULL,
  index BIGINT NOT NULL
);