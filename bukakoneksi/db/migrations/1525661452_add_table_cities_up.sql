CREATE TABLE cities (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(100) NOT NULL,
  PRIMARY KEY (id),
  INDEX name_INDEX (name ASC)
);