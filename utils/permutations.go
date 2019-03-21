package utils

func ApplyPermutations(original []string, permutations [][]int) (result [][]string) {
  for _, v := range(permutations) {
    result = append(result, applyPermutation(original, v))
  }
  return 
}

// For all permutations of N numbers, you can add another number and insert it at every point to get the list of permutions.
// So we start with 1 number, and bootstrap from there
func GeneratePermutations(permutation_size int) (result [][]int) {
  if permutation_size < 1 {
    return [][]int{}
  }

  result = [][]int{{0}} 
  for i := 1; i < permutation_size; i++ {
    var joins [][]int
    for j := 0; j < len(result); j++ {
      joins = append(joins, joinPermutation(result[j], i)...)
    }
    result = joins
  }
  return
}

func applyPermutation(original []string, permutation []int) (result []string) {
  if len(permutation) == len(original) {
    // TODO: How should this be handled?
    return original
  }

  for i, v := range(permutation) {
    result[i] = original[v]
  }
  return
}

func joinPermutation(permutation []int, new_member int) (result [][]int) {
  for i := 0; i <= len(permutation); i++ {
    new_permutation := copyAndJoin(permutation[:i], new_member, permutation[i:])
    result = append(result, new_permutation)
  }
  return
}

func copyAndJoin(arr1[]int, new_member int, arr2 []int) (arr []int) {
  tmp1 := make([]int, len(arr1))
  tmp2 := make([]int, len(arr2))
  copy(tmp1, arr1)
  copy(tmp2, arr2)
  arr = append(tmp1, new_member)
  arr = append(arr, tmp2...)
  return 
}
