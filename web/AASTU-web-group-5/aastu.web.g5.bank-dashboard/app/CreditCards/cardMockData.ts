import Chip_card1 from "@/public/assets/image/Chip_Card1.png";
import Chip_card2 from "@/public/assets/image/Chip_Card2.png";
import Chip_card3 from "@/public/assets/image/Chip_Card3.png";

const creditCardColor = [
	{
		cardBgColor: "bg-blue-500 rounded-3xl text-white",
		bottomBgColor:
			"flex justify-between p-4 bg-blue-400 rounded-bl-3xl rounded-br-3xl",
		imageCreditCard: Chip_card1,
		grayCircleColor: false,
	},
	{
		cardBgColor: "bg-blue-700 rounded-3xl text-white",
		bottomBgColor:
			"flex justify-between p-4 bg-blue-600 rounded-bl-3xl rounded-br-3xl",
		imageCreditCard: Chip_card2,
		grayCircleColor: false,
	},
	{
		cardBgColor:
			"bg-[#fff] rounded-3xl text-[#343C6A]  border-2 border-solid border-gray-200 ",
		bottomBgColor:
			"flex justify-between p-4 border-t-2 border-solid border-gray-200 rounded-bl-3xl rounded-br-3xl",
		imageCreditCard: Chip_card3,
		grayCircleColor: true,
	},
];
export default creditCardColor;
