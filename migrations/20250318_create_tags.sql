-- +goose Up
CREATE TABLE IF NOT EXISTS tags (
    tag_id INTEGER PRIMARY KEY AUTOINCREMENT,
    image_id INTEGER NOT NULL ,
    tag_name varchar(25),
    FOREIGN KEY (image_id) REFERENCES images(image_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE tags;