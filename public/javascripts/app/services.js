var recipeServices = angular.module('recipeServices', ['ngResource']);
recipeServices.factory('Recipe', function($resource) {
	return $resource('recipes/:recipeId', {}, {
		query: {method:'GET', params:{recipeId:''}, isArray:true},
		create: {method:'POST', params:{recipeId:''}}
	});
});

var dialog = angular.module('dialogService', []); 
dialog.factory('Dialog', function($http, $compile) {
	return {
		
		_dialog: undefined,
		
		_args: undefined,
		
		_scope: undefined,
		
		create: function(scope, args) {
			
			_args = args;
			_scope = scope;
			
			// Fetch the template url and add it to the body.
			var self = this;
			$http.get(args.templateUrl).success(function(data) {
				$('.' + args.elementClass).append(data);
				
				self._dialog = $('.modal-form').dialog({
					autoOpen: false,
					height: 600,
					width: 450,
					modal: true,
					buttons: args.buttons
				});
				
				// compiling must be done after we insert into the dom
				// and do the dialog setup. This ensures the controller
				// has access to the ready-made dialog.
				$compile($('.modal-form'))(scope);
			});
		},
			
		open: function() {
			$('.modal-form').dialog("open");
		},
		
		close: function() {
			this._dialog.dialog("close");
			$('.modal-form').remove();
			
			this.create(_scope, _args);
		},
		
		setButtons: function(buttons) {
			this._dialog.dialog("option", "buttons", buttons);
		}
	}
});