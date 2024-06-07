create table transaction_items
(
    id varchar(36) not null,
    quantity int not null,
    total_price varchar(100) not null,
    transaction_id varchar(36) not null,
    product_id varchar(36) not null,
    primary key (id),
    foreign key (transaction_id) references transactions(id),
    foreign key (product_id) references products(id)
) engine = InnoDB;

INSERT INTO
    transaction_items(id, quantity, total_price, transaction_id, product_id)
VALUES
    ('d1a43d9c-d767-4000-8c52-337fe18a4359', 2, 200000, 'e237913b-fb26-4390-9a76-d0f344b15c63', '95423188-e1ea-483d-b7e1-b5dd6f2b252f'),
    ('5e3f7187-8e8e-437d-906b-e2eb79fde24f', 2, 150000, '31657344-8e8b-4959-bade-75ed82190168', '071a5233-db41-4a11-bb00-7bbe9269e0c2'),
    ('f3ef6aa5-2bf6-4ef4-86e4-37d73ac670f1', 1, 50000, 'd23ae285-8210-4066-bece-45e65a823995', 'fada2eb6-644d-4ac1-af13-6633b5ef498a'),
    ('049e5a8b-9026-43c2-a925-2be04d8ed857', 1, 25000, 'd23ae285-8210-4066-bece-45e65a823995', '450a7062-0dec-42f7-b546-4c915bf83546'),
    ('8e0709cd-daca-41a6-a474-0e94aa93d93e', 1, 50000, '93dc490c-0c41-42ab-8202-4624bb5f36d8', 'fada2eb6-644d-4ac1-af13-6633b5ef498a'),
    ('8247c569-1ec1-40ee-ae18-12152ec80103', 1, 25000, 'cdb18ad2-7065-474f-b004-9a182e52cdc5', '450a7062-0dec-42f7-b546-4c915bf83546'),
    ('a19ead7b-5dab-4d33-9f5d-1c1847f9a5d3', 1, 50000, 'cdb18ad2-7065-474f-b004-9a182e52cdc5', 'fada2eb6-644d-4ac1-af13-6633b5ef498a'),
    ('f762d9db-9b0e-4604-ab96-4fba33d4a071', 1, 200000, '7610180b-2d45-4ade-ba49-1b6a5e81906b', '610f9446-fbb8-43e1-8eae-625e6e260551')
;