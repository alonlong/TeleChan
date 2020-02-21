-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE `channels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `channel_id` varchar(36) NOT NULL,
  `user_id` varchar(36) NOT NULL,
  `name` varchar(32) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE (`channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
--- SQL in this section is executed when the migration is rolled back.
