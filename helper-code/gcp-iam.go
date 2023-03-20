package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	admin "google.golang.org/api/admin/directory/v1"
	crm "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	// Authenticate with the service account
	creds, err := google.FindDefaultCredentials(ctx, admin.AdminDirectoryUserScope, admin.AdminDirectoryGroupScope, admin.AdminDirectoryGroupMemberScope, crm.CloudPlatformScope)
	if err != nil {
		log.Fatalf("Failed to create credentials: %v", err)
	}

	// Initialize the Admin SDK Directory API
	adminService, err := admin.NewService(ctx, option.WithCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to create admin service: %v", err)
	}

	// Create a user
	user := &admin.User{
		PrimaryEmail: "user@example.com",
		Name: &admin.UserName{
			GivenName:  "User",
			FamilyName: "Example",
		},
		Password: "your_password_here",
	}

	user, err = adminService.Users.Insert(user).Do()
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	// Create a group
	group := &admin.Group{
		Email: "group@example.com",
		Name:  "Example Group",
	}

	group, err = adminService.Groups.Insert(group).Do()
	if err != nil {
		log.Fatalf("Failed to create group: %v", err)
	}

	// Add user to the group
	member := &admin.Member{
		Email: user.PrimaryEmail,
		Role:  "MEMBER",
	}

	_, err = adminService.Members.Insert(group.Id, member).Do()
	if err != nil {
		log.Fatalf("Failed to add user to the group: %v", err)
	}

	// Assign a role to the user
	projectID := os.Getenv("GCP_PROJECT_ID")
	role := "roles/editor" // Example role

	crmService, err := crm.NewService(ctx, option.WithCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to create cloud resource manager service: %v", err)
	}

	policy, err := crmService.Projects.GetIamPolicy(projectID, &crm.GetIamPolicyRequest{}).Do()
	if err != nil {
		log.Fatalf("Failed to get IAM policy: %v", err)
	}

	binding := &crm.Binding{
		Role:    role,
		Members: []string{fmt.Sprintf("user:%s", user.PrimaryEmail)},
	}

	policy.Bindings = append(policy.Bindings, binding)

	_, err = crmService.Projects.SetIamPolicy(projectID, &crm.SetIamPolicyRequest{
		Policy: policy,
	}).Do()
	if err != nil {
		log.Fatalf("Failed to set IAM policy: %v", err)
	}

	fmt.Println("User, group, and role assignment completed")
}
