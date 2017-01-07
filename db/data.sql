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
INSERT INTO states VALUES (40,'ThemeSelected');

-- DIETS

INSERT INTO diets VALUES(NULL, 0, "anything");
INSERT INTO diets VALUES(NULL, 10, "vegan");
INSERT INTO diets VALUES(NULL, 20, "vegetarian");
