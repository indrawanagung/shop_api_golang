create table product_stocks (
    id varchar not null,
    product_id varchar not null,
    total int not null,
    not_reserved int not null,
    reserved int not null,
    created_at varchar not null,
    updated_at varchar null,
    constraint productStocks_pk primary key (id),
    constraint productStocks_fk foreign key (product_id) references products(id)
)