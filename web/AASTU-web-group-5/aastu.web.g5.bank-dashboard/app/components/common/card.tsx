import React from "react";
import Image from "next/image";
import Chip_card from "@/public/assets/image/Chip_Card.png";

const card = () => {
	return (
		<>
			<div className="px-5 py-2">
				<div className="bg-blue-500 w-80 rounded-3xl text-white">
					<div className="flex justify-between p-4">
						<div>
							<div className="text-sm  opacity-70">Balance</div>
							<div className="text-lg">$5,756</div>
						</div>
						<div>
							<Image src={Chip_card} alt="chip chard" />
						</div>
					</div>
					<div className="flex gap-16 p-4">
						<div>
							<div className="text-sm opacity-70 ">CARD HOLDER</div>
							<div>Eddy Cusuma</div>
						</div>
						<div>
							<div className="text-sm opacity-70">VALID THRU</div>
							<div>12/22</div>
						</div>
					</div>
					<div className="flex justify-between p-4 bg-blue-400 rounded-bl-3xl rounded-br-3xl">
						<div className="text-xl">3778 *** *** 1234</div>
						<div className="flex ">
							<div className="w-8 h-8 rounded-full bg-gray-100 opacity-50"></div>
							<div className="w-8 h-8 rounded-full bg-gray-100 -ml-4 opacity-50"></div>
						</div>
					</div>
				</div>
			</div>
		</>
	);
};

export default card;
