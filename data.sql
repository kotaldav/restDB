#
# TABLE STRUCTURE FOR: authors
#

DROP TABLE IF EXISTS `authors`;

CREATE TABLE `authors` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `last_name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `birthdate` date NOT NULL,
  `added` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

INSERT INTO `authors` (`id`, `first_name`, `last_name`, `email`, `birthdate`, `added`) VALUES (1, 'Blaise', 'Nolan', 'darien22@example.com', '1998-01-07', '2013-07-02 21:40:26');
INSERT INTO `authors` (`id`, `first_name`, `last_name`, `email`, `birthdate`, `added`) VALUES (2, 'Adrien', 'Mante', 'chloe02@example.net', '2004-11-24', '1980-03-11 12:50:00');
INSERT INTO `authors` (`id`, `first_name`, `last_name`, `email`, `birthdate`, `added`) VALUES (3, 'Brown', 'O\'Connell', 'tara.blanda@example.org', '1975-06-10', '2006-11-21 06:40:56');
INSERT INTO `authors` (`id`, `first_name`, `last_name`, `email`, `birthdate`, `added`) VALUES (4, 'Griffin', 'Farrell', 'lynn.tromp@example.net', '1993-09-04', '1995-07-13 14:33:14');
INSERT INTO `authors` (`id`, `first_name`, `last_name`, `email`, `birthdate`, `added`) VALUES (5, 'Cecelia', 'Schuster', 'russ24@example.org', '1974-07-09', '1977-09-12 09:38:00');
INSERT INTO `authors` (`id`, `first_name`, `last_name`, `email`, `birthdate`, `added`) VALUES (6, 'Lempi', 'Wolff', 'pfeffer.madonna@example.com', '1993-08-30', '2007-10-14 08:21:30');
INSERT INTO `authors` (`id`, `first_name`, `last_name`, `email`, `birthdate`, `added`) VALUES (7, 'Hobart', 'Simonis', 'dexter38@example.org', '2011-12-05', '1971-05-02 05:03:03');
INSERT INTO `authors` (`id`, `first_name`, `last_name`, `email`, `birthdate`, `added`) VALUES (8, 'Antwon', 'Funk', 'ihand@example.org', '1984-12-02', '1976-12-20 15:02:28');
INSERT INTO `authors` (`id`, `first_name`, `last_name`, `email`, `birthdate`, `added`) VALUES (9, 'Beth', 'Ruecker', 'oward@example.net', '1977-10-26', '1992-07-15 19:29:46');
INSERT INTO `authors` (`id`, `first_name`, `last_name`, `email`, `birthdate`, `added`) VALUES (10, 'Weldon', 'Schinner', 'ricky41@example.org', '1979-04-07', '2002-12-08 13:45:23');


#
# TABLE STRUCTURE FOR: posts
#

DROP TABLE IF EXISTS `posts`;

CREATE TABLE `posts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `author_id` int(11) NOT NULL,
  `title` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `description` varchar(500) COLLATE utf8_unicode_ci NOT NULL,
  `content` text COLLATE utf8_unicode_ci NOT NULL,
  `date` date NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

