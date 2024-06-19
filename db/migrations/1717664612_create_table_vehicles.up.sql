create table vehicles
(
    id varchar(36) not null,
    name varchar(100) not null,
    duration varchar(100) not null,
    description varchar(100) not null,
    url_photo varchar(100) not null,
    status varchar(100) not null,
    created_at bigint not null,
    updated_at bigint not null,
    primary key (id)
) engine = InnoDB;

INSERT INTO
    vehicles(id, name, duration, description, url_photo, status, created_at, updated_at)
VALUES
    ('2a455bee-a09e-4584-812e-e3ded337bd6e', 'Bike Pick Up', '10-30 Menit', 'Max 5kg', 'bike.jpg', 'Instant Delivery', '1717688727', '1717688727'),
    ('249986f4-f26d-48fd-937d-4cd76e938596', 'Car Pick Up', '30-60 Menit', 'Max 30kg', 'car.jpg', 'Same Day', '1717688728', '1717688728'),
    ('029ac42c-b32c-414d-88b7-bd88e4468179', 'Truck Pick Up', '1-2 Jam', 'Max 100kg', 'truck.jpg', 'Same Day', '1717688729', '1717688729')
;