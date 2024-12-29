-- Dec 08, 2024

ALTER TABLE auth.tokens
    ADD COLUMN refresh_token TEXT NOT NULL UNIQUE DEFAULT '';

ALTER TABLE auth.tokens
    ADD COLUMN refresh_token_expires_at TIMESTAMP NOT NULL DEFAULT now();

ALTER TABLE auth.tokens
    ADD COLUMN updated_at TIMESTAMP DEFAULT now() NOT NULL;
