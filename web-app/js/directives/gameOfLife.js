app.directive('gameOfLife', function() {

  return {
    restrict: 'E',
    scope: {
      content: '='
    },
    templateUrl: 'js/directives/gameOfLife.html'
  };
});
