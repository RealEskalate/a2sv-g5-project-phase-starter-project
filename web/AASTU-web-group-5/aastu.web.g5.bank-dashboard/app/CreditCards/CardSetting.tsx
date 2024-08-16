import Image from "next/image";
import React from "react";
import blockImg from "@/public/assets/image/1-block-credit-card.png";
import padLockImg from "@/public/assets/image/2-padlock.png";
import googleGlassImg from "@/public/assets/image/3-google-glass-logo.png";
import appleImg from "@/public/assets/image/apple-2-1.png";

const CardSetting = () => {
	return (
		<>
			<div className="pr-16">
				<div className="p-2 font-semibold text-blue-900 w-1/2">
					Card Setting
				</div>
				<div className="bg-white  rounded-2xl py-5 pl-4">
					<div className="flex gap-4  p-2 pb-3">
						<div className="flex items-center justify-center bg-[#FFF5D9] p-2 w-14 h-14 rounded-2xl">
							<Image src={blockImg} alt="block image" width={22} height={22} />
						</div>
						<div>
							<div className="font-semibold ">Block Card</div>
							<div className="w-60 text-blue-900 opacity-70">
								Instantly block your card
							</div>
						</div>
					</div>
					<div className="flex gap-4  p-2 pb-3">
						<div className="flex items-center justify-center bg-[#E7EDFF] p-2 w-14 h-14 rounded-2xl">
							<Image
								src={padLockImg}
								alt="block image"
								width={22}
								height={22}
							/>
						</div>
						<div>
							<div className="font-semibold ">Change Pin Code</div>
							<div className="w-60 text-blue-900 opacity-70">
								Choose another pin code
							</div>
						</div>
					</div>
					<div className="flex gap-4  p-2 pb-3">
						<div className="flex items-center justify-center bg-[#FFE0EB] p-2 w-14 h-14 rounded-2xl">
							<Image
								src={googleGlassImg}
								alt="block image"
								width={22}
								height={22}
							/>
						</div>
						<div>
							<div className="font-semibold ">Add to Google Pay</div>
							<div className="w-60 text-blue-900 opacity-70">
								Withdraw without any card
							</div>
						</div>
					</div>
					<div className="flex gap-4  p-2 pb-3">
						<div className="flex items-center justify-center bg-[#DCFAF8] p-2 w-14 h-14 rounded-2xl">
							<Image src={appleImg} alt="block image" width={22} height={22} />
						</div>
						<div>
							<div className="font-semibold ">Add to Apple Pay</div>
							<div className="w-60 text-blue-900 opacity-70">
								Withdraw without any card
							</div>
						</div>
					</div>
					<div className="flex gap-4  p-2 pb-3">
						<div className="flex items-center justify-center bg-[#DCFAF8] p-2 w-12 h-12 rounded-2xl">
							<Image src={appleImg} alt="block image" width={22} height={22} />
						</div>
						<div>
							<div className="font-semibold ">Add to Apple Store</div>
							<div className="w-60 text-blue-900 opacity-70">
								Withdraw without any card
							</div>
						</div>
					</div>
				</div>
			</div>
		</>
	);
};

export default CardSetting;
