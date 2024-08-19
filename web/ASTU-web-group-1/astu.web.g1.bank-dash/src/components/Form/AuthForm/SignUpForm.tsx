"use client";

import InputGroup from "../InputGroup";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/navigation";

const stepOneSchema = z
  .object({
    name: z.string().min(1, "Name is required"),
    email: z.string().email("Invalid email address"),
    dateOfBirth: z.string().refine((date) => !isNaN(Date.parse(date)), {
      message: "Invalid date format",
    }),
    permanentAddress: z.string().min(1, "Permanent address is required"),
    postalCode: z.string().min(1, "Postal code is required"),
    username: z.string().min(1, "Username is required"),
    password: z.string().min(6, "At least 6 characters long"),
    confirmPassword: z
      .string()
      .min(6, "At least 6 characters long"),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Passwords don't match",
    path: ["confirmPassword"],
  });

const stepTwoSchema = z.object({
  presentAddress: z.string().min(1, "Present address is required"),
  city: z.string().min(1, "City is required"),
  country: z.string().min(1, "Country is required"),

  currency: z.string().min(1, "Currency is required"),
  timeZone: z.string().min(1, "Time zone is required"),
});

const steps = ["Step 1", "Step 2"];
const stepSchemas = [stepOneSchema, stepTwoSchema];

