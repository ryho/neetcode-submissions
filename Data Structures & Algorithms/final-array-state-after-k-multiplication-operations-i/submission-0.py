class Solution:
    def getFinalState(self, nums: List[int], k: int, multiplier: int) -> List[int]:
        for _ in range(k):
            minVal = nums[0]
            minIndex = 0
            for i,val in enumerate(nums):
                if val<minVal:
                    minVal=val
                    minIndex=i
            nums[minIndex]=minVal*multiplier
        return nums