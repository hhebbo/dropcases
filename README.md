# DropCases
This is a sample Go application to showcase Dropbox's and HelloSign's APIs.
There is currently the following use case:
- Paper Wiki

# Requirements to get it running
The demo is built using Go. To get the demo running you need:
- Hosting server running a Unix-based OS
- Go v1.12+
- Dropbox Business team with Team Admin access. You'll need the member ID of the admin.
- A Dropbox Business API app with `Team member file access` token
- Define and have access to an internally-open URL to be used for the wiki

# Installation
- Clone the repo `git clone git@github.com:hhebbo/dropcases.git`
- Configure your app by adding your tokens, team admin member ID, and port environment variables. The name of these variable can be found in /apps/appConfig/appConfig.go
- On the root folder of the project run `go run main.go`
- done!

## Paper Wiki
### Why
To utilize Dropbox Paper as an internal wiki represented as a website.

### What
Paper Wiki in this use case can be built as a folder and Paper docs structure within your Dropbox Business account.

### How it works
1. Add a folder called `Paper Wiki` inside your Dropbox root.
2. Add folders inside `Paper Wiki`. Each folder is a main `section` in the wiki. The name of the folder is the name of the section.
3. Add Paper docs inside these folders. These Paper docs are the `pages` inside your wiki. The name of the Paper doc is the name of the page.

###  Where can I see a demo
Here is a link to the demo
