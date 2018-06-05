create table `domain` (
    id int(11) not null AUTO_INCREMENT primary key,
    created_at datetime not null,
    updated_at datetime not null,
    `name` varchar(256) not null default '',
    url varchar(256) not null default '',
    `private` tinyint not null default 0
);