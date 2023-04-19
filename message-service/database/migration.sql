/*lecturerID INT FOREIGN KEY REFERENCES lecturer(lecturerID)*/;
CREATE TABLE `message`
(
    messageID   bigint auto_increment,
    lecturerEmail  varchar(255) NOT NULL ,
    toEmail varchar (255) Not Null,
    content varchar(255) NOT NULL,
    PRIMARY KEY (`messageID`, lecturerEmail)
);

INSERT INTO `message` (lecturerEmail , `content`, toEmail)
VALUES ('mridulhasan157@gmail.com','I will be absent tomorrow', 'abs@gmail.com'),
       ('mridulhasan157@gmail.com','I will handout the assignment in upcoming days', 'djdsj@gmail.com');

