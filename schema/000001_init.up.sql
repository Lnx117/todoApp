CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE todo_lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255)
);

CREATE TABLE users_lists
(
    id      serial not null unique,
    user_id bigint unsigned not null,
    list_id bigint unsigned not null,
    CONSTRAINT FK_users_lists_user_id FOREIGN KEY (user_id)
        REFERENCES users (id)
        ON DELETE CASCADE,
    CONSTRAINT FK_users_lists_list_id FOREIGN KEY (list_id)
        REFERENCES todo_lists (id)
        ON DELETE CASCADE
);

CREATE TABLE todo_items
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    done        boolean      not null default false
);


CREATE TABLE lists_items
(
    id      serial not null unique,
    item_id bigint unsigned not null,
    list_id bigint unsigned not null,
    CONSTRAINT FK_list_items_item_id FOREIGN KEY (item_id)
        REFERENCES todo_items (id)
        ON DELETE CASCADE,
    CONSTRAINT FK_list_items_list_id FOREIGN KEY (list_id)
        REFERENCES todo_lists (id)
        ON DELETE CASCADE
);