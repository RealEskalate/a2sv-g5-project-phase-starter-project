"use client";
import React, { useEffect } from "react";
import BankService from "../bankService/BankService";
import BankServiceMobile, {
  BankServiceType,
} from "../bankService/BankServiceMobile";
import { useGetBankServiceQuery } from "@/lib/service/BankService";
import { useSession } from "next-auth/react";

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
      <div className="flex justify-center items-center flex-col flex-initial flex-wrap h-[225px] w-full bg-white animate-pulse rounded-[25px]">
            <div className="flex flex-row gap-2">
              <div className="w-4 h-4 rounded-full bg-blue-700 animate-bounce [animation-delay:.7s]"></div>
              <div className="w-4 h-4 rounded-full bg-blue-700 animate-bounce [animation-delay:.3s]"></div>
              <div className="w-4 h-4 rounded-full bg-blue-700 animate-bounce [animation-delay:.7s]"></div>
            </div>
          </div>
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
