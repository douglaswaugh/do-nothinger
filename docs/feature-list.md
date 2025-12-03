# Feature List

## Phase 1: Execute a script (MVP)
- [ ] **Run simple scripts** - Steps in order, no extra functions (0, 1, or multiple steps)
- [ ] **Run complex scripts** - Steps out of order and/or extra non-step functions present

## Phase 2: Script discovery and templating
- [ ] **List scripts** - Show available scripts in the scripts folder
- [ ] **Create script template** - Generate a new script with N empty steps
- [ ] **Add step template** - Add an empty step at position X in an existing script

## Phase 3: Enhanced execution
- [ ] **Text input capture** - Steps can prompt for and store user input
- [ ] **Context/variables** - Pass data between steps
- [ ] **Windows support** - PowerShell execution on Windows

## Phase 4: Script configuration
- [ ] **Script configuration** - Define required variables for a script
- [ ] **Prompt for variables** - Ask user for values at script start
- [ ] **Persistent config** - Save variable values for reuse
- [ ] **Named profiles** - Multiple configurations per script (dev/staging/prod)

## Phase 5: Reusable steps
- [ ] **Step library** - Store reusable steps separately from scripts
- [ ] **Script composition** - Scripts reference steps by name (YAML manifest)
- [ ] **Shared steps** - Use the same step across multiple scripts

---

## Script Format

Scripts are bash files in a scripts folder, with step functions following the naming convention `step_N_description`:

```bash
#!/bin/bash
# replay-dead-letters.sh

step_1_check_queue_count() {
    az servicebus queue show --name orders-dlq --query messageCount
}

step_2_sample_messages() {
    az servicebus queue peek --name orders-dlq --max-count 5
}

step_3_review_and_decide() {
    echo "Review the messages above."
    echo "Decide if it's safe to replay."
}

step_4_replay_messages() {
    az servicebus queue receive --name orders-dlq | az servicebus queue send --name orders
}
```

**Do Nothinger handles:**
- Discovering step functions by naming convention
- Sorting steps by number
- Converting function names to display text (e.g., `step_1_check_queue_count` → "Check queue count")
- Executing each step and displaying output
- Waiting for user input between steps
- Consistent UI across all scripts

**User handles (in text editor):**
- Creating and editing script content
- Deleting scripts
- Renaming scripts

---

## Example Use Cases

- Dead letter queue handling
- Incident response runbooks
- Kubernetes operations (scale, restart, drain)
- Database maintenance
- Release/deployment process
- Onboarding new team members
- Certificate renewal
- Periodic reporting tasks
