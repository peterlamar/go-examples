DROP TABLE mytable;
CREATE TABLE mytable(
   PartRequestNumber VARCHAR(14) NOT NULL PRIMARY KEY
  ,PartNumber        INTEGER  NOT NULL
  ,FulfillLocation   VARCHAR(10) NOT NULL
  ,RequestLocation   VARCHAR(8) NOT NULL
  ,TrackingNumber    INTEGER  NOT NULL
  ,RONumber          VARCHAR(16) NOT NULL
);
INSERT INTO mytable(PartRequestNumber,PartNumber,FulfillLocation,RequestLocation,TrackingNumber,RONumber) VALUES ('PRNC1234567890',1234567,'Warehouse','Boston',735514528714,'RONC9898044240');
INSERT INTO mytable(PartRequestNumber,PartNumber,FulfillLocation,RequestLocation,TrackingNumber,RONumber) VALUES ('PRNC1234567891',1234568,'Warehouse','Orlando',745514528715,'RONC9822001240');
INSERT INTO mytable(PartRequestNumber,PartNumber,FulfillLocation,RequestLocation,TrackingNumber,RONumber) VALUES ('PRNC1234567892',1234569,'Warehouse','Tampa',755514528716,'RONC1198001240');
INSERT INTO mytable(PartRequestNumber,PartNumber,FulfillLocation,RequestLocation,TrackingNumber,RONumber) VALUES ('PRNC1234567893',1234570,'Warehouse','Montreal',715514528717,'RONC6614000727');
INSERT INTO mytable(PartRequestNumber,PartNumber,FulfillLocation,RequestLocation,TrackingNumber,RONumber) VALUES ('PRNC1234567894',1234571,'Warehouse','Fremont',725514528718,'RONC4514000727');
INSERT INTO mytable(PartRequestNumber,PartNumber,FulfillLocation,RequestLocation,TrackingNumber,RONumber) VALUES ('PRNC1234567895',1234572,'Warehouse','Atlanta',735514528719,'RONC2347013607');
INSERT INTO mytable(PartRequestNumber,PartNumber,FulfillLocation,RequestLocation,TrackingNumber,RONumber) VALUES ('PRNC1234567896',1234573,'Warehouse','Boston',735514528720,'RONC9641113607');
INSERT INTO mytable(PartRequestNumber,PartNumber,FulfillLocation,RequestLocation,TrackingNumber,RONumber) VALUES ('PRNC1234567946',1234623,'Warehouse','Tampa',755514528770,'RONC9711111364');
INSERT INTO mytable(PartRequestNumber,PartNumber,FulfillLocation,RequestLocation,TrackingNumber,RONumber) VALUES ('PRNC1234567949',1234626,'Warehouse','Raleigh',765514528773,'RONC9868335527');
