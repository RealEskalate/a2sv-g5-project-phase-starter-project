// mockData.js
import Image from "next/image";
import creditCard1 from "@/public/assets/image/credit-card 1.png";
import creditCard2 from "@/public/assets/image/credit-card 2.png";
import creditCard3 from "@/public/assets/image/credit-card 3.png";

export const cardListMockData = [
	{
		bank: "DBL Bank",
		creditCardImg: creditCard1,
		bgImgColor:
			"flex justify-center items-center bg-[#E7EDFF] w-14 h-14 rounded-xl",
	},
	{
		bank: "BRC Bank",
		creditCardImg: creditCard2,
		bgImgColor:
			"flex justify-center items-center bg-[#FFE0EB] w-14 h-14 rounded-xl",
	},
	{
		bank: "ABM Bank",
		creditCardImg: creditCard3,
		bgImgColor:
			"flex justify-center items-center bg-[#FFF5D9] w-14 h-14 rounded-xl",
	},
];
