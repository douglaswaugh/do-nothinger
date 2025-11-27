Do Nothinger
===

This application is the Do Nothinger.  It is designed to create, manage, and run [do nothing scripts](https://blog.danslimmon.com/2019/07/15/do-nothing-scripting-the-key-to-gradual-automation/), which were written about by Dan Slimmon.

Do nothing scripts are scripts which start with just has a bunch of steps you need to do manually.  As you complete each manual step, you press a key to confirm the task has been completed.  Over time you automate each of the steps so that, in the end, you have a fully automated process which can run from start to finish without intervention.

This project is written in Golang, which I have never used to write an application.

This project will be written using test-driven development.

1. We will keep a test list as a physical file to use as a to do list which we will develop together at the beginning of development of a new feature or bug fix
2. Before any new behaviour is added, I will ask you to write the next unimplemented test from the test list file, which should fail because the behaviour has yet to be implemented, and then await my instruction
3. When the failing test has been implemented I will run the test to make sure it fails for the correct reason
4. Once I confirm the test is failing for the correct reason I will ask you to implement the code which will make the test pass, and then await my instruction
5. I will then run the tests to make sure they pass, and I will review the code
6. If I am happy I will ask you to make that test as complete on the test list file
7. Whenever all the tests are passing, I will look to see if there is an opportunity to refactor anything
8. Once any such refactoring is complete, I will ask you to write the next failing test from the test list file, and then await my instruction

It is vital that you don't skip ahead in the process.  Do not write the implementation which would make a test pass until I have checked the test fails for the correct reason.  Do not start writing the next failing test before I have told you that I am happy with the implementation for the last test.  Always wait for my instruction after you've written a failing test, or you've written the implementation to make a failing test pass.
