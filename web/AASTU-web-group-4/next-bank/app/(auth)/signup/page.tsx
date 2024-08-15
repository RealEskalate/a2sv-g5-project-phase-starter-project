"use client";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { User } from "@/types/index";

const SignupForm = () => {
  const [step, setStep] = useState(1);
  const [formData, setFormData] = useState<User>({
    name: "",
    email: "",
    dateOfBirth: "",
    permanentAddress: "",
    postalCode: "",
    username: "",
    password: "",
    presentAddress: "",
    city: "",
    country: "",
    profilePicture: "",
    preference: {
      currency: "",
      sentOrReceiveDigitalCurrency: false,
      receiveMerchantOrder: false,
      accountRecommendations: false,
      timeZone: "",
      twoFactorAuthentication: false,
    },
  });

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<User>();
  const onSubmit = (data: User) => {
    console.log("Final Data:", data);
  };

  const handleNextStep = () => setStep(step + 1);
  const handlePreviousStep = () => setStep(step - 1);

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <div className="flex justify-center items-center mb-8 mt-8">
        <img
          src="https://cdn.freelogovectors.net/wp-content/uploads/2024/03/chase_logo-freelogovectors.net_.png"
          alt="Logo"
          className="h-16 w-16"
        />
        <h1 className="text-2xl font-extrabold ml-4">NextBank</h1>
      </div>
      <h1 className="text-blue-500 font-bold text-center text-xl mt-3 mb-2">
        Sign Up
      </h1>
      <h1 className="text-center mb-8">Please enter your details</h1>

      {step === 1 && (
        <div className="font-bold mx-4 sm:mx-6 ml-8 md:mx-8 lg:mx-10 grid gap-6 sm:gap-8 lg:gap-10 md:grid-cols-2">
          {/* Step 1 Fields */}

          <div className="mt-4">
            <label htmlFor="name">Full Name</label>
            <input
              type="text"
              id="name"
              placeholder="ex: John"
              {...register("name", { required: "Name is required" })}
              className="border border-black font-normal rounded-md p-2 w-full"
            />
            {errors.name && (
              <p className="text-red-500 font-normal">Name is required</p>
            )}
          </div>
          <div className="mt-4">
            <label htmlFor="email">Email</label>
            <input
              type="email"
              id="email"
              placeholder="ex: example@gmail.com"
              {...register("email", {
                required: "Email is required",
                pattern: /^\S+@\S+\.\S+$/,
              })}
              className="border border-black rounded-md p-2 w-full"
            />
            {errors.email && (
              <p className="text-red-500 font-normal">{errors.email.message}</p>
            )}
          </div>

          {/* dateOfBirth */}
          <div className="mt-4 font-normal">
            <label htmlFor="dateOfBirth" className="font-bold">
              Date of Birth
            </label>
            <input
              type="date"
              id="dateOfBirth"
              {...register("dateOfBirth", {
                required: "Date of Birth is required",
              })}
              className="border border-black rounded-md p-2 w-full"
            />
            {errors.dateOfBirth && (
              <p className="text-red-500 font-normal">
                {errors.dateOfBirth.message}
              </p>
            )}
          </div>
              {/* username */}
              <div className="mt-4">
                <label htmlFor="username">Username</label>
                <input
                  type="text"
                  id="username"
                  placeholder="ex: Star"
                  {...register("username", { required: "Username is required" })}
                  className="border border-black rounded-md p-2 w-full"
                />
                {errors.username && (
                  <p className="text-red-500 font-normal">
                    {errors.username.message}
                  </p>
                )}
              </div>
    
              {/* password */}
              <div className="mt-4">
                <label htmlFor="password">Password</label>
                <input
                  type="password"
                  id="password"
                  placeholder="ex: password"
                  {...register("password", {
                    required: "Password is required",
                    minLength: 8,
                  })}
                  className="border border-black rounded-md p-2 w-full"
                />
                {errors.password && (
                  <p className="text-red-500 font-normal">
                    {errors.password.message}
                  </p>
                )}
              </div>
          {/* permanentAddress */}
          <div className="mt-4">
            <label htmlFor="permanentAddress">Permanent Address</label>
            <textarea
              id="permanentAddress"
              placeholder="ex: A.A"
              {...register("permanentAddress", {
                required: "Permanent Address is required",
              })}
              className="border border-black rounded-md p-2 w-full h-24"
            />
            {errors.permanentAddress && (
              <p className="text-red-500 font-normal">
                {errors.permanentAddress.message}
              </p>
            )}
          </div>
          {/* postalCode */}
          <div className="mt-4">
            <label htmlFor="postalCode">Postal Code</label>
            <input
              type="text"
              id="postalCode"
              placeholder="ex: 1000"
              {...register("postalCode", {
                required: "Postal Code is required",
              })}
              className="border border-black rounded-md p-2 w-full"
            />
            {errors.postalCode && (
              <p className="text-red-500 font-normal">
                {errors.postalCode.message}
              </p>
            )}
          </div>
          <div>  </div>
          {/* Add more fields as needed */}
          <button
            type="button"
            onClick={handleNextStep}
            className="flex mx-auto mt-14 text-center justify-center md:text-center bg-blue-500 hover:bg-blue-700 text-white font-bold px-4 py-3 rounded"
          >
            Continue
          </button>
        </div>
      )}

      {step === 2 && (
        <div className="font-bold mx-4 sm:mx-8 md:mx-16 lg:mx-28 grid gap-6 sm:gap-8 lg:gap-10 md:grid-cols-2">
          {/* Step 2 Fields */}
          {/* presentAddress */}
          
          <div className="mt-4">
            <label htmlFor="presentAddress">Present Address</label>
            <input
              type="text"
              id="presentAddress"
              placeholder="ex: Bola"
              {...register("presentAddress", {
                required: "Present Address is required",
              })}
              className="border border-black rounded-md p-2 w-full"
            />
            {errors.presentAddress && (
              <p className="text-red-500 font-normal">
                {errors.presentAddress.message}
              </p>
            )}
          </div>
          {/* city */}
          <div className="mt-4">
            <label htmlFor="city">City</label>
            <input
              type="text"
              id="city"
              placeholder="ex: A.A"
              {...register("city", { required: "City is required" })}
              className="border border-black rounded-md p-2 w-full"
            />
            {errors.city && (
              <p className="text-red-500 font-normal">{errors.city.message}</p>
            )}
          </div>
          {/* country */}
          <div className="mt-4">
            <label htmlFor="country">Country</label>
            <input
              type="text"
              id="country"
              placeholder="ex: Ethiopia"
              {...register("country", { required: "Country is required" })}
              className="border border-black rounded-md p-2 w-full"
            />
            {errors.country && (
              <p className="text-red-500 font-normal">
                {errors.country.message}
              </p>
            )}
</div>
          {/* profilePicture */}
          <div className="mt-4 font-normal">
            <label htmlFor="profilePicture" className="mb-3 font-bold">
              Profile Picture
            </label>
            <input
              type="file"
              id="profilePicture"
              {...register("profilePicture")}
              className="p-2 w-full"
            />
          </div>

          {/* preferences */}
          <div className="mt-4 ">
            {/* <h3 className="text-lg font-semibold">Preferences</h3> */}
            <div className="space-y-4">
              <div>
                <label htmlFor="currency" className="text-lg font-semibold">
                  Preference Currency
                </label>
                <input
                  type="text"
                  id="currency"
                  {...register("preference.currency", {
                    required: "Currency is required",
                  })}
                  className="border border-black rounded-md p-2 w-full"
                />
                {errors.preference?.currency && (
                  <p className="text-red-500 font-normal">
                    {errors.preference?.currency.message}
                  </p>
                )}
              </div>

              <div>
                <label>Sent or Receive Digital Currency</label>
                <input
                  type="checkbox"
                  id="sentOrReceiveDigitalCurrency"
                  {...register("preference.sentOrReceiveDigitalCurrency")}
                />
              </div>

              <div>
                <label>Receive Merchant Order</label>
                <input
                  type="checkbox"
                  id="receiveMerchantOrder"
                  {...register("preference.receiveMerchantOrder")}
                />
              </div>

              <div>
                <label>Account Recommendations</label>
                <input
                  type="checkbox"
                  id="accountRecommendations"
                  {...register("preference.accountRecommendations")}
                />
              </div>

              <div>
                <label htmlFor="timeZone">Time Zone</label>
                <input
                  type="text"
                  id="timeZone"
                  {...register("preference.timeZone", {
                    required: "Time Zone is required",
                  })}
                  className="border border-black rounded-md p-2 w-full"
                />
                {errors.preference?.timeZone && (
                  <p className="text-red-500 font-normal">
                    {errors.preference?.timeZone.message}
                  </p>
                )}
              </div>

              <div>
                <label>Two Factor Authentication</label>
                <input
                  type="checkbox"
                  id="twoFactorAuthentication"
                  {...register("preference.twoFactorAuthentication")}
                />
              </div>
            </div>
          </div>
          <div>   </div>
          {/* buttons */}
          <div className="flex justify-between mt-14">
            <button
              type="button"
              onClick={handlePreviousStep}
              className="bg-gray-500 hover:bg-gray-700 text-white font-bold px-4 py-3 rounded"
            >
              Back
            </button>
            <button
              type="submit"
              className="bg-blue-500 hover:bg-blue-700 text-white font-bold px-4 py-3 rounded"
            >
              Submit
            </button>
          </div>
        </div>
      )}
    </form>
  );
};

export default SignupForm;
