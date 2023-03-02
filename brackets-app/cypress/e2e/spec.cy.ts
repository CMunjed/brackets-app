describe('template spec', () => {
  // open web application on local host
  it('successfully loads', () => {
    cy.visit('http://localhost:4200/')
  })
  // test click on sign in button
  it('passes click test 1', () => {
    cy.contains('Sign in').click()
  })
})