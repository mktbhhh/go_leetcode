import collections
from typing import List


class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        def bfs(i, j):
            grid[i][j] = "0"

            if i - 1 >= 0 and grid[i - 1][j] == "1":
                bfs(i - 1, j)
            if i + 1 < nRow and grid[i + 1][j] == "1":
                bfs(i + 1, j)
            if j - 1 >= 0 and grid[i][j - 1] == "1":
                bfs(i, j - 1)
            if j + 1 < nCol and grid[i][j + 1] == "1":
                bfs(i, j + 1)

        nRow = len(grid)
        if nRow == 0:
            return 0

        nCol = len(grid[0])

        if nCol == 0:
            return 0

        ans = 0

        for i in range(nRow):
            for j in range(nCol):
                if grid[i][j] == "1":
                    ans += 1
                    bfs(i, j)

        return ans

    def numIslands1(self, grid: List[List[str]]) -> int:
        nr = len(grid)
        if nr == 0:
            return 0
        nc = len(nr[0])

        num_islands = 0
        for r in nr:
            for c in nc:
                if grid[r][c] == "1":
                    num_islands += 1
                    grid[r][c] == "0"
                    neighbors = collections.deque([r, c])
                    while neighbors:
                        row, col = neighbors.popleft()
                        for x, y in [
                            (row - 1, col),
                            (row + 1, col),
                            (row, col - 1),
                            (row, col + 1),
                        ]:
                            if 0 <= x < nr and 0 <= y < nc and grid[x][y] == "1":
                                neighbors.append((x, y))
                                grid[x][y] = "0"

        return num_islands
