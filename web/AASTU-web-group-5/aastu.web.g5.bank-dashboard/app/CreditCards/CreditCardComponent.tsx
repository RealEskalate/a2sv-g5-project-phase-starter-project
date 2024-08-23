"use client";

import React, { useEffect, useState } from "react";
import {FaGreaterThan,FaLessThan} from 'react-icons/fa'
import Card from "../components/common/card";
import CardList from "./CardList";
import CardStatistics from "./CardStatistics";
import AddNewCard from "./AddNewCard";
import CardSetting from "./CardSetting";
import creditCardColor from "./cardMockData";
import { useSession } from "next-auth/react";
interface ExtendedUser {
    name?: string;
    email?: string;
    image?: string;
    accessToken?: string;
    }
const CreditCardComponent: React.FC = () => {
	const [cardData, setCardData] = useState<any[]>([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState<string | null>(null);
	const [currentPage, setCurrentPage] = useState(0);
	const [totalPages, setTotalPages] = useState(0);
	const { data: session, status } = useSession();
	const user = session?.user as ExtendedUser;
  
	const accessToken = user.accessToken;

	const fetchCardData = async (page: number) => {
		if (!accessToken) {
			setError("No access token available");
			setLoading(false);
			return;
		}

		try {
			const response = await fetch(
				`https://bank-dashboard-1tst.onrender.com/cards?page=${page}&size=3`,
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
			setTotalPages(data.totalPages || 0);
		} catch (error) {
			setError((error as Error).message);
		} finally {
			setLoading(false);
		}
	};

	useEffect(() => {
		if (accessToken) {
			fetchCardData(currentPage);
		}
	}, [accessToken, currentPage]);

	const handleCardAdded = () => {
		fetchCardData(currentPage);
	};

	const handleNextPage = () => {
		if (currentPage < totalPages - 1) {
			setCurrentPage((prevPage) => prevPage + 1);
		}
	};

	const handlePreviousPage = () => {
		if (currentPage > 0) {
			setCurrentPage((prevPage) => prevPage - 1);
		}
	};

	if (loading) {
		return <p className="text-center py-5 text-blue-500">Loading...</p>;
	}
	if (error) {
		return <p className="py-5">Error: {error}</p>;
	}

	return (
		<div className="bg-[#F5F7FA] p-4 sm:px-8 sm:py-4">
			<div>
				<div className="font-semibold text-blue-900 p-2">My Cards</div>
				<div className="overflow-x-auto">
					<div className="flex flex-nowrap gap-6 w-[940px] sm:w-[1024px] md:w-full">
						{cardData.map((card, index) => (
							<Card
								key={index}
								cardData={card}
								cardColor={creditCardColor[index % creditCardColor.length]}
							/>
						))}
					</div>
				</div>
			</div>

			<div className="flex flex-col sm:flex-row gap-6 sm:gap-12 mt-6 pb-2">
				<div className="w-full sm:w-[33%]">
					<CardStatistics />
				</div>
				<div className="flex flex-col w-full sm:w-[67%] ">
					<div className="font-semibold text-blue-900 p-3">Card List</div>
					<CardList cardId={cardData.map((card) => card.id)} />
					<div className="flex justify-end  items-center px-3 text-sm">
						<div className="flex gap-1 items-center ">
							<FaLessThan className="text-[#1814F3]  opacity-60   " />
							<button
								onClick={handlePreviousPage}
								disabled={currentPage === 0}
								className=" text-[#1814F3] rounded "
							>
								Previous
							</button>
						</div>

						{/* Page numbers */}
						<div className="flex px-2">
							{Array.from({ length: totalPages }, (_, index) => (
								<button
									key={index}
									onClick={() => setCurrentPage(index)}
									className={`px-4 py-2 rounded-xl ${
										currentPage === index
											? "bg-blue-500 text-white"
											: " text-[#1814F3]"
									}`}
								>
									{index + 1}
								</button>
							))}
						</div>
						<div className="flex gap-1 items-center ">
							<button
								onClick={handleNextPage}
								disabled={currentPage >= totalPages - 1}
								className="  text-[#1814F3] rounded"
							>
								Next
							</button>
							<FaGreaterThan className="text-[#1814F3]  opacity-60 " />
						</div>
					</div>
				</div>
			</div>

			<div className="flex flex-col sm:flex-row gap-6 py-3">
				<div className="w-full sm:w-[67%]">
					{accessToken && (
						<AddNewCard
							accessToken={accessToken}
							onCardAdded={handleCardAdded}
						/>
					)}
				</div>
				<div className="w-full sm:w-[33%]">
					<CardSetting />
				</div>
			</div>
		</div>
	);
};

export default CreditCardComponent;
