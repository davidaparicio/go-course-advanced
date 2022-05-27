Feature: bank account
  A user's bank account must be able to withdraw and deposit cash
  Moreover, as the user is rich, with a very large estate, and a lot of vast networking,
  The bank tolerates having an overdrawn account, thank to expensive overdraft charges

  Scenario Outline: Deposit
    Given I have a bank account with <start>$
    When I deposit <deposit>$
    Then it should have a balance of <end>$
    
    Examples:
      | start | deposit | end |
      | 10    | 0       | 10  |
      | 10    | 10      | 20  |
      | 100   | 50      | 150 |
      | -100  | 50      | -50 |

  Scenario Outline: Withdrawal
    Given I have a bank account with <start>$
    When I withdraw <withdrawal>$
    Then it should have a balance of <end>$

    Examples:
      | start | withdrawal | end  |
      | 10    | 0          | 10   |
      | 20    | 10         | 10   |
      | 100   | 50         | 50   |
      | 100   | 500        | -400 |
