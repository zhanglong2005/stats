'use strict';

var statsControllers = angular.module('statsControllers', []);

statsControllers.controller('UserListCtrl', ['$scope', 'User', function($scope, User) {
	$scope.users = User.query();
}]);