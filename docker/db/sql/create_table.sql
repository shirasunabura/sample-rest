USE test

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `mail` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'メールアドレス',
  `password` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'メールアドレス',
  `password_reissue_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'パスワード再発行トークン',
  `password_reissue_limit` datetime DEFAULT NULL COMMENT 'パスワード再発行有効期限',
  `mail_modify_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'メールアドレス再発行トークン',
  `mail_modify_limit` datetime DEFAULT NULL COMMENT 'メールアドレス再発行有効期限',
  `family_name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '姓',
  `first_name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名',
  `family_name_kana` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '姓カナ',
  `first_name_kana` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名カナ',
  `is_test` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'テストユーザー',
  `enabled` tinyint(1) NOT NULL DEFAULT '1' COMMENT '0:退会、1:会員',
  `last_login_at` datetime NOT NULL COMMENT '最終ログイン日時',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='ユーザー';