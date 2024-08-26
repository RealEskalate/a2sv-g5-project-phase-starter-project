import React from "react";

const ShimmerLoanTable: React.FC = () => {
	return (
		<div className="py-5 animate-pulse">
			<div className="p-3 bg-gray-200 dark:bg-gray-600 rounded-xl w-1/4 h-6 mb-4"></div>
			<div className="overflow-x-auto">
				<div className="p-6 bg-white dark:bg-[#0f0f0f] rounded-3xl">
					<table className="w-full">
						<thead>
							<tr className="text-gray-300 border-b-2 border-solid border-gray-200 dark:border-[#0f0f0f]">
								{[
									"SL No",
									"Loan Money",
									"Left to repay",
									"Duration",
									"Interest rate",
									"Installment",
									"Repay",
								].map((header, index) => (
									<th
										key={index}
										className="p-2 font-medium hidden sm:table-cell"
									>
										<div className="h-4 bg-gray-200 dark:bg-gray-600 rounded"></div>
									</th>
								))}
							</tr>
						</thead>
						<tbody>
							{[1, 2, 3, 4, 5].map((_, index) => (
								<tr
									key={index}
									className="border-b-2 border-solid border-gray-100 dark:border-[#0f0f0f]"
								>
									{[1, 2, 3, 4, 5, 6, 7].map((cellIndex) => (
										<td key={cellIndex} className="p-3">
											<div className="h-4 bg-gray-200 dark:bg-gray-600 rounded"></div>
										</td>
									))}
								</tr>
							))}
						</tbody>
					</table>
				</div>
			</div>
		</div>
	);
};

export default ShimmerLoanTable;
