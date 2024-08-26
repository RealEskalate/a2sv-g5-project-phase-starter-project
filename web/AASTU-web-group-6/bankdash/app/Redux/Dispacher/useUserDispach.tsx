"use client";
import { useEffect } from "react";
import { useAppDispatch } from "@/app/Redux/store/store";
import CardService from "@/app/Services/api/CardService";
import { setInvestment, setStatus, setError } from "../slices/userSlice";
import InvestmentService from "@/app/Services/api/investmentApi";

const useUserDispatch = (accessToken: string) => {
  const dispatch = useAppDispatch();
  

  useEffect(() => {
    const fetchInitialCards = async () => {
      try {
        dispatch(setStatus("loading"));
        const res: any = await InvestmentService.getInvestmentData(accessToken);
        if (res) {
          dispatch(setInvestment(res));
          dispatch(setStatus("succeeded"));
        }
      } catch (error) {
        dispatch(setError("Failed to fetch cards"));
        dispatch(setStatus("failed"));
      }
    };

    fetchInitialCards();
  }, [dispatch, accessToken]);
};

export default useUserDispatch;
