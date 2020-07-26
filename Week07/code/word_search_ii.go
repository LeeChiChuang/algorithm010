package practice

// https://leetcode-cn.com/problems/word-search-ii/
// 212. 单词搜索 II
// 给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
// 同一个单元格内的字母在一个单词中不允许被重复使用。
//
// 示例:
//
// 输入:
// words = ["oath","pea","eat","rain"] and board =
// [
// ['o','a','a','n'],
// ['e','t','a','e'],
// ['i','h','k','r'],
// ['i','f','l','v']
// ]
//
// 输出: ["eat","oath"]
// 说明:
// 你可以假设所有输入都由小写字母 a-z 组成。
//
// 提示:
//
// 你需要优化回溯算法以通过更大数据量的测试。你能否早点停止回溯？
// 如果当前单词不存在于所有单词的前缀中，则可以立即停止回溯。
// 什么样的数据结构可以有效地执行这样的操作？散列表是否可行？为什么？
// 前缀树如何？如果你想学习如何实现一个基本的前缀树，请先查看这个问题： 实现Trie（前缀树）。

type Trie struct {
	children [26]*Trie
	word string
}
var dx [4]int
var dy [4]int
func findWords(board [][]byte, words []string) []string {
	root := &Trie{}
	curr := root
	for _, word := range words {
		for _, ch := range word {
			idx := ch - 'a'
			if curr.children[idx] == nil {
				curr.children[idx] = &Trie{}
			}
			curr = curr.children[idx]
		}
		curr.word = word
	}

	result := make([]string, 0)
	dx = [4]int{-1, 1, 0, 0}
	dy = [4]int{0, 0, -1, 1}
	for i:=0; i<len(board); i++ {
		for j:=0; j<len(board[0]); j++  {
			if root.children[board[i][j]] != nil {
				wordDfs(board, i, j, root, &result)
			}
		}
	}

	return result
}

func wordDfs(board [][]byte, i, j int, curr *Trie, result *[]string) {
	curr = curr.children[board[i][j]]
	if curr.word != "" {
		*result = append(*result, curr.word)
	}
	tmp := board[i][j]
	board[i][j] = '#'
	// 上下左右
	for k:=0; k<4; k++ {
		x, y := i+dx[k], j+dy[k]
		if 0 <= x && x < len(board) && 0<=y && y<len(board[0]) && board[x][y] != '#' && curr.children[board[i][j]] != nil {
			wordDfs(board, x, y, curr, result)
		}
	}
	board[i][j] = tmp
}

