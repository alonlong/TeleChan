-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE `messages` (
  `id` int NOT NULL AUTO_INCREMENT,
  `content` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
--- SQL in this section is executed when the migration is rolled back.
