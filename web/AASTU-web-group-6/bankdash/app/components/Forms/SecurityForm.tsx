'use client';
import React from 'react';
import ToggleButton from '../Button/ToggleButton';

const SecurityForm = () => {
  const handleToggle = (checked: boolean) => {
    console.log("Two-factor Authentication is now", checked ? "Enabled" : "Disabled");
  };

  return (
    <form className="mt-8 space-y-6">
      <div className="mt-6">
        <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
          Two-factor Authentication
        </label>
        <div className="flex items-center mt-4">
          <ToggleButton
            onToggle={handleToggle}
            initialChecked={true} // Set initial state based on the original checkbox state
          />
          <span className="ml-3 text-sm text-gray-700 dark:text-gray-300">
            Enable or disable two-factor authentication
          </span>
        </div>
      </div>

      <div className="mt-6">
        <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
          Change Password
        </label>
        <div className="mt-4 space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
              Current Password
            </label>
            <input
              type="password"
              value=""
              readOnly
              className="mt-1 block w-full rounded-md border-gray-300 dark:border-gray-600 shadow-sm bg-gray-100 dark:bg-gray-700"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
              New Password
            </label>
            <input
              type="password"
              value=""
              readOnly
              className="mt-1 block w-full rounded-md border-gray-300 dark:border-gray-600 shadow-sm bg-gray-100 dark:bg-gray-700"
            />
          </div>
        </div>
      </div>

      <button
        type="submit"
        className="mt-6 w-full py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 dark:bg-blue-700 hover:bg-blue-700 dark:hover:bg-blue-800"
      >
        Save
      </button>
    </form>
  );
};

export default SecurityForm;
