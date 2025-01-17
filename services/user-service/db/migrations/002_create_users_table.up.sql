CREATE TABLE IF NOT EXISTS auth.users
(
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            VARCHAR(100)        NOT NULL,
    email           VARCHAR(100) UNIQUE NOT NULL,
    password        VARCHAR(255)        NOT NULL,
    phone           VARCHAR(15),
    is_active       BOOLEAN          DEFAULT TRUE,
    is_verified     BOOLEAN          DEFAULT FALSE,
    profile_picture VARCHAR(255),
    role            VARCHAR(50)      DEFAULT 'User',
    created_at      TIMESTAMP        DEFAULT now(),
    updated_at      TIMESTAMP        DEFAULT now(),
    deleted_at      TIMESTAMP,
    last_login      TIMESTAMP,
    date_of_birth   TIMESTAMP,
    address         TEXT,
    tenant_id       UUID DEFAULT gen_random_uuid(),
    locale          VARCHAR(10)      DEFAULT 'en-US',
    timezone        VARCHAR(50)      DEFAULT 'UTC',
    mfa_enabled     BOOLEAN          DEFAULT FALSE,
    mfa_secret      VARCHAR(255)
);
