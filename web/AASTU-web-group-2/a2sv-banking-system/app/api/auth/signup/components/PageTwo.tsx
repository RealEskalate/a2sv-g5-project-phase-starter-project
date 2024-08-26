import React from "react";
import { useFormContext, Controller } from "react-hook-form";
import { BsExclamationCircle } from "react-icons/bs";

// Reusable error message component
const ErrorMessage = ({ message }: any) => (
  <div className="flex flex-row align-middle mt-2">
    <BsExclamationCircle className="text-red-500 mr-2" />
    <span className="text-red-500">{message}</span>
  </div>
);

const PageTwo = () => {
  const { control } = useFormContext();

  return (
    <div className="flex justify-center">
      <div className="flex flex-col gap-5 w-full md:w-1/2">
        <h1 className="text-4xl font-black text-[#202430] dark:text-[#cdd6f4] mt-4 text-center">
          Your Address
        </h1>
        
        {/* Reusable Field Component */}
        {[
          { name: "city", label: "City", placeholder: "Enter your city" },
          {
            name: "presentAddress",
            label: "Present Address",
            placeholder: "Enter your present address",
          },
          { name: "country", label: "Country", placeholder: "Enter your country" },
          {
            name: "permanentAddress",
            label: "Permanent Address",
            placeholder: "Enter your permanent address",
          },
          {
            name: "postalCode",
            label: "Postal Code",
            placeholder: "Enter your postal code",
          },
          { name: "timeZone", label: "Time Zone", placeholder: "Enter your time zone" },
        ].map(({ name, label, placeholder }) => (
          <div key={name} className="mb-3 flex flex-col gap-2">
            <label className="text-[#515B6F] font-semibold dark:text-[#cdd6f4]">
              {label}
            </label>
            <Controller
              name={name}
              control={control}
              defaultValue=""
              render={({ field, fieldState: { error } }) => (
                <>
                  <input
                    {...field}
                    type="text"
                    className="border border-gray-400 rounded-lg py-2 px-5 w-full dark:border-gray-600 dark:bg-[#313244] dark:text-[#cdd6f4] focus:outline-none focus:ring-2 focus:ring-indigo-500 dark:focus:border-[#4640DE]"
                    placeholder={placeholder}
                    aria-label={label}
                    aria-invalid={!!error}
                  />
                  {error && <ErrorMessage message={error.message} />}
                </>
              )}
            />
          </div>
        ))}
      </div>
    </div>
  );
};

export default PageTwo;
