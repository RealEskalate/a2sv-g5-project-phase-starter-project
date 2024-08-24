"use client";
import React from "react";
import { useForm } from "react-hook-form";
import { useRouter } from "next/navigation";
import { useUserRegistrationMutation } from "@/redux/api/authentication-controller";
import Link from "next/link";
import ErrorMessage from "@/components/Message/ErrorMessage";
import { Country } from "country-state-city";
import { useState } from "react";
import { Currencies } from "@/components/constants/currency";
import { timezones } from "@/components/constants/timezones";

interface PreferenceValues {
  currency: string;
  sentOrReceiveDigitalCurrency: boolean;
  receiveMerchantOrder: boolean;
  accountRecommendations: boolean;
  timeZone: string;
  twoFactorAuthentication: boolean;
}

interface FormValues {
  name: string;
  email: string;
  dateOfBirth: string;
  permanentAddress: string;
  postalCode: string;
  username: string;
  password: string;
  presentAddress: string;
  city: string;
  country: string;
  profilePicture: string;
  preference: PreferenceValues;
}

const Signup = () => {
  const { register, setValue, handleSubmit, formState, watch } =
    useForm<FormValues>();

  const [registerUser] = useUserRegistrationMutation();

  const { errors } = formState;
  console.log(Currencies);

  // onSubmit function
  const onSubmit = async (formData: FormValues) => {
    console.log("formData", formData);

    try {
      const { data, error } = await registerUser({
        name: formData.name,
        email: formData.email,
        dateOfBirth: formData.dateOfBirth,
        password: formData.password,
        username: formData.username,
        permanentAddress: formData.permanentAddress,
        postalCode: formData.postalCode,
        presentAddress: formData.presentAddress,
        city: formData.city,
        country: formData.country,
        profilePicture: formData.profilePicture,
        preference: formData.preference,
      }).unwrap();

      console.log("response from server upon registration", data);
    } catch (error) {
      console.log("error from server upon registration", error);
    }
  };

  const handleProfilePictureChange = (
    e: React.ChangeEvent<HTMLInputElement>
  ) => {
    const file = e.target.files?.[0];
    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => {
        const base64String = reader.result as string;
        setValue("profilePicture", base64String);
      };
      reader.readAsDataURL(file);
    }
  };
  const CountryData = Country.getAllCountries();
  const [selectedCountry, setSelectedCountry] = useState("");

  const handleCountryChange = (e: any) => {
    setSelectedCountry(e.target.value);
  };

  return (
    <div className="w-full h-screen flex justify-center">
      <div className="">
        <h1 className="text-2xl font-poppins font-bold text-center text-[#4640DE] py-10">
          Sign Up Today!
        </h1>

        <form className="flex items-start" onSubmit={handleSubmit(onSubmit)}>
          <div className="flex gap-6 px-6 mx-auto">
            <div className="">
              <div className=" flex flex-col gap-2">
                <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                  Full Name
                </label>
                <input
                  className="w-[350px] px-4 py-2 rounded-xl border border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                  placeholder="Enter your full name"
                  type="text"
                  {...register("name", {
                    required: "Full Name is required",
                    minLength: {
                      value: 2,
                      message: "Full Name must be at least 2 characters long",
                    },
                  })}
                />
                <ErrorMessage message={errors.name?.message} />
              </div>

              <div className="flex flex-col gap-2">
                <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                  Email
                </label>
                <input
                  className="w-[350px] px-4 py-2 rounded-xl border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                  type="email"
                  placeholder="Enter your email"
                  {...register("email", {
                    required: "Email is required",
                    pattern: {
                      value: /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/,
                      message: "Please enter a valid email address",
                    },
                  })}
                />
                <ErrorMessage message={errors.email?.message} />
              </div>

              <div className="w-[350px] flex flex-col gap-2">
                <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                  Date of Birth
                </label>
                <input
                  type="date"
                  {...register("dateOfBirth", {
                    required: "Date of Birth is required",
                  })}
                  className="w-[350px] text-[#515B6F] px-4 py-2 rounded-xl border border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                />
                <ErrorMessage message={errors.dateOfBirth?.message} />
              </div>

              <div className="w-[350px] flex flex-col gap-2">
                <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                  Permanent Address
                </label>
                <input
                  className="w-[350px] px-4 py-2 rounded-xl border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                  placeholder="Enter your permanent address"
                  type="text"
                  {...register("permanentAddress", {
                    required: "Permanent Address is required",
                  })}
                />
                <ErrorMessage message={errors.permanentAddress?.message} />
              </div>

              <div className="w-[350px] flex flex-col gap-2">
                <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                  Postal Code
                </label>
                <input
                  className="w-[350] px-4 py-2 rounded-xl border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                  placeholder="Enter your postal code"
                  type="text"
                  {...register("postalCode", {
                    required: "Postal Code is required",
                  })}
                />
                <ErrorMessage message={errors.postalCode?.message} />
              </div>
            </div>
            <div className="">
              <div className="w-[350px] flex flex-col gap-2">
                <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                  Username
                </label>
                <input
                  className="w-[350] px-4 py-2 rounded-xl border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                  placeholder="Enter a username"
                  type="text"
                  {...register("username", {
                    required: "Username is required",
                  })}
                />
                <ErrorMessage message={errors.username?.message} />
              </div>

              <div className="w-[350px] flex flex-col gap-2">
                <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                  Password
                </label>
                <input
                  className="w-[350] px-4 py-2 rounded-xl border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                  placeholder="Enter your password"
                  type="password"
                  {...register("password", {
                    required: "Password is required",
                    minLength: {
                      value: 6,
                      message: "Password must be at least 6 characters long",
                    },
                  })}
                />
                <ErrorMessage message={errors.password?.message} />
              </div>

              <div className="w-[350] flex flex-col gap-2">
                <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                  Present Address
                </label>
                <input
                  className="w-[350] px-4 py-2 rounded-xl border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                  placeholder="Enter your present address"
                  type="text"
                  {...register("presentAddress", {
                    required: "Present Address is required",
                  })}
                />
                <ErrorMessage message={errors.presentAddress?.message} />
              </div>

              <div className="w-[350px] flex flex-col gap-2">
                <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                  City
                </label>
                <input
                  className="w-[350] px-4 py-2 rounded-xl border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                  placeholder="Enter your city"
                  type="text"
                  {...register("city", { required: "City is required" })}
                />
                <ErrorMessage message={errors.city?.message} />
              </div>

              <div className="w-[350px] flex flex-col gap-2">
                <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                  Country
                </label>
                <select
                  {...register("country", {
                    required: {
                      value: true,
                      message: "Country is required",
                    },
                  })}
                  className="w-[350] px-4 py-2 rounded-xl border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                  value={selectedCountry}
                  onChange={handleCountryChange}
                >
                  <option value="">Select a country</option>
                  {CountryData.map((country) => (
                    <option key={country.name} value={country.name}>
                      {country.name}
                    </option>
                  ))}
                </select>
                <ErrorMessage message={errors.city?.message} />
              </div>
              <div className=" flex flex-col  gap-2 py-10">
                <button
                  type="submit"
                  className="w-full py-2 bg-[#4640DE] text-white rounded-[25px] text-center text-[16px] font-semibold"
                >
                  Continue
                </button>
                <div className="text-center text-sm text-gray-500">
                  Already have an account?{" "}
                  <Link
                    href="/auth/login"
                    className="hover:underline text-[#4640DE] font-medium"
                  >
                    Login
                  </Link>
                </div>
              </div>
            </div>
          </div>

          <div className="">
            <div className="signup-container w-fit h-fit flex flex-col gap-3">
              <h1 className="text-l font-poppins text-center text-[#25324B]">
                Set Your Preferences
              </h1>

              <div className="w-[350px] flex flex-col gap-1">
                <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                  Currency
                </label>
                <select
                  className="px-4 py-2 gap-[8px] rounded-xl border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                  {...register("preference.currency", {
                    required: {
                      value: true,
                      message: "Select a currency",
                    },
                  })}
                >
                  <option value="">Select a currency</option>
                  {Currencies.map((currency) => (
                    <option value={currency.value}>{currency.label}</option>
                  ))}
                </select>
                <ErrorMessage message={errors.preference?.currency?.message} />
              </div>

              <div className="w-[350px] h-[60px] flex flex-col gap-1">
                <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                  Time Zone
                </label>
                <select
                  className="px-4 py-2 gap-[8px] rounded-xl border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                  {...register("preference.timeZone", {
                    required: {
                      value: true,
                      message: "Select a timezone",
                    },
                  })}
                >
                  <option value="">Select a currency</option>
                  {timezones.map((time) => (
                    <option value={time.offset}>{`(${
                      time.offset === 0
                        ? "GMT"
                        : time.offset > 0
                        ? "GMT+"
                        : "GMT-"
                    }${
                      Math.abs(time.offset) > 0 ? Math.abs(time.offset) : ""
                    }) ${time.name}`}</option>
                  ))}
                </select>
                <ErrorMessage message={errors.preference?.timeZone?.message} />
              </div>

              <div className="flex flex-col items-start justify-center mt-4 gap-1">
                <div className=" flex flex-row-reverse  gap-2">
                  <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                    Send or Receive Digital Currency
                  </label>
                  <input
                    className=" rounded-[6px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px]]"
                    type="checkbox"
                    {...register("preference.sentOrReceiveDigitalCurrency")}
                  />
                </div>

                <div className=" flex flex-row-reverse gap-2">
                  <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                    Receive Merchant Order
                  </label>
                  <input
                    className="rounded-[6px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px]]"
                    type="checkbox"
                    {...register("preference.receiveMerchantOrder")}
                  />
                </div>

                <div className=" flex flex-row-reverse gap-2">
                  <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                    Account Recommendations
                  </label>
                  <input
                    className=" rounded-[6px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px]]"
                    type="checkbox"
                    {...register("preference.accountRecommendations")}
                  />
                </div>

                <div className=" flex flex-row-reverse gap-2">
                  <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                    Two-Factor Authentication
                  </label>
                  <input
                    className=" rounded-[6px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px]]"
                    type="checkbox"
                    {...register("preference.twoFactorAuthentication")}
                  />
                </div>
              </div>
              <div className="w-[340px] flex flex-col gap-1 ">
                <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                  Profile Picture
                </label>
                <div className="w-[340px] border rounded-xl px-2 flex flex-row gap-2 items-center">
                  <img src="/pubimg/picture.svg" className="size-8" />

                  <input
                    className="w-full  py-2 border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                    type="file"
                    accept="image/*"
                    onChange={handleProfilePictureChange}
                  />
                </div>
                <ErrorMessage message={errors.profilePicture?.message} />
              </div>
            </div>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Signup;
