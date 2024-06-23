create table orders
(
    id varchar(36) not null,
    invoice_number varchar(100) not null,
    total_amount float not null,
    weight float not null,
    status varchar(100) not null,
    payment_type varchar(100),
    created_at bigint not null,
    updated_at bigint not null,
    address_id varchar(36) not null,
    driver_id varchar(36) not null,
    voucher_id varchar(36),
    primary key (id),
    foreign key (address_id) references addresses(id),
    foreign key (driver_id) references drivers(id),
    foreign key (voucher_id) references vouchers(id),
    unique (invoice_number)
) engine = InnoDB;

INSERT INTO
    orders(id, invoice_number, total_amount, weight, status, payment_type, created_at, updated_at, address_id, driver_id, voucher_id)
VALUES
    ('3148b57e-284a-43da-9ee1-26aa1bec8481', 'INV/20240620/PCK/51731', 150000, 2, 'done', 'credit card', '1717688728', '1717688728', '0406fa40-dfda-469d-89c7-7b21c1f3b9b3', 'b9f79eb7-1806-4955-9c1d-8f5094f16ff4', null), -- rey
    ('f1812fb1-e069-47fd-8499-76340a5757a7', 'INV/20240620/PCK/51732', 150000, 4.4, 'done', 'credit card', '1717688729', '1717688729', '478edeb3-05df-4fc7-aac1-77e88c0f1ba0', '1e6bee79-7d0c-4850-b7e0-f00505aac3aa', null), -- ijan
    ('d46a61cd-854c-4731-91e7-2b6bd118cc6d', 'INV/20240620/PCK/51733', 150000, 5.1, 'done', 'gopay', '1717688730', '1717688730', 'ae3322a1-c825-4474-82c3-758238990a16', '23f51d46-6b31-4218-a3ae-c1b7a22eacd9', null), -- ayas
    ('89a6bf31-1824-443c-81e1-22e5323e03b7', 'INV/20240620/PCK/51734', 150000, 3.9, 'done', 'gopay', '1717688731', '1717688731', '0406fa40-dfda-469d-89c7-7b21c1f3b9b3', 'db4e7d53-7b35-4bd8-a82e-c453b90e666c', null), -- rey
    ('1ef8fdaf-7b5f-4f1d-bcf4-ece0424043fa', 'INV/20240620/PCK/51735', 150000, 4.2, 'in-progress', 'ovo', '1717688732', '1717688732', '0406fa40-dfda-469d-89c7-7b21c1f3b9b3', 'b9f79eb7-1806-4955-9c1d-8f5094f16ff4', null), -- rey
    ('80105d26-79f4-4811-8033-75d53f15825c', 'INV/20240621/PCK/51736', 150000, 10.2, 'waiting-for-payment', null, '1717688733', '1717688733', '478edeb3-05df-4fc7-aac1-77e88c0f1ba0', '23f51d46-6b31-4218-a3ae-c1b7a22eacd9', null) -- ijan
;