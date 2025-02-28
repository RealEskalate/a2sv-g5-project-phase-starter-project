import React from "react";

const ShimmerAddNewCard: React.FC = () => {
	return (
		<div className="p-4 bg-white dark:bg-[#0f0f0f] shadow rounded-lg animate-pulse">
			<div className="h-4 bg-gray-200 dark:bg-gray-600 rounded w-full mb-2"></div>
			<div className="h-4 bg-gray-200 dark:bg-gray-600 rounded w-1/2 mb-2"></div>
			<div className="h-10 bg-gray-200 dark:bg-gray-600 rounded w-full"></div>
		</div>
	);
};

export default ShimmerAddNewCard;
