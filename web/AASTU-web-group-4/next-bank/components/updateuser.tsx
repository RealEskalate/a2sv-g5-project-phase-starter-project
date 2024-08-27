import React, { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import Image from "next/image";
import { updateUserDetails, currentuser } from "@/services/userupdate";
import Cookie from "js-cookie";
import { FaPencilAlt } from "react-icons/fa";
import { getDownloadURL, getStorage, ref, uploadBytesResumable } from "firebase/storage";
import { app } from "firebase-functions";

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
  profilePicture: string; // URL
}

const EditProfileForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
    setValue,
  } = useForm<EditProfileFormData>();

  const [profileImage, setProfileImage] = useState<string>("/Images/profilepic.jpeg");
  const [file, setFile] = useState<File | null>(null); // Track profile picture changes
  const token = Cookie.get("accessToken") || "";

  // Fetch user data and prefill form
  useEffect(() => {
    const fetchUserData = async () => {
      console.log("Fetching user data...");

      try {
        const user = await currentuser();
        const userData = user.data;
        console.log("User Data:", userData);

        // Prefill the form with the fetched data
        setValue("name", userData.name);
        setValue("email", userData.email);
        setValue("dateOfBirth", userData.dateOfBirth);
        setValue("permanentAddress", userData.permanentAddress);
        setValue("postalCode", userData.postalCode);
        setValue("username", userData.username);
        setValue("presentAddress", userData.presentAddress);
        setValue("city", userData.city);
        setValue("country", userData.country);

        

      } catch (error) {
        console.error("Error fetching user data:", error);
      }
    };

    
    fetchUserData();
  }, [setValue, token]);

  // Function to handle file upload
   const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFile = e.target.files?.[0];
    if (selectedFile) {
      setFile(selectedFile);
      const imageUrl = URL.createObjectURL(selectedFile);
      setProfileImage(imageUrl); // Temporarily preview image
    }
  };


  const uploadImageToCloud = async (file: File): Promise<string> => {
    const storage = getStorage();
    const storageRef = ref(storage, `profilePictures/${file.name}`); // Store in a specific folder

    const uploadTask = uploadBytesResumable(storageRef, file);

    return new Promise((resolve, reject) => {
      uploadTask.on(
        'state_changed',
        (snapshot) => {
          const progress = (snapshot.bytesTransferred / snapshot.totalBytes) * 100;
          if (progress === 100) {
            console.log(`Upload is ${progress}% done`);
          }
        },
        (error) => {
          console.error("Error during upload:", error);
          reject(error);
        },
        async () => {
          try {
            const url = await getDownloadURL(uploadTask.snapshot.ref);
            console.log("File available at", url);
            resolve(url);
          } catch (err) {
            console.error("Error getting download URL:", err);
            reject(err);
          }
        }
      );
    });
  };

  const onSubmit = async (data: EditProfileFormData) => {
    try {
      // Ensure profilePicture is updated
      data.profilePicture = profileImage;

      console.log("Form data:", data);
      const response = await updateUserDetails(data); // Send updated data to backend
      console.log("Update User Details Response:", response);
    } catch (error) {
      console.error("Error updating user details:", error);
    }
  };

  return (
    <form
      onSubmit={handleSubmit(onSubmit)}
      className="space-y-6 md:grid md:grid-cols-3 md:gap-2"
    >
      {/* Profile picture section */}
      <div className="flex justify-center md:justify-start md:col-span-1">
        <div className="relative ml-4 h-[160px]">
          <Image
            src={profileImage}
            alt="Profile"
            width={150}
            height={150}
            className="rounded-full aspect-square object-cover"
          />
          {/* Pencil icon for changing profile picture */}
          <span className="absolute bottom-2 right-2 w-10 h-10 p-2 bg-blue-800 rounded-full cursor-pointer flex justify-center items-center">
            <FaPencilAlt
              className="text-white"
              onClick={() => document.getElementById("fileInput")?.click()}
            />
            <input
              id="fileInput"
              type="file"
              accept="image/*"
              onChange={handleFileChange} // Correctly handle file change
              className="hidden" // Hidden file input
            />
          </span>
        </div>
      </div>

      {/* Form fields */}
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
            placeholder="City"
            {...register("city")}
          />
        </div>
        <div className="w-full max-w-xs mx-auto sm:w-11/12 sm:mx-0 md:w-full">
          <input
            className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800 w-full"
            placeholder="Country"
            {...register("country")}
          />
        </div>
      </div>

      {/* Submit button */}
      <div className="md:col-span-3 flex justify-end">
        <button
          type="submit"
          className="px-4 py-2 bg-blue-800 text-white rounded-lg hover:bg-blue-700 focus:outline-none"
        >
          Save Changes
        </button>
      </div>
    </form>
  );
};

export default EditProfileForm;
