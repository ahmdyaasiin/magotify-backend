create table categories
(
    id varchar(36) not null,
    name varchar(100) not null,
    url_photo varchar(100) not null,
    created_at bigint not null,
    updated_at bigint not null,
    primary key (id)
) engine = InnoDB;

INSERT INTO
    categories(id, name, url_photo, created_at, updated_at)
VALUES
    ('f2740da4-ff8d-4143-a89d-de8e80b6dc6b', 'Pupuk', 'pupuk.jpg', '1717688728', '1717688728'),
    ('a213cace-f181-4023-b409-ac9a43b5c2d2', 'Pakan', 'pakan.jpg', '1717688728', '1717688728')
;