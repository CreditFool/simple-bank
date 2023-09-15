CREATE TABLE `account` (
  `id` bigserial PRIMARY KEY,
  `owner` varchar(255) NOT NULL,
  `balance` bigint NOT NULL,
  `created_at` timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE `entry` (
  `id` bigserial PRIMARY KEY,
  `account` bigint NOT NULL,
  `amount` bigint NOT NULL,
  `created_at` timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE `transfer` (
  `id` bigserial PRIMARY KEY,
  `from_account` bigint NOT NULL,
  `to_account` bigint NOT NULL,
  `amount` bigint NOT NULL,
  `created_at` timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX `account_index_0` ON `account` (`owner`);

CREATE INDEX `entry_index_1` ON `entry` (`account`);

CREATE INDEX `transfer_index_2` ON `transfer` (`from_account`);

CREATE INDEX `transfer_index_3` ON `transfer` (`to_account`);

CREATE INDEX `transfer_index_4` ON `transfer` (`from_account`, `to_account`);

ALTER TABLE `entry` ADD FOREIGN KEY (`account`) REFERENCES `account` (`id`);

ALTER TABLE `transfer` ADD FOREIGN KEY (`from_account`) REFERENCES `account` (`id`);

ALTER TABLE `transfer` ADD FOREIGN KEY (`to_account`) REFERENCES `account` (`id`);
