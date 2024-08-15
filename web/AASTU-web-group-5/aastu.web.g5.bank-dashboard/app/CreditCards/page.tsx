import React from "react";
import Card from "../components/common/card";
import CardList from "./CardList";
import CardStatistics from "./CardStatistics";
import AddNewCard from "./AddNewCard";
import CardSetting from "./CardSetting";
import { cardData } from "./mockData";

const Page = () => {
	return (
		<>
			<div className="p-5">
				<div>
					<div className="pl-8 font-semibold text-blue-900">My Cards</div>
					<div className="flex">
						<Card />
						<Card />
						<Card />
					</div>
				</div>
				<div className="flex gap-12">
					<CardStatistics />
					<div className="flex flex-col pt-2">
						<div className="font-semibold text-blue-900 p-1">Card List</div>
						<CardList cards={cardData} /> {/* Pass the mock data as props */}
					</div>
				</div>
				<div className="flex">
					<div>
						<AddNewCard />
					</div>
					<div>
						<CardSetting />
					</div>
				</div>
			</div>
		</>
	);
};

export default Page;
