// Security.tsx
import React from "react";
import { useForm, Controller } from "react-hook-form";
import { Switch } from "../switch";
import { SecurityFormData } from "@/types";

const Security: React.FC = () => {
  const { register, handleSubmit, control } = useForm<SecurityFormData>({
    defaultValues: {
      twoFactorAuthentication: false,
      currentPassword: "",
      newPassword: "",
    },
  });

  const onSubmit = (data: SecurityFormData) => {
    // Handle form submission here
    console.log(data);
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="mt-4">
      <div className="mb-4">
        <label
          htmlFor="twoFactorAuthentication"
          className="block text-gray-700 font-semibold mb-2"
        >
          Two-factor Authentication
        </label>

        {/* Recommendation Selector */}
        <div className="flex items-center space-x-2">
          <Controller
            name="twoFactorAuthentication"
            control={control}
            render={({ field }) => (
              <Switch
                id="twoFactorAuthentication"
                checked={field.value}
                onCheckedChange={field.onChange}
              />
            )}
          />
          <label
            htmlFor="twoFactorAuthentication"
            className="ml-3 text-sm text-gray-900"
          >
            Enable or disable two factor authentication
          </label>
        </div>
      </div>

      <div className="mb-4">
        <label
          htmlFor="currentPassword"
          className="block text-gray-700 font-semibold mb-2"
        >
          Change Password
        </label>
        <div className="mb-2">
          <label htmlFor="currentPassword" className="block text-gray-700 mb-2">
            Current Password
          </label>
          <input
            {...register("currentPassword")}
            type="password"
            id="currentPassword"
            className="appearance-none border rounded-2xl w-1/2 py-3 px-3 text-gray-700 leading-5 focus:outline-none focus:shadow-outline"
            placeholder="********"
          />
        </div>
        <div>
          <label htmlFor="newPassword" className="block text-gray-700 mb-2">
            New Password
          </label>
          <input
            {...register("newPassword")}
            type="password"
            id="newPassword"
            className="appearance-none border rounded-2xl w-1/2 py-3 px-3 text-gray-700 leading-5 focus:outline-none "
            placeholder="********"
          />
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
  );
};

export default Security;
