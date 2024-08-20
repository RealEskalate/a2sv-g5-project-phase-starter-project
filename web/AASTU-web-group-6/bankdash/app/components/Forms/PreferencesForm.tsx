"use client";
import React from "react";
import ToggleButton from "../Button/ToggleButton";
import UserService  from "@/app/Services/api/userService";

const PreferencesForm = () => {
  const [preferences, setPreferences] = React.useState({
    currency: "USD",
    sentOrReceiveDigitalCurrency: true,
    receiveMerchantOrder: false,
    accountRecommendations: true,
    timeZone: "(GMT-12:00) International Date Line West",
    twoFactorAuthentication: false,
  });

  const handleToggle = (key: keyof typeof preferences, value: boolean) => {
    setPreferences((prev) => ({ ...prev, [key]: value }));
  };

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    try {
      const response = await UserService.updatePreference(preferences,"accessToken");
      console.log("Preferences updated:", response);
    } catch (error) {
      console.error("Error updating preferences:", error);
    }
  };

  return (
    <form className="mt-8 space-y-6 px-3 py-4" onSubmit={handleSubmit}>
      <div className="flex flex-wrap gap-6">
        <div>
          <label className="block text-sm font-medium text-gray-700">
            Currency
          </label>
          <input
            type="text"
            value={preferences.currency}
            readOnly
            className="sm:w-[285px] md:w-[334px] lg:w-[510px] border-2 border-gray-300 px-5 py-4 rounded-xl"
          />
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700">
            Time Zone
          </label>
          <input
            type="text"
            value={preferences.timeZone}
            readOnly
            className="sm:w-[285px] md:w-[334px] lg:w-[510px] border-2 border-gray-300 px-5 py-4 rounded-xl"
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
            <span className="ml-3 text-sm text-gray-700">
              There are recommendations for my account
            </span>
          </div>
        </div>
      </div>

      <div className="flex lg:justify-end mt-3 sm:w-full sm:justify-center">
        <button
          type="submit"
          className="w-[192px] py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700"
        >
          Save
        </button>
      </div>
    </form>
  );
};

export default PreferencesForm;
