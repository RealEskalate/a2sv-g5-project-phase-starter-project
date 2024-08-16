import React from "react";

const AddNewCard = () => {
	return (
		<div className="px-5">
			<div className="p-2 font-semibold text-blue-900 w-full sm:w-1/2">
				Add New Card
			</div>
			<div className="bg-white rounded-2xl py-6 pr-5">
				<div className="py-3 pl-8 pr-5 sm:pr-20 text-blue-900 opacity-70">
					Credit Card generally means a plastic card issued by Scheduled
					Commercial Banks assigned to a Cardholder, with a credit limit, that
					can be used to purchase goods and services on credit or obtain cash
					advances.
				</div>
				<div>
					<form>
						<div className="flex flex-col sm:flex-row p-2 px-5 gap-4 sm:gap-8">
							<div className="flex-1">
								<div className="flex flex-col px-2">
									<label htmlFor="cardType" className="p-2">
										Card Type
									</label>
									<input
										type="text"
										placeholder="Classic"
										className="border-2 border-solid border-blue-100 p-2 rounded-xl focus:outline-none focus:border-blue-300"
									/>
								</div>
								<div className="flex flex-col px-2 mt-4">
									<label htmlFor="balance" className="p-2">
										Balance
									</label>
									<input
										type="text"
										placeholder="27,000$"
										className="border-2 border-solid border-blue-100 focus:outline-none focus:border-blue-300 p-2 rounded-xl"
									/>
								</div>
							</div>
							<div className="flex-1">
								<div className="flex flex-col px-2">
									<label htmlFor="nameOnCard" className="p-2">
										Name On Card
									</label>
									<input
										type="text"
										placeholder="My Cards"
										className="border-2 border-solid border-blue-100 focus:outline-none focus:border-blue-300 p-2 rounded-xl"
									/>
								</div>
								<div className="flex flex-col px-2 mt-4">
									<label htmlFor="expireDate" className="p-2">
										Expiration Date
									</label>
									<input
										type="date"
										value="2024-08-15"
										placeholder="25 January 2025"
										className="border-2 border-solid border-blue-100 focus:outline-none focus:border-blue-300 p-2 rounded-xl text-gray-400"
									/>
								</div>
							</div>
						</div>
						<div className="py-5 px-8">
							<div className="flex justify-center text-white font-semibold bg-blue-600 p-2.5 hover:bg-blue-700 w-full sm:w-1/5 rounded-lg">
								<button type="submit">Add Card</button>
							</div>
						</div>
					</form>
				</div>
			</div>
		</div>
	);
};

export default AddNewCard;
