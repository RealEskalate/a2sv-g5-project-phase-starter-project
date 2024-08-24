const tabs: string[] = ["Edit Profile", "Preferences"];

describe("It should test the settings page", () => {
  beforeEach(() => {
    cy.login();
    cy.visit("/settings");
    cy.wait(2000);
    cy.url().should("include", "settings");
  });

  tabs.forEach((tab) => {
    it(`should navigate to ${tab} page and submit form if applicable`, () => {
      cy.get(`[data-id="${tab}"]`).should("be.visible").click();

      if (tab === "Edit Profile") {
        cy.wait(1000);

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
        cy.wait(1000);
        cy.intercept(
          "PUT",
          "https://astu-bank-dashboard.onrender.com/user/update-preference"
        ).as("updateUserPreferenceRequest");

        cy.get('input[name="currency"]').type("USD");
        cy.get('input[name="timeZone"]').type("GMT-12:00");

        cy.get("form").submit();

        cy.wait("@updateUserPreferenceRequest")
          .its("response.statusCode")
          .should("eq", 200);

        cy.contains("Profile updated successfully").should("be.visible");
      }
    });

    it(`Should check for invalid  submission in ${tab}`, () => {
      cy.get(`[data-id="${tab}"]`).should("be.visible").click();

      if (tab === "Edit Profile") {
        cy.wait(1000);

        cy.get("form").submit();
        cy.get('input[name="name"]').clear();
        cy.contains("Name is required").should("be.visible");

        cy.get('input[name="email"]').clear();
        cy.contains("Email is required").should("be.visible");

        cy.get('input[name="postalCode"]').clear();
        cy.contains("Postal Code is required").should("be.visible");

        cy.get('input[name="username"]').clear();
        cy.contains("Username is required").should("be.visible");

        cy.get('input[name="password"]').clear();
        cy.contains("Password must be at least 6 characters").should(
          "be.visible"
        );

        cy.get('input[name="permanentAddress"]').clear();
        cy.contains("Permanent Address is required").should("be.visible");

        cy.get('input[name="presentAddress"]').clear();
        cy.contains("Present Address is required").should("be.visible");

        cy.get('input[name="city"]').clear();
        cy.contains("City is required").should("be.visible");

        cy.get('input[name="country"]').clear();
        cy.contains("Country is required").should("be.visible");
      } else if (tab === "Preferences") {
        cy.wait(1000);

        cy.get('input[name="currency"]').clear();
        cy.get('input[name="timeZone"]').clear();

        cy.contains("button", "Save Changes").click();
        cy.contains("Currency is required").should("be.visible");
        cy.contains("TimeZone is required").should("be.visible");
      }
    });
  });
});
