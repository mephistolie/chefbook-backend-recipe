# ChefBook Backend Recipe Service

The recipe service owns recipes, recipe book state, favourites, ratings, translations, recipe pictures, collections, collection sharing, recipe policy reads, and recipe-related MQ events.

## Responsibilities

- Recipe CRUD and recipe read models.
- Recipe book and favourite state.
- Collection CRUD, collection membership, and collection invite keys.
- Recipe-to-collection binding.
- Recipe rating and vote aggregation.
- Recipe translations.
- Recipe picture upload coordination through object storage.
- Recipe policy checks for other services.
- Firebase recipe import through MQ when configured.

## Main RPC Families

- `GetRecipes`, `GetRandomRecipe`, `GetRecipeBook`
- `CreateRecipe`, `GetRecipe`, `UpdateRecipe`, `DeleteRecipe`
- `GenerateRecipePicturesUploadLinks`, `SetRecipePictures`
- `RateRecipe`
- `SaveRecipeToRecipeBook`, `RemoveRecipeFromRecipeBook`
- `SaveRecipeToFavourites`, `RemoveRecipeFromFavourites`
- `AddRecipeToCollection`, `RemoveRecipeFromCollection`, `SetRecipeCollections`
- `TranslateRecipe`, `DeleteRecipeTranslation`
- `GetRecipePolicy`, `GetRecipeNames`
- `GetCollections`, `CreateCollection`, `GetCollection`, `UpdateCollection`, `DeleteCollection`
- `SaveCollectionToRecipeBook`, `RemoveCollectionFromRecipeBook`

## Dependencies

- Calls `profile` for user/profile read data.
- Calls `tag` for tag lookup and tag maps.
- Calls `encryption` for encrypted recipe key coordination.
- Uses S3-compatible object storage for recipe pictures.
- Publishes and consumes MQ messages when configured.
- Owns its PostgreSQL schema and migrations.

## Database Ownership

Owns:

- `recipes` - recipe content, visibility, encryption flag, language, pictures, rating, tags, version.
- `translations` - per-recipe translation content.
- `recipe_pictures_uploads` - pending picture upload state.
- `recipe_book` - saved recipe state per user.
- `favourites` - favourite recipe state per user.
- `scores` - rating state per user.
- `collections` - collection metadata.
- `collections_keys` - collection invite keys.
- `collections_contributors` - collection edit membership.
- `collections_users` - saved collection state per user.
- `recipes_collections` - recipe-to-collection binding.
- `inbox` and `outbox` - MQ idempotency and outgoing events.

```mermaid
erDiagram
    RECIPES {
        uuid recipe_id PK
        varchar name
        uuid owner_id
        visibility visibility
        boolean encrypted
        varchar language
        text_array translations
        jsonb ingredients
        jsonb cooking
        jsonb pictures
        decimal rating
        int votes
        text_array tags
        int version
    }

    TRANSLATIONS {
        uuid recipe_id FK
        varchar language
        uuid author_id
        varchar name
        jsonb ingredients
        jsonb cooking
    }

    RECIPE_PICTURES_UPLOADS {
        uuid recipe_id FK,UK
        jsonb pictures
    }

    RECIPE_BOOK {
        uuid recipe_id FK
        uuid user_id
    }

    FAVOURITES {
        uuid recipe_id FK
        uuid user_id
    }

    SCORES {
        uuid recipe_id FK
        uuid user_id
        smallint score
    }

    COLLECTIONS {
        uuid collection_id PK
        varchar name
        visibility visibility
    }

    COLLECTIONS_KEYS {
        uuid collection_id FK,UK
        uuid key
        timestamptz expires_at
    }

    COLLECTIONS_CONTRIBUTORS {
        uuid collection_id FK
        uuid contributor_id
        contributor_role role
    }

    COLLECTIONS_USERS {
        uuid collection_id FK
        uuid user_id
    }

    RECIPES_COLLECTIONS {
        uuid recipe_id FK
        uuid collection_id FK
        timestamptz binding_timestamp
    }

    RECIPE_INBOX {
        uuid message_id PK
        timestamptz timestamp
    }

    RECIPE_OUTBOX {
        uuid message_id PK
        varchar exchange
        varchar type
        jsonb body
    }

    RECIPES ||--o{ TRANSLATIONS : has
    RECIPES ||--o| RECIPE_PICTURES_UPLOADS : has_upload_state
    RECIPES ||--o{ RECIPE_BOOK : saved_by
    RECIPES ||--o{ FAVOURITES : favourited_by
    RECIPES ||--o{ SCORES : rated_by
    RECIPES ||--o{ RECIPES_COLLECTIONS : assigned_to
    COLLECTIONS ||--o| COLLECTIONS_KEYS : has_invite_key
    COLLECTIONS ||--o{ COLLECTIONS_CONTRIBUTORS : editable_by
    COLLECTIONS ||--o{ COLLECTIONS_USERS : saved_by
    COLLECTIONS ||--o{ RECIPES_COLLECTIONS : contains
```

Important constraints:

- Encrypted recipes cannot be public.
- Translations are unique by recipe, language, and author.
- Pair tables are unique by their pair keys.
- User IDs are logical references to auth/user/profile domains, not database foreign keys.
