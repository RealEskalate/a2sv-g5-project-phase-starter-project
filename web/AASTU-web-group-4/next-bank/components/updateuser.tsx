import React, { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import Image from "next/image";
import { updateUserDetails, fetchUserDetails } from "@/services/userupdate";
import Cookie from "js-cookie";
import { FaPencilAlt } from "react-icons/fa";

interface EditProfileFormData {
  name: string;
  email: string;
  dateOfBirth: string;
  permanentAddress: string;
  postalCode: string;
  username: string;
  presentAddress: string;
  city: string;
  country: string;
  profilePicture: File | null;
}

const EditProfileForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
    setValue,
  } = useForm<EditProfileFormData>();
  const [profileImage, setProfileImage] = useState("/Images/profilepic.jpeg");
  const [file, setFile] = useState<File | null>(null);
  const token = Cookie.get("accessToken") || "";
  // Fetch user data and prefill form
  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const userData = await fetchUserDetails(token); // Fetch user data
        setValue("name", userData.name);
        setValue("email", userData.email);
        setValue("dateOfBirth", userData.dateOfBirth);
        setValue("permanentAddress", userData.permanentAddress);
        setValue("postalCode", userData.postalCode);
        setValue("username", userData.username);
        setValue("presentAddress", userData.presentAddress);
        setValue("city", userData.city);
        setValue("country", userData.country);
        if (userData.profilePicture) {
          setProfileImage(userData.profilePicture); // Set profile picture
        }
      } catch (error) {
        console.error("Error fetching user data:", error);
      }
    };

    fetchUserData();
  }, [setValue]);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFile = e.target.files?.[0];
    if (selectedFile) {
      setFile(selectedFile);
      const imageUrl = URL.createObjectURL(selectedFile);
      setProfileImage(imageUrl); // Temporarily preview image
    }
  };

  const onSubmit = async (data: EditProfileFormData) => {
    try {
      // Append the uploaded image to the form data
      const formData = new FormData();
      formData.append("name", data.name);
      formData.append("email", data.email);
      formData.append("dateOfBirth", data.dateOfBirth);
      formData.append("permanentAddress", data.permanentAddress);
      formData.append("postalCode", data.postalCode);
      formData.append("username", data.username);
      formData.append("presentAddress", data.presentAddress);
      formData.append("city", data.city);
      formData.append("country", data.country);
      if (file) formData.append("profilePicture", file);

      const response = await updateUserDetails(formData); // Send form data to backend
      console.log("Update User Details Response:", response);
    } catch (error) {
      console.error("Error:", error);
    }
  };

  return (
    <form
      onSubmit={handleSubmit(onSubmit)}
      className="space-y-6 md:grid md:grid-cols-3 md:gap-2"
    >
      <div className="flex justify-center md:justify-start md:col-span-1">
        <div className="relative ml-4 h-[160px]">
          <Image
            src={profileImage}
            alt="Profile"
            width={150}
            height={150}
            className="rounded-full aspect-square object-cover"
          />
          {/* Pencil icon positioned bottom-right */}
          <span className="absolute  bottom-2 right-2 w-10 h-10 p-2 bg-blue-800 rounded-full cursor-pointer flex justify-center items-center">
            <FaPencilAlt className="text-white"  onClick={() => document.getElementById('fileInput')?.click()}/>
            <input
            id="fileInput"
            type="file"
            accept="image/*"
            onChange={handleFileChange}
            className="hidden"  // Hidden file input
          />
          </span>
        </div>
      </div>

      <div className="md:col-span-1 space-y-4">
        <div className="w-full max-w-xs mx-auto sm:w-11/12 sm:mx-0 md:w-full">
          <input
            className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800 w-full"
            placeholder="John Doe"
            {...register("name", { required: true })}
          />
          {errors.name && <p className="text-red-500">Name is required</p>}
        </div>
        <div className="w-full max-w-xs mx-auto sm:w-11/12 sm:mx-0 md:w-full">
          <input
            className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800 w-full"
            type="email"
            placeholder="john@example.com"
            {...register("email", { required: true })}
          />
          {errors.email && <p className="text-red-500">Email is required</p>}
        </div>
        <div className="w-full max-w-xs mx-auto sm:w-11/12 sm:mx-0 md:w-full">
          <input
            className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800 w-full"
            type="date"
            {...register("dateOfBirth")}
          />
        </div>
        <div className="w-full max-w-xs mx-auto sm:w-11/12 sm:mx-0 md:w-full">
          <input
            className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800 w-full"
            placeholder="123 Main St"
            {...register("permanentAddress")}
          />
        </div>
        <div className="w-full max-w-xs mx-auto sm:w-11/12 sm:mx-0 md:w-full">
          <input
            className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800 w-full"
            placeholder="12345"
            {...register("postalCode")}
          />
        </div>
      </div>

      <div className="md:col-span-1 space-y-4">
        <div className="w-full max-w-xs mx-auto sm:w-11/12 sm:mx-0 md:w-full">
          <input
            className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800 w-full"
            placeholder="john_doe"
            {...register("username", { required: true })}
          />
          {errors.username && (
            <p className="text-red-500">Username is required</p>
          )}
        </div>
        <div className="w-full max-w-xs mx-auto sm:w-11/12 sm:mx-0 md:w-full">
          <input
            className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800 w-full"
            placeholder="456 Another St"
            {...register("presentAddress")}
          />
        </div>
        <div className="w-full max-w-xs mx-auto sm:w-11/12 sm:mx-0 md:w-full">
          <input
            className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800 w-full"
            placeholder="Cityname"
            {...register("city")}
          />
        </div>
        <div className="w-full max-w-xs mx-auto sm:w-11/12 sm:mx-0 md:w-full">
          <input
            className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800 w-full"
            placeholder="Countryname"
            {...register("country")}
          />
        </div>
        <div className="md:pt-5">
          <button
            type="submit"
            className="w-full max-w-xs mx-auto sm:w-full bg-blue-800 text-white py-2 rounded-md"
          >
            Save
          </button>
        </div>
      </div>
    </form>
  );
};

export default EditProfileForm;
