'use strict';

var statsFilters = angular.module('statsFilters', []);

statsFilters.filter('checkmark', function() {
	return function(input) {
		return input ? '\u2713' : '\u2718';
	};
});

statsFilters.filter('timestampToDate', function() {
	return function(input) {
		var d = new Date(input/1000/1000);
		return d.toISOString().slice(0,19).replace("T"," ");
	};
});
