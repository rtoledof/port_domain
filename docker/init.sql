create table ports
(
    id varchar(36) not null,
	name varchar(50) null,
	city varchar(50) null,
	country varchar(50) null,
	alias json null,
	regions json null,
	coordinates json null,
	province varchar(100) null,
	timezone varchar(50) null,
	unlocs json null,
	code varchar(10)  null,
	constraint port_pk
		primary key (id)
);