CREATE TABLE members (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(100) NOT NULL,
  telegram_user_id varchar(100) NOT NULL,
  telegram_username varchar(100) NOT NULL,
  picture_image varchar(255) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE INDEX telegram_user_id_UNIQUE (telegram_user_id ASC),
  INDEX telegram_username_INDEX (telegram_user_id ASC),
  INDEX name_INDEX (name ASC)
);