create table drivers
(
    id varchar(36) not null,
    name varchar(100) not null,
    email varchar(100) not null,
    phone_number varchar(100) not null,
    password varchar(255) not null,
    vehicle_id varchar(36) not null,
    warehouse_id varchar(36) not null,
    plate_number varchar(100) not null,
    created_at bigint not null,
    updated_at bigint not null,
    primary key (id),
    unique(email),
    unique(phone_number),
    foreign key (vehicle_id) references vehicles(id),
    foreign key (warehouse_id) references warehouses(id)
) engine = InnoDB;

INSERT INTO
    drivers(id, name, email, phone_number, password, vehicle_id, warehouse_id, plate_number, created_at, updated_at)
VALUES
    ('a216f581-e261-4144-8329-04c7aaa16a80', 'Vega Adriansyah', 'rvegaandri@yahoo.co.id', '082832304002', '$2a$10$ym9EV7LBCfax58CigOPqlu.EjZz.9Ds/sG9fcyqXSdiXID8BS87NS', '2a455bee-a09e-4584-812e-e3ded337bd6e', '7897cfad-ae15-4b58-afa3-7f479453c65b', 'N 3060 WQS', '1717688728', '1717688728'),
    ('1e6bee79-7d0c-4850-b7e0-f00505aac3aa', 'Jarwa Nainggolan', 'jarwnaing@gmail.com', '081292644500', '$2a$10$UGWbwJ3FCp6S6k62LgCXkO2QhP/zsrBPXZO7bwaK6FhF0LjkKW3Hu', '2a455bee-a09e-4584-812e-e3ded337bd6e', '7897cfad-ae15-4b58-afa3-7f479453c65b', 'N 891 OPN', '1717688728', '1717688728'),
    ('6d2f403d-8f57-420c-830e-f231d891d38e', 'Utama Hutasoit', 'utamahuta@gmail.com', '08128019281', '$2a$10$4/i6T4Dzcda.5lZ1VoHwJuskDUd5a9zMQKFXrgpX1.ncn.mGJN/0C', '2a455bee-a09e-4584-812e-e3ded337bd6e', '7897cfad-ae15-4b58-afa3-7f479453c65b', 'N 982 IBX', '1717688728', '1717688728'),
    ('23f51d46-6b31-4218-a3ae-c1b7a22eacd9', 'Muni Pradana', 'munprad@gmail.com', '08124119860', '$2a$10$bUdKJab5sN9KCMPB6rBZgeC3f0e35j82PUAUamjMuyLLqSILI1Xwa', '249986f4-f26d-48fd-937d-4cd76e938596', '7897cfad-ae15-4b58-afa3-7f479453c65b', 'N 812 OPB', '1717688728', '1717688728'),
    ('db4e7d53-7b35-4bd8-a82e-c453b90e666c', 'Elvin Natsir', 'natsirn@gmail.com', '08135818549', '$2a$10$0ut38mRw2D8WoYZ.Yi0AOOVU0AthzyUd1z4S.2UURWJdgv8hGjR3a', '2a455bee-a09e-4584-812e-e3ded337bd6e', '2352bc28-3d2d-4bb5-a64b-c149411c97be', 'N 071 CAH', '1717688728', '1717688728'),
    ('b9f79eb7-1806-4955-9c1d-8f5094f16ff4', 'Erik Hardiansyah', 'erikhard@gmail.com', '08583414463', '$2a$10$gq9YF/cxHj95yIZ8Wt85QeKqUfiAbsp4NeycVgFqRkLLHOiG5MR2e', '2a455bee-a09e-4584-812e-e3ded337bd6e', '2352bc28-3d2d-4bb5-a64b-c149411c97be', 'N 126 QNW', '1717688728', '1717688728'),
    ('de5ad48d-fe06-4b0e-ae75-973dec9497d0', 'Dimaz Kuswoyo', 'dimazk@gmail.com', '08964204760', '$2a$10$bNDBWH69gosC82Svjuy5FeUEZfxRnUCz5DggoFozTz4CqiUayZGIy', '249986f4-f26d-48fd-937d-4cd76e938596', '2352bc28-3d2d-4bb5-a64b-c149411c97be', 'N 019 BQW', '1717688728', '1717688728'),
    ('47295be2-dcf7-4a36-94e7-8625d1296bbb', 'Darmaji Hutasoit', 'darmji@gmail.com', '089827007996', '$2a$10$fR5HtLW3Rc8TWPiU.GrDfuXce1Q2V4h7tzVJbF.aYjrAhTX20eluG', '249986f4-f26d-48fd-937d-4cd76e938596', '2352bc28-3d2d-4bb5-a64b-c149411c97be', 'N 965 NJK', '1717688728', '1717688728')
;