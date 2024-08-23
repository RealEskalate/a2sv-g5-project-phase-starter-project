
import axios from "axios";
import React from "react";
import { useForm, SubmitHandler } from "react-hook-form";
import { getServerSession } from "next-auth";
import { options } from "@/app/api/auth/[...nextauth]/options";
import { useState } from "react";


interface FormValues {
  receiverUserName: string;
  amount: number;
  description: string;
}
interface props {
  isOpen: boolean;
  onClose: () => void;
  userName : string
  amount :number
  accessToken:string
}
const ModalTrans =  ({ isOpen, onClose , userName ,amount , accessToken }: props) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormValues>();
  const [message , setMessage] = useState<string>("");
  const onSubmit: SubmitHandler<FormValues> = async (data) => {
    const formData = JSON.stringify({ ...data, type: "transfer" });
    // change to tranService
    try {
      const response = await axios.post(
        "https://bank-dashboard-rsf1.onrender.com/transactions",
        formData,
        {
          headers: {
            Authorization: `Bearer ${accessToken}`,
            "Content-Type": "application/json",
          },
        }
      );
      console.log("Transaction successful", response.data);
      setMessage(response.data.message)
      onClose();
    } catch (error) {
      console.error("Error occurred:", error);
      if (axios.isAxiosError(error) && error.response) {
        console.log(error.response.data.message); 
        setMessage(error.response.data.message)
    } else {
      setMessage("Transaction Error. Please try again")
    }
      
    }
  };
  return (
    <form
      className="flex flex-col space-y-4 p-4 bg-white rounded-lg"
      onSubmit={handleSubmit(onSubmit)}
    >
      
      <div className="flex justify-between">
        <p className="text-base font-semibold">Quick Transfer</p>
        <button className="text-right" onClick={onClose}>
          <CloseIcon />
        </button>
      </div>
      {message && <p className="text-[#1814F3] mt-2 text-center">{message}</p>}
      <div>
        <label
          htmlFor="receiverUserName"
          className="block text-md font-medium text-gray-700"
        >
          Receiver Username
        </label>
        <input
          {...register("receiverUserName", {
            required: "Receiver username is required",
          })}
          type="text"
          placeholder="Enter Receiver Username "
          defaultValue={userName}
          className="mt-1 p-3 border block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm placeholder:text-xs"
        />
        {errors.receiverUserName && (
          <span className="text-red-600 text-sm">
            {errors.receiverUserName.message}
          </span>
        )}
      </div>

      <div>
        <label
          htmlFor="amount"
          className="block text-md font-medium text-gray-700"
        >
          Amount
        </label>
        <input
          {...register("amount", { required: "Amount is required" })}
          type="number"
          placeholder="Enter Amount "
          defaultValue={amount}
          className="mt-1 p-3 border block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm placeholder:text-xs"
        />
        {errors.amount && (
          <span className="text-red-600 text-sm">{errors.amount.message}</span>
        )}
      </div>

      <div>
        <label
          htmlFor="description"
          className="block text-md font-medium text-gray-700"
        >
          Reason
        </label>
        <input
          {...register("description", { required: "Reason is required" })}
          type="text"
          placeholder="Enter Reason "
          className="mt-1 p-3 border block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm placeholder:text-xs"
        />
        {errors.description && (
          <span className="text-red-600 text-sm">
            {errors.description.message}
          </span>
        )}
      </div>

      <button
        type="submit"
        className="w-full bg-blue-600 text-white py-2 px-4 rounded-lg shadow-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
      >
        Send
      </button>
    </form>
  );
};

function CloseIcon() {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      stroke-width="1.5"
      stroke="currentColor"
      className="size-6"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="M6 18 18 6M6 6l12 12"
      />
    </svg>
  );
}

export default ModalTrans;
