"use client";
import { useEffect } from "react";
import { useAppDispatch } from "@/app/Redux/store/store";
import {
  setTran,
  setStatus,
  setError,
} from "@/app/Redux/slices/transactionSlice";

const useCardDispatch = (accessToken: string) => {
  const dispatch = useAppDispatch();

  useEffect(() => {
    const fetchInitialCards = async () => {
      // for transaction
    };

    fetchInitialCards();
  }, [dispatch, accessToken]);
};

export default useCardDispatch;
