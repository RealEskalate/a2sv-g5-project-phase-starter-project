import React from "react";
import { useForm } from "react-hook-form";
import Image from "next/image";
import { FaUserEdit } from "react-icons/fa";
import { EditProfileFormData } from "@/types";

const EditProfile: React.FC = () => {
  const { register, handleSubmit } = useForm<EditProfileFormData>({
    defaultValues: {
      name: "Charlene Reed",
      userName: "Charlene Reed",
      email: "charlenereed@gmail.com",
      dateOfBirth: "25 January 1990",
      presentAddress: "San Jose, California, USA",
      city: "San Jose",
      permanentAddress: "San Jose, California, USA",
      postalCode: "45962",
      country: "USA",
    },
  });

  const onSubmit = (data: EditProfileFormData) => {
    // Handle form submission here
    console.log(data);
  };

  return (
    <form
      onSubmit={handleSubmit(onSubmit)}
      className="grid grid-cols-12 gap-6 mt-2 w-full"
    >
      {/* User Image */}
      <div className="col-span-12 md:col-span-2">
        <div className="relative w-25 h-25">
          <Image
            src="https://res.cloudinary.com/dtt1wnvfb/image/upload/v1701954159/photo_2023-12-07%2016.02.23.jpeg.jpg"
            width={183}
            height={36}
            alt="Logo"
            className="w-full h-full rounded-full object-cover"
          />
          <button className="absolute bottom-12 right-4 p-4 bg-blue-600 text-white rounded-full md:bottom-10 md:right-3 md:p-2">
            <FaUserEdit />
          </button>
        </div>
        <div className=" bg-gray-300 mt-4">{/* Empty space with a line */}</div>
      </div>

      {/* User Column-1 */}

      <div className="col-span-12 md:col-span-5">
        {/* Name */}
        <div className="mb-2">
          <label htmlFor="name" className="block text-primary-color-900 mb-2">
            Your Name
          </label>
          <input
            {...register("name")}
            type="text"
            id="name"
            className="appearance-none border rounded-2xl w-full py-2 px-3 text-gray-400 leading-5 focus:outline-none"
          />
        </div>
        {/* Email */}
        <div className="mb-2">
          <label htmlFor="email" className="block text-primary-color-900 mb-2">
            Email
          </label>
          <input
            {...register("email")}
            type="email"
            id="email"
            className="appearance-none border rounded-2xl w-full py-2 px-3 text-gray-400 leading-5 focus:outline-none"
          />
        </div>
        {/* Birth Date */}
        <div className="mb-2">
          <label
            htmlFor="dateOfBirth"
            className="block text-primary-color-900 mb-2"
          >
            Date of Birth
          </label>
          <input
            {...register("dateOfBirth")}
            type="text"
            id="dateOfBirth"
            className="appearance-none border rounded-2xl w-full py-2 px-3 text-gray-400 leading-5 focus:outline-none"
          />
        </div>
        {/* Permanent Address */}
        <div className="mb-2">
          <label
            htmlFor="permanentAddress"
            className="block text-primary-color-900 mb-2"
          >
            Permanent Address
          </label>
          <input
            {...register("permanentAddress")}
            type="text"
            id="permanentAddress"
            className="appearance-none border rounded-2xl w-full py-2 px-3 text-gray-400 leading-5 focus:outline-none"
          />
        </div>
        {/* Postal Code */}
        <div className="mb-2">
          <label
            htmlFor="postalCode"
            className="block text-primary-color-900 mb-2"
          >
            Postal Code
          </label>
          <input
            {...register("postalCode")}
            type="text"
            id="postalCode"
            className="appearance-none border rounded-2xl w-full py-2 px-3 text-gray-400 leading-5 focus:outline-none"
          />
        </div>
      </div>

      {/* User Column-2 */}
      <div className="col-span-12 md:col-span-5">
        <div className="mb-2">
          <label
            htmlFor="userName"
            className="block text-primary-color-900 mb-2"
          >
            User Name
          </label>
          <input
            {...register("userName")}
            type="text"
            id="userName"
            className="appearance-none border rounded-2xl w-full py-2 px-3 text-gray-400 leading-5 focus:outline-none"
          />
        </div>
        <div className="mb-2">
          <label
            htmlFor="password"
            className="block text-primary-color-900 mb-2"
          >
            Password
          </label>
          <input
            type="password"
            id="password"
            className="appearance-none border rounded-2xl w-full py-2 px-3 text-gray-400 leading-5 focus:outline-none"
            placeholder="********"
            disabled
          />
        </div>
        <div className="mb-2">
          <label
            htmlFor="presentAddress"
            className="block text-primary-color-900 mb-2"
          >
            Present Address
          </label>
          <input
            {...register("presentAddress")}
            type="text"
            id="presentAddress"
            className="appearance-none border rounded-2xl w-full py-2 px-3 text-gray-400 leading-5 focus:outline-none"
          />
        </div>
        <div className="mb-2">
          <label htmlFor="city" className="block text-primary-color-900 mb-2">
            City
          </label>
          <input
            {...register("city")}
            type="text"
            id="city"
            className="appearance-none border rounded-2xl w-full py-2 px-3 text-gray-400 leading-5 focus:outline-none"
          />
        </div>
        <div className="mb-2">
          <label
            htmlFor="country"
            className="block text-primary-color-900 mb-2"
          >
            Country
          </label>
          <input
            {...register("country")}
            type="text"
            id="country"
            className="appearance-none border rounded-2xl w-full py-2 px-3 text-gray-400 leading-5 focus:outline-non"
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

export default EditProfile;
