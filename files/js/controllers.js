'use strict';

var statsControllers = angular.module('statsControllers', []);

var Nanosecond = 1
var Microsecond          = 1000 * Nanosecond
var Millisecond          = 1000 * Microsecond
var Second               = 1000 * Millisecond
var Minute               = 60 * Second
var Hour                 = 60 * Minute

statsControllers.controller('EntryCtrl', ['$scope', 'Entry', function ($scope, Entry) {
  $scope.stats = [];
  $scope.entries = Entry.query({}, function (result) {
    if (result) {
      var stats = [];
      var last;
      angular.forEach(result, function (entry) {
        if (last) {
          var valueDiff = entry.value - last.value;
          var timeDiff = (entry.timestamp - last.timestamp) / Hour;
          var stats = {
            timestamp: entry.timestamp,
            value: entry.value,
            diff: valueDiff,
            perhour: valueDiff / timeDiff,
          };
          this.push(stats);
        }
        last = entry;
      }, stats);
      $scope.stats = stats;
    }
  });
  $scope.entry = {};
  $scope.submit = function () {
    Entry.create($scope.entry);
    $scope.entry = {};
    $scope.entries = Entry.query();
  };

}]);

