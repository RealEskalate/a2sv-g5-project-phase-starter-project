"use client";
import { Switch } from "@/components/ui/switch";
import React from "react";
import { useState } from "react";
import { useForm } from "react-hook-form";
interface FormData {
  
  password: string;
  confirmPassword: string;
}
const SecuritySettings = () => {
  const [activeButton, setActiveButton] = useState("edit");
  const form = useForm<FormData>();
  const { register, handleSubmit, formState } = form;
  const { errors } = formState;
  const [passwordMismatch, setPasswordMismatch] = useState(false);
  
  const onsubmit = async (data:FormData) => {
    
    const {password,confirmPassword} =data
    
    if (confirmPassword === password){
      setPasswordMismatch(true)
    }}
  return (
    <div className="text-sm">
      <div className="flex flex-col gap-5">
        <div className="text-[#333B69] font-semibold mt-10">
          Two-factor Authentication
        </div>
        <div className="flex items-center gap-3">
          <Switch />
          <div className="text-[#232323]">
            Enable or disable two factor authentication
          </div>
        </div>
      </div>
      <div className="text-[#333B69] font-semibold py-6 mt-5">
        Change Password
      </div>
      <form onSubmit={handleSubmit(onsubmit)} onChange={() => setPasswordMismatch(false)} className="flex flex-col gap-6">
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
                value: 8,
                message: "Password must be at least 8 characters",
              },
            })}
            className="w-[510px] border border-[#DFEAF2] rounded-xl px-4 py-3"
          />
          <p
                  className="text-red-600 flex -mt-5 text-xs font-semibold gap-1
              "
                >
                  {errors.password && (
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 24 24"
                      stroke-width="1.5"
                      stroke="currentColor"
                      className="size-4"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z"
                      />
                    </svg>
                  )}
                  {errors.password?.message}{" "}
                </p>
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
              value: 8,
              message: "Password must be at least 8 characters",
            },
          })}
          className="w-[510px] border border-[#DFEAF2] rounded-xl px-4 py-3"
        />
        <p
                className="text-red-600 flex font-semibold gap-1
            "
              >
                {errors.confirmPassword && (
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    className="size-4"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z"
                    />
                  </svg>
                )}
                {errors.confirmPassword?.message}{" "}
              </p>
        {!errors.confirmPassword && passwordMismatch && (
                <div className="text-red-600 flex -mt-5 font-semibold gap-1
                ">{<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" className="size-6">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z" />
                </svg>
                }
                   New password must be different from the existing password.
                </div>
              )}
        </div>

        <div className="flex w-full justify-end pr-10 mt-10">
          <button type="submit" className="px-10 py-3 text-white rounded-xl bg-[#1814F3]">
            
            Save
          </button>
        </div>
      </form>
    </div>
  );
};

export default SecuritySettings;
