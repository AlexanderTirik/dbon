export const createDb = (name: string) => {
  cy.contains("Create database").click();
  cy.get("input").type(name);
  cy.contains("Save changes").click();
  cy.get('[data-test-id="db-selector"]').find("select").select(name);
};

export const createTable = (name: string) => {
  cy.contains("Create table").click();
  cy.get("input").type(name);
  cy.contains("+").click();
  cy.get('[data-test-id="col-name"]').type("test_col");
  cy.contains("add").click();
  cy.contains("Save changes").click();
  cy.get('[data-test-id="table-selector"]').find("select").select(name);
};

export const createData = (data: string[]) => {
  cy.contains("Add data").click();
  data.forEach((d) => {
    cy.get("input[name=test_col-input]").clear().type(d);
    cy.contains("Save changes").click();
  });
  cy.contains("Close").click();
};
