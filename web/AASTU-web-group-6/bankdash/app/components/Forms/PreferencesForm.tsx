"use client";
import React, { useEffect } from "react";
import ToggleButton from "../Button/ToggleButton";
import UserService from "@/app/Services/api/userService";
import { useAppSelector } from "@/app/Redux/store/store";

const PreferencesForm = () => {
  const userData = useAppSelector((state) => state.user);

  const [preferences, setPreferences] = React.useState({
    currency: "USD",
    sentOrReceiveDigitalCurrency: false,
    receiveMerchantOrder: false,
    accountRecommendations: false,
    timeZone: "(GMT-12:00) International Date Line West",
    twoFactorAuthentication: false,
  });

  useEffect(() => {
    if (userData.preferences) {
      setPreferences(userData.preferences);
    }
  }, [userData.preferences]);

  const handleToggle = (key: keyof typeof preferences, value: boolean) => {
    console.log("key", key);
    setPreferences((prev) => ({ ...prev, [key]: value }));
  };

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    console.log("PREFERENCE", preferences);
    try {
      const response = await UserService.updatePreference(
        preferences,
        "accessToken"
      );
      console.log("Preferences updated:", response);
    } catch (error) {
      console.error("Error updating preferences:", error);
    }
  };

  return (
    <form
      className="flex flex-col justify-start mt-6 space-y-6 py-4"
      onSubmit={handleSubmit}
    >
      <div className="flex flex-wrap gap-x-6">
        <div className="xxs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
          <label className="mb-1 text-sm font-medium text-slate-700">
            Currency
          </label>
          <select
            id="currency"
            name="currency"
            className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE] mt-1 block w-full text-sm"
            value={preferences.currency}
            onChange={(e) =>
              setPreferences({ ...preferences, currency: e.target.value })
            }
          >
            <option value="USD">USD - US Dollar</option>
            <option value="EUR">EUR - Euro</option>
            <option value="GBP">GBP - British Pound</option>
            <option value="JPY">JPY - Japanese Yen</option>
            <option value="CNY">CNY - Chinese Yuan</option>
            <option value="INR">INR - Indian Rupee</option>
            <option value="ETB">ETB - Ethiopian Birr</option>
          </select>
        </div>

        <div className="xxs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
          <label className="mb-1 text-sm font-medium text-slate-700">
            Time Zone
          </label>
          <input
            type="text"
            value={preferences.timeZone}
            readOnly
            className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
          />
        </div>
      </div>

      <div className="mt-6">
        <label className="block text-sm font-medium text-gray-700">
          Notification
        </label>
        <div className="mt-4 space-y-4">
          <div className="flex items-center">
            <ToggleButton
              onToggle={(checked) =>
                handleToggle("sentOrReceiveDigitalCurrency", checked)
              }
              initialChecked={preferences.sentOrReceiveDigitalCurrency}
            />
            <span className="ml-3 text-sm text-gray-700">
              I send or receive digital currency
            </span>
          </div>

          <div className="flex items-center">
            <ToggleButton
              onToggle={(checked) =>
                handleToggle("receiveMerchantOrder", checked)
              }
              initialChecked={preferences.receiveMerchantOrder}
            />
            <span className="ml-3 text-sm text-gray-700">
              I receive merchant order
            </span>
          </div>

          <div className="flex items-center">
            <ToggleButton
              onToggle={(checked) =>
                handleToggle("accountRecommendations", checked)
              }
              initialChecked={preferences.accountRecommendations}
            />
            <span className="ml-3 text-sm font-Inter text-wrap text-gray-700">
              There are recommendations for my account
            </span>
          </div>
        </div>
      </div>

      <div className="flex justify-end mt-3 w-full sm:justify-end">
        <button
          className="px-4 py-2 bg-[#1814F3] text-white rounded-md hover:bg-[#0702db] transition-all duration-300"
          type="submit"
        >
          Save Changes
        </button>
      </div>
    </form>
  );
};

export default PreferencesForm;
