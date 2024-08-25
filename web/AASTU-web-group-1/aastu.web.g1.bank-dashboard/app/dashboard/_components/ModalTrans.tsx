import { addTransactions } from "@/lib/api";
import React, { useState } from "react";
import { useForm, SubmitHandler } from "react-hook-form";
import { IoClose } from "react-icons/io5";
import { useUser } from "@/contexts/UserContext";
import { toast } from "@/components/ui/use-toast";

interface FormValues {
  receiverUserName: string;
  amount: number;
  description: string;
}

interface Props {
  isOpen: boolean;
  onClose: () => void;
  reciverUserName: string;
}

export const ModalTrans = ({ isOpen, onClose, reciverUserName }: Props) => {
  const { isDarkMode } = useUser();
  const [isSubmitting, setIsSubmitting] = useState(false);
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormValues>({
    defaultValues: {
      receiverUserName: reciverUserName, // Set the default value
    },
  });

  const onSubmit: SubmitHandler<FormValues> = async (data) => {
    setIsSubmitting(true);
    try {
      const success = await addTransactions({
        type: "transfer",
        description: data.description,
        amount: data.amount,
        receiverUserName: data.receiverUserName,
      });

      if (success) {
        toast({
          title: "Success",
          description: "Transaction Success",
          variant: "success",
        });
        onClose();
      } else {
        toast({
          title: "Error",
          description: "Transaction failed!",
          variant: "destructive",
        });
      }
    } catch (error) {
      console.error("Error submitting form:", error);
      toast({
        title: "Error",
        description: "An unexpected error occurred.!",
        variant: "destructive",
      });
    } finally {
      setIsSubmitting(false);
    }
  };

  if (!isOpen) return null;

  return (
    <div
      className={`fixed inset-0 z-50 flex items-center justify-center  backdrop-blur-sm ${
        isDarkMode ? "bg-black  bg-opacity-5" : "bg-black bg-opacity-50"
      } `}
    >
      <div
        className={`relative w-full max-w-lg p-8 rounded-3xl shadow-xl transition-transform transform ${
          isDarkMode
            ? "bg-black bg-opacity-100"
            : "bg-gradient-to-r from-white via-gray-100 to-white"
        } ${isSubmitting ? "scale-95" : "scale-100"}`}
      >
        <button
          type="button"
          className={`absolute top-3 right-3 p-2 rounded-full ${
            isDarkMode
              ? "text-gray-300 hover:text-white bg-gray-700 hover:bg-gray-600"
              : "text-gray-400 hover:text-gray-600 bg-gray-200 hover:bg-gray-300"
          }`}
          onClick={onClose}
        >
          <IoClose size={24} />
        </button>

        <form onSubmit={handleSubmit(onSubmit)} className="space-y-8">
          <div className="text-center">
            <h2 className="text-2xl font-bold tracking-wide text-transparent bg-clip-text bg-gradient-to-r from-sky-500 to-indigo-500">
              Send Money
            </h2>
          </div>
          <div>
            <label
              htmlFor="receiverUserName"
              className={`block text-md font-medium  ${
                isDarkMode ? "text-gray-300" : "text-gray-700"
              }`}
            >
              Receiver Username
            </label>
            <input
              {...register("receiverUserName", {
                required: "Receiver username is required",
              })}
              type="text"
              placeholder="Enter Receiver Username Here"
              className={`mt-1 p-3 border block w-full rounded-lg shadow-sm sm:text-sm ${
                isDarkMode
                  ? "bg-gray-800 border-gray-700 focus:border-blue-500 focus:ring-blue-500 text-gray-300"
                  : "bg-white border-gray-300 focus:border-blue-500 focus:ring-blue-500"
              }`}
            />

            {errors.receiverUserName && (
              <span className="text-red-600 text-sm">
                {errors.receiverUserName.message}
              </span>
            )}
          </div>
          <div className="space-y-6">
            <div>
              <label
                htmlFor="amount"
                className={`block text-md font-medium ${
                  isDarkMode ? "text-gray-300" : "text-gray-700"
                }`}
              >
                Amount
              </label>
              <input
                {...register("amount", { required: "Amount is required" })}
                type="number"
                placeholder="Enter Amount"
                className={`mt-1 p-3 border block w-full rounded-lg shadow-sm sm:text-sm ${
                  isDarkMode
                    ? "bg-gray-800 border-gray-700 focus:border-blue-500 focus:ring-blue-500 text-gray-300"
                    : "bg-white border-gray-300 focus:border-blue-500 focus:ring-blue-500"
                }`}
              />
              {errors.amount && (
                <span className="text-red-500 text-sm">
                  {errors.amount.message}
                </span>
              )}
            </div>

            <div>
              <label
                htmlFor="description"
                className={`block text-md font-medium ${
                  isDarkMode ? "text-gray-300" : "text-gray-700"
                }`}
              >
                Reason
              </label>
              <input
                {...register("description", { required: "Reason is required" })}
                type="text"
                placeholder="Enter Reason"
                className={`mt-1 p-3 border block w-full rounded-lg shadow-sm sm:text-sm ${
                  isDarkMode
                    ? "bg-gray-800 border-gray-700 focus:border-blue-500 focus:ring-blue-500 text-gray-300"
                    : "bg-white border-gray-300 focus:border-blue-500 focus:ring-blue-500"
                }`}
              />
              {errors.description && (
                <span className="text-red-500 text-sm">
                  {errors.description.message}
                </span>
              )}
            </div>
          </div>

          <button
            type="submit"
            disabled={isSubmitting}
            className={`w-full py-3 px-4 rounded-lg shadow-lg font-semibold text-white focus:outline-none focus:ring-2 transition-all duration-300 ${
              isSubmitting
                ? "bg-gradient-to-r from-blue-300 to-teal-300 cursor-not-allowed"
                : "bg-gradient-to-r from-sky-500 to-indigo-500 hover:bg-gradient-to-br focus:ring-4 focus:outline-none focus:ring-blue-300"
            }`}
          >
            {isSubmitting ? "Sending..." : "Send"}
          </button>
        </form>
      </div>
    </div>
  );
};
