CREATE TABLE offices (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(100) NOT NULL,
  address varchar(255) NOT NULL,
  city varchar(100) NOT NULL,
  location_url text NOT NULL,
  PRIMARY KEY (id),
  INDEX city_INDEX (city ASC),
  INDEX name_INDEX (name ASC)
);