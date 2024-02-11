create table shopping_carts (
    id varchar not null,
    user_id varchar not null,
    constraint shoppingCarts_pk primary key (id),
    constraint shoppingCarts_users_fk foreign key (user_id) references users(id)
)