"use client";
import React, { useState } from "react";
import { useFormContext, Controller } from "react-hook-form";
import Select from "react-select";
import countryList from "react-select-country-list";

interface StepProps {
  step: number;
}

const Step2: React.FC<StepProps> = ({ step }) => {
  const {
    register,
    control,
    formState: { errors },
  } = useFormContext();
  const countryOptions = countryList().getData();

  const [Country, setCountry] = useState("country");

  return (
    <div className="space-y-4 min-h-[350px] flex flex-col justify-between">
      <h2 className="text-2xl font-semibold dark:text-darkText">
        Step {step}: Address Details
      </h2>
      <div>
        <label className="block text-sm font-medium dark:text-darkText">
          Permanent Address <span className="text-red-500">*</span>
        </label>
        <input
          {...register("permanentAddress", {
            required: "Permanent Address is required",
          })}
          className="border rounded-md w-full h-12 border-[#CCCCF5] dark:border-darkComponent p-2 bg-white dark:bg-darkPage dark:text-darkText"
        />
        {errors.permanentAddress && (
          <p className="text-red-500 text-sm mt-1">
            {String(errors.permanentAddress.message)}
          </p>
        )}
      </div>
      <div>
        <label className="block text-sm font-medium dark:text-darkText">
          Postal Code <span className="text-red-500">*</span>
        </label>
        <input
          {...register("postalCode", { required: "Postal Code is required" })}
          className="border rounded-md w-full h-12 border-[#CCCCF5] dark:border-darkComponent p-2 bg-white dark:bg-darkPage dark:text-darkText"
        />
        {errors.postalCode && (
          <p className="text-red-500 text-sm mt-1">
            {String(errors.postalCode.message)}
          </p>
        )}
      </div>
      <div>
        <label className="block text-sm font-medium dark:text-darkText">
          Present Address <span className="text-red-500">*</span>
        </label>
        <input
          {...register("presentAddress", {
            required: "Present Address is required",
          })}
          className="border rounded-md w-full h-12 border-[#CCCCF5] dark:border-darkComponent p-2 bg-white dark:bg-darkPage dark:text-darkText"
        />
        {errors.presentAddress && (
          <p className="text-red-500 text-sm mt-1">
            {String(errors.presentAddress.message)}
          </p>
        )}
      </div>
      <div>
        <label className="block text-sm font-medium dark:text-darkText">
          City <span className="text-red-500">*</span>
        </label>
        <input
          {...register("city", { required: "City is required" })}
          className="border rounded-md w-full h-12 border-[#CCCCF5] dark:border-darkComponent p-2 bg-white dark:bg-darkPage dark:text-darkText"
        />
        {errors.city && (
          <p className="text-red-500 text-sm mt-1">
            {String(errors.city.message)}
          </p>
        )}
      </div>
      <div>
        <label className="block text-sm font-medium dark:text-darkText">
          Country <span className="text-red-500">*</span>
        </label>
        <Controller
          control={control}
          name="country"
          rules={{ required: "Country is required" }}
          render={({ field }) => (
            <Select
              {...field}
              options={countryOptions}
              className="mt-1 dark:text-darkText"
              styles={{
                control: (base, state) => ({
                  ...base,
                  padding: "0.25rem",
                  borderRadius: "0.375rem",
                  backgroundColor: state.isFocused ? "#f0f0f0" : "white", // Light mode background
                  borderColor: state.isFocused ? "#1D4ED8" : "#CCCCF5", // Light mode border
                  color: "#25324B", // Light mode text color
                  "&:hover": {
                    borderColor: "#1D4ED8", // Light mode hover border color
                  },
                  ...(field.value && {
                    backgroundColor: state.isFocused ? "#172941" : "#0f1a2b", // Dark mode background
                    borderColor: "#352F44", // Dark mode border color
                    color: "#FFFFFF", // Dark mode text color
                  }),
                }),
                menu: (base) => ({
                  ...base,
                  backgroundColor: "#FFFFFF", // Light mode menu background
                  color: "#25324B", // Light mode menu text color
                  ...(field.value && {
                    backgroundColor: "#0f1a2b", // Dark mode menu background
                    color: "#FFFFFF", // Dark mode menu text color
                  }),
                }),
                option: (base, { isFocused, isSelected }) => ({
                  ...base,
                  backgroundColor: isFocused
                    ? "#E5E7EB" // Light mode focused option background
                    : isSelected
                    ? "#1D4ED8" // Light mode selected option background
                    : "#FFFFFF", // Light mode option background
                  color: isSelected ? "#FFFFFF" : "#25324B", // Light mode option text color
                  "&:hover": {
                    backgroundColor: "#E5E7EB", // Light mode option hover background
                  },
                  ...(field.value && {
                    backgroundColor: isFocused
                      ? "#1D4ED8" // Dark mode focused option background
                      : isSelected
                      ? "#1D4ED8" // Dark mode selected option background
                      : "#0f1a2b", // Dark mode option background
                    color: isSelected ? "#FFFFFF" : "#FFFFFF", // Dark mode option text color
                    "&:hover": {
                      backgroundColor: "#1D4ED8", // Dark mode option hover background
                    },
                  }),
                }),
                singleValue: (base) => ({
                  ...base,
                  color: "#25324B", // Light mode selected value color
                  ...(field.value && {
                    color: "#FFFFFF", // Dark mode selected value color
                  }),
                }),
              }}
              onChange={(selectedOption) =>
                field.onChange(selectedOption?.label)
              }
              value={countryOptions.find(
                (option) => option.label === field.value
              )}
            />
          )}
        />

        {errors.country && (
          <p className="text-red-500 text-sm mt-1">
            {String(errors.country.message)}
          </p>
        )}
      </div>
    </div>
  );
};

export default Step2;
