"use client";

import React from "react";
import { useForm, SubmitHandler } from "react-hook-form";
import { useCreateTransactionMutation } from "@/lib/redux/api/transactionsApi";
import { toast, ToastContainer } from "react-toastify";
import 'react-toastify/dist/ReactToastify.css';

interface TransferFormInputs {
  type: string;
  description: string;
  amount: number;
  receiverUserName: string;
}

const TransferPage: React.FC = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<TransferFormInputs>();

  const [createTransaction] = useCreateTransactionMutation();

  const onSubmit: SubmitHandler<TransferFormInputs> = async (data) => {
    try {
      const response = await createTransaction(data).unwrap();
      if (response.success) {
        toast.success(`Transaction successful: ${response.message}`);
      } else {
        toast.error(`Transaction failed: ${response.message}`);
      }
    } catch (error) {
      toast.error(`Transaction failed: ${error}`);
    }
  };

  return (
    <div className="flex justify-center mt-4 h-auto">
      {/* Toast Container */}
      <ToastContainer />

      <form onSubmit={handleSubmit(onSubmit)} className="p-8 w-full max-w-md">
        <h2 className="text-2xl font-bold mb-6 text-center">Transfer Funds</h2>

        {/* Type Selection */}
        <div className="mb-4">
          <label htmlFor="type" className="block text-sm font-medium text-gray-700  dark:text-white">
            Type
          </label>
          <select
            id="type"
            {...register("type", { required: "Please select a type" })}
            className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm dark:border-white dark:bg-darkPage"
          >
            <option value="">Select a type</option>
            <option value="service">Service</option>
            <option value="shopping">Shopping</option>
            <option value="transfer">Transfer</option>
          </select>
          {errors.type && <p className="text-red-500 text-sm mt-2">{errors.type.message}</p>}
        </div>

        {/* Description */}
        <div className="mb-4">
          <label htmlFor="description" className="block text-sm font-medium text-gray-700  dark:text-white">
            Description
          </label>
          <input
            id="description"
            type="text"
            {...register("description", { required: "Description is required" })}
            className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm dark:border-white dark:bg-darkPage"
          />
          {errors.description && <p className="text-red-500 text-sm mt-2">{errors.description.message}</p>}
        </div>

        {/* Amount */}
        <div className="mb-4">
          <label htmlFor="amount" className="block text-sm font-medium text-gray-700  dark:text-white">
            Amount
          </label>
          <input
            id="amount"
            type="number"
            {...register("amount", {
              required: "Amount is required",
              valueAsNumber: true,
              min: { value: 0, message: "Amount must be a positive number" },
            })}
            className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm dark:border-white dark:bg-darkPage"
          />
          {errors.amount && <p className="text-red-500 text-sm mt-2">{errors.amount.message}</p>}
        </div>


        {/* Receiver Username */}
        <div className="mb-4">
          <label htmlFor="receiverUserName" className="block text-sm font-medium text-gray-700 dark:text-white">
            Receiver Username
          </label>
          <input
            id="receiverUserName"
            type="text"
            {...register("receiverUserName", { required: "Receiver username is required" })}
            className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm dark:border-white dark:bg-darkPage"
          />
          {errors.receiverUserName && (
            <p className="text-red-500 text-sm mt-2">{errors.receiverUserName.message}</p>
          )}
        </div>

        {/* Submit Button */}
        <div className="flex justify-end">
          <button
            type="submit"
            className="bg-indigo-600 text-white py-2 px-4 rounded-md shadow hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          >
            Submit
          </button>
        </div>
      </form>
    </div>
  );
};

export default TransferPage;
