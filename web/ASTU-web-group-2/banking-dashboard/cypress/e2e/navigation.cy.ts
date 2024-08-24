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
});
