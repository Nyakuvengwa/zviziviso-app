CREATE TABLE countries (
    id SERIAL PRIMARY KEY,
    iso_code3 CHAR(3) UNIQUE NOT NULL,
    country_name VARCHAR(255) NOT NULL,
    dialing_code VARCHAR(10)
);

CREATE TABLE provinces (
    id SERIAL PRIMARY KEY,
    country_id INTEGER REFERENCES countries(id) NOT NULL,
    province_name VARCHAR(255) NOT NULL,
    code VARCHAR(50) UNIQUE
);