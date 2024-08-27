"use client";
import React from "react";
import { useState } from "react";
import Link from "next/link";
import { useForm } from "react-hook-form";
import { useEffect } from "react";
import ErrorMessage from "@/components/Message/ErrorMessage";
import { Switch } from "@/components/ui/switch";
import { Currencies } from "@/components/constants/currency";
import { timezones } from "@/components/constants/timezones";
import { CountryData } from "@/components/constants/countries";
import { useUserRegistrationMutation } from "@/redux/api/authentication-controller";

interface FormData {
  timezone: string;
  currency: string;
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
  transaction: boolean;
  merchant: boolean;
  recommendation: boolean;
  profilePicture: File;
}

const PageLayout = () => {
  const [registerUser] = useUserRegistrationMutation();
  const [activeButton, setActiveButton] = useState("edit");
  const [profileImage, setProfileImage] = useState<File | null>(null);
  

  const form = useForm<FormData>();
  const { register, handleSubmit, formState, setValue } = form;
  const { errors } = formState;
  const [filledElements, setFilledElements] = useState(0);

  const handleSwitchChange =
    (checkParameter: any) => (e: React.FormEvent<HTMLButtonElement>) => {
      const inputElement = e.currentTarget.querySelector("input");
      if (inputElement) {
        const isChecked = inputElement.checked;
        setValue(checkParameter, isChecked);
      }
    };

  const onSubmit = async (formData: FormData) => {
    console.log("formData", formData);

    try {
      const { data, error } = await registerUser({
        name: formData.Name,
        email: formData.Email,
        dateOfBirth: formData.DOT,
        password: formData.password,
        username: formData.UN,
        permanentAddress: formData.PA,
        postalCode: formData.PC,
        presentAddress: formData.PresentAddress,
        city: formData.City,
        country: formData.Country,
        profilePicture: formData.profilePicture,
        currency: formData.currency,
        sentOrReceiveDigitalCurrency: formData.transaction,
        receiveMerchantOrder: formData.merchant,
        accountRecommendations: formData.recommendation,
        timeZone: formData.timezone,
        twoFactorAuthentication: formData.recommendation,
      }).unwrap();

      console.log("response from server upon registration", data);
    } catch (error) {
      console.log("error from server upon registration", error);
    }
  };

  const handleProfilePictureChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const file = event.target.files?.[0];
    console.log(file);

    if (file) {
      setProfileImage(file);
      setValue("profilePicture", file);

    }
  };
  useEffect(() => {
    if (filledElements === 10) {
      console.log("No errors!");
      setActiveButton(() => {
        if (filledElements === 10) {
          return "preferences";
        } else {
          return "edit";
        }
      });
    } else {
      console.log("Errors:", errors);
      setActiveButton("edit");
    }
  }, [filledElements, errors, setActiveButton]);

  const handleClick = () => {
    setFilledElements(() => (10 - Object.keys(errors).length));
  };

  return (
    <div className="px-5 py-5 flex justify-center">
      <div className="flex flex-col rounded-3xl w-fill px-10  bg-white ">
        <h1 className="text-4xl font-poppins font-bold text-center text-[#4640DE] py-5">
          Sign Up Today!
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
        <form onSubmit={handleSubmit(onSubmit)}>
          {activeButton === "edit" && (
            <div className="flex flex-col text-sm">
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
                  <ErrorMessage message={(errors.profilePicture?.message) as string} />
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
                      onClick={handleClick}
                      className="px-10 py-3 text-white rounded-xl bg-[#1814F3]"
                    >
                      Next
                    </button>
                  </div>
                </div>
              </div>
            </div>
          )}
          {activeButton === "preferences" && (
            <div className="flex flex-col mt-10 text-sm space-y-10">
              <div className="flex flex-row gap-5">
                <div className="flex flex-col gap-3">
                  <div className="text-[#232323]">Currency</div>
                  <select
                    className="text-[#718EBF] rounded-xl w-[510px]  border border-[#DFEAF2] py-3 px-5"
                    {...register("currency", {
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
                  <ErrorMessage message={errors.currency?.message} />
                </div>
                <div className="flex flex-col gap-3">
                  <div className="text-[#232323]">Time Zone</div>

                  <select
                    className="text-[#718EBF] rounded-xl w-[510px]  border border-[#DFEAF2] py-3 px-5"
                    {...register("timezone", {
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
                  <ErrorMessage message={errors.timezone?.message} />
                </div>
              </div>
              <div className="flex flex-col gap-5">
                <div className="font-semibold">Notification</div>
                <div className="flex items-center gap-4">
                  <Switch {...register("transaction")} />
                  <div>I send or receive digita currency</div>
                </div>
                <div className="flex items-center gap-4">
                  <Switch {...register("merchant")} />
                  <div>I receive merchant order</div>
                </div>
                <div className="flex items-center gap-4">
                  <Switch {...register("recommendation")} />
                  <div>There are recommendation for my account</div>
                </div>
                <div className="flex items-center gap-4">
                  <Switch
                    {...register("recommendation")}
                    onChange={handleSwitchChange("recommendation")}
                  />

                  <div>Enable or disable two factor authentication</div>
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
            </div>
          )}
        </form>
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