INSERT INTO `posts` (`id`, `author_id`, `title`, `description`, `content`, `date`) VALUES (1, 1, 'Odit est expedita eaque quam temporibus placeat asperiores.', 'Placeat ut quia dolorem voluptates culpa. Quo nostrum ut commodi aut. Temporibus blanditiis ut quam tenetur. Quo accusamus autem ut voluptas.', 'Vel voluptatem sequi similique suscipit autem sunt quia. Quidem quidem vel quia eos.', '2015-11-10');
INSERT INTO `posts` (`id`, `author_id`, `title`, `description`, `content`, `date`) VALUES (2, 2, 'Tenetur nihil occaecati enim quasi dignissimos architecto maxime.', 'Omnis maxime exercitationem nisi et. Repudiandae nesciunt placeat voluptatem alias. Corporis eum occaecati sit quas necessitatibus hic sint.', 'Eveniet omnis et porro sit ea. Voluptatem veniam dolor veritatis numquam ab rem. Nesciunt modi quo illum illo dolor eos.', '1972-07-10');
INSERT INTO `posts` (`id`, `author_id`, `title`, `description`, `content`, `date`) VALUES (3, 3, 'Perferendis qui omnis officiis temporibus voluptatem aliquid.', 'Et totam culpa sunt. Non laboriosam corrupti perspiciatis repudiandae. Qui corrupti maxime accusamus est qui reiciendis.', 'Quas omnis quo in rem est error qui quis. At neque pariatur ea id vero voluptas. Aut iure odio magni sapiente.', '1986-05-14');
INSERT INTO `posts` (`id`, `author_id`, `title`, `description`, `content`, `date`) VALUES (4, 4, 'Non ratione corrupti et aut fugiat.', 'Autem repellendus praesentium voluptatibus laborum veritatis. Assumenda placeat sint architecto repellendus odio. Rerum qui doloribus quos architecto. Sequi et possimus voluptatibus aspernatur fugiat.', 'Laborum repellat et porro. Nihil et quaerat debitis occaecati natus.\nMaiores sit ratione est ipsa. Qui non officiis adipisci praesentium. Quam sed quia consequatur quasi ea qui tempora.', '2009-09-13');
INSERT INTO `posts` (`id`, `author_id`, `title`, `description`, `content`, `date`) VALUES (5, 5, 'Molestiae ducimus laboriosam impedit dolor cupiditate ad.', 'Reprehenderit esse numquam rerum ducimus. Id laudantium officiis suscipit rerum animi ut et. Saepe vero debitis alias voluptas quia beatae reiciendis.', 'Officia harum sit quia modi consequuntur. Explicabo aut nulla inventore non amet in. Temporibus sit culpa similique quisquam ad maiores beatae. Qui iure quia nobis sint rem.', '2000-06-13');
INSERT INTO `posts` (`id`, `author_id`, `title`, `description`, `content`, `date`) VALUES (6, 6, 'Fuga in adipisci harum accusamus ullam quibusdam.', 'Autem et qui cum optio tenetur inventore fuga. Repellendus cumque sed in quisquam cumque ullam adipisci.', 'Neque cumque voluptatem ipsa et eaque id est. Hic tempora blanditiis perferendis asperiores quidem rerum dolor. Quaerat dolore ducimus commodi soluta ex. Ipsa corrupti ipsum sint aut enim eius.', '2020-01-14');
INSERT INTO `posts` (`id`, `author_id`, `title`, `description`, `content`, `date`) VALUES (7, 7, 'Rerum non et accusamus nostrum ab quidem.', 'Harum eius adipisci et et porro. Qui qui natus dolor voluptas. Ut voluptates quibusdam veniam.', 'Qui quia dolores commodi. Et ducimus eos repudiandae cupiditate. Laboriosam deleniti voluptatem et aspernatur. Molestiae praesentium ea sit.', '2008-07-31');
INSERT INTO `posts` (`id`, `author_id`, `title`, `description`, `content`, `date`) VALUES (8, 8, 'Laboriosam sint placeat provident cum.', 'Voluptatibus omnis beatae et. Quos maiores architecto maiores blanditiis. Quod veritatis ipsa eos natus.', 'Veritatis sapiente consequuntur numquam ut pariatur. Alias expedita impedit rerum eaque itaque perferendis optio. Autem velit aut doloremque.', '2010-11-19');
INSERT INTO `posts` (`id`, `author_id`, `title`, `description`, `content`, `date`) VALUES (9, 9, 'Aut eius rerum adipisci vel.', 'Sint quam id eum omnis dignissimos dolorum consectetur magni. Aspernatur ea ullam aut quisquam quae. Velit nam necessitatibus recusandae id. Ratione provident ipsum qui vel non.', 'Illo qui soluta dolorem aliquid eius at reprehenderit excepturi. Omnis ullam tempora quia quas non. Illum adipisci natus reprehenderit nulla quis sed.', '1993-06-11');
INSERT INTO `posts` (`id`, `author_id`, `title`, `description`, `content`, `date`) VALUES (10, 10, 'Eveniet sunt illum eos nihil ut.', 'Odit veniam dolor quibusdam facilis dolorem. Sit ex voluptate veritatis veritatis molestiae quo aut.', 'Non aut praesentium aut recusandae. Tempora ducimus odio eum voluptate molestiae ratione ducimus. Fugit veritatis possimus quia nihil est vitae dolor. Mollitia ex consectetur iure rerum.', '1991-01-16');
