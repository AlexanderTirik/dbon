import { createDb, createTable } from "../fixtures/helpers";

describe("Data tests", () => {
  before(() => {
    cy.visit("http://127.0.0.1:3000");
    createDb("test_data_db");
    createTable("test_data_table");
  });

  after(() => {
    cy.contains("Delete db").click();
  });

  it("Create data", () => {
    cy.contains("Add data").click();
    cy.get("input[name=test_col-input]").type("123456789");
    cy.contains("Save changes").click();
    cy.get("input[name=test_col-input]").clear().type("987654321");
    cy.contains("Save changes").click();
    cy.contains("Close").click();
    cy.get('[data-test-id="dashboard-data"]').contains("td", "123456789");
    cy.get('[data-test-id="dashboard-data"]').contains("td", "987654321");
    cy.contains("Join tables").should("exist");
  });

  it("Delete data", () => {
    cy.get('[data-test-id="dashboard-data"]')
      .find("button")
      .click({ multiple: true });
    cy.get('[data-test-id="dashboard-data"]')
      .contains("td", "123456789")
      .should("not.exist");
    cy.get('[data-test-id="dashboard-data"]')
      .contains("td", "987654321")
      .should("not.exist");
  });
});
