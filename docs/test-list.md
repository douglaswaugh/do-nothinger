# Test List

## Run simple scripts
- [x] Handle a script with 0 steps - display "Done"
- [x] Hook up run function to main so it runs when the application runs
- [x] Script with 1 step that outputs something - assert output contains what the step echoes
- [x] Display "Press Enter to continue..." before waiting for input
- [x] Wait for Enter before showing "Done"
- [x] Display step name (e.g., "Step 1: Do something")
- [x] Display step name for a differently named step (requires parsing)
- [x] Display "Step complete" after Enter
- [x] Multiple steps with functions in order, no non-step functions

## Execution
- [x] Single step script that actually does something - execute and display output

## Run complex scripts
- [x] Multiple empty steps with functions in different order, no non-step functions
- [ ] Multiple empty steps with functions in order, with non-step functions
- [ ] Multiple empty steps with functions in different order, with non-step functions

## Errors
- [ ] Handle script file not found
- [ ] Handle empty script file path
- [ ] Handle step failed (non-zero exit code)

## CLI
- [ ] Wire up CLI: `do-nothinger run <script-path>`
