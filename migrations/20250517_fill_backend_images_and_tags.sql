-- +goose Up
INSERT INTO images(image, is_dockerfile, docker_image, image_type)
VALUES
    ("PHP", 1, "php","backend"),
    ("Node.js", 1, "node","backend"),
    ("Golang", 1, "golang","backend"),
    ("Python", 1, "python","backend"),
    ("Ruby", 1, "ruby","backend");

INSERT INTO tags(image_id, tag_name)
      SELECT image_id, tag_name FROM images
      JOIN (
          SELECT "PHP" AS image, "8.4-fpm" as tag_name
          UNION ALL
          SELECT "PHP" AS image, "8.3-fpm" as tag_name
          UNION ALL
          SELECT "PHP" AS image, "8.3-apache"
          UNION ALL
          SELECT "PHP" AS image, "8.3-cli"
          union ALL
          SELECT "PHP" AS image, "8.3-fpm-alpine" as tag_name
          UNION ALL
          SELECT "PHP" AS image, "8.2-fpm"
          UNION ALL
          SELECT "PHP" AS image, "8.2-apache"
          UNION ALL
          SELECT "PHP" AS image, "7.4-fpm"
          UNION ALL
          SELECT "Node.js" AS image, "20"
          UNION ALL
          SELECT "Node.js" AS image, "20-alpine"
          UNION ALL
          SELECT "Node.js" AS image, "20-slim"
          UNION ALL
          SELECT "Node.js" AS image, "lts"
          UNION ALL
          SELECT "Node.js" AS image, "18"
          UNION ALL
          SELECT "Node.js" AS image, "18-alpine"
          UNION ALL
          SELECT "Node.js" AS image, "current"
          UNION ALL
          SELECT "Node.js" AS image, "latest"
          UNION ALL
          SELECT "Node.js" AS image, "20-bullseye"
          UNION ALL
          SELECT "Node.js" AS image, "20-bookworm"
          UNION ALL
          SELECT "Golang" AS image, "1.22"
          UNION ALL
          SELECT "Golang" AS image, "1.22-alpine"
          UNION ALL
          SELECT "Golang" AS image, "1.22-bullseye"
          UNION ALL
          SELECT "Golang" AS image, "1.21"
          UNION ALL
          SELECT "Golang" AS image, "1.21-alpine"
          UNION ALL
          SELECT "Golang" AS image, "1.22-bookworm"
          UNION ALL
          SELECT "Golang" AS image, "alpine"
          UNION ALL
          SELECT "Golang" AS image, "1.21-alpine"
          UNION ALL
          SELECT "Golang" AS image, "latest"
          UNION ALL
          SELECT "Golang" AS image, "1.22.3"
          UNION ALL
          SELECT "Python" AS image, "3.12"
          UNION ALL
          SELECT "Python" AS image, "3.12-slim"
          UNION ALL
          SELECT "Python" AS image, "3.12-alpine"
          UNION ALL
          SELECT "Python" AS image, "3.11"
          UNION ALL
          SELECT "Python" AS image, "3.11-slim"
          UNION ALL
          SELECT "Python" AS image, "3.12-bullseye"
          UNION ALL
          SELECT "Python" AS image, "slim"
          UNION ALL
          SELECT "Python" AS image, "3.12-rc"
          UNION ALL
          SELECT "Python" AS image, "latest"
          UNION ALL
          SELECT "Ruby" AS image, "3.3"
          UNION ALL
          SELECT "Ruby" AS image, "3.3-slim"
          UNION ALL
          SELECT "Ruby" AS image, "3.3-alpine"
          UNION ALL
          SELECT "Ruby" AS image, "3.2"
          UNION ALL
          SELECT "Ruby" AS image, "3.2-slim"
          UNION ALL
          SELECT "Ruby" AS image, "3.3-bullseye"
          UNION ALL
          SELECT "Ruby" AS image, "alpine"
          UNION ALL
          SELECT "Ruby" AS image, "latest"
      )AS t ON images.image = t.image;

