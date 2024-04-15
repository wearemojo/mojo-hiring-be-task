# Backend engineering task

It is crucial you read the "Your mission" section in full before you start. We recommend re-reading it before submitting.

Please do not release your solution publicly on GitHub. Instead, clone it and push it to your own private repository. Please attempt to maintain the Git history so we can quickly and easily see your changes. When you are ready to submit, please add the following users as collaborators to your repository, or download it as a ZIP for submission.

- https://github.com/DuffleOne
- https://github.com/billinghamj
- https://github.com/clem109

## Introduction

Hello and thank you for the time to contribute to the Mojo take home test! In this exercise we have a simple HTTP service that has a job to track sessions. A session represents a user on a device. The service has a few endpoints that allow us to query the sessions.

### `/list_by_platform`

This endpoint is given a platform, and returns a list of sessions associated with that platform.

### `/list_by_device_id`

This endpoint is given a device ID, and returns a list of sessions associated with that device.

## Your mission

### Part 1

The goal is to add a new endpoint that will return a list of all the sessions associated with a specific user ID. You will need to create the endpoint and register it within the service.

### Part 2

We now have 3 endpoints that all do very similiar things. What would it look like if we had just 1 endpoint with a flexible filter and pagination? Without changing the service code, what would the request and response structs look like? Write out just the two structs as you would want it for that future endpoint, along with a JSON schema for validating the input.

## Requirements & grading criteria

- We will look for code that is minimal, clean, and readable.
- We are keen on consistency and attention to detail.
- We are looking for code that stylistically matches the rest of the codebase.
- We are looking for your code to build if we run `go build ./...`
- We are looking for you to spend no more than 20 to 30 minutes on this task.

- We are not looking for any testing, or automation or proof of correctness.
- Do not worry about authentication.
- Do not worry about marshalling structs or the HTTP transport layer. If you want to know more about how that works, review the CRPC package outside of the 20 minute allotted time.
