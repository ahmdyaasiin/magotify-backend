create table media
(
    id varchar(100) not null,
    url_photo varchar(100) not null,
    product_id varchar(100) not null,
    created_at bigint not null,
    primary key (id),
    foreign key (product_id) references products(id)
) engine = InnoDB;

INSERT INTO
    media(id, url_photo, product_id, created_at)
VALUES
    ('e1aed932-d624-4b6f-baf7-6cf327ba1470', 'pupuk-kasgot-100g-1.jpg', '450a7062-0dec-42f7-b546-4c915bf83546', '1717688728'),
    ('78b43392-9cb1-4eb8-93f4-6f7010d1d346', 'pupuk-kasgot-100g-2.jpg', '450a7062-0dec-42f7-b546-4c915bf83546', '1717688728'),
    ('72c935ae-ea68-4b3a-873e-7a1957bf8a25', 'pupuk-kasgot-100g-3.jpg', '450a7062-0dec-42f7-b546-4c915bf83546', '1717688728'),
    ('38c722e3-713c-4400-9a50-fb1ddc8ccb64', 'pupuk-kasgot-100g-4.jpg', '450a7062-0dec-42f7-b546-4c915bf83546', '1717688728'),
    ('aad9c250-6705-4d9f-8c4e-1d146fcee270', 'pupuk-kasgot-250g-1.jpg', 'fada2eb6-644d-4ac1-af13-6633b5ef498a', '1717688728'),
    ('3c114304-a7eb-46e5-94d8-cca68a5a0ff4', 'pupuk-kasgot-250g-2.jpg', 'fada2eb6-644d-4ac1-af13-6633b5ef498a', '1717688728'),
    ('66943756-a36a-4671-9e23-ac54c130239d', 'pupuk-kasgot-250g-3.jpg', 'fada2eb6-644d-4ac1-af13-6633b5ef498a', '1717688728'),
    ('ef7fe538-4cf0-4913-a6eb-f83e7c50c4ce', 'pupuk-kasgot-500g-1.jpg', '071a5233-db41-4a11-bb00-7bbe9269e0c2', '1717688728'),
    ('8c78b915-47e9-4dc6-9b1c-7dc6d7334d58', 'pupuk-kasgot-500g-2.jpg', '071a5233-db41-4a11-bb00-7bbe9269e0c2', '1717688728'),
    ('8efce3ba-7bc3-422b-ad01-10065779ac5e', 'pupuk-kasgot-1kg-1.jpg', '95423188-e1ea-483d-b7e1-b5dd6f2b252f', '1717688728'),
    ('72edb104-4b6c-4f56-9e59-0a9a20fd7d48', 'pupuk-kasgot-1kg-2.jpg', '95423188-e1ea-483d-b7e1-b5dd6f2b252f', '1717688728'),
    ('1daaadf4-0580-40d8-ac8b-28738f67cc76', 'pupuk-kasgot-1kg-3.jpg', '95423188-e1ea-483d-b7e1-b5dd6f2b252f', '1717688728'),
    ('62c4465f-0c54-49e9-a3de-9e4a1e48debf', 'pupuk-kasgot-1kg-4.jpg', '95423188-e1ea-483d-b7e1-b5dd6f2b252f', '1717688728'),
    ('7aea8027-4898-4042-95d8-d53845d18eec', 'pakan-sapi-1kg-1.jpg', 'e0e7e792-c053-437b-b106-be7dbfb3b878', '1717688728'),
    ('5de6ab65-7330-45b1-b11e-b6aa15bdabde', 'pakan-sapi-1kg-2.jpg', 'e0e7e792-c053-437b-b106-be7dbfb3b878', '1717688728'),
    ('fd7d3240-f918-4310-97bb-8ca1bf64f7c5', 'pakan-sapi-1kg-3.jpg', 'e0e7e792-c053-437b-b106-be7dbfb3b878', '1717688728'),
    ('5f811826-6e28-4118-9ecf-556a08fba346', 'pakan-sapi-1kg-4.jpg', 'e0e7e792-c053-437b-b106-be7dbfb3b878', '1717688728'),
    ('97fa14f7-af51-42fd-ae47-d18aa2f24ff0', 'pakan-sapi-500g-1.jpg', '38af2583-c7be-4b3b-bb60-030083e64508', '1717688728'),
    ('daf307b7-9ada-4ebc-945a-e2e67116fdc0', 'pakan-sapi-500g-2.jpg', '38af2583-c7be-4b3b-bb60-030083e64508', '1717688728'),
    ('60ce6237-9d98-4473-801c-461fc77eb87c', 'pakan-sapi-500g-3.jpg', '38af2583-c7be-4b3b-bb60-030083e64508', '1717688728'),
    ('f984cd18-0339-474b-8502-72a258dee0fd', 'pakan-kambing-1kg-1.jpg', '610f9446-fbb8-43e1-8eae-625e6e260551', '1717688728'),
    ('c1d6c8a6-7789-4d3f-98f7-36329bb3bea7', 'pakan-kambing-1kg-2.jpg', '610f9446-fbb8-43e1-8eae-625e6e260551', '1717688728'),
    ('199dd9f8-fecc-46e7-b055-ec4398e835e8', 'pakan-kambing-500g-1.jpg', 'aa465303-585b-41b3-b000-a0f40646d66c', '1717688728'),
    ('ccac7839-7f28-4b3e-9daa-071b35e3b612', 'pakan-kambing-500g-2.jpg', 'aa465303-585b-41b3-b000-a0f40646d66c', '1717688728')
;