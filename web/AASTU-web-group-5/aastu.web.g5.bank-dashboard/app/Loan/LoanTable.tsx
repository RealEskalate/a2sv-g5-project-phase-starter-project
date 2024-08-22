import React, { useState } from "react";
import { LoanDataProps } from "./loanTypes"; // Import the shared type

interface LoanTableProps {
	data: LoanDataProps[];
}

// Function to format large numbers
const formatNumber = (num: number) => {
	if (num >= 1_000_000) {
		return num.toExponential(2); // Format as exponential with 2 decimal places
	}
	return num.toLocaleString(); // Format as regular number with thousand separators
};

const LoanTable: React.FC<LoanTableProps> = ({ data }) => {
	// State for pagination
	const [currentPage, setCurrentPage] = useState(1);
	const itemsPerPage = 5; // Number of items per page

	// Calculate total pages
	const totalPages = Math.ceil(data.length / itemsPerPage);

	// Calculate the current page's data
	const startIndex = (currentPage - 1) * itemsPerPage;
	const paginatedData = data.slice(startIndex, startIndex + itemsPerPage);

	// Calculate totals
	const totalLoanMoney = data.reduce(
		(acc, loan) => acc + (loan.loanAmount || 0),
		0
	);
	const totalLeftToRepay = data.reduce(
		(acc, loan) => acc + (loan.leftToRepay || 0),
		0
	);
	const totalInstallment = data.reduce(
		(acc, loan) => acc + (loan.installment || 0),
		0
	);

	// Function to handle page change
	const handlePageChange = (page: number) => {
		if (page >= 1 && page <= totalPages) {
			setCurrentPage(page);
		}
	};

	return (
		<div className="py-5">
			<div className="p-3 text-[#333B69] text-xl font-semibold">
				<p>Active Loans Overview</p>
			</div>
			<div className="overflow-x-auto">
				<div className="p-6 bg-white rounded-3xl">
					<table className="w-full">
						<thead>
							<tr className="text-[#718EBF] border-b-2 border-solid border-gray-200">
								<th className="p-2 font-medium hidden sm:table-cell">SL No</th>
								<th className="p-2 font-medium">Loan Money</th>
								<th className="p-2 font-medium">Left to repay</th>
								<th className="p-2 font-medium hidden sm:table-cell">
									Duration
								</th>
								<th className="p-2 font-medium hidden sm:table-cell">
									Interest rate
								</th>
								<th className="p-2 font-medium hidden sm:table-cell">
									Installment
								</th>
								<th className="p-2 font-medium">Repay</th>
							</tr>
						</thead>
						<tbody>
							{paginatedData.map((loan, index) => (
								<tr
									key={loan.id}
									className="border-b-2 border-solid border-gray-100 text-[#232323] text-center"
								>
									<td className="p-3 hidden sm:table-cell">
										{startIndex + index + 1}.
									</td>
									<td className="p-3">${formatNumber(loan.loanAmount ?? 0)}</td>
									<td className="p-3">
										${formatNumber(loan.leftToRepay ?? 0)}
									</td>
									<td className="p-3 hidden sm:table-cell">
										{loan.duration ?? "0"} month
									</td>
									<td className="p-3 hidden sm:table-cell">
										{loan.interestRate ?? "0"}%
									</td>
									<td className="p-3 hidden sm:table-cell">
										${formatNumber(loan.installment ?? 0)} / month
									</td>
									<td className="p-3">
										<div className="text-center text-[#1814F3] border-[#1814F3] opacity-70 hover:opacity-100 sm:text-[#232323] border-2 border-solid sm:border-[#232323] cursor-pointer p-1 px-3 rounded-3xl sm:hover:border-[#1814F3] sm:hover:text-[#1814F3]">
											Repay
										</div>
									</td>
								</tr>
							))}
							<tr className="border-t-2 border-solid border-gray-100 text-center text-[#FE5C73] font-medium">
								<td className="p-2">Total</td>
								<td className="p-2">${formatNumber(totalLoanMoney)}</td>
								<td className="p-2">${formatNumber(totalLeftToRepay)}</td>
								<td className="p-2 hidden sm:table-cell"></td>
								<td className="p-2 hidden sm:table-cell"></td>
								<td className="p-2 hidden sm:table-cell">
									${formatNumber(totalInstallment)} / month
								</td>
								<td className="p-2 hidden sm:table-cell"></td>
							</tr>
						</tbody>
					</table>
					{/* Pagination Controls */}
					<div className="flex justify-center mt-4">
						<button
							className={`px-3 py-1 mx-1 border rounded-md ${
								currentPage === 1
									? "bg-gray-200 cursor-not-allowed"
									: "bg-white"
							}`}
							onClick={() => handlePageChange(currentPage - 1)}
							disabled={currentPage === 1}
						>
							Previous
						</button>
						{Array.from({ length: totalPages }, (_, index) => (
							<button
								key={index + 1}
								className={`px-3 py-1 mx-1 border rounded-md ${
									currentPage === index + 1
										? "bg-blue-500 text-white"
										: "bg-white"
								}`}
								onClick={() => handlePageChange(index + 1)}
							>
								{index + 1}
							</button>
						))}
						<button
							className={`px-3 py-1 mx-1 border rounded-md ${
								currentPage === totalPages
									? "bg-gray-200 cursor-not-allowed"
									: "bg-white"
							}`}
							onClick={() => handlePageChange(currentPage + 1)}
							disabled={currentPage === totalPages}
						>
							Next
						</button>
					</div>
				</div>
			</div>
		</div>
	);
};

export default LoanTable;
