import React from "react";
import Card from "../components/common/card";
import CardList from "./CardList";
import CardStatistics from "./CardStatistics";
import AddNewCard from "./AddNewCard";
import CardSetting from "./CardSetting";
import { cardData } from "./mockData";
import Image from "next/image";
import Chip_card1 from "@/public/assets/image/Chip_Card1.png";
import Chip_card2 from "@/public/assets/image/Chip_Card2.png";
import Chip_card3 from "@/public/assets/image/Chip_Card3.png";

const creditCardColor = {
	cardOne: {
		cardBgColor: "bg-blue-500 rounded-3xl text-white",
		bottomBgColor:
			"flex justify-between p-4 bg-blue-400 rounded-bl-3xl rounded-br-3xl",
		imageCreditCard: Chip_card1,
		grayCircleColor: false,
	},
	cardTwo: {
		cardBgColor: "bg-blue-700 rounded-3xl text-white",
		bottomBgColor:
			"flex justify-between p-4 bg-blue-600 rounded-bl-3xl rounded-br-3xl",
		imageCreditCard: Chip_card2,
		grayCircleColor: false,
	},
	cardThree: {
		cardBgColor: "bg-white rounded-3xl textblack",
		bottomBgColor: "",
		imageCreditCard: Chip_card3,
		grayCircleColor: true,
	},
};

const CreditCardComponent = () => {
	return (
		<div className="bg-[#F5F7FA] p-4 sm:p-8">
			{/* Row 1: Cards */}
			<div className="pl-0 sm:pl-4">
				<div className="font-semibold text-blue-900 p-2">My Cards</div>
				<div className="flex flex-col sm:flex-row gap-4 sm:gap-10">
					<Card creditCardColor={creditCardColor.cardOne} />
					<Card creditCardColor={creditCardColor.cardTwo} />
					<Card creditCardColor={creditCardColor.cardThree} />
				</div>
			</div>

			<div className="flex flex-col sm:flex-row gap-6 sm:gap-12 mt-6">
				<div>
					<CardStatistics />
				</div>
				<div className="flex flex-col pl-0 sm:pl-5 pt-4 sm:pt-2">
					<div className="font-semibold text-blue-900 p-3">Card List</div>
					<CardList cards={cardData} />
				</div>
			</div>

			<div className="flex flex-col sm:flex-row gap-6 sm:gap-8 mt-6 pb-10">
				<div>
					<AddNewCard />
				</div>
				<div>
					<CardSetting />
				</div>
			</div>
		</div>
	);
};

export default CreditCardComponent;
