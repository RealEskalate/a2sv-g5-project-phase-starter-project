import React from 'react'

const Navigation = () => {
  return (
    <div className="flex justify-between items-center border-b pb-4">
    <div className="flex items-center">
      <img
        src="/profile-pic.png"
        alt="Profile"
        className="w-24 h-24 rounded-full mr-4"
      />
      <button className="p-2 bg-blue-600 text-white rounded-full">
        Edit
      </button>
    </div>
    <div className="flex space-x-8">
      <button className="pb-2 border-b-4 border-blue-600">
        Edit Profile
      </button>
      <button className="pb-2 border-b-4 border-transparent">
        Preferences
      </button>
      <button className="pb-2 border-b-4 border-transparent">
        Security
      </button>
    </div>
  </div>
    )
}

export default Navigation