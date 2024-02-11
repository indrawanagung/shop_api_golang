create table products (
    id varchar not null,
    name varchar not null,
    product_category_id varchar not null,
    status_id varchar not null,
    price bigint not null,
    created_at varchar not null,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint products_pk primary key (id),
    constraint products_productCategory_fk foreign key (product_category_id) references product_categories(id),
    constraint products_status_fk foreign key (status_id) references status(id)
)