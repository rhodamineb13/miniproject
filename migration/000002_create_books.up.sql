CREATE TABLE IF NOT EXISTS books(
    id BIGSERIAL PRIMARY KEY,
    isbn VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    quantity BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);
