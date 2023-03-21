CREATE TABLE IF NOT EXISTS users (
    id          SERIAL          NOT NULL PRIMARY KEY,
    username    VARCHAR(20)     NOT NULL,
    fullname    VARCHAR(255)    NOT NULL,
    password    VARCHAR(255)    NOT NULL,
    created_at  INTEGER         NOT NULL
)