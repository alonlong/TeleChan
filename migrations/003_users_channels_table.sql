-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE `users_channels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` varchar(36) NOT NULL,
  `channel_id` varchar(36) NOT NULL,
  `name` varchar(32) NOT NULL,
  `owned` boolean NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE (`user_id`, `channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
--- SQL in this section is executed when the migration is rolled back.
