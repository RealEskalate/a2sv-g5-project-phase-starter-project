"use client";

import axios from "axios";
import React, { useState, useEffect } from "react";
import { useSession } from "next-auth/react";
import { useForm, SubmitHandler } from "react-hook-form";
import Toggle from "./toogle";
import { useSelector, useDispatch } from "react-redux";
import { setUser } from "../redux/slice/userSlice";
import { RootState } from "../redux/store";
import User from "../../type/user";
interface ExtendedUser {
	name?: string;
	email?: string;
	image?: string;
	accessToken?: string;
}

interface FormData {
	currency: string;
	timeZone: string;
	sentOrReceiveDigitalCurrency: boolean;
	receiveMerchantOrder: boolean;
	accountRecommendations: boolean;
	twoFactorAuthentication: boolean;
}

function Preference() {
	const { data: session } = useSession();
	const [successMessage, setSuccessMessage] = useState("");
	const [apiError, setApiError] = useState("");

	const user = useSelector((state: RootState) => state.user as User);
	const dispatch = useDispatch();

	const {
		register,
		handleSubmit,
		setValue,
		formState: { errors },
	} = useForm<FormData>();

	const [digitalCurrency, setDigitalCurrency] = useState(false);
	const [merchantOrder, setMerchantOrder] = useState(false);
	const [accountRecommendations, setAccountRecommendations] = useState(false);

	const users = session?.user as ExtendedUser;
	const key = users?.accessToken || "";

	useEffect(() => {
		if (user) {
			setValue("currency", user.preference.currency || "");
			setValue("timeZone", user.preference.timeZone || "");
			setDigitalCurrency(user.preference.sentOrReceiveDigitalCurrency || false);
			setMerchantOrder(user.preference.receiveMerchantOrder || false);
			setAccountRecommendations(
				user.preference.accountRecommendations || false
			);
		}
	}, [user, setValue]);

	const handleDigitalCurrencyChange = () =>
		setDigitalCurrency(!digitalCurrency);
	const handleMerchantOrderChange = () => setMerchantOrder(!merchantOrder);
	const handleAccountRecommendationsChange = () =>
		setAccountRecommendations(!accountRecommendations);

	const onSubmit: SubmitHandler<FormData> = async (data) => {
		setApiError("");
		const updatedData = {
			...data,
			sentOrReceiveDigitalCurrency: digitalCurrency,
			receiveMerchantOrder: merchantOrder,
			accountRecommendations: accountRecommendations,
			twoFactorAuthentication: true,
		};

		try {
			const response = await axios.put(
				"https://bank-dashboard-rsf1.onrender.com/user/update-preference",
				updatedData,
				{
					headers: {
						"Content-Type": "application/json",
						Authorization: `Bearer ${key}`,
					},
				}
			);

			if (response.status === 200) {
				setSuccessMessage("Preferences updated successfully!");
				dispatch(setUser(updatedData));
			} else {
				throw new Error(`Failed to update preferences: ${response.statusText}`);
			}
		} catch (error: any) {
			setApiError(
				error.response?.data?.message || "Failed to update preferences."
			);
		}
	};

	return (
		<form onSubmit={handleSubmit(onSubmit)}>
			<div className="flex flex-wrap flex-col md:flex-row md:gap-10 lg:gap-12 mt-10 md:mt-12 mx-4 dark:text-gray-400">
				<div>
					<div>Currency</div>
					<input
						type="text"
						className={`border-slate-200 border-[1px] w-full h-10 mt-3 rounded-3xl md:w-[20rem] lg:w-[30rem] dark:border-[#fff] dark:focus:outline-none dark:border-opacity-50 dark:opacity-80 dark:bg-gray-500 dark:text-[#fff] ${
							errors.currency ? "border-red-500" : ""
						}`}
						style={{ paddingLeft: "1.25rem" }}
						placeholder="USD"
						{...register("currency", { required: "Currency is required" })}
					/>
					{errors.currency && (
						<div className="text-red-500 text-sm mt-2">
							{errors.currency.message}
						</div>
					)}
				</div>
				<div>
					<div>Time Zone</div>
					<input
						type="text"
						className={`border-slate-200 dark:border-[#fff] dark:focus:outline-none dark:border-opacity-50 dark:opacity-80 dark:bg-gray-500 dark:text-[#fff] border-[1px] w-full h-10 mt-3 rounded-2xl md:w-[20rem] lg:w-[30rem] ${
							errors.timeZone ? "border-red-500" : ""
						}`}
						placeholder="(GMT-12:00) International Date Line West"
						style={{ paddingLeft: "1.25rem" }}
						{...register("timeZone", { required: "Time Zone is required" })}
					/>
					{errors.timeZone && (
						<div className="text-red-500 text-sm mt-2">
							{errors.timeZone.message}
						</div>
					)}
				</div>
			</div>
			<div className="mt-6 md:mt-8 text-slate-700 text-sm md:text-base lg:text-[17px] dark:text-gray-400">
				Notification
				<div className="flex flex-col gap-4 mt-5 md:mt-6">
					<div className="flex gap-5 md:gap-6">
						<Toggle />
						<div>I send or receive digital currency</div>
					</div>
					<div className="flex gap-5 md:gap-6">
						<Toggle />
						<div>I receive merchant order</div>
					</div>
					<div className="flex gap-5 md:gap-6">
						<Toggle />
						<div>There are recommendations for my account</div>
					</div>
				</div>
			</div>
			{apiError && <div className="text-red-500 mt-4 text-sm">{apiError}</div>}
			{successMessage && (
				<div className="text-green-500 mt-2">{successMessage}</div>
			)}
			<div className="flex justify-end mt-16 md:mt-18">
				<button
					type="submit"
					className="border-none bg-blue-700 text-white w-full h-12 rounded-full md:w-[12rem] text-[13px] md:text-base"
				>
					Save
				</button>
			</div>
		</form>
	);
}

export default Preference;
