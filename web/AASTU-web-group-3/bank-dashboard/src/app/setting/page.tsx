"use client";
import React, { useEffect, useState } from "react";
import EditProfile from "../components/setting/EditProfile";
import Preference from "../components/setting/Preference";
import Security from "../components/setting/Security";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "@/lib/redux/store";
import { useGetCurrentUserQuery } from "@/lib/redux/api/settingApi";
import { setSetting,setError,setLoading } from "@/lib/redux/slices/settingSlice";


const HomePage = () => {
  const [activeTab, setActiveTab] = useState("Edit Profile");
  const dispatch = useDispatch();
  const { setting, loading, error } = useSelector(
    (state: RootState) => state.setting
  );

  console.log(setting)
  const { data, isLoading, isError } = useGetCurrentUserQuery();

  useEffect(() => {
    dispatch(setLoading(isLoading));
  
    if (data) {
      dispatch(setSetting([data]));
    }
  
    if (isError) {
      dispatch(setError("Error loading transactions"));
    }
  }, [data, isLoading, isError, dispatch]);
  
  if (loading || setting.length === 0) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;
  
  const renderContent = () => {
    switch (activeTab) {
      case "Edit Profile":
        return <EditProfile  userData = {setting}/>;
      case "Preference":
        return <Preference userPrefernce = {setting[0].data.preference} />;
      case "Security":
        return <Security />;
      default:
        return null;
    }
  };

  return (
    <div className="w-11/12 mt-3 ml-6 bg-white rounded-3xl">
      <div className="border-[#718EBF] border-b flex justify-between md:justify-start md:gap-3">
        <div
          onClick={() => setActiveTab("Edit Profile")}
          className={`cursor-pointer text-xl py-4 px-3 ${
            activeTab === "Edit Profile"
              ? "text-blue-600 border-blue-600 border-b-4"
              : "text-[#718EBF]"
          }`}
        >
          Edit Profile
        </div>
        <div
          onClick={() => setActiveTab("Preference")}
          className={`cursor-pointer text-xl py-4 px-3 ${
            activeTab === "Preference"
              ? "text-blue-600 border-blue-600 border-b-4"
              : "text-[#718EBF]"
          }`}
        >
          Preference
        </div>
        <div
          onClick={() => setActiveTab("Security")}
          className={`cursor-pointer text-xl py-4 px-3 ${
            activeTab === "Security"
              ? "text-blue-600 border-blue-600 border-b-4"
              : "text-[#718EBF]"
          }`}
        >
          Security
        </div>
      </div>

      <div>{renderContent()}</div>
    </div>
  );
};

export default HomePage;
