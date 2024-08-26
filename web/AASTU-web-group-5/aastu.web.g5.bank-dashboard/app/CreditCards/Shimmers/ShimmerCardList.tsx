import React from "react";

const ShimmerCardList: React.FC = () => {
	return (
		<div className="animate-pulse">
			<div className="flex sm:flex-row items-start sm:items-center gap-4 bg-gray-200 dark:bg-[#0f0f0f] py-5 px-5 pr-6 rounded-xl sm:justify-between">
				<div className="flex items-center gap-5">
					<div className="bg-gray-300 dark:bg-gray-600  w-8 h-8 rounded-md"></div>
					<div>
						<div className="bg-gray-300 dark:bg-gray-600 h-4 w-24 mb-2 rounded"></div>
						<div className="bg-gray-300 dark:bg-gray-600 h-3 w-16 rounded"></div>
					</div>
				</div>

				<div className="flex items-center gap-10">
					<div>
						<div className="bg-gray-300 dark:bg-gray-600 h-4 w-16 mb-2 rounded"></div>
						<div className="bg-gray-300 dark:bg-gray-600 h-3 w-20 rounded"></div>
					</div>
					<div className="hidden sm:block">
						<div className="bg-gray-300 dark:bg-gray-600 h-4 w-28 mb-2 rounded"></div>
						<div className="bg-gray-300 dark:bg-gray-600 h-3 w-24 rounded"></div>
					</div>
					<div className="hidden sm:block">
						<div className="bg-gray-300 dark:bg-gray-600 h-4 w-28 mb-2 rounded"></div>
						<div className="bg-gray-300 dark:bg-gray-600 h-3 w-20 rounded"></div>
					</div>
				</div>

				<div className="flex items-center p-2">
					<div className="bg-gray-300 dark:bg-gray-600 h-4 w-20 rounded"></div>
				</div>
			</div>
		</div>
	);
};

export default ShimmerCardList;
