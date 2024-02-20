CREATE INDEX recipes_translations_key ON recipes USING GIN (translations gin__int_ops);
CREATE INDEX recipes_tags_key ON recipes USING GIN (tags gin__int_ops);
