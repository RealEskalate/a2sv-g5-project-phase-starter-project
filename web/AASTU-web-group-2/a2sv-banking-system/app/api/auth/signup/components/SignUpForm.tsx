"use client";
import React, { useState } from "react";
import { useForm, FormProvider, SubmitHandler } from "react-hook-form";
import PageOne from "./PageOne";
import PageTwo from "./PageTwo";
import PageThree from "./PageThree";
import { register } from "@/lib/api/authenticationController";
import {
  RegisterRequest,
  RegisterResponse,
} from "@/types/authenticationController.interface";
import { useRouter } from "next/navigation";

type SignUpFormData = {
  name: string;
  email: string;
  password: string;
  city: string;
  presentAddress: string;
  country: string;
  permanentAddress: string;
  postalCode: string;
  timeZone: string;
  currency: string;
  sentOrReceiveDigitalCurrency: boolean;
  receiveMerchantOrder: boolean;
  accountRecommendations: boolean;
  twoFactorAuthentication: boolean;
  profilePicture?: File | string;
  dateOfBirth: string;
  username: string;
};

const SignUpForm = () => {
  const methods = useForm<SignUpFormData>();
  const [page, setPage] = useState(1);
  const [loading, setLoading] = useState(false); // Added loading state
  const router = useRouter();

  const onSubmit: SubmitHandler<SignUpFormData> = async (data) => {
    setLoading(true); // Start loading when form is submitted
    try {
      const registerRequest: RegisterRequest = {
        name: data.name,
        email: data.email,
        dateOfBirth: data.dateOfBirth,
        permanentAddress: data.permanentAddress,
        postalCode: data.postalCode,
        username: data.username,
        password: data.password,
        presentAddress: data.presentAddress,
        city: data.city,
        country: data.country,
        profilePicture:
          typeof data.profilePicture === "string" ? data.profilePicture : "",
        preference: {
          currency: data.currency,
          sentOrReceiveDigitalCurrency: data.sentOrReceiveDigitalCurrency,
          receiveMerchantOrder: data.receiveMerchantOrder,
          accountRecommendations: data.accountRecommendations,
          timeZone: data.timeZone,
          twoFactorAuthentication: data.twoFactorAuthentication,
        },
      };

      console.log("Final Data:", registerRequest);
      const d: RegisterResponse = await register(registerRequest);
      console.log(d);
      alert("Registered Successfully");
      router.push("/api/auth/signin");
    } catch (error) {
      console.error("Registration error:", error);
    } finally {
      setLoading(false); // Stop loading after the form is submitted
    }
  };

  const nextPage = () => setPage((prev) => prev + 1);
  const prevPage = () => setPage((prev) => prev - 1);

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-100 ">
      <div className="bg-white dark:bg-[#1e1e2e] p-8 rounded-lg shadow-lg w-full">
        <FormProvider {...methods}>
          <form onSubmit={methods.handleSubmit(onSubmit)}>
            {page === 1 && <PageOne />}
            {page === 2 && <PageTwo />}
            {page === 3 && <PageThree />}

            <div className="flex justify-center mt-4 gap-4">
              {page > 1 && (
                <button
                  type="button"
                  onClick={prevPage}
                  className="bg-[#1814F3] text-white py-2 px-5 rounded-lg flex justify-center items-center font-semibold hover:bg-blue-600 transition duration-200"
                >
                  Previous
                </button>
              )}
              {page < 3 ? (
                <button
                  type="button"
                  onClick={nextPage}
                  className="bg-[#1814F3] text-white py-2 px-5 rounded-lg flex justify-center items-center font-semibold hover:bg-blue-600 transition duration-200"
                >
                  Next
                </button>
              ) : (
                <button
                  type="submit"
                  className="bg-green-500 text-white py-2 px-5 rounded-lg flex justify-center items-center font-semibold hover:bg-green-600 transition duration-200"
                  disabled={loading} // Disable button when loading
                >
                  {loading ? (
                    <svg
                      className="animate-spin h-5 w-5 text-white mr-3"
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 24 24"
                    >
                      <circle
                        className="opacity-25"
                        cx="12"
                        cy="12"
                        r="10"
                        stroke="currentColor"
                        strokeWidth="4"
                      />
                      <path
                        className="opacity-75"
                        fill="currentColor"
                        d="M4 12a8 8 0 018-8v8l-3.5 3.5L4 12z"
                      />
                    </svg>
                  ) : (
                    "Submit"
                  )}
                </button>
              )}
            </div>
          </form>
        </FormProvider>
      </div>
    </div>
  );
};

export default SignUpForm;
