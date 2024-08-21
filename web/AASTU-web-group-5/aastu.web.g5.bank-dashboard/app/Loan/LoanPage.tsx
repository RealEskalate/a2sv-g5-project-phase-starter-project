"use client";

import React, { useEffect, useState } from "react";
import LoanCards from "./LoanCards";
import LoanTable from "./LoanTable";
import { LoanDataProps } from "./loanTypes";
import { useSession } from "next-auth/react";

interface LoanResponse {
	success: boolean;
	message: string;
	data: LoanDataProps[];
}

interface ExtendedUser {
	name?: string;
	email?: string;
	image?: string;
	accessToken?: string;
}

const LoanPage = () => {
	const [loanData, setLoanData] = useState<LoanResponse | null>(null);
	const { data: session } = useSession();
	console.log(session, "session from loan");

	const user = session?.user as ExtendedUser | undefined;
	const accessToken = user?.accessToken;

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

		if (accessToken) {
			fetchLoanData();
		}
	}, [accessToken]);

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
