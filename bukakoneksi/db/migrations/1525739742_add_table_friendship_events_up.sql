CREATE TABLE friendship_events (
  id int(11) NOT NULL AUTO_INCREMENT,
  event varchar(20) NOT NULL,
  telegram_user_id varchar(100) NOT NULL,
  telegram_username varchar(100) NOT NULL,
  city varchar(100) NOT NULL,
  day varchar(10) NOT NULL,
  PRIMARY KEY (id),
  INDEX search_INDEX (city, event, day),
  UNIQUE INDEX strict_UNIQUE (event, day, telegram_user_id)
);