-- +goose Up
INSERT INTO images(image, is_dockerfile, docker_image, image_type)
VALUES
    ('MongoDB', 0, 'mongo', 'nosql'),
    ('Redis', 0, 'redis', 'nosql');

INSERT INTO tags(image_id, tag_name)
SELECT image_id, tag_name
FROM images
         JOIN (
    SELECT 'mongo' AS docker_image, 'latest' as tag_name
    UNION ALL SELECT 'mongo', '7.0'
    UNION ALL SELECT 'mongo', '6.0'
    UNION ALL SELECT 'mongo', '7.0-jammy'
    UNION ALL SELECT 'mongo', '6.0-jammy'

    UNION ALL SELECT 'redis', 'latest'
    UNION ALL SELECT 'redis', '7.2'
    UNION ALL SELECT 'redis', '7.0'
    UNION ALL SELECT 'redis', '7.2-alpine'
    UNION ALL SELECT 'redis', '6.2'

) AS t ON images.docker_image = t.docker_image;