import React from "react";
import Image from "next/image";
import UserImg from "@/public/assets/image/user 3 2.png";
import bagImg from "@/public/assets/image/briefcase 1.png";
import graphImg from "@/public/assets/image/graph 1.png";
import supportImg from "@/public/assets/image/support 1.png";

const LoanCards = () => {
	return (
		<div className="flex flex-wrap gap-7 text-sm">
			<div className="flex items-center gap-4 bg-white w-full sm:w-[48%] md:w-[30%] lg:w-[23%] p-5 py-8 rounded-3xl">
				<div className="flex items-center justify-center rounded-full bg-[#E7EDFF] w-14 h-14">
					<Image src={UserImg} alt="user image" />
				</div>
				<div>
					<p className="text-[#718EBF]">Personal Loans</p>
					<p className="text-[#232323] text-2xl font-medium ">$50,000</p>
				</div>
			</div>
			<div className="flex items-center gap-4 bg-white w-full sm:w-[48%] md:w-[30%] lg:w-[23%] p-5 py-8 rounded-3xl">
				<div className="flex items-center justify-center rounded-full bg-[#FFF5D9] w-14 h-14">
					<Image src={bagImg} alt="user image" />
				</div>
				<div>
					<p className="text-[#718EBF]">Corporate Loans</p>
					<p className="text-[#232323] text-2xl font-medium">$100,000</p>
				</div>
			</div>
			<div className="flex items-center gap-4 bg-white w-full sm:w-[48%] md:w-[30%] lg:w-[23%] p-5 py-8 rounded-3xl">
				<div className="flex items-center justify-center rounded-full bg-[#FFE0EB] w-14 h-14">
					<Image src={graphImg} alt="user image" />
				</div>
				<div>
					<p className="text-[#718EBF]">Business Loans</p>
					<p className="text-[#232323] text-2xl font-medium ">$500,000</p>
				</div>
			</div>
			<div className="flex items-center gap-4 bg-white w-full sm:w-[48%] md:w-[30%] lg:w-[23%] p-5 py-8 rounded-3xl">
				<div className="flex items-center justify-center rounded-full bg-[#DCFAF8] w-14 h-14">
					<Image src={supportImg} alt="user image" />
				</div>
				<div>
					<p className="text-[#718EBF]">Custom Loans</p>
					<p className="text-[#232323] text-base font-semibold">Choose Money</p>
				</div>
			</div>
		</div>
	);
};

export default LoanCards;
