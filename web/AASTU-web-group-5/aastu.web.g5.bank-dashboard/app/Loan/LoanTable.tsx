import React from "react";
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
							{data.map((loan, index) => (
								<tr
									key={loan.id}
									className="border-b-2 border-solid border-gray-100 text-[#232323] text-center"
								>
									<td className="p-3 hidden sm:table-cell">{index + 1}.</td>
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
				</div>
			</div>
		</div>
	);
};

export default LoanTable;
