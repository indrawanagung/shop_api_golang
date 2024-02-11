create table variations (
    id varchar not null,
    product_category_id varchar not null,
    variation_name varchar not null,
    created_at varchar not null ,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint variations_pk primary key (id),
    constraint variations_productCategory_fk foreign key (product_category_id) references product_categories(id)
)