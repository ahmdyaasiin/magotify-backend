create table addresses
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
    is_primary tinyint not null,
    phone_number varchar(100) not null,
    created_at bigint not null,
    updated_at bigint not null,
    user_id varchar(36) not null,
    primary key (id),
    foreign key (user_id) references users(id) on delete cascade
) engine = InnoDB;

INSERT INTO
    addresses(id, name, address, district, city, state, postal_code, latitude, longitude, is_primary, phone_number, created_at, updated_at, user_id)
VALUES
    ('ae3322a1-c825-4474-82c3-758238990a16', 'Kost', 'Jl. Candi Mendut No.54a, Mojolangu', 'Kec. Lowokwaru', 'Kota Malang', 'Jawa Timur', '65141', '-7.93937731274601', '112.62622890102713', '1', '081383975000', '1717688728', '1717688728', '877e819b-adcd-4de2-b2bb-453fbad6f5b3'),
    ('0406fa40-dfda-469d-89c7-7b21c1f3b9b3', 'Rumah Ray', 'Jl. S. Supriadi, Sukun', 'Kec. Sukun', 'Kota Malang', 'Jawa Timur', '65147', '-7.993964783134341', '112.62045086406329', '1', '081383974000', '1717688728', '1717688728', 'c9a5c8cb-f1dc-43f7-85d8-2e4fe85a6a71'),
    ('cb2d4d45-25f9-43e2-ae76-9b95a266aaae', 'Kos Ijan', 'Jl. Cengger Ayam No.25, Tulusrejo', 'Kec. Lowokwaru', 'Kota Malang', 'Jawa Timur', '65141', '-7.947445870834065', '112.6317153870079', '0', '081383973000', '1717688728', '1717688728', 'e9549426-25fc-47b7-b555-4fdaec7a5989'),
    ('97d44911-3c5f-4a66-82fb-a680439e27e1', 'Rumah Ayas', 'Perum Griya Shanta, Jl. Soekarno - Hatta Blk. B No.215, Mojolangu', 'Lowokwaru', 'Malang City', 'East Java', '65141', '-7.940305718559997', '112.62119737657375', '0', '081383975000', '1717688728', '1717688728', '877e819b-adcd-4de2-b2bb-453fbad6f5b3'),
    ('478edeb3-05df-4fc7-aac1-77e88c0f1ba0', 'Kantor Ijan', 'Puri, Jl. Bantaran Bar. IV No.Kav J, Tulusrejo', 'Lowokwaru', 'Malang City', 'East Java', '65141', '-7.948876934783974', '112.633219299733', '1', '081383973000', '1717688728', '1717688728', 'e9549426-25fc-47b7-b555-4fdaec7a5989')
;