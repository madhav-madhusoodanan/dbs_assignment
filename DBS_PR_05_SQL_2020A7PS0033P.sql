create schema ADD_SUB_Window;  
USE ADD_SUB_Window;
create table student (
	studName varchar(30) not null,
    studID int not null unique,
    total_creds int not null,
    avail_creds int not null, 
    
    primary key(studID),
    check (total_creds >= 1),
    check (avail_creds >= 0 and avail_creds <= total_creds),
    check (studID >=0 and studID <= 999)
    );

create table course (
	courseName varchar(30) not null unique,
    courseID int not null unique,
    total_seats int not null,
    avail_seats int not null, 
    course_creds int not null,
    
    primary key(courseID),
    check (total_seats >= 1),
    check (avail_seats >= 0 and avail_seats <= total_seats)
    );

create table takes ( 
		studID int not null, 
		courseID int not null, 
        
        foreign key(studID) references student(studID), 
        foreign key(courseID) references course(courseID), 
        primary key(studID, courseID)
); 

INSERT INTO 
	student(studName, studID, total_creds, avail_creds)
VALUES
	("Mann Shah", 010, 10,7), 
    ("Nandlal Odedara", 012, 10, 5), 
    ("Vinayak Patel",020, 10, 4), 
    ("Nishal Shah", 104, 10, 2), 
	("Nitant Kothari", 420, 10, 1); 

INSERT INTO
	course(courseName, courseID, total_seats, avail_seats, course_creds)
VALUES
	("Database System",213, 6, 4, 4),
    ("Discrete Math", 452 , 4 , 2 , 3),
    ("General Chemistry", 240 ,5 , 2, 3),
    ("Report Writing",111, 7, 5, 2),
    ("Workshop Practice",311, 3, 1, 2);  
    
INSERT INTO 
	takes(studID,courseID)
VALUES
	(010,452), 
    (012,240), 
    (012,111), 
    (020,452),
    (020,240), 
    (104,213), 
    (104,111),
	(104,311),
    (420,213),
    (420,240), 
    (420,311); 

    
DELIMITER $$
CREATE DEFINER=`root`@`localhost` PROCEDURE `addition`(IN tstudID int, IN tcourseID int)
READS SQL DATA
NOT DETERMINISTIC
SQL SECURITY INVOKER
COMMENT 'Input - Student ID and Course ID, Output - Changes made to tables if procedure executes '
BEGIN
	DECLARE availSeats int default 0; 
    DECLARE availCreds int default 0; 
    DECLARE courseCreds int default 0; 
     
	select avail_seats 
    INTO availSeats 
    from course 
    where course.courseID = tcourseID; 
    
    select avail_creds 
    INTO availCreds 
    from student 
    where student.studID = tstudID; 
    
    select course_creds 
    INTO  courseCreds
    from course 
    where course.courseID = tcourseID; 
    
    IF (availSeats > 0 AND availCreds >= courseCreds AND NOT((tstudID,tcourseID) IN (select * from ADD_SUB_Window.takes))) THEN 
		INSERT INTO ADD_SUB_Window.takes VALUES (tstudID, tcourseID); 
        
        UPDATE ADD_SUB_Window.student 
        SET 
			student.avail_creds = student.avail_creds - courseCreds 
		WHERE 
			student.studID = tstudID; 
		
        UPDATE ADD_SUB_Window.course 
        SET 
			course.avail_seats = course.avail_seats - 1 
		WHERE 
			course.courseID = tcourseID; 
		select 'Addition Successful';
        
	ELSE
		select 'Addition not Successful';
	END IF; 
    
END$$ 
DELIMITER ;  

