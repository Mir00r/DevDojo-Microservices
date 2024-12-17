-- Dec 14, 2024
CREATE TABLE IF NOT EXISTS auth.password_reset_token
(
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    token      TEXT                           NOT NULL,
    user_id    UUID                           NOT NULL REFERENCES auth.users (id) ON DELETE CASCADE,
    used       BOOLEAN          DEFAULT FALSE,
    created_at TIMESTAMP        DEFAULT now(),
    updated_at TIMESTAMP        DEFAULT now() NOT NULL,
    expires_at TIMESTAMP                      NOT NULL,
    deleted_at TIMESTAMP                      NULL
);
