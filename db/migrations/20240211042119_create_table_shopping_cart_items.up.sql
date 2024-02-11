create table shopping_cart_items (
    id varchar not null,
    shopping_cart_id varchar not null,
    product_id varchar not null,
    qty int not null,
    created_at varchar not null ,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint shoppingCartItems_pk primary key (id),
    constraint shoppingCartItems_shoppingCarts_fk foreign key (shopping_cart_id) references shopping_carts(id),
    constraint shoppingCartItems_products_fk foreign key (product_id) references products(id)
)