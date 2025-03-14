CREATE TABLE snippets (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100),
    content TEXT,
    created DATETIME,
    expires DATETIME
);

CREATE USER 'web'@'%';
GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'%';
ALTER USER 'web'@'%' IDENTIFIED BY 'web';

FLUSH PRIVILEGES;