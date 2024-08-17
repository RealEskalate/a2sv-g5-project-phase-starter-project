import React from "react";
import Image, { StaticImageData } from "next/image";

interface CreditCardColorProps {
	cardBgColor: string;
	bottomBgColor: string;
	imageCreditCard: StaticImageData;
	grayCircleColor: boolean;
}

interface CardProps {
	balance: number;
	cardHolder: string;
	expiryDate: string;
	semiCardNumber: string;
	creditCardColor: CreditCardColorProps;
}

const Card = ({ card }: { card: CardProps }) => {
	return (
		<div className="w-[33%]">
			<div className={card.creditCardColor.cardBgColor}>
				<div className="flex justify-between p-5">
					<div>
						<div className="text-sm opacity-70">Balance</div>
						<div className="text-lg">${card.balance}</div>
					</div>
					<div>
						<Image src={card.creditCardColor.imageCreditCard} alt="chip card" />
					</div>
				</div>
				<div className="flex gap-16 p-4">
					<div className="pl-2">
						<div className="text-sm opacity-70">CARD HOLDER</div>
						<div>{card.cardHolder}</div>
					</div>
					<div>
						<div className="text-sm opacity-70">VALID THRU</div>
						<div>{card.expiryDate}</div>
					</div>
				</div>
				{card.creditCardColor.grayCircleColor ? (
					<div className={card.creditCardColor.bottomBgColor}>
						<div>
							<div className="border-b-2 border-solid border-gray-600 opacity-20"></div>
							<div className="flex justify-between p-4">
								<div className="text-xl">{card.semiCardNumber}</div>
								<div className="flex">
									<div className="w-8 h-8 rounded-full bg-gray-400 opacity-50"></div>
									<div className="w-8 h-8 rounded-full bg-gray-400 -ml-4 opacity-50"></div>
								</div>
							</div>
						</div>
					</div>
				) : (
					<div className={card.creditCardColor.bottomBgColor}>
						<div className="text-xl">{card.semiCardNumber}</div>
						<div className="flex">
							<div className="w-8 h-8 rounded-full bg-gray-100 opacity-50"></div>
							<div className="w-8 h-8 rounded-full bg-gray-100 -ml-4 opacity-50"></div>
						</div>
					</div>
				)}
			</div>
		</div>
	);
};

export default Card;
