create table product_volumes (
    id varchar not null ,
    product_id varchar not null,
    width int not null ,
    height int not null,
    length int not null ,
    created_at varchar not null,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint productVolume_pk primary key (id),
    constraint productVolume_product_fk foreign key (product_id) references products(id)
)