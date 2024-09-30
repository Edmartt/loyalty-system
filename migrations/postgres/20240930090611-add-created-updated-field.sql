
-- +migrate Up
ALTER TABLE branch
ADD created_at timestamp with time zone DEFAULT now() NOT NULL,
ADD updated_at timestamp with time zone DEFAULT now() NOT NULL;

ALTER TABLE campaign
ADD created_at timestamp with time zone DEFAULT now() NOT NULL,
ADD updated_at timestamp with time zone DEFAULT now() NOT NULL;

ALTER TABLE commerce
ADD created_at timestamp with time zone DEFAULT now() NOT NULL,
ADD updated_at timestamp with time zone DEFAULT now() NOT NULL;

ALTER TABLE customer
ADD created_at timestamp with time zone DEFAULT now() NOT NULL,
ADD updated_at timestamp with time zone DEFAULT now() NOT NULL;

ALTER TABLE purchase
ADD created_at timestamp with time zone DEFAULT now() NOT NULL,
ADD updated_at timestamp with time zone DEFAULT now() NOT NULL;

ALTER TABLE reward
ADD created_at timestamp with time zone DEFAULT now() NOT NULL,
ADD updated_at timestamp with time zone DEFAULT now() NOT NULL;


-- +migrate Down

ALTER TABLE branch
DROP COLUMN created_at,
DROP COLUMN updated_at;

ALTER TABLE campaign
DROP COLUMN created_at,
DROP COLUMN updated_at;

ALTER TABLE commerce
DROP COLUMN created_at,
DROP COLUMN updated_at;

ALTER TABLE customer
DROP COLUMN created_at,
DROP COLUMN updated_at;

ALTER TABLE purchase
DROP COLUMN created_at,
DROP COLUMN updated_at;

ALTER TABLE reward
DROP COLUMN created_at,
DROP COLUMN updated_at;
