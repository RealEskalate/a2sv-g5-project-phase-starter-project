"use client";
import React, { useState } from "react";
import Image from "next/image";
import livia from "../../public/images/Livia.svg";
import { FaPen } from "react-icons/fa6";
import { EditProfileProps, FormValues } from "@/types/index.";

const EditProfile: React.FC<EditProfileProps> = ({ isActive }) => {
  const [formValues, setFormValues] = useState<FormValues>({
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
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormValues((prevValues) => ({
      ...prevValues,
      [name]: value,
    }));
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log(formValues);
  };

  return (
    <div className="flex items-center justify-center min-h-screen">
      <form
        onSubmit={handleSubmit}
        className="relative rounded-3xl flex flex-col lg:flex-row justify-center gap-4 pt-12 w-full lg:w-[1010px] mx-auto bg-white text-slate-700"
      >
        <div className="relative w-24 h-24 lg:w-32 lg:h-32 rounded-full bg-slate-500 lg:mt-6 mb-8 lg:mb-0 mr-4 lg:mr-6">
          <Image
            src={livia}
            alt="Profile Picture"
            layout="fill"
            className="rounded-full"
          />
          <div className="absolute w-6 h-6 lg:w-8 lg:h-8 bg-[#1814f3] rounded-full border-2 border-white bottom-[-5px] right-1 flex items-center justify-center">
            <FaPen className="text-white w-2 h-2 lg:w-4 lg:h-4" />
          </div>
        </div>

        <div className="flex flex-col items-start lg:w-1/3 h-auto lg:h-auto">
          <label htmlFor="name">Your Name</label>
          <input
            placeholder="Your Name"
            type="text"
            id="name"
            name="name"
            value={formValues.name}
            onChange={handleChange}
            className="w-full p-3 mt-3 border mb-5 rounded-2xl"
          />

          <label htmlFor="email">Email</label>
          <input
            placeholder="Email"
            type="email"
            id="email"
            name="email"
            value={formValues.email}
            onChange={handleChange}
            className="w-full p-3 mt-3 border mb-5 rounded-2xl"
          />

          <label htmlFor="dateOfBirth">Date of Birth</label>
          <input
            placeholder="Date of Birth"
            type="date"
            id="dateOfBirth"
            name="dateOfBirth"
            value={formValues.dateOfBirth}
            onChange={handleChange}
            className="w-full p-3 mt-3 border mb-5 rounded-2xl"
          />

          <label htmlFor="permanentAddress">Permanent Address</label>
          <input
            placeholder="Permanent Address"
            type="text"
            id="permanentAddress"
            name="permanentAddress"
            value={formValues.permanentAddress}
            onChange={handleChange}
            className="w-full p-3 mt-3 border mb-5 rounded-2xl"
          />

          <label htmlFor="postalCode">Postal Code</label>
          <input
            placeholder="Postal Code"
            type="text"
            id="postalCode"
            name="postalCode"
            value={formValues.postalCode}
            onChange={handleChange}
            className="w-full p-3 mt-3 border mb-5 rounded-2xl"
          />
        </div>

        <div className="flex flex-col items-start lg:w-1/3 h-auto lg:h-auto">
          <label htmlFor="username">User Name</label>
          <input
            placeholder="User Name"
            type="text"
            id="username"
            name="username"
            value={formValues.username}
            onChange={handleChange}
            className="w-full p-3 mt-3 border mb-5 rounded-2xl"
          />

          <label htmlFor="password">Password</label>
          <input
            placeholder="Password"
            type="password"
            id="password"
            name="password"
            value={formValues.password}
            onChange={handleChange}
            className="w-full  p-3 mt-3 border mb-5 rounded-2xl"
          />

          <label htmlFor="presentAddress">Present Address</label>
          <input
            placeholder="Present Address"
            type="text"
            id="presentAddress"
            name="presentAddress"
            value={formValues.presentAddress}
            onChange={handleChange}
            className="w-full p-3 mt-3 border mb-5 rounded-2xl"
          />

          <label htmlFor="city">City</label>
          <input
            placeholder="City"
            type="text"
            id="city"
            name="city"
            value={formValues.city}
            onChange={handleChange}
            className="w-full p-3 mt-3 border mb-5 rounded-2xl"
          />

          <label htmlFor="country">Country</label>
          <input
            placeholder="Country"
            type="text"
            id="country"
            name="country"
            value={formValues.country}
            onChange={handleChange}
            className="w-full p-3 mt-3 border mb-5 rounded-2xl"
          />

          <button
            type="submit"
            className="w-full lg:w-52 h-14 rounded-xl bg-[#1814f3] lg:ml-auto mt-12 text-white"
          >
            Save
          </button>
        </div>
      </form>
    </div>
  );
};

export default EditProfile;
