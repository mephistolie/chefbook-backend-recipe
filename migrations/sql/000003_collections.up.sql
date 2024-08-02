CREATE TABLE collections
(
    collection_id UUID PRIMARY KEY NOT NULL,
    name          VARCHAR(64)      NOT NULL,
    visibility    visibility       NOT NULL DEFAULT 'private'
);

CREATE TABLE collections_keys
(
    collection_id uuid REFERENCES collections (collection_id) ON DELETE CASCADE NOT NULL UNIQUE,
    key           uuid                                                          NOT NULL DEFAULT gen_random_uuid(),
    expires_at    TIMESTAMP WITH TIME ZONE                                      NOT NULL
);

CREATE TYPE contributor_role as ENUM ('owner', 'coauthor');

CREATE TABLE collections_contributors
(
    collection_id  UUID REFERENCES collections (collection_id) ON DELETE CASCADE NOT NULL,
    contributor_id UUID                                                          NOT NULL,
    role           contributor_role                                              NOT NULL,
    UNIQUE (collection_id, contributor_id)
);

CREATE TABLE collections_users
(
    collection_id UUID REFERENCES collections (collection_id) ON DELETE CASCADE NOT NULL,
    user_id       UUID                                                          NOT NULL,
    UNIQUE (collection_id, user_id)
);

CREATE INDEX collections_users_user_id_key ON collections_users (user_id);

CREATE TABLE recipes_collections
(
    recipe_id         UUID REFERENCES recipes (recipe_id) ON DELETE CASCADE         NOT NULL,
    collection_id     UUID REFERENCES collections (collection_id) ON DELETE CASCADE NOT NULL,
    binding_timestamp TIMESTAMP WITH TIME ZONE                                      NOT NULL DEFAULT now()::timestamp,
    UNIQUE (recipe_id, collection_id)
);
