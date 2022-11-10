CREATE TABLE public.device (
	id serial NOT NULL,
	"name" varchar(50) NULL,
	brand varchar(50) NULL,
	creation_time varchar NULL,
	CONSTRAINT device_pk PRIMARY KEY (id)
);
