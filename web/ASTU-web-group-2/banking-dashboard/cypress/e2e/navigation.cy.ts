import { should } from "chai";

const paths: string[] = [
  "/dashboard",
  "/transactions",
  "/accounts",
  "/investments",
  "/credit-cards",
  "/loans",
  "/services",
  "/settings",
];
describe("should check for invalid login form submission", () => {
  it("should present the error message upon invalid submission", () => {
    cy.visit("/");
    cy.contains("div", "Login").should("be.visible");
    cy.contains("div", "Login").click();

    cy.get('input[name="userName"]').clear();
    cy.get('input[name="password"]').clear();

    cy.contains("button", "Login").click();
    cy.contains("User name field is required").should("be.visible");
    cy.contains("Password field is required").should("be.visible");
  });
});

describe("should navigate to the desired page", () => {
  beforeEach(() => {
    cy.login();
  });

  for (const path of paths) {
    const name = path.slice(1);

    console.log(name);
    it(`should navigate to ${name}`, () => {
      const image = cy.get(`[data-id=side-image-${path.slice(1)}]`);
      const side_name = cy.get(`[data-id=side-image-${path.slice(1)}]`);
      const link = cy.get(`[data-id=side-link-${path.slice(1)}]`);
      cy.wait(2000);

      cy.intercept("GET", `http://localhost:3000/${path}`).as("pathname");
      link.click({ multiple: true, force: true });
      cy.wait(5000);
      cy.url().should("include", path);

      const title = cy.get("[data-id=title]");
      title.should("be.visible");
      const regex = new RegExp(name, "i");

      if (path === "/credit-cards") {
        title.should("contain.text", "Credit cards");
      } else {
        title.should(
          "contain.text",
          name.charAt(0).toUpperCase() + name.slice(1)
        );
      }

      image.should("have.css", "color", "rgb(52, 60, 106)");
      side_name.should("have.css", "color", "rgb(52, 60, 106)");

      cy.wait(2000);
    });
  }

  it("Should Logout successfully", () => {
    cy.get('[data-id="Profile-image"]').should("be.visible").click();
    cy.contains("button", "Logout").should("be.visible").click();

    cy.url().should("include", "/");
  });
});
