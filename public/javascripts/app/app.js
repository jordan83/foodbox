'use strict';

/* App Module */

var app = angular.module('foodbox', ['recipeServices', 'dialogService', 'foodboxElements']);

app.config(['$routeProvider', function($routeProvider) {
	  $routeProvider.
	      when('/recipes', {templateUrl: '/partial/recipes', controller: 'RecipesCtrl'}).
	      when('/recipes/:recipeId', {templateUrl: '/partial/recipe', controller: 'RecipeCtrl'}).
	      when('/bulk', {templateUrl: '/partial/bulk', controller: 'BulkUploadCtrl'}).
	      when('/', {redirectTo: '/recipes'})
}]);