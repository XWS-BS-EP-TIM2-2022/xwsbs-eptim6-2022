INSERT INTO ROLE (name) VALUES ('ADMIN_ROLE'),('USER_ROLE'),('COMPANY_OWNER_ROLE');

INSERT INTO JOB_POSITION (name, confirmed) VALUES ('Software developer',true),('DevOps Engineer',true),('UI/UX Designer',true),
('Frontend developer',true),('Java developer',true);

INSERT INTO USERS(email, phone, name, password, surname, username, role_id)
	VALUES ('abrkljac9@gmail.com', '0628104308', 'Aleksandar', 'Lozinka123', 'Brki', 'aco11', '3');

INSERT INTO company(email, phone, culture, description, name, status, web, year_of_establishment, owner_id)
	VALUES ('info@onyx.com', '056258484', 'Kultura ne znam ni sam','Opis dodatni kompanije', 'OnyxIt',0, 'http://onyx.com', 2009, 1);