#
# @lc app=leetcode.cn id=37 lang=python3
#
# [37] 解数独
#

# @lc code=start
import subprocess
subprocess.run(['pip', 'install', 'z3-solver'], check=True, shell=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
from z3 import *
from typing import List
class Solution:
    def solveSudoku(self, board: List[List[str]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        X = [
            [Int(f'x_{i}_{j}') for j in range(9)]
            for i in range(9)
        ]
        cells = [
            And(1 <= X[i][j], X[i][j] <= 9)
            for i in range(9) for j in range(9)
        ]
        rows = [
            Distinct(X[i]) for i in range(9)
        ]
        columns = [
            Distinct([X[i][j] for i in range(9)]) for j in range(9)
        ]
        sq = [
            Distinct([X[3*i0+i][3*j0+j]
                      for i in range(3) for j in range(3)])
            for i0 in range(3) for j0 in range(3)
        ]
        sudoku = cells + rows + columns + sq
        board = self.board_str2int(board)
        board_c = [
            If(board[i][j] == 0, True, X[i][j] == board[i][j])
            for i in range(9) for j in range(9)
        ]
        s = Solver()
        s.add(sudoku + board_c)
        if s.check() == sat:
            m = s.model()
            board = [
                [str(m[X[i][j]]) for j in range(9)]
                for i in range(9)
            ]

    
    def board_str2int(self, board: List[List[str]]) -> List[List[int]]:
        X = [[0 for j in range(9)] for i in range(9)]
        for i in range(9):
            for j in range(9):
                if board[i][j] != '.':
                    X[i][j] = int(board[i][j])
        return X
                    
        
# @lc code=end

