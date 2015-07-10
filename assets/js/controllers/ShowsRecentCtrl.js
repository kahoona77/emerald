'use strict';

/* Controllers */

angular.module('xtv.controllers').
  controller('ShowsRecentCtrl', ['$scope', 'msg', '$http', '$location', '$filter', function($scope, msg, $http, $location, $filter) {


    var baseUrl = 'https://api-v2launch.trakt.tv/calendars/my/shows/';
    var config = {
      headers: {
        'Content-Type'     : 'application/json',
        // 'Authorization'    : 'Bearer 703a7584c1c6b3f50690fd2f63264f879771837b1e3bbe80912fd2968c232508',
        'trakt-api-version': '2',
        'trakt-api-key'    : '69dd71b6d1f94aabf585a0bc532e6e8fcf34033fbc044f11c6164c7a2e77d36b',
      }
    }

    $scope.loadSettings = function () {
      $http.get('data/loadSettings').success(function (response) {
        if (response.status == 'ok') {
          var settings = response.settings;

          config.headers.Authorization = "Bearer " + settings.traktToken;

          $scope.loadRecent();
        } else {
          msg.error(response.message);
        }
      });
    };
    $scope.loadSettings();

    $scope.loadRecent = function () {
      var today = new Date();
      today.setDate(today.getDate()-7);
      var date = $filter('date')(today, 'yyyy-MM-dd') + '/6';

      $http.get(baseUrl + date + '?extended=full,images', config).success(function(episodes){
        $scope.episodes = episodes;
      });
    };


    $scope.searchEpisode = function (show, episode) {
      var pad = "00";
      var season = "" + episode.season;
      season =  pad.substring(0, pad.length - season.length) + season;

      var number = "" + episode.number;
      number =  pad.substring(0, pad.length - number.length) + number;

      var query = show.title + " S" +  season + "E" + number;
      $location.path ('/search/' + query);
    };



  }]);
