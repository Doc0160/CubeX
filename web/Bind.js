
function Bind()
{
	// No closure to define?
	if (arguments.length == 0)
		return null;

	// Figure out which of the 4 call types is being used to bind
	// Locate scope, function and bound parameter start index
	if (typeof(arguments[0]) == "string"){
		var scope = window;
		var func = window[arguments[0]];
		var start = 1;
	}else if (typeof(arguments[0]) == "function"){
		var scope = window;
		var func = arguments[0];
		var start = 1;
	}else if (typeof(arguments[1]) == "string"){
		var scope = arguments[0];
		var func = scope[arguments[1]];
		var start = 2;
	}else if (typeof(arguments[1]) == "function"){
		var scope = arguments[0];
		var func = arguments[1];
		var start = 2;
	}else{
		// unknown
		console.log("Bind() ERROR: Unknown bind parameter configuration");
		return;
	}

	// Convert the arguments list to an array
	var arg_array = Array.prototype.slice.call(arguments, start);
	start = arg_array.length;

	return function(){
		// Concatenate incoming arguments
		for (var i = 0; i < arguments.length; i++)
			arg_array[start + i] = arguments[i];

		// Call the function in the given scope with the new arguments
		return func.apply(scope, arg_array);
	}
}
