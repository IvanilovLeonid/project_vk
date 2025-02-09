\connect dbname;

CREATE TABLE IF NOT EXISTS containers (
    id SERIAL PRIMARY KEY,
    ip_address VARCHAR(255) UNIQUE NOT NULL,
    last_ping_time TIMESTAMP NOT NULL,
    last_successful_ping TIMESTAMP
);

