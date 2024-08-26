import Image from "next/image";
import React from "react";
import blockImg from "@/public/assets/image/1-block-credit-card.png";
import padLockImg from "@/public/assets/image/2-padlock.png";
import googleGlassImg from "@/public/assets/image/3-google-glass-logo.png";
import appleImg from "@/public/assets/image/apple-2-1.png";

const CardSetting = () => {
	return (
		<>
			<div>
				<div className="p-2 font-semibold text-blue-900 dark:text-blue-600 w-1/2">
					Card Setting
				</div>
				<div className="flex flex-col gap-10 bg-white dark:bg-gray-800 dark:text-[#fff] rounded-2xl p-7 ">
					<div className="flex gap-4   ">
						<div>
							<div className="flex items-center justify-center bg-[#FFF5D9] p-2 w-14 h-14 rounded-2xl">
								<Image
									src={blockImg}
									alt="block image"
									width={22}
									height={22}
								/>
							</div>
						</div>
						<div>
							<div className="font-semibold ">Block Card</div>
							<div className="w-60 text-blue-900 dark:text-gray-400 opacity-70">
								Instantly block your card
							</div>
						</div>
					</div>
					<div className="flex gap-4  ">
						<div>
							<div className="flex items-center justify-center bg-[#E7EDFF] p-2 w-14 h-14 rounded-2xl">
								<Image
									src={padLockImg}
									alt="block image"
									width={22}
									height={22}
								/>
							</div>
						</div>
						<div>
							<div className="font-semibold ">Change Pin Code</div>
							<div className="w-60 text-blue-900 dark:text-gray-400 opacity-70">
								Choose another pin code
							</div>
						</div>
					</div>
					<div className="flex gap-4  ">
						<div>
							<div className="flex items-center justify-center bg-[#FFE0EB] p-2 w-14 h-14 rounded-2xl">
								<Image
									src={googleGlassImg}
									alt="block image"
									width={22}
									height={22}
								/>
							</div>
						</div>
						<div>
							<div className="font-semibold ">Add to Google Pay</div>
							<div className="w-60 text-blue-900 dark:text-gray-400 opacity-70">
								Withdraw without any card
							</div>
						</div>
					</div>
					<div className="flex gap-4  ">
						<div>
							<div className="flex items-center justify-center bg-[#DCFAF8] p-2 w-14 h-14 rounded-2xl">
								<Image
									src={appleImg}
									alt="block image"
									width={22}
									height={22}
								/>
							</div>
						</div>
						<div>
							<div className="font-semibold ">Add to Apple Pay</div>
							<div className="w-60 text-blue-900 dark:text-gray-400 opacity-70">
								Withdraw without any card
							</div>
						</div>
					</div>
					<div className="flex gap-4 ">
						<div>
							<div className="flex items-center justify-center bg-[#DCFAF8] p-2 w-14 h-14 rounded-2xl">
								<Image
									src={appleImg}
									alt="block image"
									width={22}
									height={22}
								/>
							</div>
						</div>
						<div>
							<div className="font-semibold ">Add to Apple Store</div>
							<div className="w-60 text-blue-900 dark:text-gray-400 opacity-70">
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
