CREATE TABLE office_floors (
  id int(11) NOT NULL AUTO_INCREMENT,
  office_id int(11) NOT NULL,
  name varchar(100) NOT NULL,
  floor_image varchar(100) NOT NULL,
  CONSTRAINT office_id_fk FOREIGN KEY (office_id)
    REFERENCES offices(id),
  PRIMARY KEY (id),
  INDEX name_INDEX (name ASC)
);