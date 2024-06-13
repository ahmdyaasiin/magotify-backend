create table transactions
(
    id varchar(36) not null,
    total_amount float not null,
    status varchar(100) not null,
    created_at bigint not null,
    updated_at bigint not null,
    user_id varchar(36) not null,
    voucher_id varchar(36),
    primary key (id),
    foreign key (user_id) references users(id),
    foreign key (voucher_id) references vouchers(id)
) engine = InnoDB;

INSERT INTO
    transactions(id, total_amount, status, created_at, updated_at, user_id, voucher_id)
VALUES
    ('e237913b-fb26-4390-9a76-d0f344b15c63', 200000, 'done', '1717688728', '1717688728', '877e819b-adcd-4de2-b2bb-453fbad6f5b3', null),
    ('31657344-8e8b-4959-bade-75ed82190168', 150000, 'done',  '1717688728', '1717688728', 'e9549426-25fc-47b7-b555-4fdaec7a5989', '59009d99-e8a6-41c9-b032-30dd00f26373'),
    ('d23ae285-8210-4066-bece-45e65a823995', 75000, 'done',  '1717688728', '1717688728', 'c9a5c8cb-f1dc-43f7-85d8-2e4fe85a6a71', null),
    ('93dc490c-0c41-42ab-8202-4624bb5f36d8', 45000, 'done',  '1717688728', '1717688728', 'e9549426-25fc-47b7-b555-4fdaec7a5989', '6dbe1ab1-1194-401a-bcb1-1190a7a85e3b'),
    ('cdb18ad2-7065-474f-b004-9a182e52cdc5', 75000, 'in-progress',  '1717688728', '1717688728', '877e819b-adcd-4de2-b2bb-453fbad6f5b3', null),
    ('7610180b-2d45-4ade-ba49-1b6a5e81906b', 200000, 'in-progress',  '1717688728', '1717688728', 'c9a5c8cb-f1dc-43f7-85d8-2e4fe85a6a71', null)
;