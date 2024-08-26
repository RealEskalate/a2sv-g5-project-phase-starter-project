"use client";
import React from "react";
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

  return (
    <div className="space-y-4 min-h-[350px] flex flex-col justify-between">
      <h2 className="text-2xl font-semibold">Step {step}: Address Details</h2>
      <div>
        <label className="block text-sm font-medium">
          Permanent Address <span className="text-red-500">*</span>
        </label>
        <input
          {...register("permanentAddress", {
            required: "Permanent Address is required",
          })}
          className="mt-1 p-2 block w-full border rounded-md"
        />
        {errors.permanentAddress && (
          <p className="text-red-500 text-sm mt-1">
            {String(errors.permanentAddress.message)}
          </p>
        )}
      </div>
      <div>
        <label className="block text-sm font-medium">
          Postal Code <span className="text-red-500">*</span>
        </label>
        <input
          {...register("postalCode", { required: "Postal Code is required" })}
          className="mt-1 p-2 block w-full border rounded-md"
        />
        {errors.postalCode && (
          <p className="text-red-500 text-sm mt-1">
            {String(errors.postalCode.message)}
          </p>
        )}
      </div>
      <div>
        <label className="block text-sm font-medium">
          Present Address <span className="text-red-500">*</span>
        </label>
        <input
          {...register("presentAddress", {
            required: "Present Address is required",
          })}
          className="mt-1 p-2 block w-full border rounded-md"
        />
        {errors.presentAddress && (
          <p className="text-red-500 text-sm mt-1">
            {String(errors.presentAddress.message)}
          </p>
        )}
      </div>
      <div>
        <label className="block text-sm font-medium">
          City <span className="text-red-500">*</span>
        </label>
        <input
          {...register("city", { required: "City is required" })}
          className="mt-1 p-2 block w-full border rounded-md"
        />
        {errors.city && (
          <p className="text-red-500 text-sm mt-1">
            {String(errors.city.message)}
          </p>
        )}
      </div>
      <div>
        <label className="block text-sm font-medium">
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
              className="mt-1"
              styles={{
                control: (base) => ({
                  ...base,
                  padding: "0.25rem",
                  borderRadius: "0.375rem",
                }),
              }}
              onChange={(selectedOption) =>
                field.onChange(selectedOption?.label)
              }
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
