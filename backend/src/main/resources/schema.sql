CREATE TABLE identification_numbers (
    id uuid NOT NULL,
    number VARCHAR(20) NOT NULL,
    blocked boolean,
    PRIMARY KEY (id)
);