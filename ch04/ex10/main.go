// Ex10は検索語に一致したGitHubイシューの表を1ヶ月未満、1年未満、一年以上の機関で分類して結果を報告します。
package main

import (
	"fmt"
	"log"
	"os"
	"time"
	//"../github"
	"GolangCourse/ch04/github"
)

type IssueTerm int

const (
	withinMonth IssueTerm = iota
	withinYear
	moreThanYear
)

func (term IssueTerm) String() string {
	switch term {
	case withinMonth:
		return "Within One Month"
	case withinYear:
		return "Within One Year"
	case moreThanYear:
		return "More Than A Year"
	}

	return ""
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	issues := make(map[IssueTerm][]*github.Issue)
	for _, item := range result.Items {
		if item.CreatedAt.After(time.Now().AddDate(0, -1, 0)) {
			issues[withinMonth] = append(issues[withinMonth], item)
		} else if item.CreatedAt.Equal(time.Now().AddDate(0, -1, 0)) &&
			item.CreatedAt.Before(time.Now().AddDate(0, -1, 0)) &&
			item.CreatedAt.After(time.Now().AddDate(-1, 0, 0)) {
			issues[withinYear] = append(issues[withinYear], item)
		} else {
			issues[moreThanYear] = append(issues[moreThanYear], item)
		}

	}

	for term, termIssues := range issues {
		fmt.Printf("----- [%s] -----\n", term.String())
		for _, item := range termIssues {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}
