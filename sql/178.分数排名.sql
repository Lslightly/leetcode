--
-- @lc app=leetcode.cn id=178 lang=mysql
--
-- [178] 分数排名
--
-- @lc code=start
SELECT e1.score, (
        SELECT COUNT(DISTINCT e2.score)
        FROM Scores e2
        WHERE e1.score < e2.score
    ) + 1 as 'rank'
FROM Scores e1
ORDER BY e1.score DESC
-- @lc code=end

