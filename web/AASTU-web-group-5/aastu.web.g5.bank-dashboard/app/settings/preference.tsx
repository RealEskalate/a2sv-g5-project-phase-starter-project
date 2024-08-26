"use client";

import axios from "axios";
import React, { useState, useEffect } from "react";
import { signOut, useSession } from "next-auth/react";
import { useForm, SubmitHandler } from "react-hook-form";
import Toggle from "./toogle";
import { useSelector, useDispatch } from "react-redux";
import { setUser } from "../redux/slice/userSlice";
import { RootState } from "../redux/store";
import User from "../../type/user";
import { Audio } from 'react-loader-spinner';
import { LoaderProvider } from "react-loader-ts";
import { useLoader, Loader } from "react-loader-ts";
import "react-loader-ts/lib/esm/styles/global.css";



 function Spinner() {
  return (
    <LoaderProvider>
      <Loader />
    </LoaderProvider>
  );
}

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
	const [loading, setLoading] = useState(false);
	const user = useSelector((state: RootState) => state.user as User);
	const darkMode = useSelector((state: RootState) => state.theme.darkMode);
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
	console.log(users, "users");

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
	}, [users, setValue, user]);

	if (!users) {
		return <></>;
	}

	const handleDigitalCurrencyChange = () =>
		setDigitalCurrency(!digitalCurrency);
	const handleMerchantOrderChange = () => setMerchantOrder(!merchantOrder);
	const handleAccountRecommendationsChange = () =>
		setAccountRecommendations(!accountRecommendations);

	const onSubmit: SubmitHandler<FormData> = async (data) => {
		setLoading(true);
		setSuccessMessage("");
		setApiError("");
		console.log(users.accessToken, "access token ");
		const updatedData = {
			...data,
			sentOrReceiveDigitalCurrency: digitalCurrency,
			receiveMerchantOrder: merchantOrder,
			accountRecommendations: accountRecommendations,
			twoFactorAuthentication: true,
		};
		console.log(updatedData, "updated data ");
		try {
			const response = await axios.put(
				"https://bank-dashboard-irbd.onrender.com/user/update-preference",
				updatedData,
				{
					headers: {
						"Content-Type": "application/json",
						Authorization: `Bearer ${users.accessToken}`,
					},
				}
			);

			console.log(response.status, 111111);
			if (response.status === 200) {
				setSuccessMessage("Preferences updated successfully!");
				dispatch(setUser(updatedData));
			} else if (response.status == 401) {
				signOut();
			} else {
				throw new Error(`Failed to update preferences: ${response.statusText}`);
			}
		} catch (error: any) {
			if (error.response.status == 401) {
				signOut();
			}
			setApiError(
				error.response?.data?.message || "Failed to update preferences."
			);
		}
		setLoading(false);

	};

	return (
		<form
			onSubmit={handleSubmit(onSubmit)}
			className={`${
				darkMode ? "bg-gray-900 text-white" : "bg-white text-neutral-800"
			}`}
		>
			<div className="flex flex-wrap flex-col md:flex-row md:gap-10 lg:gap-12 mt-10 md:mt-12 mx-4">
				<div>
					<div>Currency</div>
					<input
						type="text"
						className={`border-slate-200 border-[1px] w-full h-10 mt-3 rounded-3xl md:w-[20rem] lg:w-[30rem] ${
							errors.currency ? "border-red-500" : ""
						} ${darkMode ? "bg-gray-800 text-white" : ""}`}
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
						className={`border-slate-200 border-[1px] w-full h-10 mt-3 rounded-2xl md:w-[20rem] lg:w-[30rem] ${
							errors.timeZone ? "border-red-500" : ""
						} ${darkMode ? "bg-gray-800 text-white" : ""}`}
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
			<div
				className={`mt-6 md:mt-8 text-sm md:text-base lg:text-[17px] ${
					darkMode ? "text-white" : "text-slate-700"
				}`}
			>
				Notification
				<div className="flex flex-col gap-4 mt-5 md:mt-6">
					<div
						className="flex gap-5 md:gap-6"
						onClick={handleDigitalCurrencyChange}
					>
						<Toggle />
						<div>I send or receive digital currency</div>
					</div>
					<div
						className="flex gap-5 md:gap-6"
						onClick={handleMerchantOrderChange}
					>
						<Toggle />
						<div>I receive merchant order</div>
					</div>
					<div
						className="flex gap-5 md:gap-6"
						onClick={handleAccountRecommendationsChange}
					>
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
        className={`border-none w-full h-12 rounded-full md:w-[12rem] text-[13px] md:text-base ${
          darkMode ? "bg-blue-500 text-white" : "bg-blue-700 text-white"
        }`}
        disabled={loading} 
      >
        {loading ? (
			    <Spinner/>  ) : (
          'Save'
        )}
      </button>
			</div>
		</form>
	);
}

export default Preference;
