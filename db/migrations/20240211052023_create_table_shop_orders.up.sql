create table shop_orders (
    id varchar not null,
    user_id varchar not null,
    address_id varchar not null,
    payment_type varchar not null,
    total_price bigint not null,
    ordered_at varchar not null,
    deleted_at timestamp null,
    constraint shopOrders_pk primary key (id),
    constraint shopOrders_users_fk foreign key (user_id) references users(id),
    constraint shopOrders_address_fk foreign key (address_id) references address(id)
)