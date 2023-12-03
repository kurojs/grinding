func search(nums []int, target int) int {
  return binarySearch(nums, 0, len(nums) - 1, target)
}

func binarySearch(nums []int, left, right, target int) int {
  if left > right {
    return -1
  }

  pivot := (left + right) / 2
  if nums[pivot] == target {
    return pivot
  }

  if nums[pivot] > target {
    return binarySearch(nums, left, pivot - 1, target)
  }

  return binarySearch(nums, pivot + 1, right, target)
}