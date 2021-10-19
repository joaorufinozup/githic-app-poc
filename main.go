// https://www.youtube.com/watch?v=iaBEWB1As0k

package main

import (
	"context"
	"strings"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/google/go-github/github"
	"github.com/swinton/go-probot/probot"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	probot.HandleEvent("pull_request", func(ctx *probot.Context) error {
		event := ctx.Payload.(*github.PullRequestEvent)
		pr := event.GetPullRequest()
		repo := event.GetRepo()
		fmt.Printf("Got PR title: %s\n", *pr.Title)
		
		newReview := &github.PullRequestReviewRequest{CommitID: pr.Head.SHA}

		if strings.Contains(*pr.Title, "Ritchie") {
			newReview.Event = github.String("APPROVE")
			newReview.Body = github.String("LGTM")
		} else {
			newReview.Event = github.String("REQUEST_CHANGES")
			newReview.Body = github.String("This is not a Ritchie PR")
		}

		review, _, err := ctx.GitHub.PullRequests.CreateReview(context.Background(), *repo.Owner.Login, *repo.Name, pr.GetNumber(), newReview)

		if err != nil {
			return err
		}

		fmt.Printf("New review created %+v", review)

		return nil
	})

	probot.HandleEvent("repository", func(ctx *probot.Context) error {
		event := ctx.Payload.(*github.RepositoryEvent)
		action := event.GetAction()

		if action == "created" {
			repo := event.GetRepo()
			fmt.Printf("New repo created: %s", repo.GetName())
		} else {
			fmt.Printf("Repository action: %s", action);
		}

		return nil
	})

	probot.Start()
}