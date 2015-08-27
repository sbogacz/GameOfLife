app.directive('gameOfLife', function() {
  return {
    restrict: 'E',
    scope: {
      info: '='
    },
    templateUrl: 'js/directives/gameOfLife.html'
  };
});
