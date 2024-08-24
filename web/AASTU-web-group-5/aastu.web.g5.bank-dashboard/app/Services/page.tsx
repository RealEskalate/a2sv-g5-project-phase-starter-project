"use client";
import React, { useEffect, useState } from "react";
import Image from "next/image";
import axios from "axios";
import { useSession } from "next-auth/react";
import { useSelector } from "react-redux";
import { RootState } from "@/app/redux/store";
import Shimmer1 from "../Accounts/shimmer";

interface ExtendedUser {
	name?: string;
	email?: string;
	image?: string;
	accessToken?: string;
}

const Page = () => {
	const [services, setServices] = useState([]);
	const [loading, setLoading] = useState(true);
	const { data: session } = useSession();
	const user = session?.user as ExtendedUser;
	const accessToken = user.accessToken;
	const token: string = `Bearer ${accessToken}`;
	const darkMode = useSelector((state: RootState) => state.theme.darkMode);

	useEffect(() => {
		const fetchServices = async () => {
			try {
				const response = await axios.get(
					"https://bank-dashboard-rsf1.onrender.com/bank-services?page=0&size=7",
					{
						headers: {
							Authorization: token,
						},
					}
				);
				setServices(response.data.data.content);
			} catch (error) {
				console.error("Error fetching services:", error);
			} finally {
				setLoading(false);
			}
		};

		fetchServices();
	}, [token]);

	const getImageProps = (serviceName: string) => {
		switch (serviceName) {
			case "Business Loans":
				return {
					src: "/images/loan 1.png",
					alt: `${serviceName} Icon`,
					bgClass: "bg-pink-100",
				};
			case "Checking Accounts":
				return {
					src: "/images/orange-image.png",
					alt: `${serviceName} Icon`,
					bgClass: "bg-orange-100",
				};
			case "Savings Accounts":
				return {
					src: "/images/pink-image.png",
					alt: `${serviceName} Icon`,
					bgClass: "bg-pink-100",
				};
			case "Debit and credit cards":
				return {
					src: "/images/blue-image.png",
					alt: `${serviceName} Icon`,
					bgClass: "bg-blue-100",
				};
			case "Life Insurance":
				return {
					src: "/images/green-image.png",
					alt: `${serviceName} Icon`,
					bgClass: "bg-green-100",
				};
			default:
				return {
					src: "/images/default-image.png",
					alt: `${serviceName} Icon`,
					bgClass: "bg-gray-100",
				};
		}
	};

	return (
		<div
			className={`bg-[#F5F7FA] w-[95%] dark:bg-gray-900 space-y-8 pt-3 overflow-hidden mx-auto ${
				darkMode ? "text-white" : "text-black"
			}`}
		>
			{/* Life Insurance Section */}
			<div className="overflow-x-auto">
				<div className="flex flex-nowrap gap-6">
					{loading ? (
						<>
							<div className="flex-shrink-0 w-full md:w-1/3">
								<Shimmer1 />
							</div>
							<div className="flex-shrink-0 w-full md:w-1/3">
								<Shimmer1 />
							</div>
							<div className="flex-shrink-0 w-full md:w-1/3">
								<Shimmer1 />
							</div>
						</>
					) : (
						<>
							<div className="flex-shrink-0 w-full md:w-1/3">
								<div
									className={`h-32 lg:h-[90px] rounded-xl flex-grow flex justify-center items-center gap-4 ${
										darkMode ? "bg-gray-800" : "bg-white"
									} w-full`}
								>
									<div className="flex items-center justify-center bg-blue-100 rounded-full w-14 h-14">
										<Image
											src="/images/life-insurance filled 1.png"
											alt="heart Icon"
											objectFit="cover"
											width={25}
											height={25}
										/>
									</div>
									<div className="flex flex-col">
										<h3 className="font-inter text-lg font-medium leading-[16.94px] text-left">
											Life Insurance
										</h3>
										<p className="font-inter opacity-60 text-base">
											Unlimited Protection
										</p>
									</div>
								</div>
							</div>

							<div className="flex-shrink-0 w-full md:w-1/3">
								<div
									className={`h-32 lg:h-[90px] rounded-xl flex-grow flex justify-center items-center gap-4 ${
										darkMode ? "bg-gray-800" : "bg-white"
									} w-full`}
								>
									<div className="flex items-center justify-center bg-orange-100 rounded-full w-14 h-14">
										<Image
											src="/images/shopping.png"
											alt="Shopping Icon"
											objectFit="cover"
											width={20}
											height={20}
										/>
									</div>
									<div className="flex flex-col">
										<h3 className="font-inter text-lg font-medium leading-[16.94px] text-left">
											Shopping
										</h3>
										<p className="font-inter opacity-60 text-base">
											Buy. Think. Grow.
										</p>
									</div>
								</div>
							</div>

							<div className="flex-shrink-0 w-full md:w-1/3">
								<div
									className={`h-32 lg:h-[90px] rounded-xl flex-grow flex justify-center items-center gap-4 ${
										darkMode ? "bg-gray-800" : "bg-white"
									} w-full`}
								>
									<div className="flex items-center justify-center bg-green-100 rounded-full w-14 h-14">
										<Image
											src="/images/shield 1.png"
											alt="Safety Icon"
											objectFit="cover"
											width={50}
											height={50}
										/>
									</div>
									<div className="flex flex-col">
										<h3 className="font-inter text-lg font-medium leading-[16.94px] text-left">
											Safety
										</h3>
										<p className="font-inter opacity-60 text-base">
											We are all your allies
										</p>
									</div>
								</div>
							</div>
						</>
					)}
				</div>
			</div>

      {/* Bank Service List Section */}
      <div className="space-y-3">
        <div className="p-2 sm:p-1">
          <p className="ml-8 text-xl sm:text-2xl font-inter font-semibold">
            Bank Service List
          </p>
        </div>
        {loading ? (
          <>
            <Shimmer1 />
            <Shimmer1 />
            <Shimmer1 />
            <Shimmer1 />
            <Shimmer1 />
            <Shimmer1 />
          </>
        ) : (
          services.map((service) => {
            const { src, alt, bgClass } = getImageProps(service.name);
            return (
              <div key={service.id} className="w-full flex items-center justify-center">
                  <div className={`rounded-xl flex-grow items-center justify-center sm:shadow-md ${darkMode ? "bg-gray-800" : "bg-white"} w-full md:w-11/12`}>
                    <div className="flex items-center justify-center p-1 m-0">
                      <div className="flex items-center gap-3 w-full sm:flex sm:items-center sm:gap-3 sm:w-full">
                        <div className={`flex items-center justify-center rounded-xl w-12 h-12 ${bgClass}`}>
                          <Image
                            src={src}
                            alt={alt}
                            width={25}
                            height={25}
                          />
                        </div>
                        <div className="flex-1 sm:flex-none">
                          <h3 className={`font-inter text-base font-medium leading-[16.94px] text-left whitespace-nowrap ${darkMode ? "text-white" : "text-black"}`}>
                            {service.name}
                          </h3>
                          <p className={`sm:font-inter sm:text-xs sm:font-normal sm:leading-[14.52px] sm:text-left opacity-60 whitespace-nowrap ${darkMode ? "text-gray-400" : "text-gray-600"}`}>
                            {service.details.slice(0, 25)}
                          </p>
                        </div>
                      </div>

										<div className="hidden sm:flex sm:flex-row items-center justify-between w-full sm:w-auto mt-3 sm:mt-0 sm:ml-auto sm:pr-6">
											<div className="hidden sm:flex flex-row gap-4 sm:gap-12 items-start pr-6">
												<div className="flex-col items-center gap-3 w-48">
													<p
														className={`font-semibold ${
															darkMode ? "text-white" : "text-black"
														}`}
													>
														Number of Users
													</p>
													<p
														className={`font-inter opacity-60 ${
															darkMode ? "text-gray-400" : "text-gray-600"
														}`}
													>
														{service.numberOfUsers}
													</p>
												</div>
												<div className="flex-col items-center w-20">
													<p
														className={`font-semibold ${
															darkMode ? "text-white" : "text-black"
														}`}
													>
														Status
													</p>
													<p
														className={`font-inter opacity-60 ${
															darkMode ? "text-gray-400" : "text-gray-600"
														}`}
													>
														{service.status}
													</p>
												</div>
												<div className="flex-col items-center w-20">
													<p
														className={`font-semibold ${
															darkMode ? "text-white" : "text-black"
														}`}
													>
														Type
													</p>
													<p
														className={`font-inter opacity-60 ${
															darkMode ? "text-gray-400" : "text-gray-600"
														}`}
													>
														{service.type}
													</p>
												</div>
											</div>
										</div>

										<button className="text-purple-600 mr-1 hover:bg-purple-50 rounded-full sm:flex">
											View Details
										</button>
									</div>
								</div>
							</div>
						);
					})
				)}
			</div>
		</div>
	);
};

export default Page;
