CREATE TABLE IF NOT EXISTS borrow_records(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT,
    book_id BIGINT,
    borrowed_at TIMESTAMP NOT NULL,
    returned_at TIMESTAMP
);