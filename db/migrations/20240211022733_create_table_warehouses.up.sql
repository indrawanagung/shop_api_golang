create table warehouses (
    id varchar not null,
    store_id varchar not null,
    address_id varchar not null,
    status_id varchar not null,
    name varchar not null,
    created_at varchar not null,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint warehouses_pk primary key (id),
    constraint warehouses_store_fk foreign key (store_id) references stores(id),
    constraint warehouses_address_fk foreign key (address_id) references address(id),
    constraint warehouses_status_fk foreign key (status_id) references status(id)
)