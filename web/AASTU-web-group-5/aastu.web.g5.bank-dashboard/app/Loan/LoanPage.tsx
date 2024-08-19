"use client";

import React, { useEffect, useState } from "react";
import LoanCards from "./LoanCards";
import LoanTable from "./LoanTable";
import { LoanDataProps } from "./loanTypes"; // Import the shared type

interface LoanResponse {
	success: boolean;
	message: string;
	data: LoanDataProps[];
}

const LoanPage = () => {
	const [loanData, setLoanData] = useState<LoanResponse | null>(null);
	const accessToken =
		"eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJvbGlrZWwiLCJpYXQiOjE3MjQwNjg1NzcsImV4cCI6MTcyNDE1NDk3N30.Pl5I-B9Afd9HQSJ2wFNZJdJ8nZ5qzIhgvcgyxcyOZr-wz7AEhtZ2Pn--AsdiOzt7";

	useEffect(() => {
		const fetchLoanData = async () => {
			if (!accessToken) {
				console.error("Access token is missing");
				return;
			}
			try {
				const response = await fetch(
					"https://bank-dashboard-6acc.onrender.com/active-loans/all",
					{
						headers: {
							Authorization: `Bearer ${accessToken}`,
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
			}
		};

		fetchLoanData();
	}, [accessToken]);

	console.log("loannnnnnnnnn:", loanData);

	return (
		<div className="p-5 sm:px-10 sm:py-8">
			{loanData && loanData.success && (
				<>
					<LoanCards data={loanData.data} />
					<LoanTable data={loanData.data} />
				</>
			)}
		</div>
	);
};

export default LoanPage;
