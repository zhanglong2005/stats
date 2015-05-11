'use strict';

var statsControllers = angular.module('statsControllers', []);

statsControllers.controller('EntryListCtrl', ['$scope', 'Entry', function($scope, Entry) {
	$scope.entries = Entry.query();
}]);