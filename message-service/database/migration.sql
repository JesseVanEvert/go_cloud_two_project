/*lecturerID INT FOREIGN KEY REFERENCES lecturer(lecturerID)*/;
CREATE TABLE `message`
(
    messageID   bigint auto_increment,
    lecturerEmail  varchar(255) NOT NULL ,
    content varchar(255) NOT NULL,
    PRIMARY KEY (`messageID`, lecturerEmail)
);

INSERT INTO `message` (lecturerEmail , `content`)
VALUES ('mridulhasan157@gmail.com','I will be absent tomorrow'),
       ('mridulhasan157@gmail.com','I will handout the assignment in upcoming days');

