-- +goose Up
-- +goose StatementBegin
CREATE TABLE pharmacies (
    user_id BIGINT NOT NULL,
    pharmacy_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    PRIMARY KEY(user_id, pharmacy_id)
);

CREATE TABLE products (
    user_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    PRIMARY KEY(user_id, product_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pharmacies, products;
-- +goose StatementEnd