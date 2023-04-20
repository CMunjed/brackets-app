/// <reference types="cypress" />

describe('template spec', () => {
  // open web application on local host
  it('successfully loads', () => {
    cy.visit('http://localhost:4200/')
  })

  // test click on add teams button
  it('passes click test 1 (compunding)', () => {
    cy.visit('http://localhost:4200/')
    cy.contains('Add Teams').click()
    cy.contains('Edit Team Name').click()
  })


  // test click on google sign in button
  it('passes click test 2 (google)', () => {
    cy.visit('http://localhost:4200/')
    cy.contains('Sign in with Google').click()
  })

  // test click on sign in button
  it('passes click test 3 (sign in)', () => {
    cy.visit('http://localhost:4200/')
    cy.contains('Sign In').click()
  })

  // test click on github button
  it('passes click test 4 (github)', () => {
    cy.visit('http://localhost:4200/')
    cy.contains('Give our repo a star').click()
  })

  // test slide use
  it('passes slide toggle click', () => {
    cy.visit('http://localhost:4200/')
    cy.contains('Single-Elimination').click()
  })

  // test compound slide use
  it('passes compound slide toggle click', () => {
    cy.visit('http://localhost:4200/')
    cy.contains('Single-Elimination').click()
    cy.contains('Single-Elimination').click()
    cy.contains('Double-Elimination').click()
  })
  
})