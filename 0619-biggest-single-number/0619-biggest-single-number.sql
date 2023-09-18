# Write your MySQL query statement below
SELECT MAX(num) AS num FROM MyNumbers
WHERE num IN (
    SELECT MAX(n1.num) FROM MyNumbers n1
    GROUP BY n1.num
    HAVING COUNT(n1.num) = 1
)
