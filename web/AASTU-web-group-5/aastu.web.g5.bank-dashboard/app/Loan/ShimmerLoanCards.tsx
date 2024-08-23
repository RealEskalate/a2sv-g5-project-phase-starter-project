import React from "react";

const ShimmerLoanCards: React.FC = () => {
	return (
		<div className="overflow-x-auto animate-pulse">
			<div className="flex flex-nowrap gap-7 text-sm">
				{[1, 2, 3, 4].map((index) => (
					<div
						key={index}
						className="flex items-center gap-4 bg-gray-200 sm:w-[48%] md:w-[30%] lg:w-[23%] p-6 rounded-3xl"
					>
						<div className="flex items-center justify-center rounded-full bg-gray-300 w-24 h-16"></div>
						<div className="w-full">
							<div className="h-4 bg-gray-300 rounded w-3/4 mb-2"></div>
							<div className="h-6 bg-gray-300 rounded w-1/2"></div>
						</div>
					</div>
				))}
			</div>
		</div>
	);
};

export default ShimmerLoanCards;
