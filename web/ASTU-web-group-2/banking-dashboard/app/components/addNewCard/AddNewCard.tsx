import React, { useState } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useAddCreditCardMutation } from "@/lib/service/CardService";
import { useSession } from "next-auth/react";
import notify from "@/utils/notify";

// Define the Zod schema for form validation
const addCardSchema = z.object({
  cardHolder: z.string().min(1, "Card Holder is required"),
  balance: z.coerce.number().min(0, "Balance must be greater than or equal to 0"),
  expiryDate: z.string().min(1, "Expiration Date is required"),
  cardType: z.string().min(1, "Card Type is required"),
});

type AddCardFormValues = z.infer<typeof addCardSchema>;

const AddNewCard = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<AddCardFormValues>({
    resolver: zodResolver(addCardSchema),
  });
  const [addNewCard] = useAddCreditCardMutation();
  const session = useSession();
  const [loading, setLoading] = useState(false)

  const onSubmit = async (data: AddCardFormValues) => {
    setLoading(true)
    const formattedData = {
      balance: Number(data.balance),
      cardHolder: data.cardHolder,
      expiryDate: new Date(data.expiryDate).toISOString(),
      cardType: data.cardType,
    };
    const accessToken = session.data?.user.accessToken || "";
    const res = await addNewCard({
      ...formattedData,
      passcode: "56789",
      accessToken: accessToken,
    });

    if (res && res.data && res.data.id) {
      notify.success("Card successfully added!");
      // Store the card ID in local storage
    } else {
      // Handle failure case
      notify.error("Failed to add card.");
    }
    setLoading(false)
  };

  return (
    <div className="bg-white rounded-3xl grid grid-cols-1 gap-6 p-6">
      <p className="text-[#718EBF] font-normal text-base leading-6">
        Credit Card generally means a plastic card issued by Scheduled
        Commercial Banks assigned to a Cardholder, with a credit limit, that can
        be used to purchase goods and services on credit or obtain cash
        advances.
      </p>
      <form className="grid gap-6" onSubmit={handleSubmit(onSubmit)}>
        <div className="grid grid-cols-1 sm:grid-cols-2 gap-6">
          <div className="flex flex-col gap-2">
            <label className="text-base font-normal text-[#232323]">
              Card Type
            </label>
            <input
              type="text"
              className="w-full h-12 rounded-lg border border-[#E2E8F0] text-[#718EBF] pl-4"
              placeholder="Classic"
              {...register("cardType")}
            />
            {errors.cardType && (
              <span className="text-red-500">{errors.cardType.message}</span>
            )}
          </div>

          <div className="flex flex-col gap-2">
            <label className="text-base font-normal text-[#232323]">
              Name On Card
            </label>
            <input
              type="text"
              className="w-full h-12 rounded-lg border border-[#E2E8F0] pl-4"
              placeholder="My Cards"
              {...register("cardHolder")}
            />
            {errors.cardHolder && (
              <span className="text-red-500">{errors.cardHolder.message}</span>
            )}
          </div>
        </div>

        <div className="grid grid-cols-1 sm:grid-cols-2 gap-6">
          <div className="flex flex-col gap-2">
            <label className="text-base font-normal text-[#232323]">
              Balance
            </label>
            <input
              type="number"
              className="w-full h-12 rounded-lg border border-[#E2E8F0] pl-4"
              placeholder="27,000$"
              {...register("balance")}
            />
            {errors.balance && (
              <span className="text-red-500">{errors.balance.message}</span>
            )}
          </div>
          <div className="flex flex-col gap-2">
            <label className="text-base font-normal text-[#232323]">
              Expiration Date
            </label>
            <input
              type="date"
              className="w-full h-12 rounded-lg border border-[#E2E8F0] text-[#718EBF] pl-4"
              {...register("expiryDate")}
            />
            {errors.expiryDate && (
              <span className="text-red-500">{errors.expiryDate.message}</span>
            )}
          </div>
        </div>
        <div className="flex justify-start">
          {/* <button
            type="submit"
            className="w-auto h-12 rounded-lg bg-[#1814F3] text-white px-6 py-2 " disabled={loading}
          >
            Add Card
          </button> */}
          <button
            type="submit"
            className={`w-auto h-12 rounded-lg bg-[#1814F3] text-white px-6 py-2 `}
          >
            {loading ? (
              <div className="w-8 h-8 border-4 border-dashed rounded-full animate-spin [animation-duration:3s]  border-white mx-auto"></div>
            ) : (
              "Add Card"
            )}
          </button>
        </div>
      </form>
    </div>
  );
};

export default AddNewCard;
