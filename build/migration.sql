CREATE TABLE `fund`
(
    id bigint auto_increment,
    name varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `fund` (`name`)
VALUES ('CushionEquitiesFund');

CREATE TABLE `customer`
(
    id bigint auto_increment,
    name varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `customer` (`name`)
VALUES ('Jonathan'), ('David');

CREATE TABLE `customeraccount`
(
    id bigint auto_increment,
    customerID bigint,
    investmentFundID bigint,
    investmentAmount float,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`investmentFundID`) REFERENCES fund(id),
    FOREIGN KEY (`customerID`) REFERENCES customer(id)
);

INSERT INTO `customeraccount` (`customerID`, `investmentFundID`, `investmentAmount`)
VALUES (1, 1, 1.50);
