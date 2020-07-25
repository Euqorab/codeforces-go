// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`7`, `[[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]]`, `"abaedcd"`, 
			`[2,1,1,1,1,1,1]`,
		},
		{
			`4`, `[[0,1],[1,2],[0,3]]`, `"bbbb"`, 
			`[4,2,1,1]`,
		},
		{
			`5`, `[[0,1],[0,2],[1,3],[0,4]]`, `"aabab"`, 
			`[3,2,1,1,1]`,
		},
		{
			`6`, `[[0,1],[0,2],[1,3],[3,4],[4,5]]`, `"cbabaa"`, 
			`[1,2,1,1,2,1]`,
		},
		{
			`7`, `[[0,1],[1,2],[2,3],[3,4],[4,5],[5,6]]`, `"aaabaaa"`, 
			`[6,5,4,1,3,2,1]`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, countSubTrees, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-198/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/