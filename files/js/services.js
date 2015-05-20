'use strict';

angular.module('statsServices', ['ngResource']);

angular.module('statsServices').factory('Entry', ['$resource',
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


angular.module('statsServices').factory('Stats', ['$q', 'Entry',
  function ($q, Entry) {

    var Nanosecond = 1;
    var Microsecond = 1000 * Nanosecond;
    var Millisecond = 1000 * Microsecond;
    var Second = 1000 * Millisecond;
    var Minute = 60 * Second;
    var Hour = 60 * Minute;

    var stats = {};
    stats.list = function (param) {
      var deferred = $q.defer();
      Entry.query(param, function (result) {
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
          deferred.resolve(stats);
        }
      });
      return deferred.promise;
    };
    return stats;
  }]);
