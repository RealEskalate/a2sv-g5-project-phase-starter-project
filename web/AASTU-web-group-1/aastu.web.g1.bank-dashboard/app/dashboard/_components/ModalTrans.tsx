import { addTransactions } from "@/lib/api";
import React from "react";
import { useForm, SubmitHandler } from "react-hook-form";
import { toast } from "sonner";


interface FormValues {
  receiverUserName: string;
  amount: number;
  description: string;
}

interface Props {
  isOpen: boolean;
  onClose: () => void;
}

export const ModalTrans = ({ isOpen, onClose }: Props) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormValues>();

  const onSubmit: SubmitHandler<FormValues> = async (data) => {
    try {
      const success = await addTransactions({
        type: "transfer",
        description: data.description,
        amount: data.amount,
        receiverUserName: data.receiverUserName,
      });

      if (success) {
        toast("Transaction successful!");
        onClose(); 
      } else {
        toast("Transaction failed!");
      }
    } catch (error) {
      console.error("Error submitting form:", error);
      // Handle error (e.g., show a generic error message)
      alert("An unexpected error occurred.");
    }
  };

  return (
    <form
      className="flex flex-col space-y-4 p-4 bg-white rounded-lg shadow-md"
      onSubmit={handleSubmit(onSubmit)}
    >
      <div className="flex justify-between">
        <p className="text-base font-semibold">Send Money</p>
        <button type="button" className="text-right" onClick={onClose}>
          {/* Replace CloseIcon with your close icon component */}
          Close
        </button>
      </div>
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
          placeholder="Enter Receiver Username Here"
          className="mt-1 p-3 border block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
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
          placeholder="Enter Amount Here"
          className="mt-1 p-3 border block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
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
          placeholder="Enter Reason Here"
          className="mt-1 p-3 border block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
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
