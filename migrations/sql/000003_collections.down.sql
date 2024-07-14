ALTER TABLE recipes_users
    ADD COLUMN categories JSONB NOT NULL DEFAULT '[]'::jsonb;
