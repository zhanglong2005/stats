'use strict';

var statscacheServices = angular.module('statsServices', ['ngResource']);

statscacheServices.factory('Entry', ['$resource',
	function ($resource) {
		return $resource('/entry/:Id', {}, {
			query: {
				method: 'GET', params: {}, isArray: true
			},
			create: {
				method: 'POST', params: {}
			},
			delete: {
				method: 'DELETE', params: {}
			}
		});
	}]);
