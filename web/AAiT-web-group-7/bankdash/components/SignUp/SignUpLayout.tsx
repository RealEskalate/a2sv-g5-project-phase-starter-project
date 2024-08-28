"use client";
import React from "react";
import { useState } from "react";
import Link from "next/link";
import { useForm, Controller } from "react-hook-form";
import { useEffect } from "react";
import ErrorMessage from "@/components/Message/ErrorMessage";
import { Switch } from "@/components/ui/switch";
import { Currencies } from "@/components/constants/currency";
import { timezones } from "@/components/constants/timezones";
import { CountryData } from "@/components/constants/countries";
import { useUserRegistrationMutation } from "@/redux/api/authentication-controller";
import { useRouter } from "next/navigation";

interface prefData {
  timezone: string;
  currency: string;
  transaction: boolean;
  merchant: boolean;
  recommendation: boolean;
  twoFactorAuth: boolean;
}
interface FormData {
  Name: string;
  Email: string;
  DOT: string;
  PA: string;
  PC: string;
  UN: string;
  password: string;
  PresentAddress: string;
  City: string;
  Country: string;
  profilePicture: string;
}

const PageLayout = () => {
  const route=useRouter();
  const [registerUser] = useUserRegistrationMutation();
  const [activeButton, setActiveButton] = useState("edit");
  const [profileImage, setProfileImage] = useState<File | null>(null);

  const form = useForm<FormData>();
  const { register, handleSubmit, formState, setValue, getValues } = form;

  const {
    control,
    handleSubmit: handleSubmitPref,
    formState: { errors: errorsPref },
  } = useForm<prefData>({
    defaultValues: {
      currency: "",
      timezone: "",
      transaction: false,
      merchant: false,
      recommendation: false,
      twoFactorAuth: false,
    },
  });

  const { errors } = formState;

  const onSubmit = (formData: FormData) => {
    console.log("formData", formData);
    setActiveButton("preferences");
  };
  const onsubmitPref = async (formData: prefData) => {
    console.log("formData", formData);
    const allValues = getValues();
    console.log("All Form Values:", allValues);

    try {
      const { data, success } = await registerUser({
        name: allValues.Name,
        email: allValues.Email,
        dateOfBirth: allValues.DOT,
        permanentAddress: allValues.PA,
        postalCode: allValues.PC,
        username: allValues.UN,
        password: allValues.password,
        presentAddress: allValues.PresentAddress,
        city: allValues.City,
        country: allValues.Country,
        profilePicture: allValues.profilePicture,
        preference: {
          currency: formData.currency,
          sentOrReceiveDigitalCurrency: formData.transaction,
          receiveMerchantOrder: formData.merchant,
          accountRecommendations: formData.recommendation,
          timeZone: formData.timezone,
          twoFactorAuthentication: formData.recommendation,
        },
      }).unwrap();
      if (success) {
        console.log("response from server upon registration", data);
        route.push("/auth/login");
      }

      else {
        console.log("error from server upon registration");
      }
    } catch (error) {
      console.log("error from server upon registration");
    }
  };

  const handleProfilePictureChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const file = event.target.files?.[0];
    console.log(file);

    if (file) {
      setProfileImage(file);
      setValue("profilePicture", file.name);
    }
  };

  return (
    <div className="px-5 py-5 flex justify-center">
      <div className="flex flex-col rounded-3xl w-fill px-10  bg-white ">
        <h1 className="text-4xl font-poppins font-bold text-center text-[#4640DE] py-5 flex flex-row justify-center gap-2">
          Sign Up Today
          <img
            src="/pubimg/signup.svg"
            className="size-10 flex flex-row justify-center text-[#4640DE]"
          />
        </h1>
        <div className="flex flex-row font-serif   w-fill text-[#718EBF] gap-12 ">
          <div
            className={`items-center flex flex-row gap-2 p-2 font-black ${
              activeButton === "edit"
                ? "border-b-2  text-[#1814F3] border-[#1814F3]"
                : ""
            }  `}
          >
            <div className="flex border size-8 justify-center items-center rounded-full">
              1
            </div>
            <div className="">Profile</div>
          </div>
          <div
            className={`items-center flex flex-row gap-2 p-2 font-black ${
              activeButton === "preferences"
                ? "border-b-2  text-[#1814F3] border-[#1814F3]"
                : ""
            }  `}
          >
            <div className="flex border size-8 justify-center items-center rounded-full">
              2
            </div>
            <div className="">Preferences</div>
          </div>
        </div>
        <div>
          {activeButton === "edit" && (
            <form
              onSubmit={handleSubmit(onSubmit)}
              className="flex flex-col text-sm"
            >
              <div className="flex  gap-8 py-10">
                <div className="relative ">
                  {profileImage ? (
                    <img
                      src={URL.createObjectURL(profileImage)}
                      className="size-32 rounded-full"
                    />
                  ) : (
                    <img src="/pubimg/placepp.png" className="size-32" />
                  )}

                  <div className="bg-[#1814F3]  h-8 w-8 flex justify-center items-center rounded-full absolute right-0 top-20 hover:brightness-200 transition duration-200">
                    <label htmlFor="fileInput" className="cursor-pointer">
                      <img src="/pubimg/pencil.svg" />
                    </label>
                    <input
                      id="fileInput"
                      type="file"
                      accept="image/*"
                      // {...register("profilePicture", {
                      //   required: {
                      //     value: true,
                      //     message: "Picture is required",
                      //   },
                      // })}
                      onChange={handleProfilePictureChange}
                      className="hidden"
                    />
                  </div>
                  <ErrorMessage
                    message={errors.profilePicture?.message as string}
                  />
                </div>
                <div className="flex flex-col items-center gap-5">
                  <div className="flex gap-8">
                    <div className="flex flex-col gap-3">
                      <div className="flex flex-col items-start justify-center gap-2 ">
                        <label className="text-[#232323] ">Your Name</label>
                        <input
                          id="name"
                          {...register("Name", {
                            required: {
                              value: true,
                              message: "Full Name is required",
                            },
                          })}
                          placeholder="Full Name"
                          className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
                        />

                        <ErrorMessage message={errors.Name?.message} />
                      </div>
                      <div className="flex flex-col items-start justify-center gap-2">
                        <label className="text-[#232323] ">Email</label>
                        <input
                          placeholder="Email"
                          type="email"
                          id="email"
                          {...register("Email", {
                            required: {
                              value: true,
                              message: "Email is required",
                            },
                            pattern: {
                              value:
                                /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/,
                              message: "Invalid Email",
                            },
                          })}
                          className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
                        />
                        <ErrorMessage message={errors.Email?.message} />
                      </div>

                      <div className="flex flex-col items-start justify-center gap-2">
                        <label className="text-[#232323] ">Date of Birth</label>
                        <input
                          type="date"
                          {...register("DOT", {
                            required: {
                              value: true,
                              message: "Date of Birth is required",
                            },
                          })}
                          className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
                        />
                        <ErrorMessage message={errors.DOT?.message} />
                      </div>
                      <div className="flex flex-col items-start justify-center gap-2">
                        <label className="text-[#232323] ">
                          Permanent Address
                        </label>
                        <input
                          type="text"
                          placeholder="Address"
                          {...register("PA", {
                            required: {
                              value: true,
                              message: "Address is required",
                            },
                          })}
                          className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
                        />
                        <ErrorMessage message={errors.PA?.message} />
                      </div>

                      <div className="flex flex-col items-start justify-center gap-2">
                        <label className="text-[#232323] ">Postal Code</label>
                        <input
                          type="text"
                          placeholder="Code"
                          {...register("PC", {
                            required: {
                              value: true,
                              message: "Postal Code is required",
                            },
                          })}
                          className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
                        />
                        <ErrorMessage message={errors.PC?.message} />
                      </div>
                    </div>

                    <div className="flex flex-col gap-3">
                      <div className="flex flex-col items-start justify-center gap-2">
                        <label className="text-[#232323] ">User Name</label>
                        <input
                          type="text"
                          {...register("UN", {
                            required: {
                              value: true,
                              message: "Username is required",
                            },
                          })}
                          placeholder="Name"
                          className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
                        />
                        <ErrorMessage message={errors.UN?.message} />
                      </div>
                      <div className="flex flex-col items-start justify-center gap-2">
                        <label className="text-[#232323] ">Password</label>
                        <input
                          type="password"
                          {...register("password", {
                            required: {
                              value: true,
                              message: "Password is required",
                            },
                            minLength: {
                              value: 6,
                              message: "Password must be at least 6 characters",
                            },
                          })}
                          placeholder="password"
                          className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
                        />
                        <ErrorMessage message={errors.password?.message} />
                      </div>

                      <div className="flex flex-col items-start justify-center gap-2">
                        <label className="text-[#232323] ">
                          Present Address
                        </label>
                        <input
                          type="text"
                          placeholder="Address"
                          {...register("PresentAddress", {
                            required: {
                              value: true,
                              message: "Address is required",
                            },
                          })}
                          className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
                        />
                        <ErrorMessage
                          message={errors.PresentAddress?.message}
                        />
                      </div>
                      <div className="flex flex-col items-start justify-center gap-2">
                        <label className="text-[#232323] ">City</label>
                        <input
                          type="text"
                          {...register("City", {
                            required: {
                              value: true,
                              message: "City is required",
                            },
                          })}
                          placeholder="City"
                          className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
                        />
                        <ErrorMessage message={errors.City?.message} />
                      </div>

                      <div className="flex flex-col items-start justify-center gap-2">
                        <label className="text-[#232323] ">Country</label>

                        <select
                          {...register("Country", {
                            required: {
                              value: true,
                              message: "Country is required",
                            },
                          })}
                          className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
                        >
                          <option value="">Select a country</option>
                          {CountryData.map((country) => (
                            <option key={country.name} value={country.name}>
                              {country.name}
                            </option>
                          ))}
                        </select>
                        <ErrorMessage message={errors.Country?.message} />
                      </div>
                    </div>
                  </div>
                  <div className="flex w-full justify-end  ">
                    <button
                      type="submit"
                      className="px-10 py-3 text-white rounded-xl bg-[#1814F3]"
                    >
                      Next
                    </button>
                  </div>
                </div>
              </div>
            </form>
          )}
          {activeButton === "preferences" && (
            <form
              onSubmit={handleSubmitPref(onsubmitPref)}
              className="flex flex-col mt-10 text-sm space-y-10"
            >
              <div className="flex flex-row gap-5">
                <div className="flex flex-col gap-3">
                  <div className="text-[#232323]">Currency</div>
                  <Controller
                    name="currency"
                    control={control}
                    rules={{ required: "Select a currency" }}
                    render={({ field }) => (
                      <select
                        {...field}
                        className="text-[#718EBF] rounded-xl w-[510px] border border-[#DFEAF2] py-3 px-5"
                      >
                        <option value="">Select a currency</option>
                        {Currencies.map((currency) => (
                          <option key={currency.value} value={currency.value}>
                            {currency.label}
                          </option>
                        ))}
                      </select>
                    )}
                  />
                  <ErrorMessage message={errorsPref.currency?.message} />
                </div>
                <div className="flex flex-col gap-3">
                  <div className="text-[#232323]">Time Zone</div>

                  <Controller
                    name="timezone"
                    control={control}
                    rules={{ required: "Select a timezone" }}
                    render={({ field }) => (
                      <select
                        {...field}
                        className="text-[#718EBF] rounded-xl w-[510px] border border-[#DFEAF2] py-3 px-5"
                      >
                        <option value="">Select a timezone</option>
                        {timezones.map((time) => (
                          <option key={time.name} value={time.offset}>
                            {`(${
                              time.offset === 0
                                ? "GMT"
                                : time.offset > 0
                                ? "GMT+"
                                : "GMT-"
                            }${
                              Math.abs(time.offset) > 0
                                ? Math.abs(time.offset)
                                : ""
                            }) ${time.name}`}
                          </option>
                        ))}
                      </select>
                    )}
                  />
                  <ErrorMessage message={errorsPref.timezone?.message} />
                </div>
              </div>
              <div className="flex flex-col gap-5">
                <div className="font-semibold">Notification</div>
                <div className="flex items-center gap-4">
                  <Controller
                    name="transaction"
                    control={control}
                    render={({ field }) => (
                      <Switch
                        checked={field.value}
                        onCheckedChange={field.onChange}
                      />
                    )}
                  />
                  <div>Send or Receive digital currency</div>
                </div>
                <div className="flex items-center gap-4">
                  <Controller
                    name="merchant"
                    control={control}
                    render={({ field }) => (
                      <Switch
                        checked={field.value}
                        onCheckedChange={field.onChange}
                      />
                    )}
                  />
                  <div>Receive merchant order</div>
                </div>
                <div className="flex items-center gap-4">
                  <Controller
                    name="recommendation"
                    control={control}
                    render={({ field }) => (
                      <Switch
                        checked={field.value}
                        onCheckedChange={field.onChange}
                      />
                    )}
                  />
                  <div>Get recommendation</div>
                </div>
                <div className="flex items-center gap-4">
                  <Controller
                    name="twoFactorAuth"
                    control={control}
                    render={({ field }) => (
                      <Switch
                        checked={field.value}
                        onCheckedChange={field.onChange}
                      />
                    )}
                  />

                  <div>Enable two factor authentication</div>
                </div>
              </div>
              <div className="flex w-full justify-between mt-10 ">
                <button
                  className="px-10 py-3 text-white rounded-xl bg-[#1814F3]"
                  onClick={() => {
                    setActiveButton("edit");
                  }}
                >
                  Back
                </button>
                <button
                  type="submit"
                  className="px-10 py-3 text-white rounded-xl bg-[#1814F3]"
                >
                  Sign Up
                </button>
              </div>
            </form>
          )}
        </div>
        <div
          className={`text-center flex justify-center text-gray-500 ${
            activeButton === "edit" ? "-mt-8" : "mt-8"
          }`}
        >
          Already have an account?&nbsp;
          <Link
            href="/auth/login"
            className="hover:underline text-[#4640DE] font-medium"
          >
            Login
          </Link>
        </div>
      </div>
    </div>
  );
};

export default PageLayout;
