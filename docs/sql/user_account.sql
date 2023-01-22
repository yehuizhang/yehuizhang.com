DROP TABLE IF EXISTS user;

CREATE TABLE user_account
(
    id        serial PRIMARY KEY,
    uuid      UUID                        NOT NULL,
    username  varchar(20)                 NOT NULL UNIQUE,
    password  varchar(40)                 NOT NULL,
    isActive  boolean                     NOT NULL,
    creatPedAt timestamp(6) WITH TIME ZONE NOT NULL,
    updatedAt timestamp(6) WITH TIME ZONE NOT NULL
);

INSERT INTO user_account(uuid, username, password, isActive, createdAt, updatedAt)
VALUES ('607366e6-9a14-11ed-a8fc-0242ac120002', 'yehuizhang2', '123xxx', TRUE, to_timestamp(1674366277.223122),
        to_timestamp(1674366277.223122));

SELECT createdat, updatedat
FROM user_account ua;

SELECT extract(epoch from createdAt), extract(epoch from updatedat)
FROM user_account ua;
