CREATE TABLE `beers` (
  `id` serial not null
		constraint beers_pk
			primary key,
	`created_at` timestamp with time zone not null,
	`updated_at` timestamp with time zone not null,
	`deleted_at` timestamp with time zone,
  `name` varchar(50),
  `type` longtext,
  `detail` longtext,
  `image_url` longtext,
  `company_id` bigint(20)unsigned DEFAULT NULL
);

CREATE TABLE `users` (
  `id` serial not null
		constraint users_pk
			primary key,
	`created_at` timestamp with time zone not null,
	`updated_at` timestamp with time zone not null,
	`deleted_at` timestamp with time zone,
  `email` varchar(191)DEFAULT NULL,
  `password` longtext DEFAULT NULL
);

CREATE TABLE `companys` (
  `id` serial not null
		constraint company_pk
			primary key,
	`created_at` timestamp with time zone not null,
	`updated_at` timestamp with time zone not null,
	`deleted_at` timestamp with time zone,
  `name` varchar(50) DEFAULT NULL
);

CREATE TABLE `distributers` (
  `id` serial not null
		constraint distributers_pk
			primary key,
	`created_at` timestamp with time zone not null,
	`updated_at` timestamp with time zone not null,
	`deleted_at` timestamp with time zone,
  `name` varchar(50) DEFAULT NULL
);

CREATE TABLE `distributer_beers` (
  `id` serial not null
		constraint distributer_beers_pk
			primary key,
	`created_at` timestamp with time zone not null,
	`updated_at` timestamp with time zone not null,
	`deleted_at` timestamp with time zone,
  `beer_id` bigint(20) unsigned DEFAULT NULL,
  `distributer_id` bigint(20) unsigned DEFAULT NULL
);

ALTER TABLE `distributer_beers` ADD FOREIGN KEY (`beer_id`) REFERENCES `Beer` (`id`);
ALTER TABLE `distributer_beers` ADD FOREIGN KEY (`distributer_id`) REFERENCES `Distributer` (`id`);
ALTER TABLE `beers` ADD FOREIGN KEY (`company_id`) REFERENCES `companys` (`id`);
