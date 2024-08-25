"use client";
import { useUser } from "@/contexts/UserContext";
import Image from "next/image";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";

import {
  LifeIcon,
  SafteyIcon,
  CheckingIcon,
  BusinessIcon,
  DebitIcon,
  SavingIcon,
} from "../serviceIcons/icons";

interface itemProp {
  icon: string;
  name: string;
  details: string;
}
const ServiceList = ({ icon, name, details }: itemProp) => {
  const { isDarkMode } = useUser();
  const first = name.split(" ")[0].toLowerCase();
  console.log(first, " this is first");
  const iconfunc = (params: string) => {
    if (params === "life") {
      return <LifeIcon />;
    } else if (params === "saftey") {
      return <SafteyIcon />;
    } else if (params === "checking") {
      return <CheckingIcon />;
    } else if (params === "savings") {
      return <SavingIcon />;
    } else if (params === "business") {
      return <BusinessIcon />;
    } else if (params === "debit") {
      return <DebitIcon />;
    } else {
      return <CheckingIcon />;
    }
  };

  return (
    <div
      className={`flex justify-between items-center p-3 md:p5 rounded-xl ${
        isDarkMode ? "bg-gray-600 text-gray-300" : "bg-white text-gray-900"
      }`}
    >
      {" "}
      <div className="flex ml-1 gap-2">
        <div className=" flex w-12 h-12 items-center justify-center bg-green-50 rounded-2xl ">
          {iconfunc(first)}
        </div>
        <div className="px-2">
          <div className="mt-1 font-medium">{name}</div>
          <div
            className={`text-xs ${
              isDarkMode ? "text-gray-400" : "text-gray-600"
            }`}
          >
            It&apos;s a long established fact.
          </div>
        </div>
      </div>
      <div className="hidden md:flex justify-between gap-4 w-1/2">
        <div className="px-2">
          <div className="mt-1 font-medium">Business loans</div>
          <div
            className={`text-xs ${
              isDarkMode ? "text-gray-400" : "text-gray-600"
            }`}
          >
            It&apos;s a long established fact.
          </div>
        </div>
        <div className="px-2">
          <div className="mt-1 font-medium">Business loans</div>
          <div
            className={`text-xs ${
              isDarkMode ? "text-gray-400" : "text-gray-600"
            }`}
          >
            It&apos;s a long established fact.
          </div>
        </div>
        <div className="px-2">
          <div className="mt-1 font-medium">Business loans</div>
          <div
            className={`text-xs ${
              isDarkMode ? "text-gray-400" : "text-gray-600"
            }`}
          >
            It&apos;s a long established fact.
          </div>
        </div>
      </div>
      <Dialog>
        <DialogTrigger
          className={`text-sm font-bold lg:border-2 lg:px-6 lg:rounded-full lg:my-auto lg:py-1 ${
            isDarkMode
              ? "text-gray-300 border-gray-600 bg-gray-700"
              : "text-[#1814F3] border-[#1814F3] bg-white"
          }`}
        >
          View Detail
        </DialogTrigger>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Details Of The Service</DialogTitle>
            <DialogDescription>{details}</DialogDescription>
          </DialogHeader>
        </DialogContent>
      </Dialog>
    </div>
  );
};

export default ServiceList;
