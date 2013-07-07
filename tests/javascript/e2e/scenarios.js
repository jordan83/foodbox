'use strict';

/* http://docs.angularjs.org/guide/dev_guide.e2e-testing */

describe('Foodbox App', function() {

  it('should redirect / to /#/recipes', function() {
    browser().navigateTo('/');
    expect(browser().location().url()).toBe('/recipes');
  });
});