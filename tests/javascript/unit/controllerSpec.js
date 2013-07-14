'use strict';

describe('Foodbox controllers', function() {
	
	beforeEach(function(){
		this.addMatchers({
			toEqualData: function(expected) {
				return angular.equals(this.actual, expected);
			}
		});
	});
	 
	 
	beforeEach(angular.mock.module('recipeServices'));
	beforeEach(angular.mock.module('foodbox'));
	
	describe('NavCtrl', function(){
		
		var scope, ctrl, $httpBackend;

	    beforeEach(inject(function(_$httpBackend_, $rootScope, $controller) {
	      $httpBackend = _$httpBackend_;
	      $httpBackend.expectGET('/nav').
	          respond([{name: 'home', route: '/'}, {name: 'Create', route: '/create'}]);

	      scope = $rootScope.$new();
	      
	      var mockDialog = {
	    		  create: function(scope, options) {}
	      }
	      
	      ctrl = $controller('NavCtrl', {$scope: scope, Dialog: mockDialog});
	    }));
		
		it('should create "navItems" model with 2 nav items', function() {
			expect(scope.navItems).toBeUndefined();
			$httpBackend.flush();
			
			expect(scope.navItems).toEqualData([{name: 'home', route: '/'},
			                                {name: 'Create', route: '/create'}]);
	    });
	});
	
	describe('RecipesCtrl', function() {
		
		var recipeData = function() {
			return [{title: 'scallops', author: 'ramsay', id: '1', icon_location: '/path/test2'},
		            {title: 'swordfish', author: 'flay', id: '2', icon_location: '/path/test4'}];
		}
		
		var scope, ctrl, $httpBackend;
		
		beforeEach(inject(function(_$httpBackend_, $rootScope, $routeParams, $controller) {
			$httpBackend = _$httpBackend_;
			$httpBackend.expectGET('recipes').respond(recipeData());
			
			scope = $rootScope.$new();
			ctrl = $controller('RecipesCtrl', {$scope: scope});
		}));
		
		it('should create recipes model with 2 recipes', function() {
			expect(scope.recipes).toEqualData([]);
			$httpBackend.flush();
			
			expect(scope.recipes).toEqualData(recipeData());
		})
	});
	
	describe('RecipeCtrl', function(){
		var scope, $httpBackend, ctrl,
			testRecipe = function() {
				return {
					title: 'R1',
					author: 'A1',
					id: 1,
					icon_location: 'path/to/blah'
				}
		};
	 
		beforeEach(inject(function(_$httpBackend_, $rootScope, $routeParams, $controller) {
			$httpBackend = _$httpBackend_;
			$httpBackend.expectGET('recipes/1').respond(testRecipe());
 
			$routeParams.recipeId = '1';
			scope = $rootScope.$new();
			ctrl = $controller('RecipeCtrl', {$scope: scope});
		}));
		
		it('should fetch recipe detail', function() {
			expect(scope.recipe).toEqualData({});
			$httpBackend.flush();
			expect(scope.recipe).toEqualData(testRecipe());
		});
	});
	
	describe('CreateRecipeCtrl', function() {
		var scope, $httpBackend, ctrl;
		
		beforeEach(inject(function(_$httpBackend_, $rootScope, $controller) {
			$httpBackend = _$httpBackend_;
			
			scope = $rootScope.$new();
			
			var dialog = {
					setButtons: function(buttons) {
						
					}
			}
			
			ctrl = $controller('CreateRecipeCtrl', {$scope: scope, Dialog: dialog});
		}));
		
		it('should have empty title', function() {
			expect(scope.title).toEqual('');
		});
		
		it('should have empty author', function() {
			expect(scope.author).toEqual('');
		});
		
		it('should have one empty ingredient', function() {
			var ingredient = {name: '', quantity: '', unit: ''};
			expect(scope.ingredients).toEqualData([ingredient]);
		});
		
		it('ingredient added should increase ingredient size', function() {
			scope.addIngredient();
			expect(scope.ingredients.length).toEqual(2);
		});
		
		it('only ingredient is last', function() {
			var ingredient = scope.ingredients[0];
			expect(scope.isLastIngredient(ingredient)).toEqual(true);
		});
		
		it('does not remove only ingredient', function() {
			var ingredient = scope.ingredients[0];
			scope.removeIngredient(ingredient);
			
			expect(scope.ingredients.length).toEqual(1);
		});
		
		it('is last ingredient false when not last', function() {
			var ingredient = scope.ingredients[0];
			scope.addIngredient();
			expect(scope.isLastIngredient(ingredient)).toEqual(false);
		});
		
		it('is last ingredient true when more than one and last', function() {
			scope.addIngredient();
			scope.addIngredient();
			var ingredient = scope.ingredients[2];
			expect(scope.isLastIngredient(ingredient)).toEqual(true);
		});
		
	});
	
	describe('SlideShowCtrl', function(){
		var scope, $httpBackend, ctrl;
	 
		beforeEach(inject(function($rootScope, $controller) {
			scope = $rootScope.$new();
			ctrl = $controller('SlideShowCtrl', {$scope: scope});
			
			// This must be done after the controller is defined.
			scope.imageUrls = ['/some/url', '/some/other/url'];
			scope.$digest();
		}));
		
		it('should have curIndex == 0', function() {
			expect(scope.curIndex).toEqual(0);
		});
		
		it('should show first image when curIndex is 0', function() {
			expect(scope.showImage('/some/url')).toEqual(true);
		});
		
		it('should show second image when curIndex is 1', function() {
			scope.curIndex = 1;
			expect(scope.showImage('/some/other/url')).toEqual(true);
		});
		
		it('should show file upload when cur index is < 0', function() {
			scope.curIndex = -1;
			expect(scope.showFileUpload()).toEqual(true);
		});
		
		it('should show file upload when cur index is >= length of image list', function() {
			scope.curIndex = scope.imageUrls.length;
			expect(scope.showFileUpload()).toEqual(true);
		});
		
		it('should increment curIndex when curIndex is < length', function() {
			scope.next();
			expect(scope.curIndex).toEqual(1);
		});
		
		it('should set curIndex to 0 when curIndex is equal to length', function() {
			scope.curIndex = scope.imageUrls.length;
			scope.next();
			expect(scope.curIndex).toEqual(0);
		});
		
		it('should decrement curIndex when curIndex is >= 0', function() {
			scope.previous();
			expect(scope.curIndex).toEqual(-1);
		});
		
		it('should set curIndex to length - 1 when curIndex is < 0', function() {
			scope.curIndex = -1;
			scope.previous();
			expect(scope.curIndex).toEqual(scope.imageUrls.length -1);
		});
		
		it('should add url to image urls and set curIndex to length -1', function() {
			scope.postUpload({Url: 'Another/url'});
			expect(scope.curIndex).toEqual(2);
		})
	});
	
});