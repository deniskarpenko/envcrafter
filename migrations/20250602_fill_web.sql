-- +goose Up
INSERT INTO images(image, is_dockerfile, docker_image, image_type)
VALUES
    ('Nginx', 0, 'nginx', 'web'),

    ('Apache', 0, 'httpd', 'web');

INSERT INTO tags(image_id, tag_name)
SELECT image_id, tag_name
FROM images
         JOIN (
    SELECT 'nginx' AS docker_image, 'latest' as tag_name
    UNION ALL SELECT 'nginx', 'alpine'
    UNION ALL SELECT 'nginx', 'stable'

    UNION ALL SELECT 'httpd', 'latest'
    UNION ALL SELECT 'httpd', 'alpine'
    UNION ALL SELECT 'httpd', '2.4'
) AS t ON images.docker_image = t.docker_image;