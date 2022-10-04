// / <reference types="Cypress" />

context('Viewport', () => {
  beforeEach(() => {
    cy.visit('http://localhost:3000');
  });

  it('Menu work correctly', () => {
    cy.get(
      '.ant-menu-submenu.ant-menu-submenu-inline.ant-menu-submenu-open .ant-menu-submenu-title',
    )
      .first()
      .click({ force: true })
      .wait(1000)
      .click({ force: true })
      .wait(1000)
      .click({ force: true })
      .wait(1000)
      .click({ force: true });
  });
});
