"use client";
import React from "react";
import InputGroup from "./InputGroup";

import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
// import { useRouter } from "next/navigation";

const cardSchema = z.object({
  cardHolder: z.string().nonempty("Card holder name is required"),
  expiryDate: z.string().refine((date) => !isNaN(Date.parse(date)), {
    message: "Invalid date format",
  }),
  passcode: z.string().length(5, "Passcode must be exactly 5 digits"),
  cardType: z.string().nonempty("Card type is required"),
});

const AddNewCardForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(cardSchema),
    mode: "onTouched",
  });

  const onSubmit = async (data: any) => {
    const response = await fetch(`https://akil-backend.onrender.com/add-card`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });
  };
  return (
    <form action="" onSubmit={handleSubmit(onSubmit)}>
      <div className="flex flex-col md:flex-row md:gap-6">
        <InputGroup
          id="cardType"
          label="Card Type"
          inputType="text"
          registerName="cardType"
          register={register}
          placeholder="Classic"
        />
        <InputGroup
          id="cardHolder"
          label="Name On Card"
          inputType="text"
          registerName="cardHolder"
          register={register}
          placeholder="My Card"
        />
      </div>
      <div className="flex flex-col md:flex-row md:gap-6">
        <InputGroup
          id="balance"
          label="Balance"
          inputType="text"
          registerName="balance"
          register={register}
          placeholder="27,00$"
        />
        <InputGroup
          id="expiryDate"
          label="Exipiration Date"
          inputType="text"
          registerName="expiryDate"
          register={register}
          placeholder="25 January 2025"
        />
      </div>
      <button
        type="submit"
        className="bg-[#1814f3] text-white px-10 py-3 rounded-lg w-full md:w-auto mt-4"
      >
        Add Card
      </button>
    </form>
  );
};

export default AddNewCardForm;
