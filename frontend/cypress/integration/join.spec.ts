import { createDb, createTable, createData } from "../fixtures/helpers";

describe("Join tests", () => {
  before(() => {
    cy.visit("http://127.0.0.1:3000");
    createDb("test_join_db");
    createTable("test_join_table_first");
    createData(["123456789", "987654321"]);
    createTable("test_join_table_second");
    createData(["123456789", "987654321"]);
  });

  after(() => {
    cy.contains("Delete db").click();
  });

  it("Join data", () => {
    cy.contains("Join tables").click();
    cy.get('[data-test-id="first-join-table-select"]').select(
      "test_join_table_first"
    );
    cy.get('[data-test-id="first-join-col-select"]').select("test_col");
    cy.get('[data-test-id="second-join-table-select"]').select(
      "test_join_table_second"
    );
    cy.get('[data-test-id="second-join-col-select"]').select("test_col");
    cy.contains(/^Join$/).click();
    cy.get('[data-test-id="join-data"]')
      .contains("th", "test_join_table_first_test_col")
      .should("be.visible");
    cy.get('[data-test-id="join-data"]')
      .contains("th", "test_join_table_second_test_col")
      .should("be.visible");
    cy.get('[data-test-id="join-data"]')
      .contains("td", "123456789")
      .should("be.visible");
    cy.get('[data-test-id="join-data"]')
      .contains("td", "987654321")
      .should("be.visible");
    cy.contains("Close").click();
  });
});
