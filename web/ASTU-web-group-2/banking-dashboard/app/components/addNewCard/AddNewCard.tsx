import React from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useAddCreditCardMutation } from "@/lib/service/CardService";
import { useSession } from "next-auth/react";
import notify from "@/utils/notify";

// Define the Zod schema for form validation
const addCardSchema = z.object({
  cardHolder: z.string().min(1, "Card Holder is required"),
  balance: z.number().min(0, "Balance must be greater than or equal to 0"),
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

  const onSubmit = async (data: AddCardFormValues) => {
    const formattedData = {
      balance: data.balance,
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
  };

  return (
    <div className="bg-white rounded-3xl grid grid-cols-1 gap-6 p-6">
      <p className="text-[#718EBF] font-normal text-base leading-6">
        Credit Card generally means a plastic card issued by Scheduled
        Commercial Banks assigned to a Cardholder, with a credit limit, that can
        be used to purchase goods and services on credit or obtain cash
        advances.
      </p>
      <form className="grid gap-6">
        <div className="grid grid-cols-1 sm:grid-cols-2 gap-6">
          <div className="flex flex-col gap-2">
            <label className="text-base font-normal text-[#232323]">
              Card Type
            </label>
            <input
              type="text"
              className="w-full h-12 rounded-lg border border-[#E2E8F0] text-[#718EBF] pl-4"
              placeholder="Classic"
            />
          </div>

          <div className="flex flex-col gap-2">
            <label className="text-base font-normal text-[#232323]">
              Name On Card
            </label>
            <input
              type="text"
              className="w-full h-12 rounded-lg border border-[#E2E8F0] pl-4"
              placeholder="My Cards"
            />
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
            />
          </div>
          <div className="flex flex-col gap-2">
            <label className="text-base font-normal text-[#232323]">
              Expiration Date
            </label>
            <input
              type="date"
              className="w-full h-12 rounded-lg border border-[#E2E8F0] text-[#718EBF] pl-4"
            />
          </div>
        </div>
        <div className="flex justify-start">
          <button
            type="submit"
            className="w-auto h-12 rounded-lg bg-[#1814F3] text-white px-6 py-2"
          >
            Add Card
          </button>
        </div>
      </form>
    </div>
  );
};

export default AddNewCard;

