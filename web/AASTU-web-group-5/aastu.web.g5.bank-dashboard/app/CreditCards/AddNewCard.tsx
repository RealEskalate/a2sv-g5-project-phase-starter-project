import React, { useState } from "react";

interface AddNewCardProps {
	accessToken: string;
	onCardAdded: () => void; // Add this prop
}

interface CardData {
	cardType: string;
	balance: number;
	cardHolder: string;
	expiryDate: string;
	passcode: string;
}

const AddNewCard: React.FC<AddNewCardProps> = ({
	accessToken,
	onCardAdded,
}) => {
	const [cardData, setCardData] = useState<CardData>({
		cardType: "",
		balance: 0,
		cardHolder: "",
		expiryDate: "",
		passcode: "",
	});
	const [error, setError] = useState<string | null>(null);
	const [success, setSuccess] = useState<string | null>(null);

	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const { name, value } = e.target;
		setCardData((prevData) => ({
			...prevData,
			[name]: value,
		}));
	};

	const handleAddCard = async (e: React.FormEvent<HTMLFormElement>) => {
		e.preventDefault();

		// Clear previous messages
		setError(null);
		setSuccess(null);

		// Validate input
		if (
			!cardData.cardType ||
			!cardData.balance ||
			!cardData.cardHolder ||
			!cardData.expiryDate ||
			!cardData.passcode
		) {
			setError("All fields are required.");
			return;
		}

		// Convert balance to number
		const numericBalance = parseFloat(
			cardData.balance.toString().replace(/[^0-9.-]+/g, "")
		);
		if (isNaN(numericBalance)) {
			setError("Invalid balance format.");
			return;
		}

		try {
			const response = await fetch(
				"https://bank-dashboard-1tst.onrender.com/cards",
				{
					method: "POST",
					headers: {
						"Content-Type": "application/json",
						Authorization: `Bearer ${accessToken}`,
					},
					body: JSON.stringify({
						...cardData,
						balance: numericBalance,
					}),
				}
			);

			if (!response.ok) {
				const errorData = await response.json();
				throw new Error(errorData.message || "Failed to add card");
			}

			setSuccess("Card added successfully!");

			// Clear the form
			setCardData({
				cardType: "",
				balance: 0,
				cardHolder: "",
				expiryDate: "",
				passcode: "",
			});

			// Call the onCardAdded callback to update the card list
			onCardAdded();
		} catch (error) {
			setError((error as Error).message);
		}
	};

	return (
		<div>
			<div className="p-2 font-semibold text-blue-900 w-full sm:w-1/2">
				Add New Card
			</div>
			<div className="bg-white rounded-2xl py-6">
				<div className="py-3 pl-8 pr-5 sm:pr-16 text-blue-900 opacity-70">
					Credit Card generally means a plastic card issued by Scheduled
					Commercial Banks assigned to a Cardholder, with a credit limit, that
					can be used to purchase goods and services on credit or obtain cash
					advances.
				</div>

				<div>
					<form onSubmit={handleAddCard}>
						<div className="flex flex-col sm:flex-row p-2 px-5 gap-4 sm:gap-8">
							<div className="flex-1 ">
								<div className="flex flex-col px-2">
									<label htmlFor="cardType" className="p-2">
										Card Type
									</label>
									<input
										type="text"
										id="cardType"
										name="cardType"
										value={cardData.cardType}
										onChange={handleChange}
										placeholder="Classic"
										className="border-2 border-solid border-blue-100 p-2 rounded-xl focus:outline-none focus:border-blue-300"
										aria-required="true"
									/>
								</div>
								<div className="flex flex-col px-2 ">
									<label htmlFor="balance" className="p-2">
										Balance
									</label>
									<input
										type="text"
										id="balance"
										name="balance"
										value={cardData.balance}
										onChange={handleChange}
										placeholder="27,000$"
										className="border-2 border-solid border-blue-100 p-2 rounded-xl focus:outline-none focus:border-blue-300"
										aria-required="true"
									/>
								</div>
								<div className="flex flex-col px-2">
									<label htmlFor="passcode" className="p-2">
										Passcode
									</label>
									<input
										type="password"
										id="passcode"
										name="passcode"
										value={cardData.passcode}
										onChange={handleChange}
										placeholder="••••••"
										className="border-2 border-solid border-blue-100 p-2 rounded-xl focus:outline-none focus:border-blue-300"
										aria-required="true"
									/>
								</div>
							</div>
							<div className="flex-1">
								<div className="flex flex-col px-2">
									<label htmlFor="cardHolder" className="p-2">
										Name On Card
									</label>
									<input
										type="text"
										id="cardHolder"
										name="cardHolder"
										value={cardData.cardHolder}
										onChange={handleChange}
										placeholder="My Cards"
										className="border-2 border-solid border-blue-100 p-2 rounded-xl focus:outline-none focus:border-blue-300"
										aria-required="true"
									/>
								</div>
								<div className="flex flex-col px-2">
									<label htmlFor="expiryDate" className="p-2">
										Expiry Date
									</label>
									<input
										type="date"
										id="expiryDate"
										name="expiryDate"
										value={cardData.expiryDate}
										onChange={handleChange}
										className="border-2 border-solid border-blue-100 p-2 rounded-xl focus:outline-none focus:border-blue-300"
										aria-required="true"
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

				{error && <p className="text-red-500 text-center ">{error}</p>}
				{success && <p className="text-green-500 text-center ">{success}</p>}
			</div>
		</div>
	);
};

export default AddNewCard;
