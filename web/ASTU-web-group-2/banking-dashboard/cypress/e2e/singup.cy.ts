describe("Should test the signup page", () => {
  beforeEach(() => {
    cy.visit("/");
    cy.contains("div", "Login").click();

    cy.contains("a", "Register").should("be.visible");
    cy.contains("a", "Register").click();
    cy.wait(5000);
  });

  it("should successsfuly submit sign-up", () => {
    cy.intercept(
      "POST",
      "https://astu-bank-dashboard.onrender.com/auth/register"
    ).as("signUpUser");

    // make sure to change username, email, and password for each attempt
    cy.get('input[name="name"]').type("testabcdef");
    cy.get('input[name="email"]').type("testabcdef@gmail.com");
    cy.get('input[name="username"]').type("testabcdef");
    cy.get('input[name="password"]').type("testabcdef");

    cy.get('input[name="dateOfBirth"]').type("2024-01-01");
    cy.get('input[name="postalCode"]').type("12345");
    cy.get('input[name="permanentAddress"]').type("Adama Bole");
    cy.get('input[name="presentAddress"]').type("Adama Bole");
    cy.get('input[name="city"]').type("Adama");
    cy.get('input[name="country"]').type("Ethiopia");

    cy.contains("button", "Next").should("be.visible").click();

    cy.get('input[name="currency"]').type("USD");
    cy.get('input[name="timeZone"]').type("GMT-12:00");

    cy.contains("button", "Register").should("be.visible").click();

    cy.wait("@signUpUser", { timeout: 10000 })
      .its("response.statusCode")
      .should("eq", 200);
    cy.wait(1000);
    cy.url().should("include", "dashboard");
    cy.wait(2000);
  });

  it("should check for invalid form submission", () => {
    cy.contains("button", "Next").should("be.visible").click();

    cy.get('input[name="currency"]').clear();
    cy.get('input[name="timeZone"]').clear();

    cy.contains("button", "Register").should("be.visible").click();

    cy.contains("Currency is required").should("be.visible");
    cy.contains("TimeZone is required").should("be.visible");
  });
});
