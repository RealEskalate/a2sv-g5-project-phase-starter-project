"use client";
import React, { useEffect, useState } from "react";
import LoanCards from "./LoanCards";
import LoanTable from "./LoanTable";
import ShimmerLoanCards from "./ShimmerLoanCards";
import { FaGreaterThan, FaLessThan } from "react-icons/fa";
import { LoanDataProps } from "./loanTypes";
import { useSession } from "next-auth/react";
import ShimmerLoanTable from "./ShimmerLoanTable";

interface LoanResponse {
	success: boolean;
	message: string;
	data: {
		content: LoanDataProps[];
		totalPages: number;
	};
}
interface ExtendedUser {
	name?: string;
	email?: string;
	image?: string;
	accessToken?: string;
}
const LoanPage = () => {
	const [loanData, setLoanData] = useState<LoanResponse | null>(null);
	const [loading, setLoading] = useState(true);
	const [currentPage, setCurrentPage] = useState(0);

	const { data: session, status } = useSession();
	const user = session?.user as ExtendedUser;

	useEffect(() => {
		const fetchLoanData = async () => {
			if (!user) {
				console.error("Access token is missing");
				return;
			}
			try {
				const response = await fetch(
					`https://bank-dashboard-rsf1.onrender.com/active-loans/all?page=${currentPage}&size=8`,
					{
						headers: {
							Authorization: `Bearer ${user.accessToken}`,
						},
					}
				);

				if (!response.ok) {
					const errorText = await response.text();
					throw new Error(
						`Failed to fetch data: ${response.status} ${response.statusText} - ${errorText}`
					);
				}
				const data: LoanResponse = await response.json();
				setLoanData(data);
			} catch (error) {
				console.error("Error fetching loan data:", error);
			} finally {
				setLoading(false);
			}
		};

		fetchLoanData();
	}, [status, session, user, currentPage]);

	// Handle page change
	const handlePreviousPage = () => {
		if (currentPage > 0) {
			setCurrentPage(currentPage - 1);
		}
	};

	const handleNextPage = () => {
		if (loanData && currentPage < loanData.data.totalPages - 1) {
			setCurrentPage(currentPage + 1);
		}
	};

	return (
		<div className="p-5 sm:px-10 sm:py-8">
			{loading ? (
				<>
					<ShimmerLoanCards />
					<ShimmerLoanTable />
				</>
			) : (
				loanData &&
				loanData.success && (
					<>
						<LoanCards data={loanData.data.content} />
						<LoanTable data={loanData.data.content} />
						<div className="flex justify-end items-center px-3 text-sm">
							<div className="flex gap-2 items-center">
								{/* Previous Button */}
								<button
									onClick={handlePreviousPage}
									disabled={currentPage === 0}
									className={`flex items-center gap-1 text-[#1814F3] rounded ${
										currentPage === 0
											? "opacity-50 cursor-not-allowed"
											: "hover:text-blue-700"
									}`}
								>
									<FaLessThan />
									Previous
								</button>

								{/* Page numbers */}
								<div className="flex px-2 gap-1">
									{Array.from(
										{ length: loanData.data.totalPages },
										(_, index) => (
											<button
												key={index}
												onClick={() => setCurrentPage(index)}
												className={`px-4 py-2 rounded-xl ${
													currentPage === index
														? "bg-blue-500 text-white"
														: "text-[#1814F3] hover:bg-gray-200"
												}`}
											>
												{index + 1}
											</button>
										)
									)}
								</div>

								{/* Next Button */}
								<button
									onClick={handleNextPage}
									disabled={currentPage >= loanData.data.totalPages - 1}
									className={`flex items-center gap-1 text-[#1814F3] rounded ${
										currentPage >= loanData.data.totalPages - 1
											? "opacity-50 cursor-not-allowed"
											: "hover:text-blue-700"
									}`}
								>
									Next
									<FaGreaterThan />
								</button>
							</div>
						</div>
					</>
				)
			)}
		</div>
	);
};

export default LoanPage;
