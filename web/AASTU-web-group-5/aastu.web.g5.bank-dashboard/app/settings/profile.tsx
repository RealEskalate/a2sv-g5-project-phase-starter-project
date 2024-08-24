"use client";

import React, { useEffect, useState } from "react";
import Image from "next/image";
import { signOut, useSession } from "next-auth/react";
import axios from "axios";
import {
	getStorage,
	ref,
	uploadBytesResumable,
	getDownloadURL,
} from "firebase/storage";
import { app } from "../../lib/firebase"; // Import the initialized Firebase app
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../redux/store";

interface ExtendedUser {
	name?: string;
	email?: string;
	image?: string;
	accessToken?: string;
}

const EditProfile = () => {
	const [name, setName] = useState("");
	const [email, setEmail] = useState("");
	const [dateOfBirth, setDateOfBirth] = useState("");
	const [permanentAddress, setPermanentAddress] = useState("");
	const [postalCode, setPostalCode] = useState("");
	const [username, setUsername] = useState("");
	const [presentAddress, setPresentAddress] = useState("");
	const [city, setCity] = useState("");
	const [country, setCountry] = useState("");
	const reduxUser = useSelector((state: RootState) => state.user);
	const [profilePicture, setProfilePicture] = useState(
		reduxUser.profilePicture != "string"
			? reduxUser.profilePicture
			: "/images/christina.png"
	);
	const [uploading, setUploading] = useState(false);
	const [message, setMessage] = useState("");
	const [isSuccess, setIsSuccess] = useState(false);
	const dispatch = useDispatch();
	const { data: session, status } = useSession();
	const user = session?.user as ExtendedUser;
	console.log(session, "session");
	console.log(user, "this is user ");
	useEffect(() => {
		console.log(profilePicture, 1111);
		if (status === "authenticated" && user?.accessToken && !reduxUser?.name) {
			console.log("Dispatching USER_FETCH_REQUESTED");
			dispatch({
				type: "USER_FETCH_REQUESTED",
				payload: {
					username: user?.name || "",
					token: user.accessToken,
				},
			});
		}
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [status, dispatch, session, user]);

	useEffect(() => {
		setProfilePicture(
			reduxUser.profilePicture.startsWith("https")
				? reduxUser.profilePicture
				: "/images/christina.png"
		);
		console.log("Profile picture updated:", reduxUser?.profilePicture);
	}, [reduxUser?.profilePicture]);
	console.log(user?.accessToken, "user.accessToken");
	if (!user?.accessToken) {
		return <div>Loading...</div>;
	}

	const handleImageUpload = (e: React.ChangeEvent<HTMLInputElement>) => {
		const file = e.target.files?.[0];
		if (!file) return;
		console.log(file.name, 111111, user.email);

		const storage = getStorage(app);
		const storageRef = ref(
			storage,
			`profilePictures/${user?.email}/${file.name}`
		);
		const uploadTask = uploadBytesResumable(storageRef, file);

		setUploading(true);

		uploadTask.on(
			"state_changed",
			(snapshot) => {},
			(error) => {
				console.error("Upload failed:", error);
				setUploading(false);
			},
			() => {
				getDownloadURL(uploadTask.snapshot.ref).then((downloadURL) => {
					setProfilePicture(downloadURL);
					setUploading(false);
				});
			}
		);
	};

	const handleSubmit = async (e: React.FormEvent) => {
		e.preventDefault();
		const formattedDateOfBirth = new Date(dateOfBirth).toISOString();
		const data = {
			name,
			email,
			dateOfBirth: formattedDateOfBirth,
			permanentAddress,
			postalCode,
			username: user.name,
			presentAddress,
			profilePicture,
			city,
			country,
		};
		console.log(data, "savechanges1111", user.accessToken);
		try {
			const response = await axios.put(
				"https://bank-dashboard-rsf1.onrender.com/user/update",
				data,

				{
					headers: {
						"Content-Type": "application/json",
						Authorization: `Bearer ${user.accessToken}`,
					},
				}
			);

			if (response.status === 200) {
				setMessage("Profile updated successfully!");
				setIsSuccess(true);
				signOut();
			} else {
				const errorData = response.data;
				setMessage(`Error: ${errorData.message}`);
				setIsSuccess(false);
			}
		} catch (error) {
			console.error("An error occurred while updating the profile:", error);
			setMessage("An unexpected error occurred. Please try again later.");
			setIsSuccess(false);
		}
	};

	return (
		<div className="bg-white dark:bg-gray-900 dark:bg-gray-900 p-4 md:p-8">
			<form onSubmit={handleSubmit}>
				<div className="flex flex-col md:flex-row md:space-x-8">
					{/* Profile Image Section */}
					<div className="w-full md:w-[20%] flex justify-center mb-8 md:mb-0">
						<div className="w-56 h-56 md:w-40 md:h-40 rounded-full overflow-hidden flex items-center justify-center">
							<label htmlFor="profilePictureUpload" className="cursor-pointer">
								<Image
									src={profilePicture}
									alt="Profile Picture"
									width={224}
									height={224}
									className="object-cover"
								/>
							</label>
							<input
								type="file"
								id="profilePictureUpload"
								className="hidden"
								onChange={handleImageUpload}
								accept="image/*"
							/>
						</div>
					</div>

					{/* Form Section */}
					<div className="w-full md:w-[80%] flex flex-col md:flex-row gap-6">
						{/* Left Column */}
						<div className="w-full md:w-[50%] space-y-6">
							<div className="bg-white dark:bg-gray-900 dark:bg-gray-900 ">
								<p className="text-black dark:text-white font-sans text-lg mb-2">
									Your Name
								</p>
								<input
									type="text"
									className="border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 dark:border-[#fff] dark:focus:outline-none dark:border-opacity-50 dark:opacity-80 dark:bg-gray-500 dark:text-[#fff]"
									placeholder="Charlene Reed"
									value={name}
									onChange={(e) => setName(e.target.value)}
								/>
							</div>
							<div className="bg-white dark:bg-gray-900">
								<p className="text-black dark:text-white font-sans text-lg mb-2">
									Email
								</p>
								<input
									type="email"
									className="border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 dark:border-[#fff] dark:focus:outline-none dark:border-opacity-50 dark:opacity-80 dark:bg-gray-500 dark:text-[#fff]"
									placeholder="charlenereed@gmail.com"
									value={email}
									onChange={(e) => setEmail(e.target.value)}
								/>
							</div>
							<div className="bg-white dark:bg-gray-900">
								<p className="text-black dark:text-white font-sans text-lg mb-2">
									Date of Birth
								</p>
								<input
									type="date"
									className="border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 dark:border-[#fff] dark:focus:outline-none dark:border-opacity-50 dark:opacity-80 dark:bg-gray-500 dark:text-[#fff]"
									value={dateOfBirth}
									onChange={(e) => setDateOfBirth(e.target.value)}
								/>
							</div>
							<div className="bg-white dark:bg-gray-900">
								<p className="text-black dark:text-white font-sans text-lg mb-2">
									Permanent Address
								</p>
								<input
									type="text"
									className="border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 dark:border-[#fff] dark:focus:outline-none dark:border-opacity-50 dark:opacity-80 dark:bg-gray-500 dark:text-[#fff]"
									placeholder="San Jose, California, USA"
									value={permanentAddress}
									onChange={(e) => setPermanentAddress(e.target.value)}
								/>
							</div>
							<div className="bg-white dark:bg-gray-900">
								<p className="text-black dark:text-white font-sans text-lg mb-2">
									Postal Code
								</p>
								<input
									type="text"
									className="border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 dark:border-[#fff] dark:focus:outline-none dark:border-opacity-50 dark:opacity-80 dark:bg-gray-500 dark:text-[#fff]"
									placeholder="45962"
									value={postalCode}
									onChange={(e) => setPostalCode(e.target.value)}
								/>
							</div>
						</div>

						{/* Right Column */}
						<div className="w-full md:w-[50%] space-y-6 mt-8 md:mt-0">
							<div className="bg-white dark:bg-gray-900">
								<p className="text-black dark:text-white font-sans text-lg mb-2">
									Username
								</p>
								<input
									type="text"
									className="border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 dark:border-[#fff] dark:focus:outline-none dark:border-opacity-50 dark:opacity-80 dark:bg-gray-500 dark:text-[#fff]"
									value={username}
									onChange={(e) => setUsername(e.target.value)}
								/>
							</div>
							<div className="bg-white dark:bg-gray-900">
								<p className="text-black dark:text-white font-sans text-lg mb-2">
									Present Address
								</p>
								<input
									type="text"
									className="border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 dark:border-[#fff] dark:focus:outline-none dark:border-opacity-50 dark:opacity-80 dark:bg-gray-500 dark:text-[#fff]"
									value={presentAddress}
									onChange={(e) => setPresentAddress(e.target.value)}
								/>
							</div>
							<div className="bg-white dark:bg-gray-900">
								<p className="text-black dark:text-white font-sans text-lg mb-2">
									City
								</p>
								<input
									type="text"
									className="border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 dark:border-[#fff] dark:focus:outline-none dark:border-opacity-50 dark:opacity-80 dark:bg-gray-500 dark:text-[#fff]"
									value={city}
									onChange={(e) => setCity(e.target.value)}
								/>
							</div>
							<div className="bg-white dark:bg-gray-900">
								<p className="text-black dark:text-white font-sans text-lg mb-2">
									Country
								</p>
								<input
									type="text"
									className="border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 dark:border-[#fff] dark:focus:outline-none dark:border-opacity-50 dark:opacity-80 dark:bg-gray-500 dark:text-[#fff]"
									value={country}
									onChange={(e) => setCountry(e.target.value)}
								/>
							</div>
						</div>
					</div>
				</div>

				{/* Save Button */}
				<div className="flex justify-center md:justify-end mt-8 md:mt-12">
					<button
						type="submit"
						className="bg-[#605BFF] text-white rounded-lg py-3 px-10 hover:bg-[#4845d6] focus:outline-none focus:ring-2 focus:ring-[#605BFF]"
						disabled={uploading}
					>
						{uploading ? "Uploading..." : "Save Changes"}
					</button>
				</div>
				{message && (
					<div
						className={`mt-4 text-center ${
							isSuccess ? "text-green-500" : "text-red-500"
						}`}
					>
						{message}
					</div>
				)}
			</form>
		</div>
	);
};

export default EditProfile;
