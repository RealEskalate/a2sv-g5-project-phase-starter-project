import React from "react";
import activeLoansData from "./activeLoanMockData"; // Adjust the path as necessary

const LoanTable = () => {
	const totalLoanMoney = activeLoansData.reduce(
		(acc, loan) => acc + loan.money,
		0
	);
	const totalLeftToRepay = activeLoansData.reduce(
		(acc, loan) => acc + loan.leftToRepay,
		0
	);
	const totalInstallment = activeLoansData.reduce(
		(acc, loan) => acc + loan.installment,
		0
	);

	return (
		<div className="py-5 ">
			<div className="p-3 text-[#333B69] text-xl font-semibold ">
				<p>Active Loans Overview</p>
			</div>
			<div className="overflow-x-auto">
				<div className=" p-10 bg-white rounded-3xl  ">
					<table className="w-full">
						<thead>
							<tr className="text-[#718EBF] border-b-2 border-solid border-gray-200 ">
								<th className="p-2 font-medium hidden sm:table-cell ">SL No</th>
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
								<th className="p-2 font-medium ">Repay</th>
							</tr>
						</thead>
						<tbody>
							{activeLoansData.map((loan, index) => (
								<tr
									key={loan.id}
									className="border-b-2 border-solid border-gray-100  text-[#232323] text-center "
								>
									<td className="p-2 hidden sm:table-cell">
										{loan.id < 10 ? `0${loan.id}` : loan.id}.
									</td>
									<td className="p-2">${loan.money.toLocaleString()}</td>
									<td className="p-2">${loan.leftToRepay.toLocaleString()}</td>
									<td className="p-2 hidden sm:table-cell">{loan.duration}</td>
									<td className="p-2 hidden sm:table-cell">
										{loan.interestrate}%
									</td>
									<td className="p-2 hidden sm:table-cell">
										${loan.installment.toLocaleString()} / month
									</td>
									<td className="p-2">
										<div className="text-center border-2 border-solid border-[#232323] cursor-pointer p-1 px-3 rounded-3xl hover:border-[#1814F3] hover:text-[#1814F3]">
											Repay
										</div>
									</td>
								</tr>
							))}
							<tr className="border-t-2 border-solid border-gray-100 text-center text-[#FE5C73] font-medium ">
								<td className="p-2 ">Total</td>
								<td className="p-2">${totalLoanMoney.toLocaleString()}</td>
								<td className="p-2 ">${totalLeftToRepay.toLocaleString()}</td>
								<td className="p-2 hidden sm:table-cell"></td>
								<td className="p-2 hidden sm:table-cell"></td>
								<td className="p-2 hidden sm:table-cell">
									${totalInstallment.toLocaleString()} / month
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
