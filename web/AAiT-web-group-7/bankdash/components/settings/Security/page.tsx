"use client";
import { Switch } from "@/components/ui/switch";
import React from "react";
import { useState } from "react";
import { set, useForm } from "react-hook-form";
import ErrorMessage from "@/components/Message/ErrorMessage";
interface FormData {
  password: string;
  confirmPassword: string;
  twoFactorAuth: boolean;
}
const SecuritySettings = () => {
  const form = useForm<FormData>();
  const { register, handleSubmit, formState } = form;
  const { errors } = formState;
  const [passwordMismatch, setPasswordMismatch] = useState(false);
  const [twoFactorAuth, setTwoFactorAuth] = useState(false);

  const onsubmit = (data: FormData) => {
    const { password, confirmPassword } = data;
    console.log(data,twoFactorAuth)
    if (confirmPassword === password) {
      setPasswordMismatch(true);
    }
  };
  const handleSwitChange = (checked:boolean) => {
    setTwoFactorAuth(checked);
    

  }
  return (
    <div className="text-sm">
      <div className="flex flex-col gap-5">
        <div className="text-[#333B69] font-semibold mt-10">
          Two-factor Authentication
        </div>
        <div className="flex items-center gap-3">
          <Switch  
            checked={twoFactorAuth}
            onCheckedChange={handleSwitChange}
            />
          <div className="text-[#232323]">
            Enable or disable two factor authentication
          </div>
        </div>
      </div>
      <div className="text-[#333B69] font-semibold py-6 mt-5">
        Change Password
      </div>
      <form
        onSubmit={handleSubmit(onsubmit)}
        onChange={() => setPasswordMismatch(false)}
        className="flex flex-col gap-6"
      >
        <div className="flex flex-col items-start justify-center gap-2">
          <label>Current Password</label>
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
            className="w-[510px] border border-[#DFEAF2] rounded-xl px-4 py-3"
            placeholder="Enter your current password"
          />
          <ErrorMessage message={errors.password?.message} />
        </div>

        <div className="flex flex-col items-start justify-center gap-2">
          <label>New Password</label>
          <input
            type="password"
            {...register("confirmPassword", {
              required: {
                value: true,
                message: "Password is required",
              },
              minLength: {
                value: 6,
                message: "Password must be at least 6 characters",
              },
            })}
            className="w-[510px] border border-[#DFEAF2] rounded-xl px-4 py-3"
            placeholder="Enter your new password"
          />
          <ErrorMessage message={errors.confirmPassword?.message} />

          {!errors.confirmPassword && passwordMismatch && (
            <ErrorMessage
              message={
                "New password must be different from the existing password."
              }
            />
          )}
        </div>

        <div className="flex w-full justify-end pr-10 mt-10">
          <button
            type="submit"
            className="px-10 py-3 text-white rounded-xl bg-[#1814F3]"
          >
            Save
          </button>
        </div>
      </form>
    </div>
  );
};

export default SecuritySettings;
