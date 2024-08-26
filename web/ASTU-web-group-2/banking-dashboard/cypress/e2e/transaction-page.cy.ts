describe("Test the transaction side bar", () => {
  beforeEach("Should login and redirect to transactions", () => {
    cy.login();
    cy.contains("h1", "Transactions").click();
    cy.wait(4000);
  });

  describe("Tests transactions table", () => {
    it("Filters to income", () => {
      cy.contains("button", "Income").as("incomeButton");
      cy.contains("button", "Expense").as("expenseButton");
      cy.contains("button", "All Transactions").as("allTransactionsButton");

      cy.get("@incomeButton").click();
      cy.get("@incomeButton").should("not.have.class", "false");

      cy.get("@expenseButton").should("have.class", "false");
      cy.get("@allTransactionsButton").should("have.class", "false");
    });
    it("Filters to expense", () => {
      cy.contains("button", "Income").as("incomeButton");
      cy.contains("button", "Expense").as("expenseButton");
      cy.contains("button", "All Transactions").as("allTransactionsButton");

      cy.get("@expenseButton").click();
      cy.get("@expenseButton").should("not.have.class", "false");

      cy.get("@incomeButton").should("have.class", "false");
      cy.get("@allTransactionsButton").should("have.class", "false");
    });
    it("Filters to all transactions", () => {
      cy.contains("button", "Income").as("incomeButton");
      cy.contains("button", "Expense").as("expenseButton");
      cy.contains("button", "All Transactions").as("allTransactionsButton");

      cy.get("@allTransactionsButton").click();
      cy.get("@allTransactionsButton").should("not.have.class", "false");

      cy.get("@incomeButton").should("have.class", "false");
      cy.get("@expenseButton").should("have.class", "false");
    });
  });

  describe("Tests Card component", () => {
    it("Redirects to Credit Cards", () => {
      cy.contains("a", "+ Add Card").click();
      cy.wait(4000);
      cy.url().should("include", "/credit-cards");
    });
  });
});
