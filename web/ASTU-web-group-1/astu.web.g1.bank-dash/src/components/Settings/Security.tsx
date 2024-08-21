'use client'
import InputGroup from "../Form/InputGroup";
import ToggleInput from "../Form/ToggleInput";

import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";

const securitySchema = z.object({
  sendOrReceiveDigitalCurrency: z.boolean(),
  password: z
    .string()
    .min(6, "Current password must be at least 6 characters long"),
  newPassword: z
    .string()
    .min(6, "New password must be at least 6 characters long"),
});

const Titles = ({ title }: { title: string }) => {
  return <h2 className="text-17px font-semibold text-[#333b69]">{title}</h2>;
};

const Security = () => {

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(securitySchema),
    mode: "onTouched",
  });

  const onSubmit = (data: any) => {
    console.log("Form Data:", data);
  };

  return (
    <form action="" onSubmit={handleSubmit(onSubmit)}>
      <Titles title="Two-factor Authentication" />

      <ToggleInput
        label="I send or receive digital currency"
        inputType="checkbox"
        id="email"
        registerName="sendOrReceiveDigitalCurrency"
        register={register}
        placeholder="Email"
        currentState={true}
      />

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

      <div className="flex justify-end">
        <button
          type="submit"
          className="bg-[#1814f3] text-white px-10 py-2 rounded-lg w-full md:w-auto mt-4"
        >
          Submit
        </button>
      </div>
    </form>
  );
};

export default Security;
