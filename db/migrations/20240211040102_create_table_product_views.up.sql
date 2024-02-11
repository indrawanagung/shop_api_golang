create table product_views (
    product_id varchar not null,
    user_id varchar not null,
    constraint productViews_pk primary key (product_id, user_id),
    constraint productViews_user_fk foreign key (user_id) references users(id),
    constraint productViews_product_fk foreign key (product_id) references products(id)

)