describe("Test Dashboard Page", () => {
  beforeEach(() => {
    // Log in and wait for redirection
    cy.login(); // Ensure this command works as expected

    // Intercept the network requests
    cy.intercept("GET", "/cards?page=0&size=10").as("getCreditCards");
    cy.intercept(
      "GET",
      "https://astu-bank-dashboard.onrender.com/cards/66c86b944f73a73b051e9cca"
    ).as("getSpecificCard");
    cy.intercept("GET", "http://localhost:3000/api/auth/session").as(
      "getSession"
    );
    cy.intercept(
      "GET",
      "https://astu-bank-dashboard.onrender.com/user/current"
    ).as("getCurrentUser");
    cy.intercept(
      "GET",
      "https://astu-bank-dashboard.onrender.com/transactions"
    ).as("getTransactions");
    cy.intercept(
      "GET",
      "https://astu-bank-dashboard.onrender.com/transactions/balance-history"
    ).as("getBalanceHistory");
    // cy.intercept("GET", "http://localhost:3000/api/auth/providers").as("getAuthProviders");
    // cy.intercept("GET", "http://localhost:3000/api/auth/csrf").as("getCsrfData");
    // cy.intercept("GET", "http://localhost:3000/api/auth/callback/credentials").as("getCredentials");

    // Visit the dashboard page
    cy.visit("/dashboard");

    // Wait for the network requests to complete
    cy.wait("@getCreditCards");
    cy.wait("@getSession");
    cy.wait("@getCurrentUser");
    cy.wait("@getTransactions");
    cy.wait("@getBalanceHistory");
    // cy.wait("@getAuthProviders", { timeout: 10000 });
    // cy.wait("@getCsrfData");
  });

  it("should display the Credit Cards section with the correct title and button", () => {
    // Check if the Credit Cards section is visible
    cy.contains("Credit Cards").should("be.visible");

    // Verify the "See All" button is present and click it
    cy.contains("a", "See All").should("be.visible").click();
  });

  it("should navigate to the correct page when the 'See All' button is clicked", () => {
    // Click the "See All" button
    cy.contains("a", "See All").click();

    // Verify that the URL is correct (adjust the expected URL as needed)
    cy.url().should("include", "/credit-cards");

    // // Optionally, verify that the new page content is loaded correctly
    // cy.contains("Your Credit Cards").should("be.visible"); // Example assertion
  });

  it("should display Recent Transactions section", () => {
    // Check if the Recent Transactions section is visible
    cy.contains("Recent Transactions").should("be.visible");
  });

  it("should display the Weekly Activity chart", () => {
    // Check if the Weekly Activity chart section is visible
    cy.contains("Weekly Activity").should("be.visible");
  });

  it("should display the Expense Statistics chart", () => {
    // Check if the Expense Statistics chart section is visible
    cy.contains("Expense Statistics").should("be.visible");
  });

  it("should display the Quick Transfer section", () => {
    // Check if the Quick Transfer section is visible
    cy.contains("Quick Transfer").should("be.visible");
  });

  it("should display the Balance History chart", () => {
    // Check if the Balance History chart section is visible
    cy.contains("Balance History").should("be.visible");
  });

  // Add more tests as needed for other components and interactions
});
