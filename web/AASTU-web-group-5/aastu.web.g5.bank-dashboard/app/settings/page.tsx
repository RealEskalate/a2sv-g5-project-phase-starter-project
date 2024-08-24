"use client";
import React, { useState } from "react";
import Profile from "./profile";
import Preference from "./preference";
import Security from "./security";
import { useSelector } from "react-redux";
import User from "../../type/user";
import { RootState } from "@/app/redux/store"; // Adjust the import path as necessary

const activeColor = "text-blue-700";
const disabledColor = "text-slate-400";

function Settings() {
	const user = useSelector((state: { user: User }) => state.user);
	const darkMode = useSelector((state: RootState) => state.theme.darkMode);
	const [enabled, setEnabled] = useState("Edit Profile");

	const renderContent = () => {
		switch (enabled) {
			case "Edit Profile":
				return <Profile />;
			case "Preference":
				return <Preference />;
			case "Security":
				return <Security />;
			default:
				return null;
		}
	};

	return (
		<div
			className={`rounded-2xl pt-4 w-[95%] text-sm md:text-base justify-center mt-6 px-6 sm:px-10 pb-5 mx-auto ${
				darkMode ? "bg-gray-900 text-white" : "bg-white text-neutral-800"
			}`}
		>
			<div className="relative mt-8 md:mt-10">
				<div className="flex gap-10 md:flex-row md:gap-16 lg:gap-20">
					<div
						className={`cursor-pointer ${
							enabled === "Edit Profile"
								? `${activeColor} custom-underline`
								: disabledColor
						} ${darkMode ? "text-white" : ""}`}
						onClick={() => setEnabled("Edit Profile")}
					>
						Edit Profile
					</div>
					<div
						className={`cursor-pointer ${
							enabled === "Preference"
								? `${activeColor} custom-underline`
								: disabledColor
						} ${darkMode ? "text-white" : ""}`}
						onClick={() => setEnabled("Preference")}
					>
						Preference
					</div>
					<div
						className={`cursor-pointer ${
							enabled === "Security"
								? `${activeColor} custom-underline`
								: disabledColor
						} ${darkMode ? "text-white" : ""}`}
						onClick={() => setEnabled("Security")}
					>
						Security
					</div>
				</div>
				<div className="mt-10">{renderContent()}</div>
			</div>
		</div>
	);
}

export default Settings;
