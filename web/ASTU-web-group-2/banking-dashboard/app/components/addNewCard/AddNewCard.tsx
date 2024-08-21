import React from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useAddCreditCardMutation } from "@/lib/service/CardService";
import { useSession } from "next-auth/react";
import { toast, ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

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
      // Display success toast
      toast.success("Card successfully added!");
      // Store the card ID in local storage
      console.log("Card ID stored in local storage:", res.data.id);
    } else {
      // Handle failure case
      toast.error("Failed to add card.");
    }
  };

  return (
    <div className="w-[730px] h-[440px] ml-290px mt-849px bg-white rounded-[25px]">
      {/* ToastContainer to display toast messages */}
      <ToastContainer />
      <p className="pl-[30px] pt-[27px] text-[#718EBF] font-normal text-[16px] leading-[26px]">
        Credit Card generally means a plastic card issued by Scheduled
        Commercial Banks assigned to a Cardholder, with a credit limit, that can
        be used to purchase goods and services on credit or obtain cash
        advances.
      </p>
      <div className="pl-[30px] pt-[29px]">
        <form className="contact-form" onSubmit={handleSubmit(onSubmit)}>
          <div className="flex gap-[30px] pb-[22px]">
            <div className="flex flex-col gap-[11px]">
              <label className="text-[16px] leading-[19.26px] font-normal text-[#232323]">
                Card Holder
              </label>
              <input
                type="text"
                className="w-[320px] h-[50px] rounded-[15px] border-[1px] pl-[20px]"
                placeholder="John Doe"
                {...register("cardHolder")}
              />
              {errors.cardHolder && (
                <p className="text-red-500 text-sm">
                  {errors.cardHolder.message}
                </p>
              )}
            </div>

            <div className="flex flex-col gap-[11px]">
              <label className="text-[16px] leading-[19.26px] font-normal text-[#232323]">
                Card Type
              </label>
              <input
                type="text"
                className="w-[320px] h-[50px] rounded-[15px] border-[1px] pl-[20px]"
                placeholder="Classic"
                {...register("cardType")}
              />
              {errors.cardType && (
                <p className="text-red-500 text-sm">
                  {errors.cardType.message}
                </p>
              )}
            </div>
          </div>

          <div className="flex gap-[30px] pb-[30px]">
            <div className="flex flex-col gap-[11px]">
              <label className="text-[16px] leading-[19.26px] font-normal text-[#232323]">
                Balance
              </label>
              <input
                type="number"
                className="w-[320px] h-[50px] rounded-[15px] border-[1px] pl-[20px]"
                placeholder="27000"
                {...register("balance", { valueAsNumber: true })}
              />
              {errors.balance && (
                <p className="text-red-500 text-sm">{errors.balance.message}</p>
              )}
            </div>

            <div className="flex flex-col gap-[11px]">
              <label
                className="text-[16px] leading-[19.26px] font-normal text-[#232323]"
                htmlFor=""
              >
                Expiration Date
              </label>
              <input
                type="date"
                className="w-[320px] h-[50px] rounded-[15px] border-[1px] text-[#718EBF] pl-[20px]"
                {...register("expiryDate")}
              />
              {errors.expiryDate && (
                <p className="text-red-500 text-sm">
                  {errors.expiryDate.message}
                </p>
              )}
            </div>
          </div>

          <button
            type="submit"
            className="w-[160px] h-[50px] rounded-[9px] bg-[#1814F3] text-[#ffffff]"
          >
            Add Card
          </button>
        </form>
      </div>
    </div>
  );
};

export default AddNewCard;
