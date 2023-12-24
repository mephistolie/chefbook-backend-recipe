CREATE TYPE visibility as ENUM ('private', 'link', 'public');

CREATE TABLE recipes
(
    recipe_id          UUID PRIMARY KEY                                   NOT NULL UNIQUE DEFAULT gen_random_uuid(),
    name               VARCHAR(150) CHECK ( name <> '' )                  NOT NULL,

    owner_id           UUID                                               NOT NULL,

    visibility         visibility                                         NOT NULL        DEFAULT 'private',
    encrypted          BOOLEAN                                            NOT NULL        DEFAULT false,
    CHECK ( encrypted = false OR visibility <> 'public' ),

    language           VARCHAR(2)                                         NOT NULL        DEFAULT 'en',
    translations       TEXT[] NOT NULL DEFAULT '{}',
    description        VARCHAR(1500)                                                      DEFAULT NULL,

    rating             DECIMAL(7, 6) CHECK ( rating >= 0 AND rating <= 5) NOT NULL        DEFAULT 0.0,
    votes              INT CHECK ( votes >= 0 )                           NOT NULL        DEFAULT 0,

    tags               TEXT[] NOT NULL DEFAULT '{}',

    ingredients        JSONB                                              NOT NULL,
    cooking            JSONB                                              NOT NULL,
    pictures           JSONB                                              NOT NULL        DEFAULT '{}'::jsonb,

    servings           SMALLINT CHECK ( servings >= 0 )                                   DEFAULT NULL,
    cooking_time       SMALLINT CHECK ( cooking_time >= 0 )                               DEFAULT NULL,

    calories           SMALLINT CHECK ( calories >= 0 )                                   DEFAULT NULL,
    protein            SMALLINT CHECK ( protein >= 0 )                                    DEFAULT NULL,
    fats               SMALLINT CHECK ( fats >= 0 )                                       DEFAULT NULL,
    carbohydrates      SMALLINT CHECK ( carbohydrates >= 0 )                              DEFAULT NULL,

    creation_timestamp TIMESTAMP WITH TIME ZONE                           NOT NULL        DEFAULT now():: timestamp,
    update_timestamp   TIMESTAMP WITH TIME ZONE                           NOT NULL        DEFAULT now():: timestamp,
    version            INT CHECK ( recipes.version >= 0 )                                 DEFAULT 1
);

CREATE INDEX recipes_owner_id_key ON recipes (owner_id);
CREATE INDEX recipes_language_key ON recipes (language);
CREATE INDEX recipes_creation_timestamp_key ON recipes (creation_timestamp);
CREATE INDEX recipes_update_timestamp_key ON recipes (update_timestamp);
CREATE INDEX recipes_rating_key ON recipes (rating);
CREATE INDEX recipes_votes_key ON recipes (votes);

CREATE TABLE translations
(
    recipe_id   UUID REFERENCES recipes (recipe_id) ON DELETE CASCADE NOT NULL,
    language    VARCHAR(2)                                            NOT NULL DEFAULT 'en',
    author_id   UUID                                                  NOT NULL,
    name        VARCHAR(150) CHECK ( name <> '' )                     NOT NULL,
    description VARCHAR(1500)                                                  DEFAULT NULL,
    ingredients JSONB                                                 NOT NULL,
    cooking     JSONB                                                 NOT NULL,
    hidden      BOOLEAN                                               NOT NULL DEFAULT false,
    UNIQUE (recipe_id, author_id)
);

CREATE TABLE recipe_pictures_uploads
(
    recipe_id UUID REFERENCES recipes (recipe_id) ON DELETE CASCADE NOT NULL UNIQUE,
    pictures  JSONB                                                 NOT NULL
);

CREATE TABLE recipes_users
(
    recipe_id  UUID REFERENCES recipes (recipe_id) ON DELETE CASCADE NOT NULL,
    user_id    UUID                                                  NOT NULL,
    favourite  BOOLEAN                                               NOT NULL DEFAULT FALSE,
    categories JSONB                                                 NOT NULL DEFAULT '[]'::jsonb,
    UNIQUE (user_id, recipe_id)
);

CREATE TABLE scores
(
    recipe_id UUID REFERENCES recipes (recipe_id) ON DELETE CASCADE NOT NULL,
    user_id   UUID                                                  NOT NULL,
    score     SMALLINT                                              NOT NULL,
    UNIQUE (user_id, recipe_id)
);

CREATE TABLE inbox
(
    message_id UUID PRIMARY KEY         NOT NULL UNIQUE,
    timestamp  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now():: timestamp
);

CREATE TABLE outbox
(
    message_id UUID PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid(),
    exchange   VARCHAR(64)                      DEFAULT '',
    type       VARCHAR(64)      NOT NULL,
    body       JSONB            NOT NULL        DEFAULT '{}'::jsonb
);
