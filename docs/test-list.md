# Test List

## Run simple scripts
- [x] Handle a script with 0 steps - display "Done"
- [ ] Handle a script with 1 empty step - find function, execute, wait for Enter, display "Done"
- [ ] Multiple empty steps with functions in order, no non-step functions

## Run complex scripts
- [ ] Multiple empty steps with functions in different order, no non-step functions
- [ ] Multiple empty steps with functions in order, with non-step functions
- [ ] Multiple empty steps with functions in different order, with non-step functions

## Execution
- [ ] Single step script that actually does something - execute and display output

## Errors
- [ ] Handle script file not found
- [ ] Handle step failed (non-zero exit code)

## CLI
- [ ] Wire up CLI: `do-nothinger run <script-path>`
