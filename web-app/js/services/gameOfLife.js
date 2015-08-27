app.factory('gameOfLife', function($http) {
	var gameOfLife = {};

	gameOfLife.data = {};
    
	var gameId = 0;
	
	gameOfLife.getBoard = function(id) { 
		return $http.get("http://199.233.247.129:8081/getGameJSON/" + id);
	};
	
	gameOfLife.stepBoard = function(id) {
		return $http.put("http://199.233.247.129:8081/stepGame/" + id);
	}

	gameOfLife.newBoard = function() {
		return $http.post("http://199.233.247.129:8081/initGame/classic");
	}

	gameOfLife.updateGame = function(board, id) { 
		console.log(board);
		return $http.put("http://199.233.247.129:8081/updateGame/" + id, angular.toJson(board));
	}

	return gameOfLife;
});
