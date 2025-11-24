-- +goose Up
INSERT INTO images(image, is_dockerfile, docker_image, image_type)
VALUES
    ('MySQL', 0, 'mysql', 'sql'),
    ('PostgreSQL', 0, 'postgres', 'sql'),
    ('MariaDB', 0, 'mariadb', 'sql');

INSERT INTO tags(image_id, tag_name)
SELECT image_id, tag_name
FROM images
         JOIN (
    SELECT 'mysql' AS docker_image, 'latest' as tag_name
    UNION ALL SELECT 'mysql', '8.0'
    UNION ALL SELECT 'mysql', '8.0.36'
    UNION ALL SELECT 'mysql', '5.7'
    UNION ALL SELECT 'mysql', '5.7.44'

    UNION ALL SELECT 'postgres', 'latest'
    UNION ALL SELECT 'postgres', '17'
    UNION ALL SELECT 'postgres', '17.5'
    UNION ALL SELECT 'postgres', '16'
    UNION ALL SELECT 'postgres', '16.3'

    UNION ALL SELECT 'mariadb', 'latest'
    UNION ALL SELECT 'mariadb', '10.11'
    UNION ALL SELECT 'mariadb', '10.11.6'
    UNION ALL SELECT 'mariadb', '10.6'
    UNION ALL SELECT 'mariadb', '10.6.16'
) AS t ON images.docker_image = t.docker_image;
