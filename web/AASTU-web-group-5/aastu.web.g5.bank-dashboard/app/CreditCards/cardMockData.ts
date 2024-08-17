import Chip_card1 from "@/public/assets/image/Chip_Card1.png";
import Chip_card2 from "@/public/assets/image/Chip_Card2.png";
import Chip_card3 from "@/public/assets/image/Chip_Card3.png";

export const cardData = [
	{
		balance: 5678.9,
		cardHolder: "John Doe",
		expiryDate: "08/24",
		semiCardNumber: "123 **** **** 1234",
		creditCardColor: {
			cardBgColor: "bg-blue-500 rounded-3xl text-white",
			bottomBgColor:
				"flex justify-between p-4 bg-blue-400 rounded-bl-3xl rounded-br-3xl",
			imageCreditCard: Chip_card1,
			grayCircleColor: false,
		},
	},
	{
		balance: 5678.9,
		cardHolder: "John Doe",
		expiryDate: "08/24",
		semiCardNumber: "123 **** **** 1234",
		creditCardColor: {
			cardBgColor: "bg-blue-700 rounded-3xl text-white",
			bottomBgColor:
				"flex justify-between p-4 bg-blue-600 rounded-bl-3xl rounded-br-3xl",
			imageCreditCard: Chip_card2,
			grayCircleColor: false,
		},
	},
	{
		balance: 5678.9,
		cardHolder: "John Doe",
		expiryDate: "08/24",
		semiCardNumber: "123 **** **** 1234",
		creditCardColor: {
			cardBgColor: "bg-white rounded-3xl text-black",
			bottomBgColor: "",
			imageCreditCard: Chip_card3,
			grayCircleColor: true,
		},
	},
];
