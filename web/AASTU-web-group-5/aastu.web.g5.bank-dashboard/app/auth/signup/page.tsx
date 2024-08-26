"use client";
import Image from "next/image";
import { useForm } from "react-hook-form";
import axios from "axios";
import { useRouter } from "next/navigation";
import { useState } from "react";
import {
  getStorage,
  ref,
  uploadBytesResumable,
  getDownloadURL,
} from "firebase/storage";
import { app } from "../../../lib/firebase"; // Import the initialized Firebase app
// import img from "/public/images/christina.png";
interface Preference {
  currency: string;
  sentOrReceiveDigitalCurrency: boolean;
  receiveMerchantOrder: boolean;
  accountRecommendations: boolean;
  timeZone: string;
  twoFactorAuthentication: boolean;
}

interface RegistrationFormData {
  name: string;
  email: string;
  dateOfBirth: string;
  permanentAddress: string;
  postalCode: string;
  username: string;
  password: string;
  presentAddress: string;
  city: string;
  country: string;
  profilePicture: string;
  preference: Preference;
}

const Register = () => {
  // const img = 
  const img = "/public/images/christina.png";
  const [profilePicture, setProfilePicture] = useState<string>(img);
  const [uploading, setUploading] = useState(false);
  const handleImageUpload = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    const storage = getStorage(app);
    const storageRef = ref(storage, `profilePictures/$${file.name}`);
    const uploadTask = uploadBytesResumable(storageRef, file);

    setUploading(true);

    uploadTask.on(
      "state_changed",
      (snapshot) => {},
      (error) => {
        console.error("Upload failed:", error);
        setUploading(false);
      },
      () => {
        getDownloadURL(uploadTask.snapshot.ref).then((downloadURL) => {
          setProfilePicture(downloadURL);
          setUploading(false);
        });
      }
    );
  };

  console.log(profilePicture, "profile");
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegistrationFormData>();
  const router = useRouter();
  const [errorMessage, setErrorMessage] = useState<string | null>(null);

  const onSubmit = async (data: RegistrationFormData) => {
    console.log(data,'data111')
    let updatedData = data
    updatedData.profilePicture = profilePicture
    
    console.log("6666666",updatedData)
    try {
      const response = await axios.post(
        "https://bank-dashboard-irbd.onrender.com/auth/register",
        data
      );
      console.log(response.data);
      router.push("/auth/signin");
    } catch (error: any) {
      console.log('999999999999999',data)
      setErrorMessage(error.response?.data?.message || "Registration failed");
    }
  };

  return (
    <div className="flex flex-col items-center min-h-screen bg-gradient-to-br from-blue-900 via-gray-800 to-blue-950 py-10 px-5">
    
    <div className="bg-gray-900 w-full max-w-5xl p-8 rounded-lg shadow-lg">
  <h1 className="text-4xl font-extrabold mb-8 text-white text-center">
    Create Your Account
  </h1>

  <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
    {errorMessage && (
      <p className="text-red-500 text-center mb-4">{errorMessage}</p>
    )}

    <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
      <div className="flex flex-col">
        <input
          {...register("name", { required: true })}
          placeholder="Name"
          className="input-field bg-gray-800 text-white p-3 rounded-lg"
        />
        {errors.name && (
          <p className="text-red-500 mt-1">Name is required</p>
        )}
      </div>

      <div className="flex flex-col">
        <input
          {...register("email", { required: true })}
          type="email"
          placeholder="Email"
          className="input-field bg-gray-800 text-white p-3 rounded-lg"
        />
        {errors.email && (
          <p className="text-red-500 mt-1">Email is required</p>
        )}
      </div>

      <div className="flex flex-col">
        <input
          {...register("dateOfBirth", { required: true })}
          type="date"
          placeholder="Date of Birth"
          className="input-field bg-gray-800 text-white p-3 rounded-lg"
        />
        {errors.dateOfBirth && (
          <p className="text-red-500 mt-1">Date of Birth is required</p>
        )}
      </div>

      <div className="flex flex-col">
        <input
          {...register("permanentAddress", { required: true })}
          placeholder="Permanent Address"
          className="input-field bg-gray-800 text-white p-3 rounded-lg"
        />
        {errors.permanentAddress && (
          <p className="text-red-500 mt-1">Permanent Address is required</p>
        )}
      </div>

      <div className="flex flex-col">
        <input
          {...register("postalCode", { required: true })}
          placeholder="Postal Code"
          className="input-field bg-gray-800 text-white p-3 rounded-lg"
        />
        {errors.postalCode && (
          <p className="text-red-500 mt-1">Postal Code is required</p>
        )}
      </div>

      <div className="flex flex-col">
        <input
          {...register("username", { required: true })}
          placeholder="Username"
          className="input-field bg-gray-800 text-white p-3 rounded-lg"
        />
        {errors.username && (
          <p className="text-red-500 mt-1">Username is required</p>
        )}
      </div>

      <div className="flex flex-col">
        <input
          {...register("password", { required: true })}
          type="password"
          placeholder="Password"
          className="input-field bg-gray-800 text-white p-3 rounded-lg"
        />
        {errors.password && (
          <p className="text-red-500 mt-1">Password is required</p>
        )}
      </div>

      <div className="flex flex-col">
        <input
          {...register("presentAddress", { required: true })}
          placeholder="Present Address"
          className="input-field bg-gray-800 text-white p-3 rounded-lg"
        />
        {errors.presentAddress && (
          <p className="text-red-500 mt-1">Present Address is required</p>
        )}
      </div>

      <div className="flex flex-col">
        <input
          {...register("city", { required: true })}
          placeholder="City"
          className="input-field bg-gray-800 text-white p-3 rounded-lg"
        />
        {errors.city && (
          <p className="text-red-500 mt-1">City is required</p>
        )}
      </div>

      <div className="flex flex-col">
        <input
          {...register("country", { required: true })}
          placeholder="Country"
          className="input-field bg-gray-800 text-white p-3 rounded-lg"
        />
        {errors.country && (
          <p className="text-red-500 mt-1">Country is required</p>
        )}
      </div>

      <div className="w-full flex justify-center mb-8">
  <div className="w-56 h-56 rounded-full overflow-hidden flex items-center justify-center bg-gray-800">
    <label htmlFor="profilePictureUpload" className="cursor-pointer">
      <Image
        src={profilePicture}
        width={224}  
        height={224} 
        className="object-cover"
        alt="Profile picture"
      />
    </label>
    <input
      type="file"
      id="profilePictureUpload"
      className="hidden"
      onChange={handleImageUpload}
      accept="image/*"
    />
  </div>
</div>
    </div>

    <div className="flex flex-col items-center mt-8">
      <h2 className="text-2xl font-semibold text-gray-300 mb-4">
        Preferences
      </h2>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-4 w-full">
        <div className="flex flex-col">
          <select
            {...register("preference.currency", { required: true })}
            className="input-field bg-gray-800 text-white p-3 rounded-lg"
          >
            <option value="">Select Currency</option>
            <option value="USD">USD</option>
            <option value="EUR">EUR</option>
            <option value="GBP">GBP</option>
            <option value="JPY">JPY</option>
          </select>
          {errors.preference?.currency && (
            <p className="text-red-500 mt-1">Currency is required</p>
          )}
        </div>

        <div className="flex flex-col">
          <select
            {...register("preference.timeZone", { required: true })}
            className="input-field bg-gray-800 text-white p-3 rounded-lg"
          >
            <option value="">Select Time Zone</option>
            <option value="UTC-12:00">UTC-12:00</option>
            <option value="UTC-11:00">UTC-11:00</option>
            <option value="UTC-10:00">UTC-10:00</option>
            <option value="UTC-09:00">UTC-09:00</option>
            <option value="UTC-08:00">UTC-08:00</option>
            <option value="UTC-07:00">UTC-07:00</option>
            <option value="UTC-06:00">UTC-06:00</option>
            <option value="UTC-05:00">UTC-05:00</option>
            <option value="UTC-04:00">UTC-04:00</option>
            <option value="UTC-03:00">UTC-03:00</option>
            <option value="UTC-02:00">UTC-02:00</option>
            <option value="UTC-01:00">UTC-01:00</option>
            <option value="UTC+00:00">UTC+00:00</option>
            <option value="UTC+12:00">UTC+12:00</option>
          </select>
          {errors.preference?.timeZone && (
            <p className="text-red-500 mt-1">Time Zone is required</p>
          )}
        </div>
      </div>

      <div className="flex flex-wrap justify-center space-x-4 mt-4">
        <label className="flex items-center text-gray-200">
          <input
            {...register("preference.sentOrReceiveDigitalCurrency")}
            type="checkbox"
            className="mr-2 accent-blue-600"
          />
          Send or Receive Digital Currency
        </label>

        <label className="flex items-center text-gray-200">
          <input
            {...register("preference.receiveMerchantOrder")}
            type="checkbox"
            className="mr-2 accent-blue-600"
          />
          Receive Merchant Order
        </label>

        <label className="flex items-center text-gray-200">
          <input
            {...register("preference.accountRecommendations")}
            type="checkbox"
            className="mr-2 accent-blue-600"
          />
          Account Recommendations
        </label>

        <label className="flex items-center text-gray-200">
          <input
            {...register("preference.twoFactorAuthentication")}
            type="checkbox"
            className="mr-2 accent-blue-600"
          />
          Two-Factor Authentication
        </label>
      </div>
    </div>

    <div className="flex justify-center mt-8">
      <button
        type="submit"
        className="w-full md:w-1/2 py-3 bg-purple-700 text-white font-semibold rounded-full transition duration-300 hover:bg-purple-600 focus:outline-none focus:ring-2 focus:ring-purple-500"
      >
						{uploading ? "Uploading..." : "Register"}
            </button>
    </div>
  </form>
</div>
       


    </div>
  );
};

export default Register;








