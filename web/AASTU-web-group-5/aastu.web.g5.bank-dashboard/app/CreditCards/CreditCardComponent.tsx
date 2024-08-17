"use client";

import React from "react";
import Card from "../components/common/card";
import CardList from "./CardList";
import CardStatistics from "./CardStatistics";
import AddNewCard from "./AddNewCard";
import CardSetting from "./CardSetting";
import { cardData } from "./cardMockData";
import { cardListData } from "./cardListMockData";

const CreditCardComponent = () => {
	return (
		<div className="bg-[#F5F7FA] p-4 sm:p-8">
			{/* Row 1: Cards */}
			<div>
				<div className="font-semibold text-blue-900 p-2">My Cards</div>
				<div className="overflow-x-auto">
					<div className="flex flex-nowrap gap-4 sm:gap-10 w-[920px] sm:w-[1200px]">
						{cardData.map((card, index) => (
							<Card key={index} card={card} />
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
					<CardList cards={cardListData} />
				</div>
			</div>

			<div className="flex flex-col sm:flex-row gap-6 py-3">
				<div className="w-full sm:w-[67%]">
					<AddNewCard />
				</div>
				<div className="w-full sm:w-[33%]">
					<CardSetting />
				</div>
			</div>
		</div>
	);
};

export default CreditCardComponent;
