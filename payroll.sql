use empdb;
CREATE TABLE Payroll (
    payroll_id INT AUTO_INCREMENT PRIMARY KEY,
    employee_id INT,
    salary DECIMAL(10, 2),
    pay_date DATE,
    taxes DECIMAL(8, 2),
    bonuses DECIMAL(8, 2),
    net_pay DECIMAL(10, 2),
    FOREIGN KEY (employee_id) REFERENCES Employee(emp_id)
);
INSERT INTO Payroll (
        employee_id,
        salary,
        pay_date,
        taxes,
        bonuses,
        net_pay
    )
VALUES (
        101,
        5000.00,
        '2023-01-31',
        800.00,
        200.00,
        4400.00
    );
INSERT INTO Payroll (
        employee_id,
        salary,
        pay_date,
        taxes,
        bonuses,
        net_pay
    )
VALUES (
        102,
        6000.00,
        '2023-01-31',
        1200.00,
        300.00,
        5100.00
    );
INSERT INTO Payroll (
        employee_id,
        salary,
        pay_date,
        taxes,
        bonuses,
        net_pay
    )
VALUES (
        103,
        5500.00,
        '2023-01-31',
        900.00,
        100.00,
        4600.00
    );
INSERT INTO Payroll (
        employee_id,
        salary,
        pay_date,
        taxes,
        bonuses,
        net_pay
    )
VALUES (
        104,
        5200.00,
        '2023-01-31',
        1000.00,
        250.00,
        4450.00
    );
INSERT INTO Payroll (
        employee_id,
        salary,
        pay_date,
        taxes,
        bonuses,
        net_pay
    )
VALUES (
        105,
        5800.00,
        '2023-01-31',
        1100.00,
        150.00,
        4850.00
    );
INSERT INTO Payroll (
        employee_id,
        salary,
        pay_date,
        taxes,
        bonuses,
        net_pay
    )
VALUES (
        106,
        5100.00,
        '2023-01-31',
        950.00,
        180.00,
        4370.00
    );
INSERT INTO Payroll (
        employee_id,
        salary,
        pay_date,
        taxes,
        bonuses,
        net_pay
    )
VALUES (
        107,
        5900.00,
        '2023-01-31',
        1150.00,
        200.00,
        4950.00
    );
INSERT INTO Payroll (
        employee_id,
        salary,
        pay_date,
        taxes,
        bonuses,
        net_pay
    )
VALUES (
        108,
        5300.00,
        '2023-01-31',
        980.00,
        120.00,
        4420.00
    );
INSERT INTO Payroll (
        employee_id,
        salary,
        pay_date,
        taxes,
        bonuses,
        net_pay
    )
VALUES (
        109,
        5600.00,
        '2023-01-31',
        1050.00,
        220.00,
        4830.00
    );
INSERT INTO Payroll (
        employee_id,
        salary,
        pay_date,
        taxes,
        bonuses,
        net_pay
    )
VALUES (
        110,
        5400.00,
        '2023-01-31',
        1020.00,
        180.00,
        4600.00
    );
commit;