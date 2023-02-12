# Create users table

CREATE TABLE IF NOT EXISTS users
(
    id INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL UNIQUE,
    passHash VARCHAR(2048) NOT NULL,
    verified BOOL NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_modified TIMESTAMP,
    archived_at TIMESTAMP,
    banned_at TIMESTAMP,
    PRIMARY KEY(id)
);

COMMIT;
