import React from "react";

const ShimmerCardStatistics: React.FC = () => {
	return (
		<div className="p-4 bg-white dark:bg-[#0f0f0f] shadow rounded-lg animate-pulse">
			<div className="h-4 bg-gray-200 dark:bg-gray-600 rounded w-3/4 mb-2"></div>
			<div className="h-4 bg-gray-200 dark:bg-gray-600 rounded w-1/2 mb-2"></div>
			<div className="h-4 bg-gray-200 dark:bg-gray-600 rounded w-2/3"></div>
		</div>
	);
};

export default ShimmerCardStatistics;
