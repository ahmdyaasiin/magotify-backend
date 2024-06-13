create table vouchers
(
    id varchar(36) not null,
    name varchar(100) not null,
    description varchar(100) not null,
    status tinyint not null,
    amount int not null,
    is_percent tinyint not null,
    url_logo varchar(100) not null,
    min_amount float not null,
    created_at bigint not null,
    updated_at bigint not null,
    user_id varchar(36) not null,
    primary key (id),
    foreign key (user_id) references users(id)
) engine = InnoDB;

INSERT INTO
    vouchers(id, name, description, status, amount, is_percent, min_amount, url_logo, created_at, updated_at, user_id)
VALUES
    ('8777b995-ab33-493d-86f7-b4561a7d24af', 'Perayaan Hari Pakan', 'Voucher Hari Pakan', 1, 5, 1, 100000, 'hari-pakan.jpg', '1717688728', '1717688728', '877e819b-adcd-4de2-b2bb-453fbad6f5b3'),
    ('6797538a-175f-4103-a1ea-e47c0ca6dc1a', 'Perayaan Hari Pakan', 'Voucher Hari Pakan', 1, 5, 1, 100000, 'hari-pakan.jpg', '1717688728', '1717688728', 'c9a5c8cb-f1dc-43f7-85d8-2e4fe85a6a71'),
    ('837ddaef-ad01-4dc9-8ebd-400653c9baea', 'Perayaan Hari Pakan', 'Voucher Hari Pakan', 1, 5, 1, 100000, 'hari-pakan.jpg', '1717688728', '1717688728', 'e9549426-25fc-47b7-b555-4fdaec7a5989'),
    ('7231011e-35a5-4742-8a53-229e8e0b8f3d', 'User terpilih', 'Voucher User Terpilih Bulan Juni', 1, 100000, 0, 1000000, 'user-terpilih.jpg', '1717688728', '1717688728', 'e9549426-25fc-47b7-b555-4fdaec7a5989'),
    ('59009d99-e8a6-41c9-b032-30dd00f26373', 'Voucher Test 1', 'Ini adalah voucher test pertama', 0, 10000, 0, 100000, 'voucher-test-1.jpg', '1717688728', '1717688728', 'e9549426-25fc-47b7-b555-4fdaec7a5989'),
    ('6dbe1ab1-1194-401a-bcb1-1190a7a85e3b', 'Voucher Test 2', 'Ini adalah voucher test kedua', 10, 5, 1, 10000, 'voucher-test-2.jpg', '1717688728', '1717688728', 'e9549426-25fc-47b7-b555-4fdaec7a5989')
;