import React from "react";

const Shimmer = () => {
	return (
		<div className="relative overflow-hidden bg-gray-300 dark:bg-gray-800 rounded-lg w-full h-24">
			<div className="absolute inset-0 bg-gradient-to-r dark:bg-gray-900 from-gray-300 via-gray-200 to-gray-300 animate-shimmer"></div>
		</div>
	);
};

export default Shimmer;
