CREATE TABLE snippets (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100),
    content TEXT,
    created DATETIME,
    expires DATETIME
);

CREATE TABLE sessions (
    token CHAR(43) PRIMARY KEY,
    data BLOB NOT NULL,
    expiry TIMESTAMP(6) NOT NULL
);
CREATE INDEX sessions_expiry_idx ON sessions (expiry);

CREATE USER 'web'@'%';
GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'%';
ALTER USER 'web'@'%' IDENTIFIED BY 'web';

FLUSH PRIVILEGES;