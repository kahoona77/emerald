angular.module ('xtv.services').factory('msg', [ '$rootScope', '$timeout', function($rootScope, $timeout) {
    return {
        show: function (message, type, timeout) {
            $rootScope.message = {};
            $rootScope.message.text = message;
            $rootScope.message.type = type ? type : 'info';

            timeout = timeout ? timeout : 3000;

            if ($rootScope.message.timeout) {
                $timeout.cancel ($rootScope.message.timeout);
            }
            $rootScope.message.timeout = $timeout (function () {
                $rootScope.message.text = undefined;
            }, timeout);
        },

        error: function (message) {
          this.show(message, 'error');
        },

        hideMessage: function () {
            if ($rootScope.message.timeout) {
                $timeout.cancel ($rootScope.message.timeout);
            }
            $rootScope.message.text = undefined;
        }
    }
}]);