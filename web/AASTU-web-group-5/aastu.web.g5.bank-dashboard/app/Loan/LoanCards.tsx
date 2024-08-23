import React from "react";
import Image from "next/image";
import UserImg from "@/public/assets/image/user 3 2.png";
import bagImg from "@/public/assets/image/briefcase 1.png";
import graphImg from "@/public/assets/image/graph 1.png";
import supportImg from "@/public/assets/image/support 1.png";
import { LoanDataProps } from "./loanTypes"; // Import the shared type

interface LoanCardsProps {
	data: LoanDataProps[];
}

const LoanCards: React.FC<LoanCardsProps> = ({ data = [] }) => {
	if (!Array.isArray(data)) {
		console.error("Expected 'data' to be an array but got:", data);
		return <div>Error loading loan data</div>;
	}

	const personalLoan = data.find((loan) => loan.type === "personal");
	const corporateLoan = data.find((loan) => loan.type === "corporate");
	const businessLoan = data.find((loan) => loan.type === "business");
	const customLoan = data.find((loan) => loan.type === "custom");

	return (
		<div className="overflow-x-auto">
			<div className="flex flex-nowrap gap-7 text-sm">
				{/* Personal Loan Card */}
				<div className="flex items-center gap-4 bg-white dark:bg-gray-900 sm:w-[48%] md:w-[30%] lg:w-[23%] p-6 rounded-3xl">
					<div className="flex items-center justify-center rounded-full bg-[#E7EDFF] dark:opacity-90 w-24 h-16">
						<Image src={UserImg} alt="Personal Loan" />
					</div>
					<div className="w-full">
						<p className="text-[#718EBF] dark:text-gray-400 ">Personal Loans</p>
						<p className="text-[#232323] dark:text-[#fff] text-2xl font-medium">
							${personalLoan ? personalLoan.loanAmount.toLocaleString() : "0"}
						</p>
					</div>
				</div>
				{/* Corporate Loan Card */}
				<div className="flex items-center gap-4 bg-white dark:bg-gray-900 sm:w-[48%] md:w-[30%] lg:w-[23%] p-6 rounded-3xl">
					<div className="flex items-center justify-center rounded-full bg-[#FFF5D9] dark:opacity-90 w-24 h-16">
						<Image src={bagImg} alt="Corporate Loan" />
					</div>
					<div className="w-full">
						<p className="text-[#718EBF] dark:text-gray-400 ">
							Corporate Loans
						</p>
						<p className="text-[#232323] dark:text-[#fff] text-2xl font-medium">
							${corporateLoan ? corporateLoan.loanAmount.toLocaleString() : "0"}
						</p>
					</div>
				</div>
				{/* Business Loan Card */}
				<div className="flex items-center gap-4 bg-white dark:bg-gray-900 sm:w-[48%] md:w-[30%] lg:w-[23%] p-6 rounded-3xl">
					<div className="flex items-center justify-center rounded-full bg-[#FFE0EB] dark:opacity-90 w-24 h-16">
						<Image src={graphImg} alt="Business Loan" />
					</div>
					<div className="w-full">
						<p className="text-[#718EBF] dark:text-gray-400 ">Business Loans</p>
						<p className="text-[#232323] dark:text-[#fff] text-2xl font-medium">
							${businessLoan ? businessLoan.loanAmount.toLocaleString() : "0"}
						</p>
					</div>
				</div>
				{/* Custom Loan Card */}
				<div className="flex items-center gap-4 bg-white dark:bg-gray-900 sm:w-[48%] md:w-[30%] lg:w-[23%] p-6 rounded-3xl">
					<div className="flex items-center justify-center rounded-full bg-[#DCFAF8] dark:opacity-90 w-24 h-16">
						<Image src={supportImg} alt="Custom Loan" />
					</div>
					<div className="w-full">
						<p className="text-[#718EBF] dark:text-gray-400 ">Custom Loans</p>
						<p className="text-[#232323] dark:text-[#fff] text-base font-semibold">
							Choose Money
						</p>
					</div>
				</div>
			</div>
		</div>
	);
};

export default LoanCards;
