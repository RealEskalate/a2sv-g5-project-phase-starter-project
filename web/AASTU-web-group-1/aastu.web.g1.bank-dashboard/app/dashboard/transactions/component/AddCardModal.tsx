'use client'
import React, { useState } from "react";
import { IoClose } from "react-icons/io5";
import { useUser } from "@/contexts/UserContext";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { cardSchema } from "@/schema";
import { postCards } from "@/lib/api";
import { toast } from "sonner";

interface Props {
  isOpen: boolean;
  onClose: () => void;
}

type FormData = z.infer<typeof cardSchema>;
export const AddCardModal = ({ isOpen, onClose }: Props) => {
    const [isSubmitting, setIsSubmitting] = useState(false);
  const { isDarkMode } = useUser();

  if (!isOpen) return null;

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    resolver: zodResolver(cardSchema),
  });

  const onSubmit = async (data: FormData) => {
    setIsSubmitting(true);
    try {
      const updatedForm = { ...data, balance: 300 };
      console.log(updatedForm);

      await postCards(updatedForm);
      toast("Card was submitted successfully");
    } catch (error) {
      console.error("Failed to submit form:", error);
      toast("Card submission failed");
    }
  };

  return (
    <div
      className={`fixed inset-0 z-50 flex items-center justify-center ${
        isDarkMode
          ? "bg-black bg-opacity-5 backdrop-blur-md"
          : "bg-black bg-opacity-50"
      }`}
    >
      <div
        className={`relative w-full max-w-lg p-8 rounded-3xl shadow-xl transition-transform transform ${
          isDarkMode
            ? "bg-black bg-opacity-100"
            : "bg-gradient-to-r from-white via-gray-100 to-white"
        }`}
      >
        <button
          onClick={onClose}
          className={`absolute top-4 right-4 text-gray-600 ${
            !isDarkMode ? "hover:text-gray-900" : "hover:text-white"
          } `}
        >
          <IoClose size={24} />
        </button>

        <div className="text-center">
          <h2 className="text-2xl font-bold tracking-wide text-transparent bg-clip-text bg-gradient-to-r from-sky-500 to-indigo-500">
            Add New Card
          </h2>
        </div>

        <form onSubmit={handleSubmit(onSubmit)}>
          <div className="md:grid md:grid-cols-2 gap-4">
            <div className="flex flex-col my-4">
              <label
                className={`${
                  isDarkMode ? "text-gray-300" : "text-[#515B6F]"
                } font-semibold`}
              >
                Card Type
              </label>
              <input
                type="text"
                placeholder="Classic"
                className={`inputField mb-2 rounded-xl py-2 px-2 border ${
                  isDarkMode ? "border-gray-600" : "border-gray-300"
                }`}
                {...register("cardType")}
              />
              {errors.cardType && (
                <p className="text-red-500 text-center mt-2">
                  {errors.cardType?.message}
                </p>
              )}
            </div>
            <div className="flex flex-col my-4">
              <label
                className={`${
                  isDarkMode ? "text-gray-300" : "text-[#515B6F]"
                } font-semibold`}
              >
                Name On Card
              </label>
              <input
                type="text"
                placeholder="John Doe"
                className={`inputField mb-2 rounded-xl py-2 px-2 border ${
                  isDarkMode ? "border-gray-600" : "border-gray-300"
                }`}
                {...register("cardHolder")}
              />
              {errors.cardHolder && (
                <p className="text-red-500 text-center mt-2">
                  {errors.cardHolder?.message}
                </p>
              )}
            </div>
          </div>
          <div className="md:grid md:grid-cols-2 gap-4">
            <div className="flex flex-col my-4">
              <label
                className={`${
                  isDarkMode ? "text-gray-300" : "text-[#515B6F]"
                } font-semibold`}
              >
                Card Number
              </label>
              <input
                type="text"
                placeholder="**** **** **** ****"
                className={`inputField mb-2 rounded-xl py-2 px-2 border ${
                  isDarkMode ? "border-gray-600" : "border-gray-300"
                }`}
                {...register("passcode")}
              />
              {errors.passcode && (
                <p className="text-red-500 text-center mt-2">
                  {errors.passcode?.message}
                </p>
              )}
            </div>
            <div className="flex flex-col my-4">
              <label
                className={`${
                  isDarkMode ? "text-gray-300" : "text-[#515B6F]"
                } font-semibold`}
              >
                Expiration Date
              </label>
              <input
                type="date"
                className={`inputField mb-2 rounded-xl py-2 px-2 border ${
                  isDarkMode ? "border-gray-600" : "border-gray-300"
                }`}
                {...register("expiryDate")}
              />
              {errors.expiryDate && (
                <p className="text-red-500 text-center mt-2">
                  {errors.expiryDate?.message}
                </p>
              )}
            </div>
          </div>
          <button
            type="submit"
            className={`bg-[#1814F3] sm:w-[100%] text-white p-2 sm:rounded-full w-full md:rounded-md  ${
              isSubmitting
                ? "bg-gradient-to-r from-blue-300 to-teal-300 cursor-not-allowed"
                : "bg-gradient-to-r from-sky-500 to-indigo-500 hover:bg-gradient-to-br focus:ring-4 focus:outline-none focus:ring-blue-300"
            }`}
          >
            {isSubmitting ? "Adding to Card..." : "  Add to Card"}
          </button>
        </form>
      </div>
    </div>
  );
};
