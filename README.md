[![Work in MakeCode](https://classroom.github.com/assets/work-in-make-code-46eb539bcdc54ff4682c9f84a178d570a59fd821693cb33b02a3e5220eed4e48.svg)](https://classroom.github.com/online_ide?assignment_repo_id=15260045&assignment_repo_type=AssignmentRepo)


## Objective

Design and implement a backend service for options contracts risk and reward analysis using GoLang.

## Brief

Aries Financial is looking for a Lead GoLang Developer that can create a backend service to generate a risk & reward graph for options contracts. The service should accept an input of up to four options contracts and output X & Y values for a risk & reward graph where X is the price of the underlying at the time of expiry and Y is the profit/loss at that price. It should also return the following: max profit, max loss, and all break even points.

Here are the tasks you need to accomplish:

1. **Task 1: Options Contract Model**
    - Implement an OptionsContract model with the following fields: type (call or put), strike_price, bid, ask, expiration_date, long_short

2. **Task 2: Analysis Endpoint**
    - Implement an endpoint that accepts an array of up to four options contracts and returns the following:
        - An array of X & Y values for the risk & reward graph
        - The maximum possible profit
        - The maximum possible loss
        - All break even points

3. **Task 3: Analysis Logic**
    - Implement logic to calculate the X & Y values for the risk & reward graph
        - X values should be the price of the underlying at the time of expiry
        - Y values should be the profit or loss at that price
    - Implement logic to calculate the maximum possible profit, maximum possible loss, and all break even points

4. **Task 4: Testing**
    - Write unit tests for the options contract model validation
    - Write unit tests for the analysis endpoint
    - Write integration tests that simulate a user submitting options contracts and receiving the analysis results

### Evaluation Criteria

- Correctness and completeness of the code.
- Use of GoLang idioms and best practices.
- Structure and organization of the code.
- Quality of the tests and coverage of the code.

### CodeSubmit 

Please organize, design, test, and document your code as if it were
going into production - then push your changes to the master branch.

Have fun coding! 🚀

The Aries Financial Team

