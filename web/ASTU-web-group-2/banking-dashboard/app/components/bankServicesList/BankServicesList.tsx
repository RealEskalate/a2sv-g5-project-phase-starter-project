import React from "react";
import BankService from "../bankService/BankService";
import BankServiceMobile from "../bankService/BankServiceMobile";
import { useGetBankServiceQuery } from "@/lib/service/BankService";
import { useSession } from "next-auth/react";
import { useRouter } from "next/navigation";

const BankServicesList = () => {
  const router = useRouter();

  const { data: session, status } = useSession();

  if (status == "unauthenticated") router.push("/login");

  const accessToken = session?.user.accessToken;

  const {data } = useGetBankServiceQuery({
    accessToken: accessToken,
    size: 6,
    page: 1,
  });
  console.log(data);
  
  return (
    <div>
      <div className="flex flex-col gap-5 max-md:hidden">
        {data.map((bankService) => (
          <BankService
            logoLink={bankService.logo}
            category={bankService.category}
            description={bankService.description}
            details={bankService.details}
            action={bankService.action}
            key={bankService.id}
          />
        ))}
      </div>

      {/* Mobile view */}
      <div className="flex flex-col gap-5 md:hidden">
        {data.map((bankService) => (
          <BankServiceMobile
            logoLink={bankService.logo}
            category={bankService.category}
            description={bankService.description}
            details={bankService.details}
            action={bankService.action}
            key={bankService.id}
          />
        ))}
      </div>
    </div>
  );
};

export default BankServicesList;
