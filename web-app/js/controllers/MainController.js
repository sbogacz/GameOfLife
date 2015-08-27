app.controller('MainController', ['$scope','gameOfLife', function ($scope, gameOfLife){
   /*gameOfLife.getBoard().success(function(data) {
		$scope.board = data;
		console.log(data)
   })*/
	gameOfLife.getBoard().then(function(response){
		$scope.board = response.data;
		console.log($scope.board);
	});
}]);
