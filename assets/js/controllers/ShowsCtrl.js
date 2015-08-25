'use strict';

/* Controllers */

angular.module('emerald.controllers').
  controller('ShowsCtrl', ['$scope', 'msg', '$http', '$location', function($scope, msg, $http, $location) {

    $scope.loadShows = function () {
      $http.get('/shows/load').success(function(response){
        if (response.success) {
          $scope.shows = response.data;

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
        if (response.success) {
          $scope.searchResults = response.data;
        } else {
          msg.error (response.message);
        }
      });
    };

    $scope.showAddShowDialog = function () {
      $scope.query = null;
      $scope.searchResults = null;
      $('#addShowDialog').modal('show');
    };

    $scope.selectShow = function (show) {
      $scope.selectedShow = show;
    };

    $scope.saveShow = function (show) {
      $http.post ('/shows/save', show).success (function (response) {
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
      $http.post ('shows/delete', $scope.showToDelete).success (function (response) {
        if (response.success) {
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
        if (response.success) {
          var result = [];
          angular.forEach(response.data, function(value, key) {
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
      $('#episodesDialog').modal ('hide');
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
