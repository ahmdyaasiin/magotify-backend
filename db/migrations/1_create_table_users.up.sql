create table users
(
    id varchar(36) not null,
    name varchar(100) not null,
    email varchar(100) not null,
    password varchar(100) not null,
    photo_profile varchar(100) not null,
    created_at bigint not null,
    updated_at bigint not null,
    primary key (id)
) engine = InnoDB;