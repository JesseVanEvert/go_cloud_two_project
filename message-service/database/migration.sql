/*lecturerID INT FOREIGN KEY REFERENCES lecturer(lecturerID)*/;
CREATE TABLE `message`
(
    messageID   bigint auto_increment,
    lecturerID  bigint,
    content varchar(255) NOT NULL,
    PRIMARY KEY (`messageID`, `lecturerID`)
);

INSERT INTO `message` (`lecturerID`,`content`)
VALUES (3,'I will be absent tomorrow'),
       (4,'I will handout the assignment in upcoming days');

/*CREATE TABLE `classMessage`
(
    classID   bigint auto_increment,
    PRIMARY KEY (`classID`)
    messageID INT FOREIGN KEY REFERENCES message(messageID)
);*/
