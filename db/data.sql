-- STATES

INSERT INTO states VALUES (10,'SelectDiet');
INSERT INTO states VALUES (12,'SelectDietOrNot');
INSERT INTO states VALUES (14,'SelectedAnythingDiet');
INSERT INTO states VALUES (16,'SelectedVeganDiet');
INSERT INTO states VALUES (18,'SelectedVegetarianDiet');
INSERT INTO states VALUES (19,'DidNotSelectedDiet');

INSERT INTO states VALUES (20,'SelectSensitivity');
INSERT INTO states VALUES (22,'SelectSensitivityOrNot');
INSERT INTO states VALUES (24,'SelectedGlutenSensitivity');
INSERT INTO states VALUES (26,'SelectedMilkSensitivity');
INSERT INTO states VALUES (28,'SelectedNoSensitivity');
INSERT INTO states VALUES (29,'NotSelectedSensitivity');

INSERT INTO states VALUES (30,'SelectTheme');
INSERT INTO states VALUES (32,'SelectThemeOrNot');
INSERT INTO states VALUES (34,'SelectedAsianTheme');
INSERT INTO states VALUES (36,'SelectedMoroccanTheme');
INSERT INTO states VALUES (38,'SelectedMoroccasianTheme');
INSERT INTO states VALUES (39,'NotSelectedTheme');

-- DIETS

INSERT INTO diets VALUES(NULL, 0, "anything");
INSERT INTO diets VALUES(NULL, 10, "vegan");
INSERT INTO diets VALUES(NULL, 20, "vegetarian");

-- SENSITIVITIES

INSERT INTO sensitivities VALUES(NULL, 0, "no");
INSERT INTO sensitivities VALUES(NULL, 10, "gluten");
INSERT INTO sensitivities VALUES(NULL, 20, "milk");

-- THEMES

INSERT INTO themes VALUES(NULL, 10, "asian");
INSERT INTO themes VALUES(NULL, 20, "moroccan");
INSERT INTO themes VALUES(NULL, 30, "moroccasian");

-- COURSES

INSERT INTO courses VALUES(NULL, 'דג מרוקאי','', 'https://s30.postimg.org/49fugw4z5/fish.jpg');
INSERT INTO courses VALUES(NULL, 'דג בקארי','', 'https://s28.postimg.org/u1hmefp1p/malai_not_grained.jpg');
INSERT INTO courses VALUES(NULL, 'קובה בקרס קוקוס וקארי','', 'https://s23.postimg.org/ccl4mikfv/kuba1.jpg');

-- METADATA_TYPES

INSERT INTO metadata_types VALUES(NULL, 'diet');
INSERT INTO metadata_types VALUES(NULL, 'sensitivity');
INSERT INTO metadata_types VALUES(NULL, 'theme');
