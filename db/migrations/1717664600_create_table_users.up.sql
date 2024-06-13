create table users
(
    id varchar(36) not null,
    name varchar(100) not null,
    email varchar(100) not null,
    password varchar(100) not null,
    url_photo varchar(100) not null,
    balance varchar(100) not null default 0,
    created_at bigint not null,
    updated_at bigint not null,
    primary key (id)
) engine = InnoDB;

INSERT INTO
    users(id, name, email, password, url_photo, balance, created_at, updated_at)
VALUES
    ('877e819b-adcd-4de2-b2bb-453fbad6f5b3', 'Ayas', 'ayas@student.ub.ac.id', '$2a$10$FEKBwDZ1t/DXd0YLSWAAz.Lq96ZuTAQsE4qMv4v3TnxE6fAp0V6Nm', 'default.jpg', '0', '1717688728', '1717688728'),
    ('c9a5c8cb-f1dc-43f7-85d8-2e4fe85a6a71', 'Ray', 'ray@student.ub.ac.id', '$2a$10$ZqueWhlrHgiogER2IEfIU.6TeRENLGgbme.HFvx3oEtjcqOZ9u7fW', 'default.jpg', '0', '1717688728', '1717688728'),
    ('e9549426-25fc-47b7-b555-4fdaec7a5989', 'Ijan', 'ijan@student.ub.ac.id', '$2a$10$HSntVyDh2Z4USIW.kbhD6uwveU449jnwlUm3HEajsfdDTwYQ9Cgsy', 'default.jpg', '0', '1717688728', '1717688728')
;