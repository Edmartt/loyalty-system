
-- +migrate Up

ALTER TABLE commerce
ALTER COLUMN points_x_buy SET DEFAULT 0,
ALTER COLUMN value_x_point SET DEFAULT 0;


-- +migrate Down

ALTER TABLE commerce
ALTER COLUMN points_x_buy DROP DEFAULT,
ALTER COLUMN value_x_point DROP DEFAULT;
