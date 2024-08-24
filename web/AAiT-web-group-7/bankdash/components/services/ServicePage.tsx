import React from "react";
import ServiceCard from "./serviceCard";
import ServiceList from "./ServiceList";

const ServicePage = () => {
  return (
    <div className="flex flex-col py-10 space-y-10">
      <div className="flex justify-center gap-5">
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
          bgColor="#FFF5D9"
        />
        <ServiceCard
          imageSrc="pubimg/shield.svg"
          title="Safety"
          decription="We are your allies"
          bgColor="#DCFAF8"
        />
      </div>
      <div className="p-5 space-y-5">
        <div className="text-[#333B69] font-bold">Bank Services List</div>
        <div className="flex flex-col gap-4 justify-center">
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
    </div>
  );
};

export default ServicePage;
