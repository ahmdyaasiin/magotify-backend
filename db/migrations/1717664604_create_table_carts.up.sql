create table carts
(
    id varchar(36) not null,
    quantity int not null,
    created_at bigint not null,
    updated_at bigint not null,
    user_id varchar(36) not null,
    product_id varchar(36) not null,
    primary key (id),
    foreign key (user_id) references users(id),
    foreign key (product_id) references products(id)
) engine = InnoDB;

INSERT INTO
    carts(id, quantity, created_at, updated_at, user_id, product_id)
VALUES
    ('e3bcb236-cd0c-4772-a5ca-6bba8484f819', '2', '1717688728', '1717688728', 'e9549426-25fc-47b7-b555-4fdaec7a5989', 'fada2eb6-644d-4ac1-af13-6633b5ef498a'),
    ('c9256fc0-908e-4083-be12-17960aa286ec', '1', '1717688728', '1717688728', 'e9549426-25fc-47b7-b555-4fdaec7a5989', '450a7062-0dec-42f7-b546-4c915bf83546'),
    ('add6a96f-0eca-4ec5-b1b0-06ead99e49aa', '1', '1717688728', '1717688728', 'e9549426-25fc-47b7-b555-4fdaec7a5989', 'e0e7e792-c053-437b-b106-be7dbfb3b878'),
    ('83469834-63da-4128-ad70-dbaca0e19b59', '2', '1717688728', '1717688728', 'c9a5c8cb-f1dc-43f7-85d8-2e4fe85a6a71', '610f9446-fbb8-43e1-8eae-625e6e260551'),
    ('e9575e57-9afe-47ea-87d1-48b9804537fb', '3', '1717688728', '1717688728', 'c9a5c8cb-f1dc-43f7-85d8-2e4fe85a6a71', '38af2583-c7be-4b3b-bb60-030083e64508'),
    ('e1aed932-d624-4b6f-baf7-6cf327ba1470', '1', '1717688728', '1717688728', 'e9549426-25fc-47b7-b555-4fdaec7a5989', 'aa465303-585b-41b3-b000-a0f40646d66c')
;