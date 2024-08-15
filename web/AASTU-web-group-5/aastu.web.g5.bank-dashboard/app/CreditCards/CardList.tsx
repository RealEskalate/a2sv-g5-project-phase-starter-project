import React from "react";
import Image from "next/image";
import creditCardIcon from "@/public/assets/image/credit-card 1.png";

interface CardData {
	cardType: string;
	bank: string;
	cardNumber: string;
	namingCard: string;
}

interface CardListProps {
	cards: CardData[];
}

const CardList: React.FC<CardListProps> = ({ cards }) => {
	return (
		<>
			{cards.map((card, index) => (
				<div key={index} className="list pb-5">
					<div className="flex items-center gap-10 bg-white py-5 pr-12 pl-3 rounded-xl text-sm">
						<div className="flex gap-3">
							<div className="flex justify-center items-center bg-blue-100 w-12 h-12 rounded-xl ">
								<Image src={creditCardIcon} alt="credit Card Icon" />
							</div>
							<div>
								<div className="font-medium py-1">{card.cardType}</div>
								<div className="text-blue-900 opacity-70">Secondary</div>
							</div>
						</div>
						<div>
							<div className="font-medium py-1">Bank</div>
							<div className="text-blue-900 opacity-70">{card.bank}</div>
						</div>
						<div>
							<div className="font-medium py-1">Card Number</div>
							<div className="text-blue-900 opacity-70">{card.cardNumber}</div>
						</div>
						<div>
							<div className="font-medium py-1">Naming Card</div>
							<div className="text-blue-900 opacity-70">{card.namingCard}</div>
						</div>
						<div className="flex items-center">
							<span className="text-blue-600 font-medium">View Details</span>
						</div>
					</div>
				</div>
			))}
		</>
	);
};

export default CardList;
