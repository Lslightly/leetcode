--
-- @lc app=leetcode.cn id=175 lang=mysql
--
-- [175] 组合两个表
--

-- @lc code=start
SELECT Person.FirstName, Person.LastName, Address.City, Address.State
FROM Person
LEFT OUTER JOIN Address on Person.PersonId = Address.PersonId

-- @lc code=end

