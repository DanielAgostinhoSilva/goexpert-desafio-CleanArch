-- +goose Up
CREATE TABLE orders
(
    id          CHAR(36) PRIMARY KEY,
    price       FLOAT,
    tax         FLOAT,
    final_price FLOAT
);

-- +goose Down
DROP TABLE orders;
