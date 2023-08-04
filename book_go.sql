CREATE TABLE authors(
    id SERIAL PRIMARY KEY,
    full_name TEXT NOT NULL
);

CREATE TABLE books(
    id TEXT PRIMARY KEY,
    ISBF TEXT NOT NULL,
    title TEXT NOT NULL,
    page INT NOT NULL,
    author_id INT,
    FOREIGN KEY(author_id) REFERENCES authors(id)
);