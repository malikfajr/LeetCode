CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  DECLARE LINE INT ;
  SET LINE = N - 1 ;
  RETURN (
      # Write your MySQL query statement below.      
      SELECT  salary FROM Employee 
      GROUP BY salary
      ORDER BY salary DESC
      LIMIT 1 
      OFFSET LINE
  );
END