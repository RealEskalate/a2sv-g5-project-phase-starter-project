import React from 'react'

const EditProfileForm = () => {
  return (
    <div className="max-w-3xl mx-auto w-full mt-2  bg-white rounded-lg shadow-md">

      <form className="mt-8 space-y-6">
        <div className="flex flex-col gap-6">
        <div className='flex'>
            <div>
                <label className="block text-sm font-medium text-gray-700">
                Your Name
                </label>
                <input
                type="text"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
                />
            </div>

            <div>
                <label className="block text-sm font-medium text-gray-700">
                User Name
                </label>
                <input
                type="text"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
                />
            </div>
        </div>

        <div className='flex'>
            <div>
                <label className="block text-sm font-medium text-gray-700">
                Email
                </label>
                <input
                type="email"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
                />
            </div>

            <div>
                <label className="block text-sm font-medium text-gray-700">
                Password
                </label>
                <input
                type="password"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
                />
            </div>
        </div>

        <div className='flex'>
          <div>
            <label className="block text-sm font-medium text-gray-700">
              Date of Birth
            </label>
            <input
              type="text"
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
              />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">
              Present Address
            </label>
            <input
              type="text"
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
            />
          </div>
        </div>
        
        <div className='flex'>
            <div>
                <label className="block text-sm font-medium text-gray-700">
                Permanent Address
                </label>
                <input
                type="text"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
                />
            </div>

            <div>
                <label className="block text-sm font-medium text-gray-700">
                City
                </label>
                <input
                type="text"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
                />
            </div>
        </div>
        
        <div className='flex'>
            <div>
                <label className="block text-sm font-medium text-gray-700">
                Postal Code
                </label>
                <input
                type="text"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
                />
            </div>

            <div>
                <label className="block text-sm font-medium text-gray-700">
                Country
                </label>
                <input
                type="text"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
                />
            </div>
            </div>
        </div>

        <button
          type="submit"
          className="w-full py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700"
        >
          Save
        </button>
      </form>
    </div>
  )
}

export default EditProfileForm