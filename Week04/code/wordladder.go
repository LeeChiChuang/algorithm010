package main

import "fmt"

// 127. 单词接龙
// 给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：
//
// 每次转换只能改变一个字母。
// 转换过程中的中间单词必须是字典中的单词。
// 说明:
//
// 如果不存在这样的转换序列，返回 0。
// 所有单词具有相同的长度。
// 所有单词只由小写字母组成。
// 字典中不存在重复的单词。
// 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
// 示例 1:
//
// 输入:
// beginWord = "hit",
// endWord = "cog",
// wordList = ["hot","dot","dog","lot","log","cog"]
//
// 输出: 5
//
// 解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
// 返回它的长度 5。
// 示例 2:
//
// 输入:
// beginWord = "hit"
// endWord = "cog"
// wordList = ["hot","dot","dog","lot","log"]
//
// 输出: 0
//
// 解释: endWord "cog" 不在字典中，所以无法进行转换。

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	ret := ladderLength(beginWord, endWord, wordList)
	fmt.Println(ret)
}

// BFS Time Complexity: O(n*26^l), l = len(word), n=|wordList| Space Complexity: O(n)
func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict := make(map[string]bool) // 把word存入字典
	for _, word := range wordList {
		dict[word] = true // 可以利用字典快速添加、删除和查找单词
	}
	if _, ok := dict[endWord]; !ok {
		return 0
	}
	// queue := []string{beginWord} 定义辅助队列
	var queue []string
	queue = append(queue, beginWord)

	l := len(beginWord)
	steps := 0

	for len(queue) > 0 {
		steps++
		size := len(queue)
		for i := size; i > 0; i-- { // 当前层级节点
			s := queue[0] // 原始单词
			queue = queue[1:]
			chs := []rune(s)
			for i := 0; i < l; i++ { // 对单词的每一位进行修改
				ch := chs[i]                  // 对当前单词的一位
				for c := 'a'; c <= 'z'; c++ { // 尝试从a-z
					if c == ch { // 跳过本身比如hot修改为hot
						continue
					}
					chs[i] = c
					t := string(chs)
					if t == endWord { // 找到结果
						return steps + 1
					}
					if _, ok := dict[t]; !ok { // 不在dict中，跳过
						continue
					}
					delete(dict, t)          // 从字典中删除该单词，因为已经访问过，若重复访问路径一定不是最短的
					queue = append(queue, t) // 将新的单词添加到队列
				}
				chs[i] = ch // 单词的第i位复位，再进行下面的操作
			}
		}
	}
	return 0
}
