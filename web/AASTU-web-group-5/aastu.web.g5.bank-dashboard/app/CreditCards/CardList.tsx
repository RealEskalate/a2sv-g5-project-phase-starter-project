import React, { useState, useEffect } from "react";
import Image from "next/image";
import { useSession } from "next-auth/react";
import { cardListMockData } from "./cardListMockData";
import ShimmerCardList from "./Shimmers/ShimmerCardList";

interface CardListData {
	id: string;
	balance: number;
	cardHolder: string;
	expiryDate: string;
	cardNumber: string;
	passcode: number;
	cardType: string;
	userId: string;
}

interface CardListProps {
	cardId: string[];
}

interface ExtendedUser {
	name?: string;
	email?: string;
	image?: string;
	accessToken?: string;
}

const CardList: React.FC<CardListProps> = ({ cardId }) => {
	const { data: session } = useSession();
	const user = session?.user as ExtendedUser;
	const [cardListData, setCardListData] = useState<CardListData[]>([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState<string | null>(null);
	const accessToken = user?.accessToken;

	const fetchCardListData = async () => {
		if (!accessToken) {
			setError("No access token available");
			setLoading(false);
			return;
		}

		try {
			const data = await Promise.all(
				cardId.map(async (id) => {
					const response = await fetch(
						`https://bank-dashboard-irbd.onrender.com/cards/${id}`,
						{
							headers: {
								Authorization: `Bearer ${accessToken}`,
							},
						}
					);

					if (!response.ok) {
						throw new Error(`Failed to fetch card with ID: ${id}`);
					}

					return response.json();
				})
			);

			setCardListData(data);
		} catch (error) {
			setError((error as Error).message);
		} finally {
			setLoading(false);
		}
	};

	useEffect(() => {
		fetchCardListData();
	}, [cardId]);

	const cardNumberFormat = (cardnumber: string) => {
		const lastFourNumber = cardnumber.slice(-4);
		return `**** **** **** ${lastFourNumber}`;
	};

	if (loading) {
		return (
			<div>
				{Array.from({ length: 3 }).map((_, index) => (
					<div key={index} className="list pb-2 p-1">
						<ShimmerCardList />
					</div>
				))}
			</div>
		);
	}

	if (error) {
		return <p className="py-5">Error: {error}</p>;
	}

	return (
		<div className="">
			{cardListData.map((card, index) => (
				<div key={card.id} className="list pb-2 p-1">
					<div className="flex sm:flex-row items-start sm:items-center gap-4 bg-white dark:bg-gray-800 dark:text-white py-5 px-5 pr-6 rounded-xl sm:justify-between">
						<div className="flex items-center gap-5">
							<div
								className={
									cardListMockData[index % cardListMockData.length].bgImgColor
								}
							>
								<Image
									src={
										cardListMockData[index % cardListMockData.length]
											.creditCardImg
									}
									alt="Credit Card Icon"
									width={30}
									height={30}
								/>
							</div>
							<div>
								<div className="font-medium py-1 ">{card.cardType}</div>
								<div className="text-blue-900 dark:text-gray-300 opacity-70">
									Secondary
								</div>
							</div>
						</div>

						<div className="flex items-center gap-10">
							<div>
								<div className="font-medium py-1">Bank</div>
								<div className="text-blue-900 dark:text-gray-300 opacity-70">
									{cardListMockData[index % cardListMockData.length].bank}
								</div>
							</div>
							<div className="hidden sm:block">
								<div className="font-medium py-1">Card Number</div>
								<div className="text-blue-900 dark:text-gray-300 opacity-70">
									{cardNumberFormat(card.cardNumber)}
								</div>
							</div>
							<div className="hidden sm:block">
								<div className="font-medium py-1">Naming Card</div>
								<div className="text-blue-900 dark:text-gray-300 opacity-70">
									{card.cardHolder}
								</div>
							</div>
						</div>

						<div className="flex items-center p-2">
							<p className="text-blue-600 font-medium">View Details</p>
						</div>
					</div>
				</div>
			))}
		</div>
	);
};

export default CardList;
