import React from "react";

const SecurityPage: React.FC = () => {
  return (
    <div className="max-w-3xl mx-auto mt-10 p-8 bg-white rounded-lg shadow-md">
      <div className="flex justify-between items-center border-b pb-4">
        <div className="flex space-x-8">
          <button className="pb-2 border-b-4 border-transparent text-gray-500">
            Edit Profile
          </button>
          <button className="pb-2 border-b-4 border-transparent text-gray-500">
            Preferences
          </button>
          <button className="pb-2 border-b-4 border-blue-600 text-blue-600">
            Security
          </button>
        </div>
      </div>

      <form className="mt-8 space-y-6">
        <div className="mt-6">
          <label className="block text-sm font-medium text-gray-700">
            Two-factor Authentication
          </label>
          <div className="flex items-center mt-4">
            <input
              type="checkbox"
              checked
              className="h-6 w-6 text-blue-600 border-gray-300 rounded"
            />
            <span className="ml-3 text-sm text-gray-700">
              Enable or disable two-factor authentication
            </span>
          </div>
        </div>

        <div className="mt-6">
          <label className="block text-sm font-medium text-gray-700">
            Change Password
          </label>
          <div className="mt-4 space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-700">
                Current Password
              </label>
              <input
                type="password"
                value="********"
                readOnly
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm bg-gray-100"
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700">
                New Password
              </label>
              <input
                type="password"
                value="********"
                readOnly
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm bg-gray-100"
              />
            </div>
          </div>
        </div>

        <button
          type="submit"
          className="mt-6 w-full py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700"
        >
          Save
        </button>
      </form>
    </div>
  );
};

export default SecurityPage;
