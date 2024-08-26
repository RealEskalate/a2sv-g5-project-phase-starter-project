import React, { useState } from "react";
import { useFormContext, Controller } from "react-hook-form";
import Image from "next/image";
import { FaPencilAlt } from "react-icons/fa";
import { ref, uploadBytes, getDownloadURL } from "firebase/storage";
import { v4 } from "uuid";
import { storage } from "@/app/firebase";
import { BsExclamationCircle } from "react-icons/bs";

type ToggleType = {
  name: string;
  label: string;
  control: any;
};

// Reusable Toggle Switch Component
const ToggleSwitch = ({ name, label, control }: ToggleType) => (
  <div className="flex items-center justify-between mb-3">
    <label className="text-[#515B6F] font-semibold dark:text-[#cdd6f4]">
      {label}
    </label>
    <Controller
      name={name}
      control={control}
      defaultValue={false}
      render={({ field: { onChange, value } }) => (
        <div
          className={`relative inline-flex items-center h-6 rounded-full w-11 cursor-pointer transition-colors duration-200 ${
            value ? "bg-indigo-600" : "bg-gray-400"
          }`}
          onClick={() => onChange(!value)}
        >
          <span
            className={`transform transition-transform duration-200 ease-in-out ${
              value ? "translate-x-6" : "translate-x-1"
            } inline-block w-4 h-4 bg-white rounded-full`}
          />
        </div>
      )}
    />
  </div>
);

// Reusable error message component
const ErrorMessage = ({ message }: any) => (
  <div className="flex flex-row align-middle mt-2">
    <BsExclamationCircle className="text-red-500 mr-2" />
    <span className="text-red-500">{message}</span>
  </div>
);

const PageThree = () => {
  const { control, setValue, formState: { errors } } = useFormContext();
  const [user, setUser] = useState({
    profilePicture: "",
  });

  const [loading, setLoading] = useState(false); // Loading state for image upload
  const [error, setError] = useState(""); // Error state for image upload

  return (
    <div className="flex flex-col gap-5 w-full md:w-1/2 mx-auto">
      <h1 className="text-4xl font-black text-[#202430] dark:text-[#cdd6f4] mt-4 text-center">
        Account Settings
      </h1>
      
      {/* Profile Picture Upload */}
      <div className="relative mt-4 w-24 h-24 mx-auto">
        {loading && (
          <div className="absolute inset-0 flex items-center justify-center bg-opacity-75 rounded-full ">
            <div className="loader min-w-fit flex-nowrap text-blue-500">Uploading...</div>
          </div>
        )}
        <Image
          src={
            user?.profilePicture && user?.profilePicture !== "string"
              ? user.profilePicture
              : "https://firebasestorage.googleapis.com/v0/b/a2sv-wallet.appspot.com/o/images%2Fminions-removebg-preview.png-99cefd58-79e9-408d-b747-94bcb3bb16ab?alt=media&token=5822c470-99fb-4875-a4fc-425a64bf1473"
          }
          alt="Profile Picture"
          width={96} // Smaller width
          height={96} // Smaller height
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
              setLoading(true);

              // Validate file type and size before upload
              if (!file.type.startsWith("image/")) {
                setError("Please select a valid image file.");
                setLoading(false);
                return;
              }
              if (file.size > 5 * 1024 * 1024) {
                setError("File size should be less than 5MB.");
                setLoading(false);
                return;
              }

              // Immediately upload the image to Firebase
              try {
                const imageRef = ref(storage, `images/${file.name}-${v4()}`);
                await uploadBytes(imageRef, file);

                // Get the download URL after the image is uploaded
                const downloadUrl = await getDownloadURL(imageRef);

                // Update the profile picture in the form state and in the UI
                setValue("profilePicture", downloadUrl);
                setUser((prev) => ({ ...prev, profilePicture: downloadUrl }));
                setError(""); // Clear any previous errors
              } catch (error) {
                setError("Error uploading image. Please try again.");
                console.error("Error uploading image:", error);
              } finally {
                setLoading(false);
              }
            }
          }}
          style={{ display: "none" }} // Hide the input
          id="profilePictureInput"
        />

        {/* Label for the file input, styled as an edit icon */}
        <label htmlFor="profilePictureInput">
          <span className="absolute bottom-0 right-0 bg-[#1814F3] rounded-full w-6 h-6 flex items-center justify-center text-white cursor-pointer">
            <FaPencilAlt size={12} />
          </span>
        </label>
      </div>

      {/* Error Message for Image Upload */}
      {error && <ErrorMessage message={error} />}

      {/* Currency Field */}
      <div className="mb-3 flex flex-col gap-2">
        <label className="text-[#515B6F] font-semibold dark:text-[#cdd6f4]">
          Currency
        </label>
        <Controller
          name="currency"
          control={control}
          defaultValue=""
          rules={{ required: "Currency is required" }} // Making currency field required
          render={({ field, fieldState: { error } }) => (
            <>
              <input
                {...field}
                type="text"
                className="border border-gray-400 rounded-lg py-2 px-5 w-full dark:border-gray-600 dark:bg-[#313244] dark:text-[#cdd6f4] focus:outline-none focus:ring-2 focus:ring-indigo-500 dark:focus:border-[#4640DE]"
                placeholder="Enter your preferred currency"
                aria-label="Currency"
              />
              {error && <ErrorMessage message={error.message} />}
            </>
          )}
        />
      </div>

      {/* Toggles */}
      <ToggleSwitch
        name="sentOrReceiveDigitalCurrency"
        label="Send or Receive Digital Currency"
        control={control}
      />
      <ToggleSwitch
        name="receiveMerchantOrder"
        label="Receive Merchant Order"
        control={control}
      />
      <ToggleSwitch
        name="accountRecommendations"
        label="Account Recommendations"
        control={control}
      />
      <ToggleSwitch
        name="twoFactorAuthentication"
        label="Two-Factor Authentication"
        control={control}
      />
    </div>
  );
};

export default PageThree;
