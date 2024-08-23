describe("Test Account Page", () => {
  beforeEach(() => {
    cy.login();
    cy.contains("h1", "Account").should("exist");
    cy.wait(5000); // Or use more reliable waiting strategies
    cy.visit("/accounts");
  });

  describe("Test See All Button", () => {
    it("should contiain My Balance", () => {
      // Check if the Balance History chart section is visible
      cy.contains("My Balance").should("be.visible");
    });
    it('should redirect to the credit cards page when "See All" button is clicked', () => {
      // Debugging: Check if button exists and is visible
      cy.get("a").contains("See All").should("be.visible").click();

      // Verify that the URL changes to '/credit-cards'
      cy.url().should("include", "/credit-cards");
    });
  });
});
