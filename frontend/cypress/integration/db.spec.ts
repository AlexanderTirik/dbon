describe("DB tests", () => {
  before(() => {
    cy.visit("http://127.0.0.1:3000");
  });

  it("Create DB", () => {
    cy.contains("Create database").click();
    cy.get("input").type("test_db");
    cy.contains("Save changes").click();
    cy.get('[data-test-id="db-selector"]')
      .find("select")
      .select("test_db")
      .should("have.value", "test_db");
    cy.contains("Delete db").should("exist");
  });

  it("Delete DB", () => {
    cy.contains("Delete db").click();
    cy.get('[data-test-id="db-selector"]')
      .find("option")
      .contains("test_db")
      .should("not.exist");
  });
});
export {}
