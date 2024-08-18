"use client";
import React, { useRef, useState } from "react";
import { useForm } from "react-hook-form";
import SwitchButton from "./switch";

interface FormInput {
  Password: string;
  newPassword: string;
  twoFactorAuthentication: boolean;
}

const Security = () => {
  const { register, handleSubmit, setValue } = useForm<FormInput>({
    defaultValues: {
      Password: "",
      newPassword: "",
      twoFactorAuthentication: false,
    },
  });

  const [tFactorAuthentication, setTFactorAuthentication] = useState(true);

  const fileInputRef = useRef<HTMLInputElement>(null);

  const onSubmit = async (data: FormInput) => {
    console.log("Form submitted:", data);
    const formData = new FormData();
    // Submit the formData object to your server or API
  };

  return (
    <div className="p-4 ">
      <div className=" flex flex-col py-3">
        <div className="font-semibold text-xl py-2 text-[#333B69] ">
          Two-factor Authentication
        </div>
        <div className="flex gap-3 w-full items-center py-2 ">
          <div className="w-1/5">
            <SwitchButton
              isOn={tFactorAuthentication}
              onToggle={setTFactorAuthentication}
            />
          </div>
          Enable or disable two factor authentication
        </div>
      </div>

      <div className="items-end justify-between w-full font-semibold">
        <div className="font-semibold text-xl p-3 text-[#333B69] ">
          Change Password
        </div>

        <div className="mb-3 w-full md:w-[45%]">
          <label className="block text-black text-sm mb-2">Password</label>
          <input
            className="w-full p-3 text-[#718EBF] border-2 text-sm border-[#DFEAF2] rounded-lg focus:outline-none"
            type="password"
            id="password"
            placeholder="*********"
            {...register("Password", {
              required: {
                value: true,
                message: `password is required`,
              },
            })}
          />
        </div>
        <div className="mb-3 w-full md:w-[45%]">
          <label className="block text-black text-sm mb-2">New Password</label>
          <input
            className="w-full p-3 text-[#718EBF] border-2 text-sm border-[#DFEAF2] rounded-lg focus:outline-none"
            type="password"
            id="newPassword"
            placeholder="*********"
            {...register("newPassword", {
              required: {
                value: true,
                message: `newPassword is required`,
              },
            })}
          />
        </div>
      </div>
      <div className="flex justify-end w-full my-6">
        <button
          className=" w-full md:w-1/5 bg-[#1814F3] text-white font-semibold py-2 px-4 rounded-lg focus:outline-none"
          type="submit"
          onClick={handleSubmit(onSubmit)}
        >
          Save
        </button>
      </div>
    </div>
  );
};

export default Security;
