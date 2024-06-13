create table ratings
(
    id varchar(36) not null,
    content varchar(255) not null,
    star int not null,
    created_at bigint not null,
    updated_at bigint not null,
    user_id varchar(36) not null,
    transaction_item_id varchar(36) not null,
    primary key (id),
    foreign key (user_id) references users(id),
    foreign key (transaction_item_id) references transaction_items(id)
) engine = InnoDB;

INSERT INTO
    ratings(id, content, star, created_at, updated_at, user_id, transaction_item_id)
VALUES
    ('6b1c58a7-00eb-450e-9ae3-a13b46c56822', 'bagus banget', 5, '1717688728', '1717688728', '877e819b-adcd-4de2-b2bb-453fbad6f5b3', 'd1a43d9c-d767-4000-8c52-337fe18a4359'),
    ('3916a64e-96a1-44b4-a743-43cc60301a1c', 'berkualitas tinggi', 5, '1717688728', '1717688728', 'c9a5c8cb-f1dc-43f7-85d8-2e4fe85a6a71', '049e5a8b-9026-43c2-a925-2be04d8ed857'),
    ('55b1eada-5ff0-4b2a-8e94-85db9b065b49', 'sangat terjamin kualitasnya', 5, '1717688728', '1717688728', 'c9a5c8cb-f1dc-43f7-85d8-2e4fe85a6a71', 'f3ef6aa5-2bf6-4ef4-86e4-37d73ac670f1'),
    ('ba5f7f69-0760-4a0e-b89f-58f974eb0cd4', 'barangnya kurang manjur', 3, '1717688728', '1717688728', 'e9549426-25fc-47b7-b555-4fdaec7a5989', '8e0709cd-daca-41a6-a474-0e94aa93d93e'),
    ('30f7bdbd-9d6e-4165-81cb-018b91833b54', 'sama aja seperti pada umumnya', 4, '1717688728', '1717688728', '877e819b-adcd-4de2-b2bb-453fbad6f5b3', 'a19ead7b-5dab-4d33-9f5d-1c1847f9a5d3')
;