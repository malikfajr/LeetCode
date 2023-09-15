# Write your MySQL query statement below
SELECT salary as SecondHighestSalary
    FROM Employee 
    UNION (SELECT null)
ORDER BY SecondHighestSalary DESC
LIMIT 1, 1



