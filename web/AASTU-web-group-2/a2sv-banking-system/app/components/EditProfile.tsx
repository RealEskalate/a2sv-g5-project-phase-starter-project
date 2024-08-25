"use client";
import React, { useEffect, useState } from "react";
import Image from "next/image";
import { FaPencilAlt, FaCaretDown } from "react-icons/fa";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import { useForm, Controller } from "react-hook-form";
import User, { UserInfo } from "@/types/userInterface";
import {
  getCurrentUser,
  getUserByUsername,
  userUpdate,
} from "@/lib/api/userControl";
import Refresh from "../api/auth/[...nextauth]/token/RefreshToken";
import { storage } from "../firebase";
import {
  ref,
  uploadBytes,
  getDownloadURL,
  deleteObject,
} from "firebase/storage";
import { v4 } from "uuid";
interface FormData {
  name: string;
  username: string;
  email: string;
  city: string;
  dateOfBirth: Date | null;
  presentAddress: string;
  permanentAddress: string;
  country: string;
  postalCode: string;
  profilePicture: string | File;
}

const EditProfile = () => {
  const { control, handleSubmit, setValue } = useForm<FormData>();
  const [user, setUser] = useState<UserInfo | null>(null);
  const [accessToken, setAccessToken] = useState<string>("");
  useEffect(() => {
    const fetchData = async () => {
      try {
        const accessToken = await Refresh();
        setAccessToken(accessToken);
      } catch (error) {
        console.error("Error fetching token:", error);
      }
    };

    fetchData();
  }, []);

  useEffect(() => {
    const fetchData = async () => {
      try {
        if (accessToken) {
          const currentUser = await getCurrentUser(accessToken);
          const userData = await getUserByUsername(
            currentUser.username,
            accessToken
          );
          setUser(userData);

          // Populate form fields
          setValue("name", userData.name || "");
          setValue("username", userData.username || "");
          setValue("email", userData.email || "");
          setValue("city", userData.city || "");
          setValue(
            "dateOfBirth",
            userData.dateOfBirth ? new Date(userData.dateOfBirth) : null
          );
          setValue("presentAddress", userData.presentAddress || "");
          setValue("permanentAddress", userData.permanentAddress || "");
          setValue("country", userData.country || "");
          setValue("postalCode", userData.postalCode || "");
        }
      } catch (error) {
        console.error("Error fetching user data:", error);
      }
    };

    fetchData();
  }, [accessToken, setValue]);

  const onSubmit = async (data: FormData) => {
    try {
      if (data.profilePicture != null) {
        // Update the data with the download URL for profilePicture
        const updatedData = {
          ...data,
          profilePicture: String(data.profilePicture),
          dateOfBirth: data.dateOfBirth
            ? data.dateOfBirth.toISOString().split("T")[0]
            : null,
        };

        // Send the updated data to userUpdate
        await userUpdate(updatedData, accessToken);
      }

      console.log("Profile updated successfully");
      alert("Profile Edited Successfully")
    } catch (error) {
      console.error("Error updating profile:", error);
    }
  };

  return (
    <>
      <div className="flex flex-col items-center py-10 w-full dark:bg-[#020817]">
        <div className="relative">
          <Image
            src={
              user?.profilePicture ||
              "https://firebasestorage.googleapis.com/v0/b/a2sv-wallet.appspot.com/o/images%2Fminions-removebg-preview.png-99cefd58-79e9-408d-b747-94bcb3bb16ab?alt=media&token=5822c470-99fb-4875-a4fc-425a64bf1473" || "/ProfilePicture"
            } // Fallback to a default image if userData.profilePicture is null or undefined
            alt="Profile Picture"
            width={170}
            height={170}
            className="rounded-full"
          />

          {/* Hidden file input for selecting a new profile picture */}
          <input
            type="file"
            accept="image/*"
            onChange={async (e: React.ChangeEvent<HTMLInputElement>) => {
              const file = e.target.files?.[0]; // Check if files exist
              if (file) {
                setValue("profilePicture", file);

                // Immediately upload the image to Firebase
                try {
                  const imageRef = ref(storage, `images/${file.name}-${v4()}`);
                  await uploadBytes(imageRef, file);

                  // Get the download URL after the image is uploaded
                  const downloadUrl = await getDownloadURL(imageRef);
                  console.log("dowloaded the url ", downloadUrl);

                  // Update the profile picture in the form state and in the UI
                  setValue("profilePicture", downloadUrl);
                  setUser(
                    (prev) => prev && { ...prev, profilePicture: downloadUrl }
                  );
                } catch (error) {
                  console.error("Error uploading image:", error);
                }
              }
            }}
            style={{ display: "none" }} // Hide the input
            id="profilePictureInput"
          />

          {/* Label for the file input, styled as an edit icon */}
          <label htmlFor="profilePictureInput">
            <span className="absolute bottom-5 right-0 bg-[#1814F3] rounded-full w-10 h-10 flex items-center justify-center text-white cursor-pointer">
              <FaPencilAlt />
            </span>
          </label>
        </div>
      </div>

      <div>
        <form onSubmit={handleSubmit(onSubmit)}>
          <div className="flex flex-col gap-3 dark:bg-[#020817]">
            <div className="flex flex-col md:flex md:flex-row">
              {/* Name Goes In Here */}
              <div className="flex flex-col gap-3 px-6 py-3 md:w-[48%]">
                <label
                  htmlFor="name"
                  className="text-[#232323] font-semibold px-1 dark:text-[#9faaeb]"
                >
                  Your Name
                </label>
                <Controller
                  name="name"
                  control={control}
                  defaultValue=""
                  render={({ field }) => (
                    <input
                      {...field}
                      type="text"
                      id="name"
                      className="border border-[#DFEAF2] focus:outline-[#DFEAF2] focus:border-[#DFEAF2] rounded-xl py-3 px-6 placeholder:text-[#718EBF] dark:border-gray-600 dark:focus:outline-none dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
                      placeholder="Rebuma Tadele"
                    />
                  )}
                />
              </div>

              {/* Username Goes In Here */}
              <div className="flex flex-col gap-3 px-6 py-3 md:w-[48%]">
                <label
                  htmlFor="username"
                  className="text-[#232323] font-semibold px-1 dark:text-[#9faaeb]"
                >
                  Username
                </label>
                <Controller
                  name="username"
                  control={control}
                  defaultValue=""
                  render={({ field }) => (
                    <input
                      {...field}
                      type="text"
                      id="username"
                      className="border border-[#DFEAF2] focus:outline-[#DFEAF2] focus:border-[#DFEAF2] rounded-xl py-3 px-6 placeholder:text-[#718EBF] dark:border-gray-600 dark:focus:outline-none dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
                      placeholder="rebuma"
                    />
                  )}
                />
              </div>
            </div>

            <div className="flex flex-col md:flex md:flex-row">
              {/* Email Goes In Here */}
              <div className="flex flex-col gap-3 px-6 py-3 md:w-[48%]">
                <label
                  htmlFor="email"
                  className="text-[#232323] font-semibold px-1 dark:text-[#9faaeb]"
                >
                  Email
                </label>
                <Controller
                  name="email"
                  control={control}
                  defaultValue=""
                  render={({ field }) => (
                    <input
                      {...field}
                      type="email"
                      id="email"
                      className="border border-[#DFEAF2] focus:outline-[#DFEAF2] focus:border-[#DFEAF2] rounded-xl py-3 px-6 placeholder:text-[#718EBF] dark:border-gray-600 dark:focus:outline-none dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
                      placeholder="john@example.com"
                    />
                  )}
                />
              </div>

              {/* City Goes In Here */}
              <div className="flex flex-col gap-3 px-6 py-3 md:w-[48%]">
                <label
                  htmlFor="city"
                  className="text-[#232323] font-semibold px-1 dark:text-[#9faaeb]"
                >
                  City
                </label>
                <Controller
                  name="city"
                  control={control}
                  defaultValue=""
                  render={({ field }) => (
                    <input
                      {...field}
                      type="text"
                      id="city"
                      className="border border-[#DFEAF2] focus:outline-[#DFEAF2] focus:border-[#DFEAF2] rounded-xl py-3 px-6 placeholder:text-[#718EBF] dark:border-gray-600 dark:focus:outline-none dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
                      placeholder="Addis Ababa"
                    />
                  )}
                />
              </div>
            </div>

            <div className="flex flex-col md:flex md:flex-row">
              {/* Date Of Birth Goes In Here */}
              <div className="flex flex-col gap-3 px-6 py-3 md:w-[48%]">
                <label
                  htmlFor="datePicker"
                  className="text-[#232323] font-semibold px-1 dark:text-[#9faaeb]"
                >
                  Date Of Birth
                </label>
                <div className="relative w-full">
                  <Controller
                    name="dateOfBirth"
                    control={control}
                    defaultValue={null}
                    render={({ field: { onChange, value } }) => (
                      <DatePicker
                        selected={value}
                        onChange={(date: Date | null) => onChange(date)}
                        placeholderText="Date Of Birth"
                        className="w-full border border-[#DFEAF2] focus:outline-[#DFEAF2] focus:border-[#DFEAF2] rounded-xl py-3 px-6 placeholder:text-[#718EBF] bg-white dark:border-gray-600 dark:focus:outline-none dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
                        dateFormat="MMMM d, yyyy"
                        id="datePicker"
                      />
                    )}
                  />
                  <FaCaretDown className="absolute top-1/2 right-4 transform -translate-y-1/2 text-gray-500 dark:text-[#9faaeb]" />
                </div>
              </div>

              {/* Present Address Goes In Here */}
              <div className="flex flex-col gap-3 px-6 py-3 md:w-[48%]">
                <label
                  htmlFor="presentAddress"
                  className="text-[#232323] font-semibold px-1 dark:text-[#9faaeb]"
                >
                  Present Address
                </label>
                <Controller
                  name="presentAddress"
                  control={control}
                  defaultValue=""
                  render={({ field }) => (
                    <input
                      {...field}
                      type="text"
                      id="presentAddress"
                      className="border border-[#DFEAF2] focus:outline-[#DFEAF2] focus:border-[#DFEAF2] rounded-xl py-3 px-6 placeholder:text-[#718EBF] dark:border-gray-600 dark:focus:outline-none dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
                      placeholder="Present Address"
                    />
                  )}
                />
              </div>
            </div>

            <div className="flex flex-col md:flex md:flex-row">
              {/* Permanent Address Goes In Here */}
              <div className="flex flex-col gap-3 px-6 py-3 md:w-[48%]">
                <label
                  htmlFor="permanentAddress"
                  className="text-[#232323] font-semibold px-1 dark:text-[#9faaeb]"
                >
                  Permanent Address
                </label>
                <Controller
                  name="permanentAddress"
                  control={control}
                  defaultValue=""
                  render={({ field }) => (
                    <input
                      {...field}
                      type="text"
                      id="permanentAddress"
                      className="border border-[#DFEAF2] focus:outline-[#DFEAF2] focus:border-[#DFEAF2] rounded-xl py-3 px-6 placeholder:text-[#718EBF] dark:border-gray-600 dark:focus:outline-none dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
                      placeholder="Permanent Address"
                    />
                  )}
                />
              </div>

              {/* Country Goes In Here */}
              <div className="flex flex-col gap-3 px-6 py-3 md:w-[48%]">
                <label
                  htmlFor="postalCode"
                  className="text-[#232323] font-semibold px-1 dark:text-[#9faaeb]"
                >
                  Postal Code
                </label>
                <Controller
                  name="postalCode"
                  control={control}
                  render={({ field }) => (
                    <input
                      {...field}
                      type="text"
                      id="postalCode"
                      className="border border-[#DFEAF2] focus:outline-[#DFEAF2] focus:border-[#DFEAF2] rounded-xl py-3 px-6 placeholder:text-[#718EBF] dark:border-gray-600 dark:focus:outline-none dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
                      placeholder="postal Code"
                    />
                  )}
                />
              </div>

              {/* Country Goes In Here */}
              <div className="flex flex-col gap-3 px-6 py-3 md:w-[48%]">
                <label
                  htmlFor="country"
                  className="text-[#232323] font-semibold px-1 dark:text-[#9faaeb]"
                >
                  Country
                </label>
                <Controller
                  name="country"
                  control={control}
                  defaultValue=""
                  render={({ field }) => (
                    <input
                      {...field}
                      type="text"
                      id="country"
                      className="border border-[#DFEAF2] focus:outline-[#DFEAF2] focus:border-[#DFEAF2] rounded-xl py-3 px-6 placeholder:text-[#718EBF] dark:border-gray-600 dark:focus:outline-none dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
                      placeholder="Country"
                    />
                  )}
                />
              </div>
            </div>

            <div className="text-center">
              <button
                type="submit"
                className="bg-[#1814F3] text-white rounded-xl py-3 px-6"
              >
                Save Changes
              </button>
            </div>
          </div>
        </form>
      </div>
    </>
  );
};

export default EditProfile;
