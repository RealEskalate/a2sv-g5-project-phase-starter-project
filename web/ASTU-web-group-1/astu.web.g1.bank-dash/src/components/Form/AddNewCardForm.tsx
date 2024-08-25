"use client";
import React, { use, useState } from "react";
import InputGroup from "./InputGroup";

import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useSession } from "next-auth/react";
import { useToast } from "../ui/use-toast";
// import { useRouter } from "next/navigation";

const cardSchema = z.object({
  cardHolder: z.string().nonempty("Card holder name is required"),
  expiryDate: z.string().nonempty("Expiry date is required"),
  passcode: z
    .string()
    .nonempty("passcode is required")
    .length(4, "Passcode must be exactly 4 digits"),
  cardType: z.string().nonempty("Card type is required"),
});

const AddNewCardForm = () => {
  const session = useSession();
  const { toast } = useToast();
  const [isLoading, setIsLoading] = useState(false);
  const { register, handleSubmit, formState, reset } = useForm({
    resolver: zodResolver(cardSchema),
    mode: "onTouched",
  });
  const { errors } = formState;
  const onSubmit = async (data: any) => {
    setIsLoading(true);
    const accessToken = session.data?.accessToken;
    const Balance = { balance: 0 };
    const cardData = { ...data, ...Balance };
    console.log(cardData);
    const response = await fetch(
      `https://astu-bank-dashboard.onrender.com/cards`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${accessToken}`,
        },
        body: JSON.stringify(cardData),
      }
    );
    const result = await response.json();
    console.log(result);
    setIsLoading(false);
    reset();
    console.log(result);
    if (result) {
      toast({
        description: "Your card has been added successfully.",
        className: "bg-green-600 text-white font-bold",
        duration: 3000,
      });
    }
  };
  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <div className="flex flex-col md:flex-row md:gap-6">
        <InputGroup
          id="cardHolder"
          label="Name On Card"
          inputType="text"
          registerName="cardHolder"
          register={register}
          placeholder="John Doe"
          errorMessage={errors.cardHolder?.message as string}
        />
        <InputGroup
          id="cardType"
          label="Card Type"
          inputType="text"
          registerName="cardType"
          register={register}
          placeholder="Classic"
          errorMessage={errors.cardType?.message as string}
        />
      </div>
      <div className="flex flex-col md:flex-row md:gap-6">
        <InputGroup
          id="passcode"
          label="Passcode"
          inputType="number"
          registerName="passcode"
          register={register}
          placeholder="0000"
          errorMessage={errors.passcode?.message as string}
        />
        <InputGroup
          id="expiryDate"
          inputType="date"
          label="Exipiration Date"
          registerName="expiryDate"
          register={register}
          placeholder="25 January 2025"
          min={
            new Date(new Date().setFullYear(new Date().getFullYear() + 1))
              .toISOString()
              .split("T")[0]
          }
          errorMessage={errors.expiryDate?.message as string}
        />
      </div>
      {isLoading ? (
        <button
          type="submit"
          className="bg-blue-steel flex justify-center items-center gap-4 text-white rounded-xl px-10 py-3 mt-5"
          disabled
        >
          <div className="h-4 w-4 animate-spin rounded-full bg-blue-steel border-white border-2 border-t-blue-steel "></div>
          Loading...
        </button>
      ) : (
        <button
          type="submit"
          className="bg-blue-bright text-white rounded-xl px-10 py-3 mt-5"
        >
          Add Card
        </button>
      )}
    </form>
  );
};

export default AddNewCardForm;
