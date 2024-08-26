"use client"
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { cardSchema } from "@/schema";
import { postCards } from "@/lib/api";
import { useToast } from "@/components/ui/use-toast";






type FormData = z.infer<typeof cardSchema>;

const InputForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    resolver: zodResolver(cardSchema),
  });
  const {toast} = useToast();
  const onSubmit = async (data: FormData) => {
    
    try {
      const updatedForm = {...data,balance:300}
      console.log(updatedForm);
      
      await postCards(updatedForm);
      toast({
        title: "Success",
        description: "Card was submitted succesfully",
        variant: "success",
      });
    } catch (error) {
      console.error("Failed to submit form:", error);
      toast({
        title: "Error",
        description: "Card submission failed",
        variant: "destructive",
      });
    }
  };


  return (
    <div>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="md:grid md:grid-cols-2 gap-4">
          <div className="flex flex-col my-4">
            <label className="text-[#515B6F] font-semibold">Card Type</label>
            <input
              type="text"
              placeholder="Classic"
              className="inputField mb-2 rounded-xl py-2 px-2 border border-gray-300"
              {...register("cardType")}
            />
            {errors.cardType && (
              <p className="text-red-500 text-center mt-2">
                {errors.cardType?.message}
              </p>
            )}
          </div>
          <div className="flex flex-col my-4">
            <label className="text-[#515B6F] font-semibold">Name On Card</label>
            <input
              type="text"
              placeholder="John Doe"
              className="inputField mb-2 rounded-xl py-2 px-2 border border-gray-300"
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
            <label className="text-[#515B6F] font-semibold">Card Number</label>
            <input
              type="text"
              placeholder="**** **** **** ****"
              className="inputField mb-2 rounded-xl py-2 px-2 border border-gray-300"
              {...register("passcode")}
            />
            {errors.passcode && (
              <p className="text-red-500 text-center mt-2">
                {errors.passcode?.message}
              </p>
            )}
          </div>
          <div className="flex flex-col my-4">
            <label className="text-[#515B6F] font-semibold">
              Expiration Date
            </label>
            <input
              type="date"
              className="inputField mb-2 rounded-xl py-2 px-2 border border-gray-300"
              {...register("expiryDate")}
            />
            {errors.expiryDate && (
              <p className="text-red-500 text-center mt-2">
                {errors.expiryDate?.message}
              </p>
            )}
          </div>
        </div>
        <button className="bg-[#1814F3] sm:w-[100%] text-white p-2 sm:rounded-full md:max-w-[160px] md:rounded-md">
          Add to Cart
        </button>
      </form>
    </div>
  );
};

export default InputForm;