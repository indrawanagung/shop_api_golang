create table user_reviews (
    id varchar not null,
    user_id varchar not null,
    order_product_id varchar not null ,
    rating_product_id varchar not null,
    comment varchar not null,
    created_at varchar not null,
    constraint userReviews_pk primary key (id),
    constraint userReviews_users foreign key (user_id) references users(id),
    constraint userReviews_orderProducts_fk foreign key (order_product_id) references shop_order_items(id),
    constraint userReviews_ratingProducts_fk foreign key (rating_product_id) references rating_products(id)
)