CREATE TABLE cart_items (
                            id SERIAL PRIMARY KEY,
                            user_id INT REFERENCES users(id) ON DELETE CASCADE,
                            dessert_id INT REFERENCES desserts(id),
                            quantity INT NOT NULL,
                            created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                            deleted_at    TIMESTAMP
);