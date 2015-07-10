'use strict';

/* Controllers */

angular.module('xtv.controllers').
  controller('ShowsRecentCtrl', ['$scope', 'msg', '$http', '$location', '$filter', function($scope, msg, $http, $location, $filter) {

    $scope.loadRecent = function () {
      $http.get('/shows/recentEpisodes', {params : {duration: 7}}).success(function(response){
         if (response.success) {
           $scope.episodes = response.data;
         } else {
           msg.error (response.message);
         }
       });
    };

  $scope.loadRecent();
    $scope.searchEpisode = function (show, episode) {
      var pad = "00";
      var season = "" + episode.seasonNumber;
      season =  pad.substring(0, pad.length - season.length) + season;

      var number = "" + episode.episodeNumber;
      number =  pad.substring(0, pad.length - number.length) + number;

      var query = show.name + " S" +  season + "E" + number;
      $location.path ('/search/' + query);
    };



  }]);
