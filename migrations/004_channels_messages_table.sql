-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE `channels_messages` (
  `id` int NOT NULL AUTO_INCREMENT,
  `channel_id` varchar(36) NOT NULL,
  `message_id` int NOT NULL,
  `brief` varchar(256) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE (`channel_id`, `message_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
--- SQL in this section is executed when the migration is rolled back.
