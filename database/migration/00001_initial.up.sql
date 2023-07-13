create table if not exists users(
                                    id serial primary key,
                                    user_name text not null,
                                    user_email text not null,
                                    password text not null,
                                    joined_at timestamp default current_timestamp
);
create table if not exists user_role(
                                        id serial primary key,
                                        user_id int references users(id),
                                        user_type text not null
);
create table if not exists address(
                                      id serial primary key,
                                      user_id int references users(id),
                                      latitude decimal not null,
                                      longitude decimal not null,
                                      location text not null
);
create table if not exists restaurants(
                                          id serial primary key,
                                          restaurant_name text not null,
                                          longitude decimal not null,
                                          latitude decimal not null,
                                          address text not null,
                                          opening_time timestamp not null,
                                          closing_time timestamp not null,
                                          user_id int references users(id)
);
create table if not exists dishes(
                                     id serial primary key,
                                     dish_name text not null,
                                     dish_type text not null,
                                     restaurant_id int references restaurants(id),
                                     start_serve timestamp not null,
                                     finish_serve timestamp not null
);
INSERT INTO users (user_name, user_email, password)
VALUES ('Gautam Jain', 'gautamjain@example.com', 'password123');

WITH users_id AS (
    SELECT id
    FROM users
    WHERE user_name = 'Gautam Jain'
    LIMIT 1
)
INSERT INTO user_role (user_id, user_type)
SELECT id, 'admin'
FROM users_id;


WITH users_id AS (
    SELECT id
    FROM users
    WHERE user_name = 'Gautam Jain'
    LIMIT 1
)
INSERT INTO address (user_id, latitude,longitude,location)
SELECT id,64.4556,54.8645,'Delhi'
FROM users_id;