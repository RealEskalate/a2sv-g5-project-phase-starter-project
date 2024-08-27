"use client";
import { Country } from "country-state-city";
import React from "react";
import { useForm } from "react-hook-form";
import { useState } from "react";
import ErrorMessage from "@/components/Message/ErrorMessage";
type Form = {
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
};
const EditProfile = () => {
  const form = useForm<Form>();
  const { register, handleSubmit, formState } = form;
  const { errors } = formState;
  const onSubmit = (data: Form) => {
    console.log(data);
  };
  const CountryData = Country.getAllCountries();
  const [selectedCountry, setSelectedCountry] = useState("");
  const [profileImage, setProfileImage] = useState<File | null>(null);

  const handleCountryChange = (e: any) => {
    setSelectedCountry(e.target.value);
  };
  const handleProfilePictureChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const file = event.target.files?.[0];
    console.log(file);

    if (file) {
      setProfileImage(file);
    }
  };
  return (
    <form  onSubmit={handleSubmit(onSubmit)} className="flex flex-col text-sm">
      <div className="flex  gap-8 py-10">
        <div className="relative">
          {profileImage ? (
            <img
              src={URL.createObjectURL(profileImage)}
              className="size-32 rounded-full"
            />
          ) : (
            <img src="/pubimg/placepp.png" className="size-32" />
          )}
          <div className="bg-[#1814F3] h-8 w-8 flex justify-center items-center rounded-full absolute right-0 top-20 hover:brightness-200 transition duration-200">
            <label htmlFor="fileInput" className="cursor-pointer">
              <img src="/pubimg/pencil.svg" />
            </label>
            <input
              id="fileInput"
              type="file"
              accept="image/*"
              onChange={handleProfilePictureChange}
              className="hidden"
            />
          </div>
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
                      value: /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/,
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
                <label className="text-[#232323] ">Permanent Address</label>
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
                <label className="text-[#232323] ">Present Address</label>
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
                <ErrorMessage message={errors.PresentAddress?.message} />
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
                <ErrorMessage message={errors.Country?.message} />
              </div>
            </div>
          </div>

          <div className="flex w-full justify-end mt-5 px-[30px] ">
            <button
              type="submit"
              className="px-10 py-3 text-white rounded-xl bg-[#1814F3]"
            >
              Save
            </button>
          </div>
        </div>
      </div>
    </form>
  );
};

export default EditProfile;