const SignUpForm = () => {
  const [step, setStep] = useState(0);
  const [prevFormData, setprevFormData] = useState({});
  const router = useRouter()

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(stepSchemas[step]),
    mode: "onTouched",
  });

  const onSubmit = async (data: any) => {
    if (step < steps.length - 1) {
      setprevFormData((prevFormData) => ({ ...prevFormData, ...data }));
      console.log("Form Data at step", prevFormData);
      setStep(step + 1);
    } else {
      setprevFormData((prevFormData) => ({ ...prevFormData, ...data }));

      const finalData = {
        ...prevFormData,
        presentAddress: data.presentAddress,
        city: data.city,
        country: data.country,
        profilePicture: "/images/67sdfsd6f7s8d6fa8s6fsgf_s6fs7",
        preference: {
          currency: data.currency,
          sentOrReceiveDigitalCurrency: true,
          receiveMerchantOrder: true,
          accountRecommendations: true,
          timeZone: data.timeZone,
          twoFactorAuthentication: true,
        },
      };
      // console.log("Returned and combined values", finalData);

      const res = fetch('https://bank-dashboard-6acc.onrender.com/auth/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(finalData),
      });
      res.then((res) => {
        if (res.ok) {
          console.log('User created successfully');
          router.push('/api/auth/signin')
        } 
        if (!res.ok) {
          console.log('Failed to create user');
        }
      })
    }
  };

  const onBack = () => setStep(step - 1);

  return (
    <form
      className="flex flex-col items-center w-full md:w-10/12 justify-center p-6 rounded-2xl bg-slate-50"
      onSubmit={handleSubmit(onSubmit)}
    >
      <p className='text-[#333B69] pb-3 text-20px text-left font-semibold w-full'>Register</p>
      {step == 0 && (
        <>
          <div className="w-full md:flex md:gap-4 ">
            <InputGroup
              id="name"
              label="Full Name"
              inputType="text"
              registerName="name"
              register={register}
              placeholder="Enter Full Name"
              errorMessage={errors?.name?.message as string}
            />

            <InputGroup
              id="email"
              label="Email"
              inputType="email"
              registerName="email"
              register={register}
              placeholder="Enter Your Email"
              errorMessage={errors?.email?.message as string}
            />
          </div>

          <div className="w-full md:flex md:gap-4 ">
            <InputGroup
              id="dateOfBirth"
              label="Date Of Birth"
              inputType="date"
              registerName="dateOfBirth"
              register={register}
              placeholder="Enter Date Of Birth"
              errorMessage={errors?.dateOfBirth?.message as string}
            />

            <InputGroup
              id="permanentAddress"
              label="Permanent Address"
              inputType="text"
              registerName="permanentAddress"
              register={register}
              placeholder="Enter Permanent Address"
              errorMessage={errors?.permanentAddress?.message as string}
            />
          </div>

          <div className="w-full md:flex md:gap-4 ">
            <InputGroup
              id="postalCode"
              label="Postal Code"
              inputType="text"
              registerName="postalCode"
              register={register}
              placeholder="Enter Postal Code"
              errorMessage={errors?.postalCode?.message as string}
            />

            <InputGroup
              id="username"
              label="Username"
              inputType="text"
              registerName="username"
              register={register}
              placeholder="Enter Username"
              errorMessage={errors?.username?.message as string}
            />
          </div>
          <div className="w-full md:flex md:gap-4 ">
            <InputGroup
              id="password"
              label="Password"
              inputType="password"
              registerName="password"
              register={register}
              placeholder="Enter Password"
              errorMessage={errors?.password?.message as string}
            />

            <InputGroup
              id="confirmPassword"
              label="Confirm Password"
              inputType="password"
              registerName="confirmPassword"
              register={register}
              placeholder="RE-Enter password"
              errorMessage={errors?.confirmPassword?.message as string}
            />
          </div>
        </>
      )}

      {step == 1 && (
        <>
          <div className="w-full md:flex md:gap-4 ">
            <InputGroup
              id="presentAddress"
              label="Present Address"
              inputType="text"
              registerName="presentAddress"
              register={register}
              placeholder="Enter Present Address"
              errorMessage={errors?.presentAddress?.message as string}
            />

            <InputGroup
              id="city"
              label="City"
              inputType="text"
              registerName="city"
              register={register}
              placeholder="Enter City"
              errorMessage={errors?.city?.message as string}
            />
          </div>

          <div className="w-full md:flex md:gap-4 ">
            <InputGroup
              id="country"
              label="Country"
              inputType="text"
              registerName="country"
              register={register}
              placeholder="Enter Country"
              errorMessage={errors?.country?.message as string}
            />
          </div>

          <div className="w-full md:flex md:gap-4 ">
            <div className=" w-full lg:w-6/12 space-y-3 my-3">
              <label htmlFor="timeZone" className="gray-dark text-16px">
                Currency
              </label>
              <select
                id="timeZone"
                {...register("currency")}
                className="w-full border-2 border-[#DFEAF2] p-5 py-3 rounded-xl placeholder:text-blue-steel focus:border-blue-steel outline-none"
              >
                <option value="USD">USD</option>
                <option value="US">Birr</option>
                <option value="CA">Birr</option>
                <option value="FR">Birr</option>
              </select>
              {errors?.currency && (
                <p className="text-red-400"> {errors?.currency?.message as string} </p>
              )}
            </div>

            <div className=" w-full lg:w-6/12 space-y-3 my-3">
              <label htmlFor="timeZone" className="gray-dark text-16px">
                Time Zone
              </label>
              <select
                id="timeZone"
                {...register("timeZone")}
                className="w-full border-2 border-[#DFEAF2] p-5 py-3 rounded-xl placeholder:text-blue-steel focus:border-blue-steel outline-none"
              >
                <option value="GMT 3+">GMT 3+</option>
                <option value="US">Birr</option>
                <option value="CA">Birr</option>
                <option value="FR">Birr</option>
              </select>
              {errors?.timeZone && (
                <p className="text-red-400"> {errors?.timeZone?.message as string} </p>
              )}
            </div>
          </div>
        </>
      )}

      <div
        className={`w-full flex ${
          step == 0 ? "justify-end" : "justify-between"
        }`}
      >
        {step > 0 && (
          <button
            type="submit"
            className="bg-[#1814f3] text-white px-10 py-3 font-Lato font-bold rounded-lg mt-4"
            onClick={onBack}
          >
            Back
          </button>
        )}

        <button
          type="submit"
          className="bg-[#1814f3] text-white px-10 py-3 font-Lato font-bold rounded-lg mt-4"
        >
          {step < steps.length - 1 ? "Next" : "Register"}
        </button>
      </div>
    </form>
  );
};

export default SignUpForm;
