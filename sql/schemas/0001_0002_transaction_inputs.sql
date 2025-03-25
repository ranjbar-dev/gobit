CREATE TABLE transaction_inputs (
  id SERIAL PRIMARY KEY,
  transaction_id VARCHAR(100) NOT NULL,
  hash VARCHAR(100) NOT NULL,
  address VARCHAR(100) NOT NULL,
  value BIGINT NOT NULL,
  index BIGINT NOT NULL
);