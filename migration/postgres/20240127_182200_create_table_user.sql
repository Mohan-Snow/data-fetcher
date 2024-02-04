CREATE TABLE IF NOT EXISTS currency (
    id BIGSERIAL PRIMARY KEY,
    coin_name VARCHAR(255) NOT NULL,
    price_usd VARCHAR(255) NOT NULL,
    last_updated TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TRIGGER update_users_updated_at_trigger
    BEFORE UPDATE ON test_data
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_field();