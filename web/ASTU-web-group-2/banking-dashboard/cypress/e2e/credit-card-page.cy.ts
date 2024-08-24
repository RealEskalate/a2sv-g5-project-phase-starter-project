describe("Tests Credit Cards page", () => {
  beforeEach(() => {
    cy.login();
    cy.wait(4000);
    cy.contains("h1", "Credit Cards").click();
    cy.wait(4000);
  });
  it("Add credit card", () => {
    cy.get("input[name='cardType']").type("modern");
    cy.get("input[name='cardHolder']").type("ayele ayele");
    cy.get("input[name='balance']").type("0");
    cy.get("input[name='expiryDate']").type("2025-02-06");
    cy.contains("button", "Add Card").click();
    cy.wait(1000);
    cy.get(".Toastify__toast--success").should("be.visible");
  });
  it("Diplay error message", () => {
    cy.get("input[name='cardType']").clear();
    cy.get("input[name='cardHolder']").clear();
    cy.get("input[name='expiryDate']").clear();
    cy.contains("button", "Add Card").click();

    cy.contains("div", "Card Type is required").should("be.visible");
    cy.contains("div", "Card Holder is required").should("be.visible");
    cy.contains("div", "Expiration Date is required").should("be.visible");
  });
  it("Error credit card", () => {
    cy.get("input[name='cardType']").type("modern");
    cy.get("input[name='cardHolder']").type("ayele ayele");
    cy.get("input[name='balance']").type("899");
    cy.get("input[name='expiryDate']").type("2025-02-06");
    cy.contains("button", "Add Card").click();
    cy.wait(1000);
    cy.get(".Toastify__toast--error").should("be.visible");
  });
});
