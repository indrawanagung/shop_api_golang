create table product_images (
    id varchar not null,
    product_id varchar not null,
    image_url varchar not null,
    constraint productimages_pk primary key (id),
    constraint productImages_product_fk foreign key (product_id) references products(id)
)