CREATE INDEX recipes_translations_key ON recipes USING GIN (translations);
CREATE INDEX recipes_tags_key ON recipes USING GIN (tags);
