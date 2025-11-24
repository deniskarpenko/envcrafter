-- +goose Up
CREATE TABLE IF NOT EXISTS images (
                       image_id INTEGER PRIMARY KEY AUTOINCREMENT,
                       image varchar(60),
                       docker_image varchar(60),
                       is_dockerfile TINYINT(1) NOT NULL,
                       image_type VARCHAR(10) NOT NULL
);

-- +goose Down
DROP TABLE images;