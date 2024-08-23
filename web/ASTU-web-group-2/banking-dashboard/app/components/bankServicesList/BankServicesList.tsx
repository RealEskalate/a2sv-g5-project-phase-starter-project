"use client";
import React, { useEffect } from "react";
import BankService from "../bankService/BankService";
import BankServiceMobile, {
  BankServiceType,
} from "../bankService/BankServiceMobile";
import { useGetBankServiceQuery } from "@/lib/service/BankService";
import { useSession } from "next-auth/react";
import BankServiceSkeleton from "./BankServiceSkeleton";

const BankServicesList = () => {
  const { data: session, status } = useSession();

  useEffect(() => {}, [session, status]);

  const accessToken = session?.user.accessToken!;

  const { data: res, isLoading } = useGetBankServiceQuery({
    accessToken: accessToken,
    size: 10,
    page: 0,
  });

  if (isLoading) {
    return (
      <BankServiceSkeleton />
    );
  }
  const data = res.data!.content!;

  return (
    <div>
      {data.length == 0 && (
        <img src="/assets/bankService/empty-image.png" alt="list empty" />
      )}
      <div className="flex flex-col gap-5 max-md:hidden">
        {data.map((bankService: BankServiceType) => (
          <BankService {...bankService} key={bankService.id} />
        ))}
      </div>

      {/* Mobile view */}
      <div className="flex flex-col gap-5 md:hidden">
        {data.map((bankService: BankServiceType) => (
          <BankServiceMobile {...bankService} key={bankService.id} />
        ))}
      </div>
    </div>
  );
};

export default BankServicesList;
