'use strict';

angular.module('statsControllers', []);

angular.module('statsControllers').controller('EntryCtrl', ['$scope', 'Entry', 'Stats', function ($scope, Entry, Stats) {
  var limit = 20;
  $scope.stats = [];
  Stats.list({'limit': limit}).then(function (stats) {
    $scope.stats = stats;
  });
  $scope.entry = {};
  $scope.submit = function () {
    Entry.create($scope.entry, function () {
      Stats.list({'limit': limit}).then(function (stats) {
        $scope.stats = stats;
      });
      $scope.entry = {};
    });
  };
}]);

