CREATE TABLE IF NOT EXISTS facilities (
    code        VARCHAR(10) PRIMARY KEY,
    name        TEXT        NOT NULL,
    address     TEXT        NOT NULL,
    created_at  TIMESTAMP   DEFAULT NOW()
);