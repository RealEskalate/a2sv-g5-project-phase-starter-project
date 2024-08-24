"use client";

import React, { useState, useEffect } from "react";
import { FaBars, FaTimes } from "react-icons/fa";
import Image from "next/image";
import { signOut, useSession } from "next-auth/react";
import { useRouter, usePathname } from "next/navigation";
import { useSelector, useDispatch } from "react-redux";
import { RootState } from "@/app/redux/store";
import { toggleDarkMode } from "@/app/redux/slice/themeSlice";
import seetings from "/public/assets/icons/Group417.png";
import notification from "/public/assets/icons/Group418.png";
import magnifying from "/public/assets/icons/magnifying-glass.png";
import lightModeImg from "@/public/assets/image/lightmode.png";
import darkModeImg from "@/public/assets/image/night-mode.png";
import person from "/public/assets/icons/MaskGroup.png";

interface ExtendedUser {
	name?: string;
	email?: string;
	image?: string;
	accessToken?: string;
}

const NavBar = ({ toggleSidebar, isSidebarVisible }) => {
	const { data: session, status } = useSession();
	const router = useRouter();
	const [searchText, setSearchText] = useState("");
	const pathname = usePathname();

	const darkmode = useSelector((state: RootState) => state.theme.darkMode);
	const reduxUser = useSelector((state: RootState) => state.user);
	const [profilePicture, setProfilePicture] = useState(
		reduxUser?.profilePicture && reduxUser?.profilePicture.startsWith("https")
			? reduxUser?.profilePicture
			: "/images/christina.png"
	);

	const user = session?.user as ExtendedUser;
	const dispatch = useDispatch();

	useEffect(() => {
		if (status === "authenticated" && user?.accessToken && !reduxUser?.name) {
			dispatch({
				type: "USER_FETCH_REQUESTED",
				payload: {
					username: user?.name || "",
					token: user.accessToken,
				},
			});
		}
	}, [status, dispatch, user, reduxUser?.name]);

	useEffect(() => {
		if (
			reduxUser?.profilePicture &&
			reduxUser?.profilePicture.startsWith("https")
		) {
			setProfilePicture(reduxUser.profilePicture);
		} else {
			setProfilePicture("/images/christina.png");
		}
	}, [reduxUser?.profilePicture]);

	const handleToggleDarkMode = () => {
		dispatch(toggleDarkMode());
	};

	const handleSearchSubmit = (e: React.FormEvent) => {
		e.preventDefault();
		// Implement search functionality here
		console.log("Search submitted:", searchText);
	};

<<<<<<< HEAD
  return (
    <div className={`shadow-md bg-white dark:bg-gray-900 `}>
      {/* Mobile view */}
      <div className="flex justify-between items-center p-6 sm:hidden">
        <button onClick={toggleSidebar} aria-label="Toggle sidebar">
          {isSidebarVisible ? <FaTimes size={24} /> : <FaBars size={24} />}
        </button>
        <div
          className={`font-semibold text-lg ${
            darkmode ? "text-gray-200" : "text-primary-2"
          }`}
        >
          Overview
        </div>
        <Image src="/images/christina.png" width={50} height={50}  alt="User Icon" className="h-12 w-12" />
        <button
          onClick={() => signOut()}
          className="bg-slate-500 text-white rounded-full px-4 py-2 ml-4"
        >
          Logout
        </button>
      </div>
=======
	if (status !== "authenticated") {
		return null; // Do not render NavBar if not authenticated
	}
>>>>>>> aastu.web.g5.main

	return (
		<div className={`shadow-md bg-white dark:bg-gray-900`}>
			{/* Mobile view */}
			<div className="flex justify-between items-center p-6 sm:hidden">
				<button onClick={toggleSidebar} aria-label="Toggle sidebar">
					{isSidebarVisible ? <FaTimes size={24} /> : <FaBars size={24} />}
				</button>
				<div
					className={`font-semibold text-lg ${
						darkmode ? "text-gray-200" : "text-primary-2"
					}`}
				>
					Overview
				</div>
				<Image
					src={profilePicture}
					width={50}
					height={50}
					alt="User Icon"
					className="h-12 w-12"
				/>
				<button
					onClick={() => signOut()}
					className="bg-slate-500 text-white rounded-full px-4 py-2 ml-4"
				>
					Logout
				</button>
			</div>

<<<<<<< HEAD
      {/* Desktop view */}
      <div className="hidden sm:flex justify-between items-center h-24 px-12">
        <div
          className={`font-semibold text-xl ${
            darkmode ? "text-gray-200" : "text-primary-2"
          }`}
        >
          {name !== `/auth/signup` && name !== `/auth/signin`
            ? name.slice(1, name.length)
            : "Overview"}
        </div>
        <div className="flex gap-5 items-center">
          <div className="relative w-full max-w-xs">
            <input
              type="text"
              placeholder="Search..."
              value={searchText}
              onChange={(e) => setSearchText(e.target.value)}
              className={`pl-12 pr-12 py-2 w-full ${
                darkmode ? "bg-gray-700 text-white" : "bg-gray-100 text-black"
              } rounded-full border border-gray-300`}
            />
            <div className="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-500">
              <Image src={magnifying} alt="Search Icon" width={16} height={16} />
            </div>
            {searchText && (
              <button
                onClick={() => setSearchText("")}
                className="absolute right-4 top-1/2 transform -translate-y-1/2 text-gray-500"
              >
                <FaTimes />
              </button>
            )}
          </div>
          <button onClick={handleToggleDarkMode}>
            <Image
              src={darkmode ? lightModeImg : darkModeImg}
              alt="Toggle Dark Mode"
              width={45}
              height={45}
            />
          </button>
          <Image src={notification} alt="Notification Icon" className="h-10 w-10" />
          <Image src={seetings} alt="Settings Icon" className="h-10 w-10" />
          <Image width={50} height={50}   src="/images/christina.png" alt="User Icon" className="h-12 rounded-full w-12" />
          <button
            onClick={() => signOut()}
            className="bg-slate-500 text-white rounded-full p-4 py-2 ml-4"
          >
            Logout
          </button>
        </div>
      </div>
    </div>
  );
=======
			{/* Mobile search bar */}
			<div className="sm:hidden p-6">
				<form onSubmit={handleSearchSubmit} className="relative w-full">
					<input
						type="text"
						placeholder="Search..."
						value={searchText}
						onChange={(e) => setSearchText(e.target.value)}
						className="pl-12 pr-12 py-2 w-full bg-gray-100 rounded-full border border-gray-300"
					/>
					<div className="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-500">
						<Image src={magnifying} alt="Search Icon" width={16} height={16} />
					</div>
					{searchText && (
						<button
							type="button"
							onClick={() => setSearchText("")}
							className="absolute right-4 top-1/2 transform -translate-y-1/2 text-gray-500"
						>
							<FaTimes />
						</button>
					)}
				</form>
			</div>

			{/* Desktop view */}
			<div className="hidden sm:flex justify-between items-center h-24 px-12">
				<div
					className={`font-semibold text-xl ${
						darkmode ? "text-gray-200" : "text-primary-2"
					}`}
				>
					{pathname !== `/auth/signup` && pathname !== `/auth/signin`
						? pathname.slice(1)
						: "Overview"}
				</div>
				<div className="flex gap-5 items-center">
					<form
						onSubmit={handleSearchSubmit}
						className="relative w-full max-w-xs"
					>
						<input
							type="text"
							placeholder="Search..."
							value={searchText}
							onChange={(e) => setSearchText(e.target.value)}
							className={`pl-12 pr-12 py-2 w-full ${
								darkmode ? "bg-gray-700 text-white" : "bg-gray-100 text-black"
							} rounded-full border border-gray-300`}
						/>
						<div className="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-500">
							<Image
								src={magnifying}
								alt="Search Icon"
								width={16}
								height={16}
							/>
						</div>
						{searchText && (
							<button
								type="button"
								onClick={() => setSearchText("")}
								className="absolute right-4 top-1/2 transform -translate-y-1/2 text-gray-500"
							>
								<FaTimes />
							</button>
						)}
					</form>
					<button onClick={handleToggleDarkMode}>
						<Image
							src={darkmode ? lightModeImg : darkModeImg}
							alt="Toggle Dark Mode"
							width={45}
							height={45}
						/>
					</button>
					<Image
						src={notification}
						alt="Notification Icon"
						className="h-10 w-10"
					/>
					<Image src={seetings} alt="Settings Icon" className="h-10 w-10" />
					<Image
						width={50}
						height={50}
						src={profilePicture}
						alt="User Icon"
						className="h-12 rounded-full w-12"
					/>
					<button
						onClick={() => signOut()}
						className="bg-slate-500 text-white rounded-full p-4 py-2 ml-4"
					>
						Logout
					</button>
				</div>
			</div>
		</div>
	);
>>>>>>> aastu.web.g5.main
};

export default NavBar;