DELIMITER $$
CREATE DEFINER=`root`@`localhost` PROCEDURE `substitution`(IN tstudID int, IN toldCourseID int, IN tnewCourseID int)
READS SQL DATA
NOT DETERMINISTIC
SQL SECURITY INVOKER
COMMENT 'Input - Student ID, Old Course ID and New Course ID, Output - Changes made to tables if procedure executes '
BEGIN 
	DECLARE newCourseCreds int default 0;
    DECLARE oldCourseCreds int default 0;
    DECLARE availSeats int default 0;
    DECLARE availCreds int default 0;
    
    select course_creds 
    INTO  oldCourseCreds
    from course 
    where course.courseID = toldCourseID; 
    
    select course_creds 
    INTO  newCourseCreds
    from course 
    where course.courseID = tnewCourseID; 
    
    select avail_seats 
    INTO availSeats 
    from course 
    where course.courseID = tnewCourseID;
    
    select avail_creds 
    INTO availCreds 
    from student 
    where student.studID = tstudID; 
    
	IF (((tstudID,toldCourseID) IN (select * from ADD_SUB_Window.takes)) AND NOT((tstudID,tnewCourseID) IN (select * from ADD_SUB_Window.takes))) THEN  
		IF((availSeats > 0) and (newCourseCreds <= (oldCourseCreds + availCreds))) THEN
			
            DELETE from ADD_SUB_Window.takes 
            where (takes.studID = tstudID and takes.courseID = toldCourseID);
            
            UPDATE ADD_SUB_Window.student 
			SET 
				student.avail_creds = student.avail_creds + oldCourseCreds 
			WHERE 
				student.studID = tstudID; 
			
			UPDATE ADD_SUB_Window.course 
			SET 
				course.avail_seats = course.avail_seats + 1 
			WHERE 
				course.courseID = toldCourseID;
			
            call addition(tstudID,tnewCourseID);
            select 'Substitution Successful';
        ELSE 
			select 'Substitution not Successful'; 
		END IF;
    ELSE 
		select 'Substitution not Successful';
	END IF;
END$$ 
DELIMITER ;


DELIMITER $$
CREATE DEFINER=`root`@`localhost` PROCEDURE `displayStud`(IN tstudID int)
READS SQL DATA
NOT DETERMINISTIC
SQL SECURITY INVOKER
COMMENT 'Input - Student ID , Output - Print all the courses that a given student has taken '
BEGIN
	select takes.courseID
    from ADD_SUB_Window.takes 
    where takes.studID = tstudID;  
    
    select * 
    from ADD_SUB_Window.student
    where student.studID = tstudID;
    
END$$ 
DELIMITER ;
 

DELIMITER $$
CREATE DEFINER=`root`@`localhost` PROCEDURE `displayCourse`(IN tcourseID int)
READS SQL DATA
NOT DETERMINISTIC
SQL SECURITY INVOKER
COMMENT 'Input - Course ID , Output - Give all details related to a course from course ID '
BEGIN
	select * 
    from ADD_SUB_Window.course 
    where course.courseID = tcourseID;
    
END$$ 
DELIMITER ; 

select * from takes; 
select * from course; 
select * from student; 

start transaction; 
	call addition(010,240); 
commit; 
start transaction; 
	call addition(010,452); 
commit;
start transaction; 
	 call addition(010,213);
commit;
start transaction; 
	 call substitution(010,213,111);
commit;
start transaction; 
	call addition(012,452); 
commit;
start transaction; 
	 call substitution(012,240,213); 
commit;
start transaction; 
	call addition(020,213);
commit; 
start transaction; 
	call substituition(020,240,111); 
commit; 
start transaction; 
	call substituition(020,452,311); 
commit; 
start transaction; 
	call substituition(020,213,452); 
commit; 
start transaction; 
	call addition(104,240); 
commit; 
start transaction; 
	call substitution(104,111,452); 
commit; 
start transaction; 
	call substitution(104,213,240); 
commit;
start transaction; 
	call addition(420,111);
commit; 
start transaction; 
	call substitution(420, 311, 240);
commit;
start transaction; 
	call substitution(420, 311, 213);
commit; 

select * from takes; 
select * from course; 
select * from student;

-- drop table takes; 
-- drop table course;
-- drop table student; 
-- drop database ADD_SUB_Window;