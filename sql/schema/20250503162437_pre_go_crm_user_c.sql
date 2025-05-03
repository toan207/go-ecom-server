-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS go_db_users (
  `usr_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Account ID',
  `usr_email` varchar(30) NOT NULL DEFAULT '' COMMENT 'Email',
  `usr_phone` varchar(15) NOT NULL DEFAULT '' COMMENT 'Phone number',
  `usr_name` varchar(30) NOT NULL DEFAULT '' COMMENT 'Username',
  `usr_password` varchar(32) NOT NULL DEFAULT '' COMMENT 'Password',
  `usr_created_at` int(11) NOT NULL DEFAULT '0' COMMENT 'Created time',
  `usr_updated_at` int(11) NOT NULL DEFAULT '0' COMMENT 'Updated time',
  `usr_created_ip_at` varchar(12) NOT NULL DEFAULT '' COMMENT 'Created IP',
  `usr_last_login_at` int(12) NOT NULL DEFAULT '0' COMMENT 'Last login time',
  `usr_last_login_ip_at` varchar(12) NOT NULL DEFAULT '' COMMENT 'Last login IP',
  `usr_login_at` int(12) NOT NULL DEFAULT '0' COMMENT 'Login time',
  `usr_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Status: 1: enable, 0: disable, -1: delete',
  PRIMARY KEY(usr_id),
  KEY `idx_email` (`usr_email`),
  KEY `idx_phone` (`usr_phone`),
  KEY `idx_name` (`usr_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Account';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `go_db_users`;
-- +goose StatementEnd
