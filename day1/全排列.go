package day1

var (
	res  [][]int
	path []int
	st   []bool // state的缩写
)

func init() {
	res = make([][]int, 0)
	path = make([]int, 0)
	st = make([]bool, 0)
}
func Permute(nums []int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, len(nums))
	st = make([]bool, len(nums))
	dfs(nums, 0)
	return res
}

func dfs(nums []int, cur int) {
	if cur == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}
	for i := 0; i < len(nums); i++ {
		if !st[i] {
			path = append(path, nums[i])
			st[i] = true
			dfs(nums, cur+1)
			st[i] = false
			path = path[:len(path)-1]
		}
	}
}
