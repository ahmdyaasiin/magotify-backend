create table banners
(
    id varchar(36) not null,
    url_photo varchar(100) not null,
    is_clickable tinyint not null,
    destination varchar(100),
    primary key (id)
) engine = InnoDB;

INSERT INTO
    banners(id, url_photo, is_clickable, destination)
VALUES
    ('23deab01-a371-43e7-ba30-095887c3ebeb', 'banner-1.jpg', 0, ''),
    ('9040160b-ebc5-4553-9fc5-2ec217d815e4', 'banner-2.jpg', 0, ''),
    ('7cc3ee63-de19-496f-a72a-ef33941b62ec', 'banner-3.jpg', 0, ''),
    ('6c00a47d-f259-483b-89de-c89b6b7cf667', 'banner-4.jpg', 0, '')
;