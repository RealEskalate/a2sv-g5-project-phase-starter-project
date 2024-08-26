"use client";
import React, { useState } from "react";
import Toggle from "./toogle";
import { signOut, useSession } from "next-auth/react";
import { useForm } from "react-hook-form";
import axios from "axios";
import Error from "../error";
import { useSelector } from "react-redux";
import { RootState } from "../redux/store";
import { LoaderProvider } from "react-loader-ts";

import { useLoader, Loader } from "react-loader-ts";

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

export default function Security() {
	const {
		register,
		formState: { errors },
	} = useForm();
	const [apiError, setApiError] = useState("");
	const [currentPassword, setCurrentPassword] = useState("");
	const [newPassword, setNewPassword] = useState("");
	const { data: session } = useSession();
	const user = session.user as ExtendedUser;
	const [loading, setLoading] = useState(false);
	const [successMessage, setSuccessMessage] = useState("");
	const darkMode = useSelector((state: RootState) => state.theme.darkMode);

	const key: string = user.accessToken || "";
	const handleSubmit = async (e) => {
		setLoading(true);
		e.preventDefault();
		setSuccessMessage("");
		setApiError("");

		const data = { password: currentPassword, newPassword: newPassword };

		try {
			const response = await axios.post(
				`https://bank-dashboard-irbd.onrender.com/auth/change_password`,
				data,
				{
					headers: {
						"Content-Type": "application/json",
						Authorization: `Bearer ${key}`,
					},
				}
			);

			if (response.status === 200) {
				const result = response.data;
				if (result.success) {
					setSuccessMessage("Password changed successfully!");
					// signOut()
				} else {
					setApiError(result.message || "Failed to change password.");
				}
			} else {
				setApiError(`Failed to change password: ${response.statusText}`);
			}
		} catch (error) {
			setApiError(
				error.response?.data?.message || "Failed to change password."
			);
			console.error("Error changing password:", error);
		}
		setLoading(false)
	};

	return (
		<div
			className={`text-[16px] ${
				darkMode ? "bg-gray-900 text-white" : "bg-white text-neutral-800"
			}`}
		>
			<Error session={session} />

			<div
				className={`text-sm md:text-base lg:text-[17px] ${
					darkMode ? "text-white" : "text-slate-700"
				}`}
			>
				Two-factor Authentication
			</div>
			<div className="flex gap-5 md:gap-6 mt-4">
				<Toggle />
				<div
					className={`text-sm md:text-base lg:text-[17px] ${
						darkMode ? "text-white" : "text-slate-700"
					}`}
				>
					Enable or disable two-factor authentication
				</div>
			</div>

			<form onSubmit={handleSubmit} className="mt-10">
				<div
					className={`text-sm md:text-base lg:text-[17px] ${
						darkMode ? "text-white" : "text-slate-700"
					}`}
				>
					Change Password
				</div>
				{apiError && <div className="text-red-500 mt-2">{apiError}</div>}
				{successMessage && (
					<div className="text-green-500 mt-2">{successMessage}</div>
				)}

				<div className="mt-4">
					<div className="text-slate-700 text-sm md:text-base lg:text-[17px]">
						Current Password
					</div>
					<input
						type="password"
						{...register("currentPassword", {
							required: "Current password is required",
						})}
						className={`border-slate-200 border-[1px] w-full h-10 mt-3 rounded-3xl md:w-[20rem] lg:w-[30rem] ${
							errors.currentPassword ? "border-red-500" : ""
						}`}
						style={{ paddingLeft: "1.25rem" }}
						placeholder="************"
						onChange={(e) => setCurrentPassword(e.target.value)}
					/>
					{errors.currentPassword && (
						<div className="text-red-500 text-sm mt-2">
							{errors.currentPassword.message.toString()}
						</div>
					)}
				</div>

				<div className="mt-4">
					<div className="text-slate-700 text-sm md:text-base lg:text-[17px]">
						New Password
					</div>
					<input
						type="password"
						{...register("newPassword", {
							required: "New password is required",
						})}
						value={newPassword}
						onChange={(e) => setNewPassword(e.target.value)}
						className="border-slate-200 border-[1px] w-full h-10 mt-3 rounded-3xl md:w-[20rem] lg:w-[30rem]"
						style={{ paddingLeft: "1.25rem" }}
						placeholder="************"
					/>
					{errors.newPassword && (
						<div className="text-red-500 text-sm mt-2">
							{errors.newPassword.message.toString()}
						</div>
					)}
				</div>

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
		</div>
	);
}
