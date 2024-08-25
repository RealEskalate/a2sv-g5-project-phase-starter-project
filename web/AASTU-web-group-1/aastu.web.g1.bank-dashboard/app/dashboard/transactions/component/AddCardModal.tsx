"use client";
import React, { useEffect, useState } from "react";
import { IoClose } from "react-icons/io5";
import { useUser } from "@/contexts/UserContext";
import { FieldErrors, SubmitHandler, useForm } from "react-hook-form";
import { AddCard } from "./utils";
import { toast } from "@/components/ui/use-toast";

interface Props {
  isOpen: boolean;
  onClose: () => void;
}
export interface FormValues {
    cardType:string;
    cardHolder:string;
    passcode:string;
    expiryDate:string;
}



export const AddCardModal = ({ isOpen, onClose }: Props) => {
  const [isSubmitting, setIsSubmitting] = useState(false);
  const { isDarkMode } = useUser();

  const {register,handleSubmit,formState} =useForm<FormValues>({defaultValues:{cardHolder:""} 
  });
  const {errors} = formState;
  
  const onSubmit:SubmitHandler<FormValues> = async (data) => {
    setIsSubmitting(true);
    try{
      const success = await AddCard(data)
      if(success){
        toast({
          title:"Success",
          description:"Card Added Successfully",
          variant:"success"});
          onClose();}
          else{
            toast({
              title:"Error",
              description:"Failed to add card",
              variant:"destructive"
            });
          }
        }catch(error){
            console.error("Error submitting form:",error);
            toast({
              title:"Error",
              description:"An unexpected error occurred.",
              variant:"destructive"
            });
          }
          finally{setIsSubmitting(false);}}
  if (!isOpen) return null;
  return (
    <div
      className={`fixed inset-0 z-50 flex items-center justify-center backdrop-blur-sm ${
        isDarkMode ? "bg-black bg-opacity-5 " : "bg-black bg-opacity-50"
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
              Add New Card
            </h2>
          </div>
          <div className="md:grid md:grid-cols-2 gap-4">
            <div className="flex flex-col my-4">
              <label
                htmlFor="cardType"
                className={`block text-md font-medium  ${
                  isDarkMode ? "text-gray-300" : "text-gray-700"
                }`}
              >
                Card Type
              </label>
              <input
                type="text"
                placeholder="Visa"
                className={`mt-1 p-3 border block w-full rounded-lg shadow-sm sm:text-sm ${
                  isDarkMode
                    ? "bg-gray-800 border-gray-700 focus:border-blue-500 focus:ring-blue-500 text-gray-300"
                    : "bg-white border-gray-300 focus:border-blue-500 focus:ring-blue-500"
                }`}
                {...register("cardType", { required: "cardType is required" })}
              />
              {errors.cardType && (
                <span className="text-red-600 mt-2 text-xs">
                  {errors.cardType?.message}
                </span>
              )}
            </div>
            <div className="flex flex-col my-4">
              <label
                className={`block text-md font-medium  ${
                  isDarkMode ? "text-gray-300" : "text-gray-700"
                }`}
              >
                Name On Card
              </label>
              <input
                type="text"
                placeholder="John Doe"
                {...register("cardHolder", { required: "name is required" })}
                className={`mt-1 p-3 border block w-full rounded-lg shadow-sm sm:text-sm ${
                  isDarkMode
                    ? "bg-gray-800 border-gray-700 focus:border-blue-500 focus:ring-blue-500 text-gray-300"
                    : "bg-white border-gray-300 focus:border-blue-500 focus:ring-blue-500"
                }`}
              />
              {errors.cardHolder && (
                <p className="text-red-600 mt-2 text-xs">
                  {errors.cardHolder?.message}
                </p>
              )}
            </div>
          </div>
          <div className="md:grid md:grid-cols-2 gap-4">
            <div className="flex flex-col my-4">
              <label
                className={`block text-md font-medium  ${
                  isDarkMode ? "text-gray-300" : "text-gray-700"
                }`}
              >
                Passcode
              </label>
              <input
                type="number"
                placeholder="**** ****"
                className={`mt-1 p-3 border block w-full rounded-lg shadow-sm sm:text-sm ${
                  isDarkMode
                    ? "bg-gray-800 border-gray-700 focus:border-blue-500 focus:ring-blue-500 text-gray-300"
                    : "bg-white border-gray-300 focus:border-blue-500 focus:ring-blue-500"
                }`}
                {...register("passcode", { required: "passcode is required" })}
              />
              {errors.passcode && (
                <p className="text-red-600 mt-2 text-xs">
                  {errors.passcode?.message}
                </p>
              )}
            </div>
            <div className="flex flex-col my-4">
              <label
                className={`block text-md font-medium  ${
                  isDarkMode ? "text-gray-300" : "text-gray-700"
                }`}
              >
                Expiration Date
              </label>
              <input
                type="date"
                className={`mt-1 p-3 border block w-full rounded-lg shadow-sm sm:text-sm ${
                  isDarkMode
                    ? "bg-gray-800 border-gray-700 focus:border-blue-500 focus:ring-blue-500 text-gray-300"
                    : "bg-white border-gray-300 focus:border-blue-500 focus:ring-blue-500"
                }`}
                {...register("expiryDate", {
                  required: "expiryDate is required",
                })}
              />
              {errors.expiryDate && (
                <span className="text-red-600 mt-2 text-xs">
                  {errors.expiryDate?.message}
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
            {isSubmitting ? "Adding to Card..." : "  Add to Card"}
          </button>
        </form>
      </div>
    </div>
  );
};


export default AddCardModal;

