create table `domain` (
    id int(11) not null AUTO_INCREMENT primary key,
    created_at datetime not null,
    updated_at datetime not null,
    `name` varchar(256) not null default '',
    url varchar(256) not null default '',
    `private` tinyint not null default 0
);

create table `computer` (
    id int(11) not null AUTO_INCREMENT primary key,
    created_at datetime not null,
    updated_at datetime not null,
    cpu int(11) not null  default 0,
    ram int(11) not null  default 0,
    private_ip varchar(32) not null default '',
    public_ip varchar(32) not null default ''
);

create table `disk` (
    id int(11) not null AUTO_INCREMENT primary key,
    size int(11) not null  default 0,
    `left` int(11) not null  default 0,
    computer_id int(11) not null default 0,
    created_at datetime not null,
    updated_at datetime not null
);