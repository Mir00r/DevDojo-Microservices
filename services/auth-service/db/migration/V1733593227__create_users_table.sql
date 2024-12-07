-- Dec 08, 2024

CREATE TABLE IF NOT EXISTS auth.users (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- Unique user ID
                       name VARCHAR(100) NOT NULL,                   -- User's full name
                       email VARCHAR(100) UNIQUE NOT NULL,           -- User's email address (unique)
                       password VARCHAR(255) NOT NULL,               -- Hashed password
                       created_at TIMESTAMP DEFAULT now() NOT NULL,  -- Timestamp when the user was created
                       updated_at TIMESTAMP DEFAULT now() NOT NULL,   -- Timestamp when the user was last updated
                       deleted_at TIMESTAMP NULL   -- Timestamp when the user was last deleted
);
