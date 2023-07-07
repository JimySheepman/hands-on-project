SELECT 'CREATE DATABASE er'
WHERE NOT EXISTS(SELECT FROM pg_database WHERE datname = 'er')
\gexec

DROP TABLE IF EXISTS currency CASCADE;
CREATE TABLE currency (
	id serial PRIMARY KEY,
	base_code VARCHAR ( 50 ) NOT NULL,
	target_code VARCHAR ( 50 ) NOT NULL,
	conversion_rate REAL NOT NULL,
    created_at BIGINT NOT NULL
);