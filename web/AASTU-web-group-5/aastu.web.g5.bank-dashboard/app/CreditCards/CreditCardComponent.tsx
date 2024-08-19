"use client";

import React, { useEffect, useState } from "react";
import Card from "../components/common/card";
import CardList from "./CardList";
import CardStatistics from "./CardStatistics";
import AddNewCard from "./AddNewCard";
import CardSetting from "./CardSetting";
import creditCardColor from "./cardMockData";

const CreditCardComponent: React.FC = () => {
	const [cardData, setCardData] = useState<any[]>([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState<string | null>(null);
	const accessToken =
		"eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJvbGlrZWwiLCJpYXQiOjE3MjQwNjg1NzcsImV4cCI6MTcyNDE1NDk3N30.Pl5I-B9Afd9HQSJ2wFNZJdJ8nZ5qzIhgvcgyxcyOZr-wz7AEhtZ2Pn--AsdiOzt7";

	const fetchCardData = async () => {
		if (!accessToken) {
			setError("No access token available");
			setLoading(false);
			return;
		}

		try {
			const response = await fetch(
				"https://bank-dashboard-6acc.onrender.com/cards",
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
			setCardData(data);
		} catch (error) {
			setError((error as Error).message);
		} finally {
			setLoading(false);
		}
	};

	useEffect(() => {
		fetchCardData();
	}, []);

	const handleCardAdded = () => {
		// Re-fetch the card data to update the list
		fetchCardData();
	};

	if (loading) {
		return <p className="text-center py-5 text-blue-500 ">Loading...</p>;
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
						{creditCardColor.map((cardColor, index) => (
							<Card key={index} cardData={cardData[0]} cardColor={cardColor} />
						))}
					</div>
				</div>
			</div>

			<div className="flex flex-col sm:flex-row gap-6 sm:gap-12 mt-6">
				<div className="w-full sm:w-[33%]">
					<CardStatistics />
				</div>
				<div className="flex flex-col w-full sm:w-[67%]">
					<div className="font-semibold text-blue-900 p-3">Card List</div>
					<CardList cardId={cardData.map((card) => card.id)} />
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
