import React from "react";

const ShimmerCardSetting: React.FC = () => {
	return (
		<div className="p-4 bg-white dark:bg-[#0f0f0f] shadow rounded-lg animate-pulse">
			{Array.from({ length: 2 }).map((_, index) => (
				<div
					key={index}
					className="h-4 bg-gray-200 dark:bg-gray-600 rounded w-3/4 mb-3"
				></div>
			))}
			<div className="h-4 bg-gray-200 dark:bg-gray-600 rounded w-full"></div>
		</div>
	);
};

export default ShimmerCardSetting;
