'use strict';

var statsServices = angular.module('statsServices', ['ngResource']);

statsServices.factory('User', ['$resource',
	function ($resource) {
		return $resource(':userId.json', {}, {
			query: {method: 'GET', params: {userId: 'users'}, isArray: true}
		});
	}]);