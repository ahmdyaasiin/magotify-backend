create table warehouses
(
    id varchar(36) not null,
    name varchar(100) not null,
    address varchar(100) not null,
    district varchar(100) not null,
    city varchar(100) not null,
    state varchar(100) not null,
    postal_code varchar(100) not null,
    latitude float(10, 6) not null,
    longitude float(10, 6) not null,
    created_at bigint not null,
    updated_at bigint not null,
    primary key (id)
) engine = InnoDB;

INSERT INTO
    warehouses(id, name, address, district, city, state, postal_code, latitude, longitude, created_at, updated_at)
VALUES
    ('7897cfad-ae15-4b58-afa3-7f479453c65b', 'Warehouse Suhat', '3J7F+HV9, Jl. Puncak Borobudur, Mojolangu', 'Kec. Lowokwaru', 'Kota Malang', 'Jawa Timur', '65142', '-7.936125568356787', '112.62478071845362', '1717688728', '1717688728'),
    ('2352bc28-3d2d-4bb5-a64b-c149411c97be', 'Warehouse Sukun', 'Jl. S. Supriadi Gg. 4 No.91, Tanjungrejo, Sukun', 'Kec. Sukun', 'Kota Malang', 'Jawa Timur', '65147', '-7.995419963700882', '112.62032929545205', '1717688728', '1717688728')
;