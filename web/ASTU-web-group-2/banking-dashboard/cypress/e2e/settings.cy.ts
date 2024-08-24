const tabs: string[] = ["Edit Profile", "Preferences", "Security"];

describe("It should test the settings page", () => {
  beforeEach(() => {
    cy.login();
    cy.visit("/settings");
    cy.wait(2000);
    cy.url().should("include", "/settings");
  });

  tabs.forEach((tab) => {
    it(`should navigate to ${tab} page and submit form if applicable`, () => {
      cy.get(`[data-id="${tab}"]`).should("be.visible").click();

      if (tab === "Edit Profile") {
        cy.wait(10000);
        cy.contains("hello theer").should("be.visible");
        cy.intercept(
          "PUT",
          "https://astu-bank-dashboard.onrender.com/user/update"
        ).as("updateUserRequest");

        cy.get('input[name="name"]').type("testjunior");
        cy.get('input[name="email"]').type("testjuinor@gmail.com");
        cy.get('input[name="dateOfBirth"]').type("2024-01-01");
        cy.get('input[name="postalCode"]').type("12345");
        cy.get('input[name="username"]').type("testjunior");
        cy.get('input[name="password"]').type("testjunior");
        cy.get('input[name="permanentAddress"]').type("Adama Bole");
        cy.get('input[name="presentAddress"]').type("Adama Bole");
        cy.get('input[name="city"]').type("Adama");
        cy.get('input[name="country"]').type("Ethiopia");

        cy.get("form").submit();

        cy.wait("@updateUserRequest")
          .its("response.statusCode")
          .should("eq", 200);

        cy.contains("Profile updated successfully").should("be.visible");
      } else if (tab === "Preferences") {
        // Add preferences-related tests here
      } else if (tab === "Security") {
        // Add security-related tests here
      }
    });
  });
});
