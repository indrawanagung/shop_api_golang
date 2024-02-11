INSERT INTO public.cities
(id, city, created_at, updated_at, deleted_at)
VALUES('1', 'makassar', '1707660761', NULL, NULL);
INSERT INTO public.cities
(id, city, created_at, updated_at, deleted_at)
VALUES('2', 'surabaya', '1707660761', NULL, NULL);


INSERT INTO public.users
(id, fullname, email_address, phone_number, "password", created_at, updated_at, deleted_at)
VALUES('1', 'John Doe', 'johndoe@gmail.com', '0872121234', '123', '1707660761', NULL, NULL);


INSERT INTO public.address
(id, "name", city_id, postal_code, created_at, updated_at, deleted_at, user_id)
VALUES('1', 'Jalan Landak No 123', '1', 90111, '1707660761', NULL, NULL, '1');


INSERT INTO public.category_status
(id, "name")
VALUES('1', 'warehouse');


INSERT INTO public.status
(id, category_status_id, "name", created_at, updated_at, deleted_at)
VALUES('1', '1', 'active', '1707660761', NULL, NULL);

INSERT INTO public.stores
(id, user_id, "name", points, created_at, updated_at, deleted_at)
VALUES('1', '1', 'Landak Store', 100, '1707660761', NULL, NULL);

INSERT INTO public.warehouses
(id, store_id, address_id, status_id, "name", created_at, updated_at, deleted_at)
VALUES('1', '1', '1', '1', 'landak warehouse 1', '1707660761', NULL, NULL);


INSERT INTO public.product_categories
(id, category_name)
VALUES('1', 'electronic');
INSERT INTO public.product_categories
(id, category_name)
VALUES('2', 'fashion');
INSERT INTO public.product_categories
(id, category_name)
VALUES('3', 'toy');
INSERT INTO public.product_categories
(id, category_name)
VALUES('4', 'accessoris');


