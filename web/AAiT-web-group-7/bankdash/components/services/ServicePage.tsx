import React from "react";
import ServiceCard from "./serviceCard";
import ServiceList from "./ServiceList";

const ServicePage = () => {
  return (
    <div className="w-[1191px] h-[1000px] flex flex-col gap-[26px] p-10">
      <div className="w-[1100px] h-[120px] flex flex-row gap-[26px] ">
        <ServiceCard
          imageSrc="pubimg/key.svg"
          title="Life Insurance"
          decription="Unlimited protection"
          bgColor="#E7EDFF"
        />
        <ServiceCard
          imageSrc="pubimg/bag.svg"
          title="Shopping"
          decription="Buy. Think. Grow."
          bgColor="#FFE0EB"
        />
        <ServiceCard
          imageSrc="pubimg/shield.svg"
          title="Safety"
          decription="We are your allies"
          bgColor="#DCFAF8"
        />
      </div>

      <div className="text-[#333B69] w-fit font-bold">Bank Services List</div>

      <div className="flex flex-col gap-[20px] justify-center">
        <ServiceList
          imageSrc="pubimg/dollar.svg"
          title="Bussiness loans"
          bgColor="#FFE0EB"
        />
        <ServiceList
          imageSrc="pubimg/briefcase.svg"
          title="Checking accounts"
          bgColor="#FFF5D9"
        />
        <ServiceList
          imageSrc="pubimg/Group.svg"
          title="Savings accounts"
          bgColor="#FFE0EB"
        />
        <ServiceList
          imageSrc="pubimg/user.svg"
          title="Debit and credit cards"
          bgColor="#E7EDFF"
        />
        <ServiceList
          imageSrc="pubimg/shield.svg"
          title="Life Insurance"
          bgColor="#DCFAF8"
        />
        <ServiceList
          imageSrc="pubimg/dollar.svg"
          title="Bussiness loans"
          bgColor="#FFE0EB"
        />
      </div>
    </div>
  );
};

export default ServicePage;
