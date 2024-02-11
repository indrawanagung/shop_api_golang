create table shop_order_items (
    id varchar not null,
    product_id varchar not null,
    shop_order_id varchar not null,
    qty int not null,
    price bigint not null,
    status_id varchar not null,
    constraint shopOrderItems_pk primary key (id),
    constraint shopOrderItems_products_fk foreign key (product_id) references products(id),
    constraint shopOrderItems_shopOrders_fk foreign key (shop_order_id) references shop_orders(id),
    constraint shopOrderItems_status_fk foreign key (status_id) references status(id)
)