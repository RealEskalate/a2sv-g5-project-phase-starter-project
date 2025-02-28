"use client";
import React, { useEffect, useState } from "react";
import Image from "next/image";
import axios from "axios";
import TotalAmmount_img from "@/public/assests/icon/Investments/Group303.png";
import Number_img from "@/public/assests/icon/Investments/Group305.png";
import Rate_img from "@/public/assests/icon/Investments/Group307.png";
import ChartCard_Invest from "./ChartCard_Invest";
import MonthlyRevenueChart from "./MonthlyRevenueChart";
import { tradingStockData, investmentsData } from "./mockData";
import { useSession } from "next-auth/react";
import Shimmer1 from "../Accounts/shimmer";
import { useSelector } from "react-redux";
import { RootState } from "@/app/redux/store";
interface ExtendedUser {
	name?: string;
	email?: string;
	image?: string;
	accessToken?: string;
}

const Investments = () => {
	const [loading, setLoading] = useState(true);
	const [investmentOverview, setInvestmentOverview] = useState({
		totalAmount: 0,
		numberOfInvestments: 0,
		rateOfReturn: 0,
	});

	const [yearlyTotalInvestment, setYearlyTotalInvestment] = useState([]);
	const [monthlyRevenue, setMonthlyRevenue] = useState([]);
	const { data: session } = useSession();
	const user = session.user as ExtendedUser;
	const darkmode = useSelector((state: RootState) => state.theme.darkMode);

	const token = user?.accessToken ? `Bearer ${user.accessToken}` : "";

	useEffect(() => {
		const fetchInvestmentData = async () => {
			if (!session || !session.user || !token) {
				setLoading(false);
				return;
			}

			setLoading(true); // Set loading to true before fetching data
			try {
				const response = await axios.get(
					"https://bank-dashboard-irbd.onrender.com/user/random-investment-data?years=3&months=5",
					{
						headers: {
							Authorization: token, // Make sure to add your token here
						},
					}
				);

				const {
					totalInvestment,
					rateOfReturn,
					yearlyTotalInvestment,
					monthlyRevenue,
				} = response.data.data;

				setInvestmentOverview({
					totalAmount: totalInvestment,
					numberOfInvestments: yearlyTotalInvestment.length,
					rateOfReturn: rateOfReturn,
				});

				setYearlyTotalInvestment(yearlyTotalInvestment);
				setMonthlyRevenue(monthlyRevenue);
			} catch (error) {
				console.error("Error fetching investment data:", error);
			} finally {
				setLoading(false); // Set loading to false after fetching data
			}
		};

		if (session && session.user && token) {
			fetchInvestmentData();
		}
	}, [session, token]); // Depend on session and token to re-fetch when they change

	return (
		<div className="bg-[#F5F7FA] dark:bg-gray-900  space-y-8 w-[95%] pt-3 overflow-hidden mx-auto">
			{/* Row 1: Investment Overview */}
			<div className="grid grid-cols-1 md:grid-cols-3 gap-4">
				{loading ? (
					<Shimmer1 />
				) : (
					<div
						className={`p-4 rounded-lg flex items-center justify-center space-x-4 ${
							darkmode ? "bg-gray-800 text-white" : "bg-white text-black"
						}`}
					>
						<Image
							height={44}
							width={44}
							src={TotalAmmount_img}
							alt="balance"
						/>
						<div>
							<p>Total Invested Amount</p>
							<p className="text-xl font-semibold">
								${investmentOverview.totalAmount.toFixed(2)}
							</p>
						</div>
					</div>
				)}

				{loading ? (
					<Shimmer1 />
				) : (
					<div
						className={`p-4 rounded-lg flex items-center justify-center space-x-4 ${
							darkmode ? "bg-gray-800 text-white" : "bg-white text-black"
						}`}
					>
						<Image height={44} width={44} src={Number_img} alt="balance" />
						<div>
							<p>Number of Investments</p>
							<p className="text-xl font-semibold">
								{investmentOverview.numberOfInvestments.toFixed(2)}
							</p>
						</div>
					</div>
				)}

				{loading ? (
					<Shimmer1 />
				) : (
					<div
						className={`p-4 rounded-lg flex items-center justify-center space-x-4 ${
							darkmode ? "bg-gray-800 text-white" : "bg-white text-black"
						}`}
					>
						<Image height={44} width={44} src={Rate_img} alt="balance" />
						<div>
							<p>Rate of Return</p>
							<p className="text-xl font-semibold">
								{investmentOverview.rateOfReturn.toFixed(2)}%
							</p>
						</div>
					</div>
				)}
			</div>

			{/* Row 2: Yearly Total Investment and Monthly Revenue */}
			<div className="grid grid-cols-1 md:grid-cols-2 gap-4">
				<div
					className={`p-4 rounded-lg ${
						darkmode ? "bg-gray-800 text-white" : "bg-gray-100 text-black"
					}`}
				>
					<p>Yearly Total Investment</p>

					{loading ? (
						<Shimmer1 />
					) : (
						<div
							className={`h-36 rounded mt-4 ${
								darkmode ? "bg-gray-700" : "bg-white"
							}`}
							style={{ width: "100%", height: 329 }}
						>
							<ChartCard_Invest data={yearlyTotalInvestment} />
						</div>
					)}
				</div>
				<div
					className={`p-4 rounded-lg ${
						darkmode ? "bg-gray-800 text-white" : "bg-gray-100 text-black"
					}`}
				>
					<p>Monthly Revenue</p>
					{loading ? (
						<Shimmer1 />
					) : (
						<div
							className={`h-36 rounded mt-4 ${
								darkmode ? "bg-gray-700" : "bg-white"
							}`}
							style={{ width: "100%", height: 329 }}
						>
							<MonthlyRevenueChart data={monthlyRevenue} />
						</div>
					)}
				</div>
			</div>

			{/* Row 3: Investments and Trading Stock */}
			<div className="flex flex-col md:flex-row gap-4">
				{/* Investments Section */}
				<div
					className={`md:w-[58%] p-4 rounded-lg min-h-[345px] ${
						darkmode ? "bg-gray-800 text-white" : "bg-gray-100 text-black"
					}`}
				>
					<p className="text-lg font-semibold">My Investments</p>
					<div className="space-y-4 mt-4">
						{investmentsData.slice(0, 3).map((investment) => (
							<div
								key={investment.id}
								className={`flex items-center space-x-4 p-2 rounded-lg shadow ${
									darkmode ? "bg-gray-700" : "bg-white"
								}`}
							>
								{loading ? (
									<Shimmer1 />
								) : (
									<>
										<Image
											src={investment.image}
											alt={investment.name}
											width={44}
											height={44}
											className="rounded-full object-cover"
										/>
										<div className="flex-1">
											<p className="font-semibold">{investment.name}</p>
											<p
												className={`text-gray-500 ${
													darkmode ? "text-gray-400" : ""
												}`}
											>
												{investment.service}
											</p>
										</div>
										<div>
											<p className="text-sm font-semibold">
												{investment.value}
											</p>
											<p
												className={`text-xs ${
													darkmode ? "text-gray-400" : "text-gray-500"
												}`}
											>
												Investment value
											</p>
										</div>
										<div>
											<div>
												{investment.return < 0 ? (
													<p className="text-red-500">{investment.return}%</p>
												) : (
													<p className="text-green-500">{investment.return}%</p>
												)}
											</div>
											<p
												className={`text-xs ${
													darkmode ? "text-gray-400" : "text-gray-500"
												}`}
											>
												Return
											</p>
										</div>
									</>
								)}
							</div>
						))}
					</div>
				</div>

				{/* Trading Stock Section */}
				<div
					className={`md:w-[42%] p-4 rounded-lg min-h-[345px] ${
						darkmode ? "bg-gray-800 text-white" : "bg-gray-100 text-black"
					}`}
				>
					<p className="text-lg font-semibold">Trading Stock</p>
					<div className="mt-4">
						<table
							className={`w-full rounded-lg shadow ${
								darkmode ? "bg-gray-700" : "bg-white"
							}`}
						>
							<thead>
								<tr className={`${darkmode ? "bg-gray-600" : "bg-gray-200"}`}>
									<th className="p-2">Sl.No</th>
									<th className="p-2">Name</th>
									<th className="p-2">Price</th>
									<th className="p-2">Return</th>
								</tr>
							</thead>
							<tbody>
								{tradingStockData.map((stock, index) => (
									<tr key={stock.id}>
										<td className="p-2">{index + 1}</td>
										<td className="p-2">{stock.name}</td>
										<td className="p-2">{stock.price}</td>
										<td className="p-2">
											{stock.return < 0 ? (
												<p className="text-red-500">{stock.return}%</p>
											) : (
												<p className="text-green-500">{stock.return}%</p>
											)}
										</td>
									</tr>
								))}
							</tbody>
						</table>
					</div>
				</div>
			</div>
		</div>
	);
};

export default Investments;
