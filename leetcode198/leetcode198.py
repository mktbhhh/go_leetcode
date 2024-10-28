from typing import List


class Solution:
    def rob(self, nums: List[int]) -> int:
        n = len(nums)

        if n <= 0:
            return 0
        if n == 1:
            return nums[0]

        dp = [[0, 0] for _ in range(n)]
        dp[0][0] = 0
        dp[0][1] = nums[0]

        for i in range(1, n):
            dp[i][1] = dp[i - 1][0] + nums[i]
            dp[i][0] = max(dp[i - 1][1], dp[i - 1][0])

        return max(dp[n - 1][0], dp[n - 1][1])


nums = [1, 2, 3, 1]
s = Solution()
print(s.rob(nums))
