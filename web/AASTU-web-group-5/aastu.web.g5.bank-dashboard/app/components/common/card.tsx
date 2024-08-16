import React from "react";
import Image, { StaticImageData } from "next/image";

interface creditCardColorProps {
	cardBgColor: string;
	bottomBgColor: string;
	imageCreditCard: StaticImageData;
	grayCircleColor: boolean;
}

const Card = ({
	creditCardColor,
}: {
	creditCardColor: creditCardColorProps;
}) => {
	return (
		<div>
			<div className={creditCardColor.cardBgColor}>
				<div className="flex justify-between p-4 px-6">
					<div>
						<div className="text-sm opacity-70">Balance</div>
						<div className="text-lg">$5,756</div>
					</div>
					<div>
						<Image src={creditCardColor.imageCreditCard} alt="chip card" />
					</div>
				</div>
				<div className="flex gap-16 p-4 pr-28">
					<div className="pl-2">
						<div className="text-sm opacity-70">CARD HOLDER</div>
						<div>Eddy Cusuma</div>
					</div>
					<div>
						<div className="text-sm opacity-70">VALID THRU</div>
						<div>12/22</div>
					</div>
				</div>
				{creditCardColor.grayCircleColor ? (
					<div className={creditCardColor.bottomBgColor}>
						<div className="">
							<div className="border-2 border-solid border-gray-600 opacity-20"></div>
							<div className="flex justify-between p-4">
								<div className="text-xl">3778 *** *** 1234</div>
								<div className="flex">
									<div className="w-8 h-8 rounded-full bg-gray-400 opacity-50"></div>
									<div className="w-8 h-8 rounded-full bg-gray-400 -ml-4 opacity-50"></div>
								</div>
							</div>
						</div>
					</div>
				) : (
					<div className={creditCardColor.bottomBgColor}>
						<div className="text-xl">3778 *** *** 1234</div>
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
