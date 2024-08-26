"use client";
import { useEffect } from "react";
import { useAppDispatch } from "@/app/Redux/store/store";
import {
  setUserData,
  setPreferences,
  setStatus,
  setError,
} from "../slices/userSlice";
import UserService from "@/app/Services/api/userService";

const useUserDataDispatch = (accessToken: string) => {
  const dispatch = useAppDispatch();

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        dispatch(setStatus("loading"));

        const userData = await UserService.current(accessToken);

        if (userData) {
          dispatch(setUserData(userData));
          dispatch(setPreferences(userData.preference));
        }

        dispatch(setStatus("succeeded"));
      } catch (error) {
        dispatch(setError("Failed to fetch user data and preferences"));
        dispatch(setStatus("failed"));
      }
    };

    fetchUserData();
  }, [dispatch, accessToken]);
};

export default useUserDataDispatch;
