CREATE TABLE IF NOT EXISTS orders (
    id            VARCHAR(15)      PRIMARY KEY,
    facility_code VARCHAR(10) NOT NULL,
    status        TEXT        NOT NULL,
    created_at    TIMESTAMP   DEFAULT NOW(),
    CONSTRAINT fk_facility
        FOREIGN KEY (facility_code)
        REFERENCES facilities(code)
        ON DELETE RESTRICT
);