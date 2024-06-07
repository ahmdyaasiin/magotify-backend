create table wishlists
(
    id varchar(36) not null,
    created_at bigint not null,
    user_id varchar(36) not null,
    product_id varchar(36) not null,
    primary key (id),
    foreign key (user_id) references users(id),
    foreign key (product_id) references products(id)
) engine = InnoDB;

INSERT INTO
    wishlists(id, created_at, user_id, product_id)
VALUES
    ('41e120e6-6772-445b-ac52-fb518638fa6d', '1717688728', 'e9549426-25fc-47b7-b555-4fdaec7a5989', '450a7062-0dec-42f7-b546-4c915bf83546'),
    ('f99f19e0-6519-450d-9749-12ec33f2e19a', '1717688728', '877e819b-adcd-4de2-b2bb-453fbad6f5b3', '610f9446-fbb8-43e1-8eae-625e6e260551'),
    ('bcadff26-2cb4-46d8-a36b-34565dedf162', '1717688728', 'c9a5c8cb-f1dc-43f7-85d8-2e4fe85a6a71', '610f9446-fbb8-43e1-8eae-625e6e260551'),
    ('e8155111-adaf-42d8-965e-68e8399d93d7', '1717688728', 'c9a5c8cb-f1dc-43f7-85d8-2e4fe85a6a71', 'aa465303-585b-41b3-b000-a0f40646d66c'),
    ('ba6d2265-a704-4449-b0fe-962eeecb6c97', '1717688728', 'e9549426-25fc-47b7-b555-4fdaec7a5989', 'aa465303-585b-41b3-b000-a0f40646d66c')
;