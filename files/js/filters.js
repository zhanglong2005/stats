'use strict';

angular.module('statsFilters', []);

angular.module('statsFilters').filter('checkmark', function () {
  return function (input) {
    return input ? '\u2713' : '\u2718';
  };
});

angular.module('statsFilters').filter('timestampToDate', function () {
  return function (input) {
    var d = new Date(input / 1000 / 1000);
    return d.toISOString().slice(0, 19).replace("T", " ");
  };
});

angular.module('statsFilters').filter('timeToString', function () {
  return function (input, l) {
    var Nanosecond = 1;
    var Microsecond = 1000 * Nanosecond;
    var Millisecond = 1000 * Microsecond;
    var Second = 1000 * Millisecond;
    var Minute = 60 * Second;
    var Hour = 60 * Minute;

    var format = function (value) {
      if (l) {
        return value.toFixed(l);
      }
      return value;
    };

    if (input > Hour) {
      return format(input / Hour) + ' hour';
    }

    if (input > Minute) {
      return format(input / Minute) + ' min';
    }

    if (input > Second) {
      return format(input / Second) + ' sec';
    }

    if (input > Millisecond) {
      return format(input / Millisecond) + ' ms';
    }

    if (input > Microsecond) {
      return format(input / Microsecond).toFixed(l) + ' µs';
    }
    return format(input) + 'ns';
  };
});

angular.module('statsFilters').filter('toFixed', function () {
  return function (input, l) {
    return input.toFixed(l);
  };
});

angular.module('statsFilters').filter('reverse', function () {
  return function (items) {
    if (angular.isArray(items)) {
      return items.slice().reverse();
    } else {
      return items;
    }
  };
});