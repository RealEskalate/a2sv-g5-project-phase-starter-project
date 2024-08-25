"use client";
import InputGroup from "../Form/InputGroup";
import ToggleInput from "../Form/ToggleInput";

import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { toastError, toastSuccess } from "../Toastify/Toastify";
import { useState } from "react";
import { Loader } from "lucide-react";
import { useAppDispatch } from "@/hooks/hoooks";

import { useChangePassowordMutation } from "@/lib/redux/api/profileAPI";

const securitySchema = z
  .object({
    password: z
      .string()
      .min(6, "Current password must be at least 6 characters long"),
    newPassword: z
      .string()
      .min(6, "New password must be at least 6 characters long"),
    conNewPassword: z
      .string()
      .min(6, "Confirm password must be at least 6 characters long"),
  })
  .refine((data) => data.newPassword === data.conNewPassword, {
    message: "New passwords do not match",
    path: ["conNewPassword"], // This indicates where the error message should appear
  });

type FormData = z.infer<typeof securitySchema>;

const Titles = ({ title }: { title: string }) => {
  return <h2 className="text-17px font-semibold text-[#333b69]">{title}</h2>;
};

const Security = () => {
  const dispatch = useAppDispatch();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    resolver: zodResolver(securitySchema),
    mode: "onTouched",
  });

  const [changePassword, { isLoading, isSuccess, isError }] =
    useChangePassowordMutation();

  const onSubmit = async (data: FormData) => {
    console.log("chenge password data", data);
    const res = await changePassword({
      password: data.password,
      newPassword: data.newPassword,
    });

    console.log("res for change password", res)

    if (res?.data?.success) {
      toastSuccess("Password changed successfully");
    }

    if (!res?.data?.success) {

      toastError(res?.data?.message || "Failed to change password");
    }
  };

  return (
    <>
      <Titles title="Two-factor Authentication" />
      <ToggleInput
        label="Enable or disable two-factorauthentication"
        inputType="checkbox"
        id="email"
        registerName="sendOrReceiveDigitalCurrency"
        register={register}
        placeholder="Email"
        currentState={true}
      />

      <form action="" onSubmit={handleSubmit(onSubmit)}>
        <Titles title="Change Password" />

        <div className="w-full md:w-1/2">
          <InputGroup
            id="password"
            label="Current Password"
            inputType="password"
            registerName="password"
            register={register}
            placeholder="*********************"
          />
        </div>
        <div className="w-full md:w-1/2">
          <InputGroup
            id="newPassword"
            label="New Password"
            inputType="password"
            registerName="newPassword"
            register={register}
            placeholder="*********************"
          />
        </div>

        <div className="w-full md:w-1/2">
          <InputGroup
            id="conNewPassword"
            label="Confirm New Password"
            inputType="password"
            registerName="conNewPassword"
            register={register}
            placeholder="*********************"
          />
        </div>

        <div className="flex justify-end">
          <button
            type="submit"
            className="bg-[#1814f3] text-white px-10 py-2 rounded-lg w-full md:w-auto mt-4"
          >
            {isLoading ? <Loader className="animate-spin" /> : "Save"}
          </button>
        </div>
      </form>
    </>
  );
};

export default Security;
