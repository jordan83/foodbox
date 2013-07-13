
app.controller('NavCtrl', function($scope, $http, Dialog) {
	$http.get('/nav').success(function(data) {
	    $scope.navItems = data;
	});
	
	Dialog.create($scope, {
		templateUrl: '/partial/createRecipe',
		elementClass: 'create-recipe-placeholder'
	});
	
	$scope.openDialog = function(){
		Dialog.open();
	};
});

app.controller('CreateRecipeCtrl', function($scope, Dialog, Recipe, $route) {
	
	function submit() {
		Recipe.create({
			title: $scope.title,
			author: $scope.author,
			ingredients: $scope.ingredients
		}, function() {
			$route.reload();
		});
		Dialog.close();
	}
	
	Dialog.setButtons({
		Save: submit
	});
	
	$scope.title = '';
	
	$scope.author = '';
	
	$scope.ingredients = [];
	
	$scope.addIngredient = function() {
		$scope.ingredients.push({
			name: '',
			quantity: '',
			unit: ''
		});
	}
	
	$scope.removeIngredient = function(ingredient) {
		if ($scope.ingredients.length <= 1) {
			return;
		}
		
		var index = $scope.ingredients.indexOf(ingredient);
		if (index >= 0) {
			$scope.ingredients.splice(index, 1);
		}
	}
	
	$scope.isLastIngredient = function(ingredient) {
		var index = $scope.ingredients.indexOf(ingredient);
		var length = $scope.ingredients.length;
		return index == length - 1;
	}
	
	$scope.addIngredient();
});

app.controller('RecipesCtrl', function($scope, Recipe) {
	
	var placeholderUrl = "/images/placeholder_rev.png";
	
	$scope.recipes = Recipe.query();
	
	
	$scope.getImageUrl = function(recipe) {
		if (recipe.ImageUrls.length > 0) {
			return recipe.ImageUrls[0];
		}
		return placeholderUrl;
	}
	
	$scope.$on("RecipeAdded", function() {
		$scope.recipes = Recipe.query();
	});
});

app.controller('RecipeCtrl', function($scope, $routeParams, Recipe, $location) {
	
	$scope.recipe = Recipe.get({recipeId:$routeParams.recipeId}, function() {
		$scope.imageUrls = $scope.recipe.ImageUrls;
	});
	
	$scope.remove = function() {
		Recipe.remove({recipeId:$routeParams.recipeId}, function() {
			$location.path("/recipes");
		});
	}
});

app.controller("SlideShowCtrl", function($scope) {
	
	$scope.$watch('imageUrls', function(imageUrls) {
	
		if (imageUrls == undefined) {
			return;
		}
		
		$scope.showImage = function(image) {
			return imageUrls.indexOf(image) == $scope.curIndex;
		}
		
		$scope.showFileUpload = function() {
			return $scope.curIndex < 0 || $scope.curIndex >= imageUrls.length;
		}
		
		$scope.next = function() {
			if ($scope.curIndex < imageUrls.length) {
				$scope.curIndex++;
			} else {
				$scope.curIndex = 0;
			}
		}
		
		$scope.previous = function() {
			if ($scope.curIndex >= 0) {
				$scope.curIndex--;
			} else {
				$scope.curIndex = imageUrls.length -1;
			}
		}
		
		$scope.curIndex = 0;
		
		$scope.postUpload = function(data) {
			imageUrls.push(data.Url);
			$scope.curIndex = imageUrls.length - 1;
		}
		
		$scope.counterText = function() {
			var numerator = $scope.showFileUpload() ? "-" : ($scope.curIndex + 1) + "";
			return numerator + " of " + imageUrls.length;
		}
	});
	
});
