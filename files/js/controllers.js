'use strict';

var statsControllers = angular.module('statsControllers', []);

statsControllers.controller('EntryCtrl', ['$scope', 'Entry', function ($scope, Entry) {

  var Nanosecond = 1;
  var Microsecond = 1000 * Nanosecond;
  var Millisecond = 1000 * Microsecond;
  var Second = 1000 * Millisecond;
  var Minute = 60 * Second;
  var Hour = 60 * Minute;

  $scope.stats = [];
  var statsFunc = function (result) {
    if (result) {
      var stats = [];
      var last;
      angular.forEach(result, function (entry) {
        if (last) {
          var valueDiff = entry.value - last.value;
          var timeDiff = entry.timestamp - last.timestamp;
          var stats = {
            timestamp: entry.timestamp,
            value: entry.value,
            valueDiff: valueDiff,
            timeDiff: timeDiff,
            diffPerHour: valueDiff / (timeDiff / Hour),
          };
          if (valueDiff != 0) {
            this.push(stats);
          }
        }
        last = entry;
      }, stats);
      $scope.stats = stats;
    }
  };
  $scope.entries = Entry.query({'limit':10}, statsFunc);
  $scope.entry = {};
  $scope.submit = function () {
    Entry.create($scope.entry, function () {
      $scope.entries = Entry.query({'limit':10}, statsFunc);
      $scope.entry = {};
    });
  };
}]);

