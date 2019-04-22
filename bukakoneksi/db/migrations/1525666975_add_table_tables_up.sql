CREATE TABLE tables (
  id int(11) NOT NULL AUTO_INCREMENT,
  office_floor_id int(11) NOT NULL,
  name varchar(100) NOT NULL,
  status varchar(50) NOT NULL,
  workspace_size int(11) NOT NULL,
  empty_workspace int(11) NOT NULL,
  point_x float NULL,
  point_y float NULL,
  PRIMARY KEY (id),
  CONSTRAINT table_office_floor_id_fk FOREIGN KEY (office_floor_id)
    REFERENCES office_floors(id),
  INDEX name_INDEX (name ASC),
  INDEX status_INDEX (status ASC),
  INDEX empty_workspace_INDEX (empty_workspace ASC)
);