
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS customer(
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	user_dni VARCHAR(30),
	name VARCHAR(50),
	last_name VARCHAR(60),
	phone VARCHAR(60),
	email VARCHAR(60),
	points_collected NUMERIC(10,2),
	cashback_collected NUMERIC(10,2)
);

CREATE INDEX idx_user_dni ON customer(user_dni);

CREATE TABLE IF NOT EXISTS commerce(
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	name VARCHAR(50),
	points_x_buy NUMERIC(10,2),
	value_x_point NUMERIC(10,2)
);

CREATE TABLE IF NOT EXISTS branch(
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	commerce_id UUID REFERENCES commerce(id),
	address VARCHAR(60)
);

CREATE TABLE IF NOT EXISTS purchase(
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	branch_id UUID REFERENCES branch(id),
	customer_id UUID REFERENCES customer(id),
	points_per_purchase NUMERIC(10,2),
	cashback_per_purchase NUMERIC(10,2),
	purchased_amount NUMERIC(10,2)
);

CREATE TABLE IF NOT EXISTS campaign(
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	commerce_id  UUID REFERENCES commerce(id),
	branch_id UUID REFERENCES branch(id),
	campaign_name VARCHAR(60),
	campaign_type VARCHAR(60),
	campaign_multiplier NUMERIC(5,2),
	campaign_percent_bonus NUMERIC(5,2),
	start_date DATE,
	end_date DATE
);

CREATE TABLE IF NOT EXISTS reward(
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	name VARCHAR(60),
	commerce_id UUID REFERENCES commerce(id),
	points_redeem NUMERIC(10,2)
);


-- +migrate Down

DROP INDEX IF EXISTS idx_user_dni;

-- sorted according to relationship
DROP TABLE IF EXISTS purchase;
DROP TABLE IF EXISTS campaign;
DROP TABLE IF EXISTS purchase;
DROP TABLE IF EXISTS reward;
DROP TABLE IF EXISTS commerce;
DROP TABLE IF EXISTS branch;

DROP TABLE IF EXISTS customer;

