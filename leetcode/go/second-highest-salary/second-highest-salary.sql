# Write your MySQL query statement below
SELECT
	MAX(salary) as SecondHighestSalary
FROM
	Employee as e1
WHERE 
    salary < (SELECT MAX(salary) FROM Employee)
