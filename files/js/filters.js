'use strict';

angular.module('statsFilters', []).filter('checkmark', function() {
	return function(input) {
		return input ? '\u2713' : '\u2718';
	};
});
