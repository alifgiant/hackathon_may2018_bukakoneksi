CREATE TABLE workspaces (
  id int(11) NOT NULL AUTO_INCREMENT,
  table_id int(11) NOT NULL,
  member_id int(11) NOT NULL,
  position int(11) NOT NULL,
  amenities_tags text NULL,
  PRIMARY KEY (id),
  INDEX position_INDEX (position ASC)
);