'use strict';

var statsControllers = angular.module('statsControllers', []);

statsControllers.controller('EntryListCtrl', ['$scope', 'Stats', function($scope, Stats) {
	$scope.entries = Entry.query();
}]);