use go_schema;
/*用户信息表*/
create table if not exists `user_info`(
  `user_id` int unsigned auto_increment primary key,
  `user_realname` varchar(50),
  `user_nickname` varchar(50) not null,
  `user_password` varchar(50) default '123456',
  `user_age` int not null default 25,
  `user_sex` int check(user_sex in(0,1)),
  `user_adress` varchar(100),
  `user_phone` varchar(50),
  `user_qq` int,
  `user_wechat` varchar(25)
);
alter table user_info auto_increment=10000;
/**/