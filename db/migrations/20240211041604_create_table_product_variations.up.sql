create table product_variations (
    product_id varchar not null,
    variation_option_id varchar not null,
    constraint productVariations_pk primary key (product_id, variation_option_id),
    constraint productVariations_products_fk foreign key (product_id) references products(id),
    constraint productVariations_variationOptions_fk foreign key (variation_option_id) references variation_options(id)
)