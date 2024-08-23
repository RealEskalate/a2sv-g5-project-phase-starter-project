import React from "react";
import { useForm, Controller } from "react-hook-form";
import { Switch } from "../switch";
import { PreferenceFormData } from "@/types";

const Preferences: React.FC = () => {
  const { register, handleSubmit, control } = useForm<PreferenceFormData>({
    defaultValues: {
      currency: "USD",
      timeZone: "(GMT-12:00) International Date Line West",
      digitalCurrency: true,
      merchantOrder: false,
      recommendations: true,
    },
  });

  const onSubmit = (data: PreferenceFormData) => {
    // TODO: Handle form submission here
    console.log(data);
  };

  return (
    <div className="flex">
      <form
        onSubmit={handleSubmit(onSubmit)}
        className=" grid grid-cols-12 gap-6 mt-2 w-full"
      >
        {/* Currency */}
        <div className="col-span-12 sm:col-span-6">
          <label
            htmlFor="currency"
            className="block text-sm font-medium text-gray-700"
          >
            Currency
          </label>
          <input
            {...register("currency")}
            type="text"
            id="currency"
            className="mt-1 block w-full rounded-md text-gray-500 border-gray-300 py-3 shadow-sm sm:text-sm"
          />
        </div>

        {/* Time Zone */}
        <div className="col-span-12 sm:col-span-6">
          <label
            htmlFor="timeZone"
            className="block text-sm font-medium text-gray-700"
          >
            Time Zone
          </label>
          <input
            {...register("timeZone")}
            type="text"
            id="timeZone"
            className="mt-1 block w-full rounded-md text-gray-500 border-gray-300 py-3 shadow-sm sm:text-sm"
          />
        </div>

        {/* Notifications */}
        <div className="col-span-12 sm:col-span-6">
          <h3 className="text-lg font-medium leading-6 text-gray-900">
            Notification
          </h3>

          <div className="mt-2 space-y-2">
            {/* Digital Currency Selector */}
            <div className="flex items-center space-x-2">
              <Controller
                name="digitalCurrency"
                control={control}
                render={({ field }) => (
                  <Switch
                    id="digitalCurrency"
                    checked={field.value}
                    onCheckedChange={field.onChange}
                  />
                )}
              />
              <label
                htmlFor="digitalCurrency"
                className="ml-3 text-sm text-gray-900"
              >
                I send or receive digital currency
              </label>
            </div>

            {/* Merchant Order Selector */}
            <div className="flex items-center space-x-2">
              <Controller
                name="merchantOrder"
                control={control}
                render={({ field }) => (
                  <Switch
                    id="merchantOrder"
                    checked={field.value}
                    onCheckedChange={field.onChange}
                  />
                )}
              />
              <label
                htmlFor="merchantOrder"
                className="ml-3 text-sm text-gray-900"
              >
                I receive merchant order
              </label>
            </div>

            {/* Recommendation Selector */}
            <div className="flex items-center space-x-2">
              <Controller
                name="recommendations"
                control={control}
                render={({ field }) => (
                  <Switch
                    id="recommendations"
                    checked={field.value}
                    onCheckedChange={field.onChange}
                  />
                )}
              />
              <label
                htmlFor="recommendations"
                className="ml-3 text-sm text-gray-900"
              >
                There are recommendation for my account
              </label>
            </div>
          </div>
        </div>

        {/* Save Button */}
        <div className="col-span-12 flex justify-end">
          <button
            type="submit"
            className="bg-blue-500 hover:bg-blue-700 text-white py-2 px-10 rounded-xl focus:outline-none"
          >
            Save
          </button>
        </div>
      </form>
    </div>
  );
};

export default Preferences;
