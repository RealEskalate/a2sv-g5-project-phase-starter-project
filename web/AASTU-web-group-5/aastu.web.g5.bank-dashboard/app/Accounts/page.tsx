"use client";
import React, { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import Image from "next/image";
import axios from "axios";
import Card from "../components/common/card";
import Shimmer1 from "./shimmer"; // Import the SkeletonCard component
import ChartCard from "./ChartCard";
import {
	balance as mockBalance,
	invoicesData as mockInvoicesData,
} from "./mockData";
import Balance_img from "@/public/assests/icon/Accounts/Group494.png";
import Income_img from "@/public/assests/icon/Accounts/Group400.png";
import Expense_img from "@/public/assests/icon/Accounts/Group402.png";
import TotalSaving_img from "@/public/assests/icon/Accounts/Group401.png";
import Shopping from "@/public/assests/icon/Accounts/Group328.png";
import Service from "@/public/assests/icon/Accounts/Group327.png";
import Transfer from "@/public/assests/icon/Accounts/user2.png";
import { useSession } from "next-auth/react";
import creditCardColor from "../CreditCards/cardMockData";
import { ClipLoader } from "react-spinners";
import { useSelector, useDispatch } from "react-redux";
import { RootState } from "@/app/redux/store";
interface ExtendedUser {
	name?: string;
	email?: string;
	image?: string;
	accessToken?: string;
}

const Accounts = () => {
	const [transactions, setTransactions] = useState<any[]>([]);
	const [balance, setBalance] = useState<number>(0);
	const [income, setIncome] = useState<number>(0);
	const [expense, setExpense] = useState<number>(0);
	const [netBalance, setNetBalance] = useState<number>(0);
	const [invoicesData, setInvoicesData] = useState<any[]>(mockInvoicesData); // Use mock data
	const [cardData, setCardData] = useState<any[]>([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState<string | null>(null);
	const { data: session, status } = useSession();
	const darkmode = useSelector((state: RootState) => state.theme.darkMode);

	const router = useRouter();

	useEffect(() => {
		const fetchData = async () => {
			if (!session || !session.user) {
				setError("No session or user available");
				setLoading(false);
				return;
			}

			const user = session.user as ExtendedUser;
			const accessToken = user.accessToken;

			try {
				await Promise.all([
					fetchTransactions(accessToken),
					fetchBalanceData(accessToken),
					fetchIncomes(accessToken),
					fetchExpenses(accessToken),
					fetchCardData(0, accessToken),
				]);
			} catch (error) {
				setError((error as Error).message);
			} finally {
				setLoading(false);
			}
		};

		if (session) {
			fetchData();
		} else {
			setLoading(true); // Ensure loading state is set while waiting for session
		}
	}, [session, income, expense]);

	const fetchTransactions = async (accessToken: string) => {
		try {
			const response = await axios.get(
				"https://bank-dashboard-rsf1.onrender.com/transactions?page=0&size=10",
				{
					headers: {
						Authorization: `Bearer ${accessToken}`,
					},
				}
			);

			if (response.data.success) {
				setTransactions(response.data.data.content.reverse().slice(0, 3));
			}
		} catch (error) {
			console.error("Error fetching transactions:", error);
		}
	};

	const fetchBalanceData = async (accessToken: string) => {
		try {
			const response = await axios.get(
				"https://bank-dashboard-rsf1.onrender.com/transactions/random-balance-history?monthsBeforeFirstTransaction=3",
				{
					headers: {
						Authorization: `Bearer ${accessToken}`,
					},
				}
			);

			if (response.data.success) {
				const latestBalance =
					response.data.data[response.data.data.length - 1]?.value;
				setBalance(Math.trunc(latestBalance));
				setNetBalance(Math.trunc(latestBalance + income - expense));
			}
		} catch (error) {
			console.error("Error fetching balance data:", error);
		}
	};

	const fetchIncomes = async (accessToken: string) => {
		try {
			const response = await axios.get(
				"https://bank-dashboard-rsf1.onrender.com/transactions/incomes?page=0&size=5",
				{
					headers: {
						Authorization: `Bearer ${accessToken}`,
					},
				}
			);

			if (response.data.success) {
				setIncome(
					response.data.data.content.reduce(
						(total: number, income: any) => total + income.amount,
						0
					)
				);
			}
		} catch (error) {
			console.error("Error fetching incomes:", error);
		}
	};

	const fetchExpenses = async (accessToken: string) => {
		try {
			const response = await axios.get(
				"https://bank-dashboard-rsf1.onrender.com/transactions/expenses?page=0&size=3",
				{
					headers: {
						Authorization: `Bearer ${accessToken}`,
					},
				}
			);

			if (response.data.success) {
				setExpense(
					response.data.data.content.reduce(
						(total: number, expense: any) => total + expense.amount,
						0
					)
				);
			}
		} catch (error) {
			console.error("Error fetching expenses:", error);
		}
	};

	const fetchCardData = async (page: number, accessToken: string) => {
		try {
			const response = await fetch(
				`https://bank-dashboard-rsf1.onrender.com/cards?page=${page}&size=3`,
				{
					headers: {
						Authorization: `Bearer ${accessToken}`,
					},
				}
			);

			if (!response.ok) {
				throw new Error("Failed to fetch cards");
			}

			const data = await response.json();
			setCardData(data.content || []);
		} catch (error) {
			setError((error as Error).message);
		} finally {
			setLoading(false);
		}
	};

	if (error) {
		return <p className="py-5">Error: {error}</p>;
	}

	return (
		<div className="bg-[#F5F7FA] dark:bg-gray-900 space-y-8 w-[95%] pt-3 overflow-hidden mx-auto">
			{/* Balance and Overview */}
			<div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
				{loading ? (
					<Shimmer1 />
				) : (
					<div
						className={`p-4 rounded-lg flex items-center justify-center space-x-4 ${
							darkmode ? "bg-gray-800 text-white" : "bg-white text-black"
						}`}
					>
						<Image height={44} width={44} src={Balance_img} alt="balance" />
						<div>
							<p className="text-sm md:text-lg lg:text-xl font-semibold">
								My balance
							</p>
							<p className="text-base md:text-xl lg:text-2xl break-all">
								${Math.round(balance)}
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
						<Image height={44} width={44} src={Income_img} alt="income" />
						<div>
							<p className="text-sm md:text-lg lg:text-xl font-semibold">
								Income
							</p>
							<p className="text-base md:text-xl lg:text-2xl break-all">
								${Math.round(income)}
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
						<Image height={44} width={44} src={Expense_img} alt="expense" />
						<div>
							<p className="text-sm md:text-lg lg:text-xl font-semibold">
								Expense
							</p>
							<p className="text-base md:text-xl lg:text-2xl break-all">
								${Math.round(expense)}
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
						<Image
							height={44}
							width={44}
							src={TotalSaving_img}
							alt="total saving"
						/>
						<div>
							<p className="text-sm md:text-lg lg:text-xl font-semibold">
								Total saving
							</p>
							<p className="text-base md:text-xl lg:text-2xl break-all">
								${Math.round(netBalance)}
							</p>
						</div>
					</div>
				)}
			</div>

			{/* Last Transactions and Card */}
			<div className="grid grid-cols-1 lg:grid-cols-3 gap-4">
				{/* Last Transactions */}
				<div
					className={`p-4 rounded-lg lg:col-span-2 space-y-7 ${
						darkmode ? "bg-gray-800 text-white" : "bg-white text-black"
					}`}
				>
					<p className="text-lg font-semibold">Last transactions</p>
					{loading ? (
						<>
							<Shimmer1 />
							<Shimmer1 />
							<Shimmer1 />
						</>
					) : (
						<>
							{transactions.map((transaction, index) => (
								<div
									key={index}
									className="flex items-center pr-4 space-x-4 mb-4"
								>
									<Image
										height={44}
										width={44}
										src={
											transaction.type === "shopping"
												? Shopping
												: transaction.type === "service"
												? Service
												: Transfer
										}
										alt="transaction"
										className="rounded-full object-cover"
									/>
									<div className="flex-1">
										<p className="font-semibold text-sm md:text-base">
											{transaction.description}
										</p>
										<p className="text-xs md:text-sm text-gray-500">
											{transaction.date}
										</p>
									</div>
									<p className="flex-1 text-xs md:text-sm break-words">
										{transaction.type}
									</p>
									<p className="flex-1 text-xs md:text-sm break-words">
										{transaction.transactionId
											? `${transaction.transactionId.slice(0, 4)}***`
											: "N/A"}
									</p>
									<p className="flex-1 text-xs md:text-sm break-words">
										{transaction.receiverUserName}
									</p>
									<p className="flex-1 text-xs md:text-sm break-words">
										{transaction.amount < 0 ? (
											<span className="text-red-500">
												-${Math.abs(transaction.amount)}
											</span>
										) : (
											<span className="text-green-500">
												+${transaction.amount}
											</span>
										)}
									</p>
								</div>
							))}
						</>
					)}
				</div>

				{/* Hidden on small screens */}
				{/* Card */}
				<div
					className={`rounded-lg ${
						darkmode ? "bg-gray-800 text-white" : "bg-white text-black"
					}`}
				>
					<div className="flex justify-between p-2">
						<p className="text-lg font-semibold">My card</p>
						<button
							onClick={() => router.push("/CreditCards")}
							className="text-blue-500"
						>
							See All
						</button>
					</div>
					{loading ? (
						<Shimmer1 />
					) : (
						<div className="flex flex-col">
							{cardData.length > 0 ? (
								<Card
									key={0}
									cardData={cardData[0]}
									cardColor={creditCardColor[0 % creditCardColor.length]}
								/>
							) : (
								<div className="flex justify-center items-center h-full">
									<ClipLoader size={50} color={"#123abc"} loading={loading} />
								</div>
							)}
						</div>
					)}
				</div>
			</div>

			{/* Debit and Credit Overview */}
			<div className="grid grid-cols-1 lg:grid-cols-3 gap-4">
				<div
					className={`rounded-lg lg:col-span-2 ${
						darkmode ? "bg-gray-800 text-white" : "bg-gray-100 text-black"
					}`}
				>
					<p className="text-lg font-semibold">Debit and credit overview</p>
					{loading ? (
						<Shimmer1 />
					) : (
						<>
							<div
								className={`p-1 mt-3 rounded-lg shadow ${
									darkmode ? "bg-gray-700" : "bg-white"
								}`}
							>
								<ChartCard />
							</div>
						</>
					)}
				</div>

				{/* Invoices Sent */}
				<div
					className={`w-full rounded-lg ${
						darkmode ? "bg-gray-800 text-white" : "bg-gray-100 text-black"
					}`}
				>
					<p className="text-lg font-semibold">Invoices sent</p>
					{loading ? (
						<Shimmer1 />
					) : (
						<>
							<div
								className={`p-4 mt-4 rounded-lg shadow h-[364px] ${
									darkmode ? "bg-gray-700" : "bg-white"
								}`}
							>
								{invoicesData.slice(0, 4).map((data, index) => (
									<div
										key={index}
										className="flex items-center justify-between space-x-8 mb-4"
									>
										<div className="flex items-center justify-between space-x-8 mb-4">
											<Image
												height={44}
												width={44}
												src={data.image}
												alt="invoice"
												className="rounded-full object-cover"
											/>
											<div>
												<p className="font-semibold text-sm md:text-base">
													{data.name}
												</p>
												<p className="text-xs md:text-sm text-gray-500">
													{data.date}
												</p>
											</div>
										</div>
										<p className="font-semibold text-sm md:text-base">
											{data.amount}
										</p>
									</div>
								))}
							</div>
						</>
					)}
				</div>
			</div>
		</div>
	);
};

export default Accounts;
