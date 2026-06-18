# Do Nothinger

A tool for running [do-nothing scripts](https://blog.danslimmon.com/2019/07/15/do-nothing-scripting-the-key-to-gradual-automation/) — a pattern by Dan Slimmon for gradually automating operational processes.

## What is a do-nothing script?

A do-nothing script is a bash script that breaks a process into named steps. Initially each step just prompts you to perform the action manually. Over time you replace the manual prompts with real automation, one step at a time, until the whole process runs unattended.

## Script format

Scripts are bash files with step functions named `step_N_description`:

```bash
#!/bin/bash

step_1_check_queue_count() {
    az servicebus queue show --name orders-dlq --query messageCount
}

step_2_review_and_decide() {
    echo "Review the messages above and decide if it's safe to replay."
}

step_3_replay_messages() {
    az servicebus queue receive --name orders-dlq | az servicebus queue send --name orders
}
```

Do Nothinger discovers the step functions by naming convention, runs them in order, and waits for you to press Enter between each one.

## Prerequisites

- [Go](https://go.dev/dl/) 1.25.4 or later

## Usage

```
go run . <path-to-script>
```

For each step it will:
1. Print the step name (e.g. `Step 1: Check queue count`)
2. Execute the step function and show its output
3. Wait for you to press Enter
4. Confirm the step is complete before moving on

## Use cases

- Incident response runbooks
- Dead letter queue handling
- Deployment and release processes
- Database maintenance
- Kubernetes operations
- Certificate renewal
- Onboarding new team members
