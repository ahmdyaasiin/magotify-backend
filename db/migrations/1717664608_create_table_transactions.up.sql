create table transactions
(
    id varchar(36) not null,
    invoice_number varchar(100) not null,
    total_amount float not null, -- real price without discount
    shipping_costs float not null,
    status varchar(100) not null,
    service_name varchar(100) not null,
    service_type varchar(100) not null,
    receipt_number varchar(100),
    payment_type varchar(100),
    created_at bigint not null,
    updated_at bigint not null,
    address_id varchar(36) not null,
    voucher_id varchar(36),
    primary key (id),
    foreign key (address_id) references addresses(id),
    foreign key (voucher_id) references vouchers(id),
    unique (invoice_number)
) engine = InnoDB;

INSERT INTO
    transactions(id, invoice_number, total_amount, shipping_costs, status, service_name, service_type, receipt_number, payment_type, created_at, updated_at, address_id, voucher_id)
VALUES
    ('e237913b-fb26-4390-9a76-d0f344b15c63', 'INV/20240620/SHP/51731', 200000, 7000, 'done', 'jne', 'REG', 'NO_RESI', 'credit card', '1717688728', '1717688728', 'ae3322a1-c825-4474-82c3-758238990a16', null),
    ('31657344-8e8b-4959-bade-75ed82190168', 'INV/20240620/SHP/51732', 150000, 7000, 'done', 'jne', 'REG', 'NO_RESI', 'gopay', '1717688729', '1717688729', '478edeb3-05df-4fc7-aac1-77e88c0f1ba0', '59009d99-e8a6-41c9-b032-30dd00f26373'),
    ('d23ae285-8210-4066-bece-45e65a823995', 'INV/20240620/SHP/51733', 75000, 7000, 'done', 'jne', 'REG', 'NO_RESI', 'bank bca', '1717688730', '1717688730', '0406fa40-dfda-469d-89c7-7b21c1f3b9b3', null),
    ('93dc490c-0c41-42ab-8202-4624bb5f36d8', 'INV/20240620/SHP/51734', 50000, 7000, 'done', 'jne', 'REG', 'NO_RESI', 'bank mandiri', '1717688731', '1717688731', 'cb2d4d45-25f9-43e2-ae76-9b95a266aaae', '6dbe1ab1-1194-401a-bcb1-1190a7a85e3b'),
    ('cdb18ad2-7065-474f-b004-9a182e52cdc5', 'INV/20240620/SHP/51735', 75000, 7000, 'in-progress', 'jne', 'REG', 'NO_RESI', 'shopee pay', '1717688732', '1717688732', '97d44911-3c5f-4a66-82fb-a680439e27e1', null),
    ('7610180b-2d45-4ade-ba49-1b6a5e81906b', 'INV/20240620/SHP/51736', 200000, 7000, 'in-progress', 'jne', 'REG', 'NO_RESI', 'bank mandiri', '1717688733', '1717688733', '0406fa40-dfda-469d-89c7-7b21c1f3b9b3', null)
;