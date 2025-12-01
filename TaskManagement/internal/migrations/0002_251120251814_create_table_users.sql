-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    userid          SERIAL PRIMARY KEY,
    firstname       VARCHAR(100) NOT NULL,
    lastname        VARCHAR(100) NOT NULL,
    passport_number VARCHAR(20) NOT NULL
);

-- Добавление двух записей в таблицу users
INSERT INTO users (userid, firstname, lastname, passport_number)
VALUES
    (1, 'Иван', 'Иванов', '1234567890'),
    (2, 'Петр', 'Петров', '0987654321');

-- +goose StatementEnd    

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd