package main

import (
	"fmt"
	"log"

	"github.com/xanzy/go-gitlab"
)

func main() {

	git, err := gitlab.NewClient(T, gitlab.WithBaseURL(U))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	opt := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 200,
			Page:    1,
		},
	}
	// users, _, err := git.Users.ListUsers(&gitlab.ListUsersOptions{})
	// if err != nil {
	// 	log.Fatalf("Hata var %v", err)
	// }

	// empJSON, err := json.MarshalIndent(users, "", "  ")
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// fmt.Printf("MarshalIndent funnction output\n %s\n", string(empJSON))

	// projects, _, err := git.Projects.ListProjects(&gitlab.ListProjectsOptions{})
	// if err != nil {
	// 	log.Fatalf("Hata var %v", err)
	// }

	// bosJson, err := json.MarshalIndent(projects, "", "  ")
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// // if(projects.runnerToken)
	// fmt.Printf("%+v", string(bosJson))

	for {
		// Get the first page with projects.
		ps, resp, err := git.Projects.ListProjects(opt)
		if err != nil {
			log.Fatal(err)
		}

		// List all the projects we've found so far.
		for _, p := range ps {
			fmt.Printf("Found project: %s  ve shareRunner enable mi ? %s\n", p.Name, p.ID)
		}

		// Exit the loop when we've seen all pages.
		if resp.NextPage == 0 {
			break
		}

		// Update the page number to get the next page.
		opt.Page = resp.NextPage
	}

}
