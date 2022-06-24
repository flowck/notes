-- +goose Up
-- +goose StatementBegin
insert into entries (id, content, user_id, created_at, updated_at) values ('a7b473d2-cbb3-4d71-9e0a-e42bacb5d22f', 'Proin interdum mauris non ligula pellentesque ultrices. Phasellus id sapien in sapien iaculis congue. Vivamus metus arcu, adipiscing molestie, hendrerit at, vulputate vitae, nisl.', 'fe1433f8-8576-4e04-87df-031778028bd5', '9/1/2021', '12/26/2021');
insert into entries (id, content, user_id, created_at, updated_at) values ('64086fd9-2eba-4e8d-972e-e79e10c74d42', 'Fusce consequat. Nulla nisl. Nunc nisl. Duis bibendum, felis sed interdum venenatis, turpis enim blandit mi, in porttitor pede justo eu massa. Donec dapibus. Duis at velit eu est congue elementum.', 'fe1433f8-8576-4e04-87df-031778028bd5', '2/10/2022', '2/26/2022');
insert into entries (id, content, user_id, created_at, updated_at) values ('a092a523-0bca-4fdd-a804-a131139ee7d8', 'Nulla ut erat id mauris vulputate elementum. Nullam varius. Nulla facilisi. Cras non velit nec nisi vulputate nonummy. Maecenas tincidunt lacus at velit. Vivamus vel nulla eget eros elementum pellentesque. Quisque porta volutpat erat. Quisque erat eros, viverra eget, congue eget, semper rutrum, nulla. Nunc purus. Phasellus in felis.', 'fe1433f8-8576-4e04-87df-031778028bd5', '4/29/2022', '10/24/2021');
insert into entries (id, content, user_id, created_at, updated_at) values ('6c89ff17-43d9-4206-a8f9-ab8c7201165c', 'Aliquam non mauris. Morbi non lectus. Aliquam sit amet diam in magna bibendum imperdiet. Nullam orci pede, venenatis non, sodales sed, tincidunt eu, felis.', 'fe1433f8-8576-4e04-87df-031778028bd5', '9/24/2021', '2/2/2022');
insert into entries (id, content, user_id, created_at, updated_at) values ('b5c51832-4969-49ba-9b6f-f75b4a566612', 'Nam congue, risus semper porta volutpat, quam pede lobortis ligula, sit amet eleifend pede libero quis orci. Nullam molestie nibh in lectus. Pellentesque at nulla. Suspendisse potenti. Cras in purus eu magna vulputate luctus.', 'fe1433f8-8576-4e04-87df-031778028bd5', '10/3/2021', '7/8/2021');
insert into entries (id, content, user_id, created_at, updated_at) values ('3b945283-7464-4747-9875-6d7c35ccccb6', 'Ut tellus. Nulla ut erat id mauris vulputate elementum. Nullam varius. Nulla facilisi. Cras non velit nec nisi vulputate nonummy. Maecenas tincidunt lacus at velit. Vivamus vel nulla eget eros elementum pellentesque. Quisque porta volutpat erat. Quisque erat eros, viverra eget, congue eget, semper rutrum, nulla.', 'fe1433f8-8576-4e04-87df-031778028bd5', '11/1/2021', '8/29/2021');
insert into entries (id, content, user_id, created_at, updated_at) values ('ffd4f409-d52a-4608-aeaa-e26b60263ca9', 'Morbi non lectus.', 'fe1433f8-8576-4e04-87df-031778028bd5', '10/28/2021', '12/28/2021');
insert into entries (id, content, user_id, created_at, updated_at) values ('e4fe6f84-6aec-493c-a717-5764d84a50ff', 'Maecenas rhoncus aliquam lacus. Morbi quis tortor id nulla ultrices aliquet. Maecenas leo odio, condimentum id, luctus nec, molestie sed, justo. Pellentesque viverra pede ac diam. Cras pellentesque volutpat dui. Maecenas tristique, est et tempus semper, est quam pharetra magna, ac consequat metus sapien ut nunc. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Mauris viverra diam vitae quam. Suspendisse potenti. Nullam porttitor lacus at turpis.', 'fe1433f8-8576-4e04-87df-031778028bd5', '8/12/2021', '12/3/2021');
insert into entries (id, content, user_id, created_at, updated_at) values ('947f20a6-5c7b-4515-aa49-7c2e22d96662', 'Nam dui. Proin leo odio, porttitor id, consequat in, consequat ut, nulla. Sed accumsan felis. Ut at dolor quis odio consequat varius. Integer ac leo. Pellentesque ultrices mattis odio. Donec vitae nisi. Nam ultrices, libero non mattis pulvinar, nulla pede ullamcorper augue, a suscipit nulla elit ac nulla. Sed vel enim sit amet nunc viverra dapibus.', 'fe1433f8-8576-4e04-87df-031778028bd5', '3/4/2022', '11/21/2021');
insert into entries (id, content, user_id, created_at, updated_at) values ('27b6fa58-caaf-4759-adc5-8d7f573aee19', 'Pellentesque at nulla. Suspendisse potenti. Cras in purus eu magna vulputate luctus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus vestibulum sagittis sapien. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam vel augue. Vestibulum rutrum rutrum neque. Aenean auctor gravida sem. Praesent id massa id nisl venenatis lacinia.', 'fe1433f8-8576-4e04-87df-031778028bd5', '3/14/2022', '10/18/2021');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FROM entries;
-- +goose StatementEnd
