-- Dec 08, 2024

CREATE TABLE IF NOT EXISTS auth.tokens (
                        id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- Unique token ID
                        user_id UUID NOT NULL REFERENCES auth.users(id) ON DELETE CASCADE, -- Foreign key to the users table
                        token TEXT NOT NULL,                            -- The token string
                        type VARCHAR(50) NOT NULL,                     -- Type of token (e.g., 'access', 'refresh')
                        expires_at TIMESTAMP NOT NULL,                 -- Expiration timestamp of the token
                        created_at TIMESTAMP DEFAULT now() NOT NULL,    -- Timestamp when the token was created
                        deleted_at TIMESTAMP NULL   -- Timestamp when the token was last deleted
);
