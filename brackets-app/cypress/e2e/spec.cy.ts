/// <reference types="cypress" />

describe('template spec', () => {
  // open web application on local host
  it('successfully loads', () => {
    cy.visit('http://localhost:4200/')
  })

  
  /*
  // test text enter in box
  it('passes text enter test', () => {
    cy.visit('http://localhost:4200/')
    cy.get('input').type('8')
  })
  */

  // test click on add teams button
  it('passes click test 1', () => {
    cy.visit('http://localhost:4200/')
    cy.contains('Add Teams').click()
  })

  // test click on sign in button
  it('passes click test 2', () => {
    cy.visit('http://localhost:4200/')
    cy.contains('Sign in with Google').click()
  })
})