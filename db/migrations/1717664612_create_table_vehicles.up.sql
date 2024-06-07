create table vehicles
(
    id varchar(36) not null,
    name varchar(100) not null,
    description varchar(100) not null,
    url_photo varchar(100) not null,
    status varchar(100) not null,
    created_at bigint not null,
    updated_at bigint not null,
    warehouse_id varchar(36) not null,
    primary key (id),
    foreign key (warehouse_id) references warehouses(id)
) engine = InnoDB;

INSERT INTO
    vehicles(id, name, description, url_photo, status, created_at, updated_at, warehouse_id)
VALUES
    ('2a455bee-a09e-4584-812e-e3ded337bd6e', 'Bike Pick Up', 'For Weights Under 5kg', 'bike.jpg', 'Instant Delivery', '1717688728', '1717688728', '7897cfad-ae15-4b58-afa3-7f479453c65b'),
    ('249986f4-f26d-48fd-937d-4cd76e938596', 'Car Pick Up', 'For Weights 5-10kg', 'car.jpg', 'Same Day', '1717688728', '1717688728', '7897cfad-ae15-4b58-afa3-7f479453c65b'),
    ('029ac42c-b32c-414d-88b7-bd88e4468179', 'Truck Pick Up', 'For Weights 10-50kg', 'truck.jpg', 'Same Day', '1717688728', '1717688728', '7897cfad-ae15-4b58-afa3-7f479453c65b'),
    ('624074f9-e2b0-40a7-b35f-e34381596c55', 'Bike Pick Up', 'For Weights Under 5kg', 'bike.jpg', 'Instant Delivery', '1717688728', '1717688728', '2352bc28-3d2d-4bb5-a64b-c149411c97be'),
    ('5c52f5e8-c55f-4ae3-89e8-472d1d84a3c6', 'Car Pick Up', 'For Weights 5-10kg', 'car.jpg', 'Same Day', '1717688728', '1717688728', '2352bc28-3d2d-4bb5-a64b-c149411c97be')
;