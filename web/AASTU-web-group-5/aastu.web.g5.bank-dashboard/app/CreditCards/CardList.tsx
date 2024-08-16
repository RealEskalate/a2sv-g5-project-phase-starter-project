import React from "react";
import Image, { StaticImageData } from "next/image";

interface CardData {
	cardType: string;
	bank: string;
	cardNumber: string;
	namingCard: string;
	creditCardImg: StaticImageData;
	bgImgColor: string;
}

interface CardListProps {
	cards: CardData[];
}

const CardList: React.FC<CardListProps> = ({ cards }) => {
	return (
		<>
			{cards.map((card, index) => (
				<div key={index} className="list pb-5">
					<div className="flex sm:flex-row items-start sm:items-center gap-4 lg:gap-24 bg-white py-5 px-5 pr-6 rounded-xl sm:justify-between">
						<div className="flex items-center gap-5">
							<div className={card.bgImgColor}>
								<Image src={card.creditCardImg} alt="credit Card Icon" />
							</div>
							<div>
								<div className="font-medium">{card.cardType}</div>
								<div className="text-blue-900 opacity-70 ">Secondary</div>
							</div>
						</div>

						<div className="flex items-center gap-5">
							<div>
								<div className="font-medium">Bank</div>
								<div className="text-blue-900 opacity-70">{card.bank}</div>
							</div>
							<div className="hidden sm:block">
								<div className="font-medium">Card Number</div>
								<div className="text-blue-900 opacity-70">
									{card.cardNumber}
								</div>
							</div>
							<div className="hidden sm:block">
								<div className="font-medium py-1">Naming Card</div>
								<div className="text-blue-900 opacity-70">
									{card.namingCard}
								</div>
							</div>
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
