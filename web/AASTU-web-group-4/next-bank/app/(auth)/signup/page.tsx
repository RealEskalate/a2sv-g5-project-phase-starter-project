"use client";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { User } from "@/types/index";
import { creditcardstyles, colors, logo } from "../../../constants/index";
import Image from "next/image";
import Link from "next/link";
import { registerUser } from "@/services/authentication";
import Cookie from "js-cookie";
import { uploadImage } from "@/components/Imageupload"; // Import your upload function

const SignupForm = () => {
  const [step, setStep] = useState(1);
  const [formData, setFormData] = useState<User>({
    name: "",
    email: "",
    dateOfBirth: "",
    permanentAddress: "",
    postalCode: "",
    username: "",
    password: "",
    presentAddress: "",
    city: "",
    country: "",
    profilePicture: "",
    preference: {
      currency: "",
      sentOrReceiveDigitalCurrency: false,
      receiveMerchantOrder: false,
      accountRecommendations: false,
      timeZone: "",
      twoFactorAuthentication: false,
    },
  });

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<User>();

  const [imageUrl, setImageUrl] = useState("");

  const handleFileChange = async (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      try {
        const url = await uploadImage(file);
        setImageUrl(url);
        console.log("Image uploaded successfully:", url);
      } catch (error) {
        console.error("Error uploading image:", error);
      }
    }
  };

  const onSubmit = async (data: User) => {
    data.profilePicture = imageUrl ? imageUrl : data.profilePicture = imageUrl;
    console.log(data, imageUrl);
    try {
      const registeredUser = await registerUser(data);
      console.log("Registered User:", registeredUser);
      Cookie.set("accessToken", registeredUser.data.access_token);
      Cookie.set("refreshToken", registeredUser.data.refresh_token);
    } catch (error) {
      console.error("Registration Error:", error);
    }
  };
  const handleNextStep = () => setStep(step + 1);
  const handlePreviousStep = () => setStep(step - 1);

  return (
    <form
      onSubmit={handleSubmit(onSubmit)}
      className="max-w-3xl mx-auto p-8 bg-white shadow-md rounded-lg "
    >
      <div className="flex justify-center items-center mb-8">
        <Image src={logo.icon} alt="Logo" height={60} width={60} />
        <h1 className="font-bold text-3xl text-gray-700 font-serif p-2">
          {" "}
          <p className=" text-gray-600">NEXT BANK</p>
        </h1>
      </div>
      <h2 className="text-blue-600 font-semibold text-2xl text-center mb-4">
        Sign Up
      </h2>
      <p className="text-center text-gray-600 mb-8">
        Please enter your details
      </p>

      {step === 1 && (
        <div className="grid gap-6 sm:grid-cols-1 md:grid-cols-2">
          {/* Step 1 Fields */}

          <div>
            <label htmlFor="name" className="block font-medium text-gray-700">
              Full Name
            </label>
            <input
              type="text"
              id="name"
              placeholder="ex: John"
              {...register("name", { required: "Name is required" })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.name && (
              <p className="text-red-500 mt-1">{errors.name.message}</p>
            )}
          </div>
          <div>
            <label htmlFor="email" className="block font-medium text-gray-700">
              Email
            </label>
            <input
              type="email"
              id="email"
              placeholder="ex: example@gmail.com"
              {...register("email", {
                required: "Email is required",
                pattern: /^\S+@\S+\.\S+$/,
              })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.email && (
              <p className="text-red-500 mt-1">{errors.email.message}</p>
            )}
          </div>

          {/* dateOfBirth */}
          <div>
            <label
              htmlFor="dateOfBirth"
              className="block font-medium text-gray-700"
            >
              Date of Birth
            </label>
            <input
              type="date"
              id="dateOfBirth"
              {...register("dateOfBirth", {
                required: "Date of Birth is required",
              })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.dateOfBirth && (
              <p className="text-red-500 mt-1">{errors.dateOfBirth.message}</p>
            )}
          </div>

          {/* username */}
          <div>
            <label
              htmlFor="username"
              className="block font-medium text-gray-700"
            >
              Username
            </label>
            <input
              type="text"
              id="username"
              placeholder="ex: Star"
              {...register("username", {
                required: "Username is required",
              })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.username && (
              <p className="text-red-500 mt-1">{errors.username.message}</p>
            )}
          </div>

          {/* password */}
          <div>
            <label
              htmlFor="password"
              className="block font-medium text-gray-700"
            >
              Password
            </label>
            <input
              type="password"
              id="password"
              placeholder="ex: password"
              {...register("password", {
                required: "Password is required",
                minLength: 8,
              })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.password && (
              <p className="text-red-500 mt-1">{errors.password.message}</p>
            )}
          </div>

          {/* permanentAddress */}
          <div className="md:row-span-2">
            <label
              htmlFor="permanentAddress"
              className="block font-medium text-gray-700"
            >
              Permanent Address
            </label>
            <textarea
              id="permanentAddress"
              placeholder="ex: A.A"
              {...register("permanentAddress", {
                required: "Permanent Address is required",
              })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.permanentAddress && (
              <p className="text-red-500 mt-1">
                {errors.permanentAddress.message}
              </p>
            )}
          </div>

          {/* postalCode */}
          <div>
            <label
              htmlFor="postalCode"
              className="block font-medium text-gray-700"
            >
              Postal Code
            </label>
            <input
              type="text"
              id="postalCode"
              placeholder="ex: 1000"
              {...register("postalCode", {
                required: "Postal Code is required",
              })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.postalCode && (
              <p className="text-red-500 mt-1">{errors.postalCode.message}</p>
            )}
          </div>

          <div className="md:col-span-2 flex justify-center">
            <button
              type="button"
              onClick={handleNextStep}
              className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-3 px-6 rounded-lg mt-8 transition duration-300"
            >
              Continue
            </button>
          </div>
          {/* Link to Sign In Page */}
          <div className="md:col-span-2 flex justify-center">
            <p className="mx-2">Already have an account?</p>
            <a href="/signin" className="text-blue-800">
              Login
            </a>
          </div>
        </div>
      )}

      {step === 2 && (
        <div className="grid gap-6 sm:grid-cols-1 md:grid-cols-2">
          {/* Step 2 Fields */}
          {/* presentAddress */}
          <div>
            <label
              htmlFor="presentAddress"
              className="block font-medium text-gray-700"
            >
              Present Address
            </label>
            <input
              type="text"
              id="presentAddress"
              placeholder="ex: Bola"
              {...register("presentAddress", {
                required: "Present Address is required",
              })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.presentAddress && (
              <p className="text-red-500 mt-1">
                {errors.presentAddress.message}
              </p>
            )}
          </div>

          {/* city */}
          <div>
            <label htmlFor="city" className="block font-medium text-gray-700">
              City
            </label>
            <input
              type="text"
              id="city"
              placeholder="ex: A.A"
              {...register("city", { required: "City is required" })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.city && (
              <p className="text-red-500 mt-1">{errors.city.message}</p>
            )}
          </div>

          {/* country */}
          <div>
            <label
              htmlFor="country"
              className="block font-medium text-gray-700"
            >
              Country
            </label>
            <input
              type="text"
              id="country"
              placeholder="ex: Ethiopia"
              {...register("country", { required: "Country is required" })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.country && (
              <p className="text-red-500 mt-1">{errors.country.message}</p>
            )}
          </div>

          {/* profilePicture */}
          <div>
            <label
              htmlFor="profilePicture"
              className="mb-3 font-medium text-gray-700"
            >
              Profile Picture
            </label>
            <input
              type="file"
              id="profilePicture"
              {...register("profilePicture")}
              onChange={handleFileChange}
              className="p-2 w-full"
            />
          </div>

          {/* Currency Preference */}
          <div>
            <label
              htmlFor="currency"
              className="block font-medium text-gray-700"
            >
              Preferred Currency
            </label>
            <select
              id="currency"
              {...register("preference.currency", {
                required: "Currency is required",
              })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            >
              <option value="">Select Currency</option>
              <option value="USD">USD</option>
              <option value="EUR">EUR</option>
              <option value="ETB">ETB</option>
              {/* Add more currencies as needed */}
            </select>
            {errors.preference?.currency && (
              <p className="text-red-500 mt-1">
                {errors.preference.currency.message}
              </p>
            )}
          </div>

          <div>
            <label htmlFor="timeZone">Time Zone</label>
            <select
              id="timeZone"
              {...register("preference.timeZone", {
                required: "Time Zone is required",
              })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            >
              <option value="">Select Time Zone</option>
              <option value="GMT">GMT</option>
              <option value="UTC+05:00">EST</option>
              <option value="UTC-08:00">PST</option>
              <option value="UTC+01:00">JST</option>
              {/* Add more time zone as needed */}
            </select>
            {errors.preference?.timeZone && (
              <p className="text-red-500 mt-1">
                {errors.preference?.timeZone.message}
              </p>
            )}
          </div>

          <div>
            <label
              htmlFor="sentOrReceiveDigitalCurrency"
              className="font-medium pr-3 text-gray-700"
            >
              Sent or Receive Digital Currency
            </label>
            <input
              type="checkbox"
              id="sentOrReceiveDigitalCurrency"
              {...register("preference.sentOrReceiveDigitalCurrency")}
            />
          </div>

          <div>
            <label
              htmlFor="receiveMerchantOrder"
              className="font-medium pr-3 text-gray-700"
            >
              Receive Merchant Order
            </label>
            <input
              type="checkbox"
              id="receiveMerchantOrder"
              {...register("preference.receiveMerchantOrder")}
            />
          </div>

          <div>
            <label
              htmlFor="accountRecommendations"
              className="font-medium pr-3 text-gray-700"
            >
              Account Recommendations
            </label>
            <input
              type="checkbox"
              id="accountRecommendations"
              {...register("preference.accountRecommendations")}
            />
          </div>

          {/* Two-Factor Authentication */}
          <div className="md:row">
            <label
              htmlFor="twoFactorAuthentication"
              className="pr-3 font-medium text-gray-700"
            >
              Enable Two-Factor Authentication
            </label>
            <input
              type="checkbox"
              id="twoFactorAuthentication"
              {...register("preference.twoFactorAuthentication")}
            />
          </div>

          <div className="md:col-span-2 flex justify-between">
            <button
              type="button"
              onClick={handlePreviousStep}
              className="bg-gray-500 hover:bg-gray-600 text-white font-bold py-3 px-6 rounded-lg mt-8 transition duration-300"
            >
              Back
            </button>
            <button
              type="submit"
              className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-3 px-6 rounded-lg mt-8 transition duration-300"
            >
              Sign Up
            </button>
          </div>
        </div>
      )}
      {/* Link to Sign In Page */}
      {/* <div className="flex m-4 text-center justify-center">
          <p className='mx-2'>Already have an account?</p>
          <a  href="/signin" className="text-blue-800">Login</a>
        </div> */}
    </form>
  );
};

export default SignupForm;
