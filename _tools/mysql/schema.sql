CREATE TABLE `user`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'User ID',
    `username` VARCHAR(255) NOT NULL COMMENT 'Username',
    `password` VARCHAR(255) NOT NULL COMMENT 'Password',
    `email`    VARCHAR(255) NOT NULL COMMENT 'Email',
    `created`  DATETIME NOT NULL COMMENT 'レコード作成日時',
    `updated`  DATETIME NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uix_name` (`username`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT = 'ユーザー情報';

CREATE TABLE `task`
(
    `id`          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Task ID',
    `user_id`     BIGINT UNSIGNED NOT NULL COMMENT 'タスクを作成したUser ID',
    `title`       VARCHAR(255) NOT NULL COMMENT 'タスクのタイトル',
    `status`      VARCHAR(20) NOT NULL COMMENT 'タスクの状態',
    `created`     DATETIME NOT NULL COMMENT 'レコード作成日時',
    `updated`     DATETIME NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`) USING BTREE,
    CONSTRAINT `fk_user_id`
        FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
            ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT = 'タスク情報';
