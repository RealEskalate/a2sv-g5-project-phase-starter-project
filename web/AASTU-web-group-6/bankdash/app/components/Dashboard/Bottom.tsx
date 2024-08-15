import React from "react";
import Image from "next/image";
import x from "../../../public/assets/profile-1.png";
import { AreaComp } from "../Charts/AreaComp";
const Bottom = () => {
  const people: any = [1, 2, 3];
  return (
    <section className="Botom w-full flex gap-6 ">
      <div className="cards-container w-2/5 center-content flex flex-col gap-6">
        <h1 className="flex grow page text-xl font-semibold text-colorBody-1">
          Expense Statistics
        </h1>

        <div className="flex gap-6 bg-white rounded-3xl border-solid border-2 border-gray-200 p-6">
          <div className="profle-box flex flex-col gap-4">
            <div className="flex gap-2">
              <div className="profile-item flex flex-col gap-1 p-6 items-center justify-center">
                <Image
                  className=" rounded-full "
                  src={"/assets/profile-1.png"}
                  width={70}
                  height={70}
                  alt=""
                />
                <div className="name text-base font-semibold">LIvia Bator</div>
                <div className="role text-base text-[#718EBF]">CEO</div>
              </div>
              <div className="profile-item flex flex-col gap-1 p-6 items-center justify-center">
                <Image
                  className=" rounded-full "
                  src={"/assets/profile-1.png"}
                  width={70}
                  height={70}
                  alt=""
                />
                <div className="name text-base font-semibold">LIvia Bator</div>
                <div className="role text-base text-[#718EBF]">CEO</div>
              </div>
              <div className="profile-item flex flex-col gap-1 p-6 items-center justify-center">
                <Image
                  className=" rounded-full "
                  src={"/assets/profile-1.png"}
                  width={70}
                  height={70}
                  alt=""
                />
                <div className="name text-base font-semibold">LIvia Bator</div>
                <div className="role text-base text-[#718EBF]">CEO</div>
              </div>
            </div>

            <div className="bottom flex gap-4">
              <div className="text-gray- text-base text-blue-900">
                Write Amount
              </div>
              <div className="flex items-center text-base text-[#718EBF] bg-[#EDF1F7] rounded-[50px] px-4 py-2">
                <span>525.50</span>
                <button className="p-4 rounded-[50px] bg-[#1814F3] text-white">
                  {" "}
                  Send
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div className="cards-container center-content flex flex-col gap-6 grow">
        <h1 className="flex grow page text-xl font-semibold text-colorBody-1">
          Balance History
        </h1>

        <div className="flex w-full h-80 gap-6 bg-white rounded-3xl border-solid border-2 border-gray-200">
          <AreaComp />
        </div>
      </div>
    </section>
  );
};

export default Bottom;
