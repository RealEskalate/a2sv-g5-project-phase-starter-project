"use client";
import React from "react";
import { useForm } from "react-hook-form";

interface EditProfileProps {
  isActive: boolean;
}

type FormValues = {
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
};

const EditProfile: React.FC<EditProfileProps> = ({ isActive }) => {
  const form = useForm<FormValues>();
  const { register, handleSubmit, formState } = form;

  const { errors } = formState;

  const onSubmit = (data: FormValues) => {
    console.log(data);
  };

  return (
    <form
      onSubmit={handleSubmit(onSubmit)}
      className="relative flex justify-center gap-3 pt-10 w-[1110px] bg-white text-slate-700"
    >
      <div className="relative w-[130px] h-[132px] rounded-full bg-slate-500 mr-7 mt-8">
        <img src="" alt="" className="" />
        <div className="absolute w-[40px] h-[40px] bg-[#1814f3] rounded-full border-2 border-white bottom-[-5px] right-[-1px]"></div>
      </div>

      <div className="flex flex-col items-start w-[418px] h-[600px] mt-8">
        <label htmlFor="name">Your Name</label>
        <input
          placeholder="Your Name"
          type="text"
          id="name"
          {...register("name", { required: true })}
          className="w-full p-2 mt-2 border mb-4 rounded-md"
        />
        {errors.name && (
          <span className="text-red-500">This field is required</span>
        )}

        <label htmlFor="email">Email</label>
        <input
          placeholder="Email"
          type="email"
          id="email"
          {...register("email", { required: true })}
          className="w-full p-2 mt-2 border mb-4 rounded-md"
        />
        {errors.email && (
          <span className="text-red-500">This field is required</span>
        )}

        <label htmlFor="dateOfBirth">Date of Birth</label>
        <input
          placeholder="Date of Birth"
          type="date"
          id="dateOfBirth"
          {...register("dateOfBirth", { required: true })}
          className="w-full p-2 mt-2 border mb-4 rounded-md"
        />
        {errors.dateOfBirth && (
          <span className="text-red-500">This field is required</span>
        )}

        <label htmlFor="permanentAddress">Permanent Address</label>
        <input
          placeholder="Permanent Address"
          type="text"
          id="permanentAddress"
          {...register("permanentAddress", { required: true })}
          className="w-full p-2 mt-2 border mb-4 rounded-md"
        />
        {errors.permanentAddress && (
          <span className="text-red-500">This field is required</span>
        )}

        <label htmlFor="postalCode">Postal Code</label>
        <input
          placeholder="Postal Code"
          type="text"
          id="postalCode"
          {...register("postalCode", { required: true })}
          className="w-full p-2 mt-2 border mb-4 rounded-md"
        />
        {errors.postalCode && (
          <span className="text-red-500">This field is required</span>
        )}
      </div>

      <div className="flex flex-col items-start  w-[418px] h-[600px] mt-8">
        <label htmlFor="username">User Name</label>
        <input
          placeholder="User Name"
          type="text"
          id="username"
          {...register("username", { required: true })}
          className="w-full p-2 mt-2 border mb-4 rounded-md"
        />
        {errors.username && (
          <span className="text-red-500">This field is required</span>
        )}

        <label htmlFor="password">Password</label>
        <input
          placeholder="Password"
          type="password"
          id="password"
          {...register("password", { required: true })}
          className="w-full p-2 mt-2 border mb-4 rounded-md"
        />
        {errors.password && (
          <span className="text-red-500">This field is required</span>
        )}

        <label htmlFor="presentAddress">Present Address</label>
        <input
          placeholder="Present Address"
          type="text"
          id="presentAddress"
          {...register("presentAddress", { required: true })}
          className="w-full p-2 mt-2 border mb-4 rounded-md"
        />
        {errors.presentAddress && (
          <span className="text-red-500">This field is required</span>
        )}

        <label htmlFor="city">City</label>
        <input
          placeholder="City"
          type="text"
          id="city"
          {...register("city", { required: true })}
          className="w-full p-2 mt-2 border mb-4 rounded-md"
        />
        {errors.city && (
          <span className="text-red-500">This field is required</span>
        )}

        <label htmlFor="country">Country</label>
        <input
          placeholder="Country"
          type="text"
          id="country"
          {...register("country", { required: true })}
          className="w-full p-2 mt-2 border mb-4 rounded-md"
        />
        {errors.country && (
          <span className="text-red-500">This field is required</span>
        )}

        <button
          type="submit"
          className="w-[190px] h-[50px] rounded-[15px] bg-[#1814f3] ml-auto mt-10 text-white"
        >
          Save
        </button>
      </div>
    </form>
  );
};

export default EditProfile;
