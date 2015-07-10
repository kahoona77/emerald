'use strict';

/* Controllers */

angular.module('xtv.controllers').
  controller('ShowsCtrl', ['$scope', 'msg', '$http', '$location', function($scope, msg, $http, $location) {

    $scope.loadShows = function () {
      $http.get('/shows/load').success(function(response){
        if (response.status == 'ok') {
          $scope.shows = response.shows;

          //reselect server
          angular.forEach ($scope.shows, function (show) {
            if ($scope.selectedShow) {
              if (show._id == $scope.selectedShow._id) {
                $scope.selectedShow = show;
              }
            }
          });
        } else {
          msg.error (response.message);
        }
      });
    };
    $scope.loadShows();

    //search
    $scope.searchShow = function () {

      $http.get('/shows/search', {params : {query: $scope.query}}).success(function(response){
        if (response.status == 'ok') {
          $scope.searchResults = response.shows;
        } else {
          msg.error (response.message);
        }
      });
    };

    $scope.showAddShowDialog = function () {
      $('#addShowDialog').modal('show');
    };

    $scope.selectShow = function (show) {
      $scope.selectedShow = show;
    };

    $scope.saveShow = function (show) {
      $http.post ('/shows/save', {data: show}).success (function (response) {
        if (response.status = 'ok') {
          $('#addShowDialog').modal('hide');
          $scope.query = undefined;
          $scope.loadShows();
        } else {
          msg.error (response.message);
        }
      });
    };

    $scope.showDeleteShowConfirm = function (show) {
      $scope.showToDelete = show;
      $('#deleteShowConfirmDialog').modal ('show');
    };

    $scope.deleteShow = function () {
      $http.post ('shows/delete', {data: $scope.showToDelete}).success (function (response) {
        if (response.status = 'ok') {
          $('#deleteShowConfirmDialog').modal('hide');
          $scope.showToDelete = undefined;
          $scope.selectedShow = undefined;
          $scope.seasons = undefined;
          $scope.loadShows();
        } else {
          msg.error (response.message);
        }
      });
    };

    $scope.loadEpisodes = function (show) {
     $http.get('/shows/loadEpisodes', {params : {showId: show.id}}).success(function(response){
        if (response.status == 'ok') {
          var result = [];
          angular.forEach(response.episodes, function(value, key) {
            this.push({seasonNumber: key, episodes: value});
          }, result);
          $scope.seasons = result;
          $('#episodesDialog').modal ('show');
        } else {
          msg.error (response.message);
        }
      });
    };

    $scope.searchEpisode = function (show, episode) {
      var pad = "00";
      var season = "" + episode.seasonNumber;
      season =  pad.substring(0, pad.length - season.length) + season;

      var number = "" + episode.episodeNumber;
      number =  pad.substring(0, pad.length - number.length) + number;

      var query = show.searchName + " S" +  season + "E" + number;
      $location.path ('/search/' + query);
    };

    $scope.updateEpisodes = function () {
     $http.get('/shows/updateEpisodes').success(function(response){
        if (response.status == 'ok') {
          msg.error ("Updating episodes started...");
        } else {
          msg.error (response.message);
        }
      });
    };

  }]);
