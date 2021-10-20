import { createDb } from "../fixtures/helpers";

describe('Table tests', () => {
    before(() => {
      cy.visit('http://127.0.0.1:3000')
      createDb('test_table_db')
    })

    after(() => {
        cy.contains('Delete db').click()
    })

    it('Create table', () => {
      cy.contains('Create table').click()
      cy.get('input').type('test_table')
      cy.contains('+').click()
      cy.get('[data-test-id="col-name"]').type('test_col')
      cy.contains('add').click()
      cy.contains('Save changes').click()
      cy.get('[data-test-id="table-selector"]').find('select').select('test_table')
      cy.get('[data-test-id="table-cols"]').contains('td', 'test_col').should('be.visible');
      cy.get('[data-test-id="table-cols"]').contains('td', 'integer').should('be.visible');
      cy.contains('Delete table').should('exist')
    })
  
    it('Delete table', () => {
        cy.contains('Delete table').click()
      cy.get('[data-test-id="table-selector"]').find('option').contains('test_table').should('not.exist')
    })
  })
