'use strict';

/* Filters */

angular.module('emerald.filters', []).
  filter('megaBytes', [function() {
    return function(input) {

      if (angular.isNumber (input)) {
        input = parseInt (input / 1048576);
      }
      return input + ' MB';
    }
  }]);
