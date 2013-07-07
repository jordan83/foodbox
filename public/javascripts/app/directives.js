var module = angular.module("foodboxElements", []);

module.directive("foodboxButton", function() {
	
	return function(scope, element, attrs) {
		var button = $(element);
		button.button();
	}
});

module.directive("foodboxFileDrop", function($http) {
	return {
		
		scope: {
			callback: "=fileUploadFinished",
			route: "=fileUploadUrl"
		},
	
		link: function(scope, element, attrs) {
			var fileDrop = $(element);
			fileDrop.addClass('foodie-file-drop');
			
			function stopEvent(e) {
				e.stopPropagation();
				e.preventDefault();
			}
			
			function uploadFile(file) {
				var formData = new FormData();
				formData.append("file", file);
				$http({
					method: 'POST',
		            url: scope.route,
		            headers: { 'Content-Type': false },
		            //This method will allow us to change how the data is sent up to the server
		            // for which we'll need to encapsulate the model data in 'FormData'
		            transformRequest: function (data) {
		                var formData = new FormData();
		                formData.append("file", data);
		                
		                return formData;
		            },
		            data: file
		        }).
		        success(function (data, status, headers, config) {
		            fileDrop.removeClass('foodie-file-drop-enter');
		            scope.callback(data);
		        }).
		        error(function (data, status, headers, config) {
		            fileDrop.removeClass('foodie-file-drop-enter');
		        });
			}
			
			element.bind('dragenter', function(e) {
				stopEvent(e);
				fileDrop.addClass('foodie-file-drop-enter');
			});
			
			element.bind('dragleave', function(e) {
				stopEvent(e);
				fileDrop.removeClass('foodie-file-drop-enter');
			});
			
			element.bind('dragover', stopEvent);
			
			element.bind('dragexit', stopEvent);
			
			element.bind('drop', function(e) {
				stopEvent(e);
				
				var files = e.originalEvent.dataTransfer.files;
				var count = files.length;
				
				if (count > 0) {
					uploadFile(files[0]);
				}
			});
		}
	}
});