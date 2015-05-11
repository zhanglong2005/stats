'use strict';

var statsControllers = angular.module('statsControllers', []);

statsControllers.controller('EntryListCtrl', ['$scope', 'Entry', function ($scope, Entry) {
  $scope.entries = Entry.query();
}]);

statsControllers.controller('EntryCreateCtrl', ['$scope', 'Entry', function ($scope, Entry) {
  $scope.entry = {};
  $scope.submit = function () {
    Entry.create($scope.entry);
    $scope.entry = {};
  };
}]);